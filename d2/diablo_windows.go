// +build windows

package d2

import (
	"bufio"
	"bytes"
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
	"unicode/utf8"

	"golang.org/x/sys/windows/registry"
)

// SHA1 of the different versions of Diablo Game.exe.
var hashList = map[string]string{
	"a875b98fa3a8b9300bcc04c84be1fa057eb277b5": "1.12",
	"af2b33c90b50ede8d9a8bca9b8d9720c87f78641": "1.13c",
	"27ddadbc457affed122564ae7a4bd2223181e15a": "1.13c", // Custom 1.13c build with HD icon.
	"11cd918cb6906295769d9be1b3e349e02af6b229": "1.13d",
	"3e64f12c6ef72847f49d301c2472280d4460589d": "1.14a",
	"11e940266c6838414c2114c2172227f982d4054e": "1.14b",
	"255691dd53e3bcd646e5c6e1e2e7b16da745b706": "1.14c",
	"af0ea93d2a652ceb11ac01ee2e4ae1ef613444c2": "1.14d",
}

const (
	// ModMaphackIdentifier is the identifier we use to look for installs of maphack.
	ModMaphackIdentifier = "BH.dll"

	// ModHDIdentifier is the identifier we use to look for installs of hd mod.
	ModHDIdentifier = "D2HD.dll"

	errRegistryKeyNotFound = "The system cannot find the file specified."
)

// validate113cVersion will check the given installations Diablo II version.
func validate113cVersion(path string) (bool, error) {
	// Open local Game.exe.
	content, err := ioutil.ReadFile(localizePath(path) + "\\Game.exe")
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	// Hash the content of the Game.exe.
	hashed := fmt.Sprintf("%x", sha1.Sum(content))

	// Check the game version.
	version, ok := hashList[hashed]

	// Unknown game version.
	if !ok {
		return false, nil
	}

	return version == "1.13c", nil
}

// launch will execute the Diablo II.exe in the given directory.
func launch(path string, flags []string, done chan execState) (*int, error) {
	// Localize the path.
	localized := localizePath(path)

	// Make sure Windows registry keys are correctly setup before launch.
	setDiabloRegistryKeys()

	// Exec the Diablo II.exe with the given command line args.
	cmd := exec.Command(localized+"\\Diablo II.exe", flags...)
	cmd.Dir = localized

	// Collect the output from the command.
	var stderr bytes.Buffer

	// Pipe errors to our buffer.
	cmd.Stderr = &stderr

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	// Wait on separate thread.
	go func() {
		if err := cmd.Wait(); err != nil {
			if exiterr, ok := err.(*exec.ExitError); ok {
				// The program has exited with an exit code != 0
				if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
					done <- execState{pid: &cmd.Process.Pid, err: fmt.Errorf("Exit status: %d : %s", status.ExitStatus(), stderr.String())}
				}
			} else {
				// Was some other wait error such as permissions, return the err.
				done <- execState{pid: &cmd.Process.Pid, err: fmt.Errorf("cmd.Wait: %d : %s", err, stderr.String())}
			}
		}

		done <- execState{pid: &cmd.Process.Pid, err: nil}
	}()

	return &cmd.Process.Pid, nil
}

// configureForOS will set specific configurations, such as compatibility mode.
func configureForOS(path string) error {
	// The key name is the localized path for the Diablo II directory.
	keyName := fmt.Sprintf("%s\\%s", localizePath(path), "Game.exe")

	// Open the compatibility key directory.
	compatibilityKey, err := registry.OpenKey(registry.CURRENT_USER,
		`Software\Microsoft\Windows NT\CurrentVersion\AppCompatFlags\Layers`,
		registry.QUERY_VALUE|registry.SET_VALUE,
	)
	if err != nil {
		return err
	}

	// Set Windows XP Service Pack 2 compatibility mode.
	if err := compatibilityKey.SetStringValue(keyName, "~ WINXPSP2"); err != nil {
		return err
	}

	// Close the registry when we're done.
	if err := compatibilityKey.Close(); err != nil {
		return err
	}

	return nil
}

// applyDEP will run a fix to disable DEP.
func applyDEP(path string) error {
	// Localize the path.
	localized := localizePath(path)

	// Channel to send that we're done on.
	done := make(chan bool, 1)

	// Use cmd.exe to call the bat file.
	cmd := exec.Command("cmd.exe", "/C", "call", "DEP_fix_v2.bat")
	cmd.Dir = localized

	// Capture stdin for the command, so we can send data on it.
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	r, w, err := os.Pipe()
	if err != nil {
		return err
	}

	cmd.Stdout = w
	cmd.Stderr = w

	// Log output of the bat file.
	go func() {
		scanner := bufio.NewScanner(r)
		var linesWritten int

		for scanner.Scan() {
			linesWritten++

			// Kill the go routine when all the output has been captured.
			if linesWritten >= 5 {
				// Write any key to the cmd to exit.
				_, err = io.WriteString(stdin, "a")
				if err != nil {
					done <- false
					return
				}

				// Notify main thread that the DEP change has been applied successfully.
				done <- true
				return
			}
		}
	}()

	// Start won't wait for the program to finish,
	// this is because we need to send input to the program.
	err = cmd.Start()
	if err != nil {
		return err
	}

	// Wait for output on the done channel or simply timeout the action after 5 sec.
	for {
		select {
		case result := <-done:
			if !result {
				return errors.New("error while applying DEP change")
			}
			return nil
		case <-time.After(time.Second * 5):
			return errors.New("timeout while applying DEP change")
		}
	}
}

// setDiabloRegistryKeys will remove the registry for BNETIP and set CmdLine options.
func setDiabloRegistryKeys() error {
	// Open the Battle net configuration directory.
	confKey, err := registry.OpenKey(registry.CURRENT_USER, `Software\Blizzard Entertainment\Diablo II`, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		return err
	}

	bnetIP, _, err := confKey.GetStringValue("BNETIP")
	// If getting the string value errors and the error is something other than not found, return err.
	if err != nil && err.Error() != errRegistryKeyNotFound {
		return err
	}

	// If the user has the BNETIP set, let's remove it, not to mess other installs.
	if bnetIP != "" {
		err = confKey.DeleteValue("BNETIP")
		if err != nil {
			return err
		}
	}

	// Set the Command line args when starting.
	if err := confKey.SetStringValue("CmdLine", "-skiptobnet"); err != nil {
		return err
	}

	// Close the registry when we're done.
	if err := confKey.Close(); err != nil {
		return err
	}

	return nil
}

// setGateway will set the gateway for Diablo II.
func setGateway(gateway string) error {
	var gatewayHex []byte

	switch gateway {
	case GatewaySlashdiablo:
		gatewayHex = getSlashGateway()
	case GatewayBattleNet:
		gatewayHex = getBattleNetGateway()
	}

	// Open the Battle.net configuration registry directory.
	gatewayKey, err := registry.OpenKey(registry.CURRENT_USER, `Software\Battle.net\Configuration`, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		return err
	}

	// Set the gateway hex.
	if err := gatewayKey.SetBinaryValue("Diablo II Battle.net Gateways", gatewayHex); err != nil {
		return err
	}

	// Close the registry when we're done.
	if err := gatewayKey.Close(); err != nil {
		return err
	}

	return nil
}

func isModInstalled(path string, identifier string, manifest *Manifest) (bool, error) {
	filePath := localizePath(fmt.Sprintf("%s/%s", path, identifier))

	// Get the checksum from the file on disk.
	hashed, err := hashCRC32(filePath, polynomial)
	if err != nil {
		// The file doesn't exist on disk, so it's not installed.
		if err == ErrCRCFileNotFound {
			return false, nil
		}

		return false, err
	}

	var crc string
	// File exists on disk, find the CRC.
	for _, f := range manifest.Files {
		if f.Name == identifier {
			crc = f.CRC
			break
		}
	}

	if crc == hashed {
		return true, nil
	}

	return false, nil
}

// localizePath will localize the path for the OS.
func localizePath(path string) string {
	// Windows uses backslashes for paths, so we'll reverse them.
	reversed := strings.Replace(path, "/", "\\", -1)

	// Remove the heading backslash.
	_, i := utf8.DecodeRuneInString(reversed)

	return reversed[i:]
}

func getSlashGateway() []byte {
	return []byte{
		0x31, 0x30, 0x30, 0x32,
		0x00, 0x30, 0x31, 0x00,
		0x70, 0x6c, 0x61, 0x79,
		0x2e, 0x73, 0x6c, 0x61,
		0x73, 0x68, 0x64, 0x69,
		0x61, 0x62, 0x6c, 0x6f,
		0x2e, 0x6e, 0x65, 0x74,
		0x00, 0x2d, 0x36, 0x00,
		0x73, 0x6c, 0x61, 0x73,
		0x68, 0x00, 0x65, 0x76,
		0x6e, 0x74, 0x2e, 0x73,
		0x6c, 0x61, 0x73, 0x68,
		0x64, 0x69, 0x61, 0x62,
		0x6c, 0x6f, 0x2e, 0x6e,
		0x65, 0x74, 0x00, 0x38,
		0x00, 0x45, 0x76, 0x65,
		0x6e, 0x74, 0x00, 0x00,
	}
}

func getBattleNetGateway() []byte {
	return []byte{
		0x31, 0x30, 0x30, 0x32,
		0x00, 0x30, 0x31, 0x00,
		0x75, 0x73, 0x77, 0x65,
		0x73, 0x74, 0x2e, 0x62,
		0x61, 0x74, 0x74, 0x6c,
		0x65, 0x2e, 0x6e, 0x65,
		0x74, 0x00, 0x38, 0x00,
		0x55, 0x2e, 0x53, 0x2e,
		0x20, 0x57, 0x65, 0x73,
		0x74, 0x00, 0x75, 0x73,
		0x65, 0x61, 0x73, 0x74,
		0x2e, 0x62, 0x61, 0x74,
		0x74, 0x6c, 0x65, 0x2e,
		0x6e, 0x65, 0x74, 0x00,
		0x36, 0x00, 0x55, 0x2e,
		0x53, 0x2e, 0x20, 0x45,
		0x61, 0x73, 0x74, 0x00,
		0x61, 0x73, 0x69, 0x61,
		0x2e, 0x62, 0x61, 0x74,
		0x74, 0x6c, 0x65, 0x2e,
		0x6e, 0x65, 0x74, 0x00,
		0x2d, 0x39, 0x00, 0x41,
		0x73, 0x69, 0x61, 0x00,
		0x65, 0x75, 0x72, 0x6f,
		0x70, 0x65, 0x2e, 0x62,
		0x61, 0x74, 0x74, 0x6c,
		0x65, 0x2e, 0x6e, 0x65,
		0x74, 0x00, 0x2d, 0x31,
		0x00, 0x45, 0x75, 0x72,
		0x6f, 0x70, 0x65, 0x00,
		0x00,
	}
}
