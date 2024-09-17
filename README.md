# Resurgence launcher

![launcher screenshot](/docs/launcher.png)

## About the project

This project is a launcher for Diablo II Resurgence. It was built to help new players install patches, updating registries and help with other technical issues to lower the barrier of entry into the Resurgence community, while also assisting more experienced players with more advanced settings such as HD mods and launching multiple boxes.  This project is a copy-paste of Nokka's SlashDiablo launcher with a few additions.

## Installing
If you want to use the launcher and play, grab the msi from the latest version on the [releases page]([url](https://github.com/DoctorWoot420/resurgence-launcher/releases)).

For any qustions about using or installing see the [Resurgence Launcher FAQs]([url](https://diablo2resurgence.fandom.com/wiki/Resurgence_Launcher_FAQs)) wiki page.

## Important credits

This project distributes the work of many folk in the community, primary credits go to Nokka for the launcher, Mir Drualga for the HD code, the many folk who have contributed to BH, and Bayaraa/Jarcho for d2gl components.

## Features

- [x] Patching Diablo II up to 1.13c from previous game versions
- [x] Applying Resurgence patches automatically
- [x] Allows for multiple installs of Diablo II with different settings (such as Maphack & HD)
- [x] Automatically installs and updates Maphack & HD mod
- [x] Launch multiple Diablo II boxes from multiple installs
- [x] Help with OS specific configuration such as DEP issues
- [x] Apply custom bh loot filter and customization for default game server, game name, and password

### Full OS support

- [x] Windows
- [ ] OSX (missing some D2 specific features)
- [ ] Linux (missing some D2 specific features)

## Development

### Go

Install Go 1.12 or higher by following [install instructions](http://golang.org/doc/install.html) for your OS.

### Qt bindings for Go

Before you can build you need to install the [Go/Qt bindings](https://github.com/therecipe/qt/wiki/Installation#regular-installation).

### Install Qt5

#### OSX

On OSX using brew is by far the most simple way of installing Qt5.

```bash
$ brew install qt
```

#### Windows

Use the [installer](https://download.qt.io/official_releases/qt/5.13/5.13.0/qt-opensource-windows-x86-5.13.0.exe) provided by Qt (Make sure you install the MinGW build of Qt).

#### Building Slashdiablo launcher

```bash
# Get binding source
$ go get -u -v -tags=no_env github.com/therecipe/qt/cmd/...

# Download the repository with dependencies
$ go get -d -u -v github.com/nokka/slashdiablo-launcher

# Build the launcher
$ cd $(go env GOPATH)/src/github.com/nokka/slashdiablo-launcher
$ qtdeploy build

# Start launcher (different depending on OS)
$ ./deploy/darwin/slashdiablo-launcher.app/Contents/MacOS/slashdiablo-launcher
```

## Deploying

Deploying to a target can be done from any host OS if there's a docker image available,
otherwise the target OS and the host must be the same.

### Windows

#### Build in docker

```bash
$ docker pull therecipe/qt:windows_64_static
$ qtdeploy -docker build windows_64_static
```

#### Build on local machine

```bash
$ qtdeploy build desktop
```

#### Updating application binary version and manifest data

```bash
# Download goversioninfo tool
$ go get github.com/josephspurrier/goversioninfo/cmd/goversioninfo

# Make your changes to the manifest file.
$ vim versioninfo.json

# Generate a new resource.syso including manifest.
$  go generate
```

### MacOS (from MacOS only)

```bash
$ qtdeploy build darwin github.com/nokka/slashdiablo-launcher
```

### Additional windows environment setup notes
Might clean this up later, but want to at least leave breadcrumbs.  The above instructions are direct copies from Nokka's repo, but I struggled a lot to build the application and get an environment running.  I do not recall the exact steps required but was eventually able to get this working.  Hopefully these help someone else (or myself down the line):

General install steps from raw notes, some of these steps may not have been required, but it's what I did:
- Download go from go website - ended up  with go version go1.22.2 windows/amd64
- Download Qt5.13.0 from QT website site - open source version
- In QT installer make sure to install qscript and mingw64 options, they are disabled by default
- Do the go bindings env variables from therecipe repo's installation instructions for windows

        These are the the relevent environment variables I ended up with:
        QT_API=5.13.0
        QT_DIR=C:\Qt
        QT_VERSION=5.13.0
        GOPATH=C:\Users\woahc\go
        Path = ... C:\Program Files\Go\bin ... (only relevant one I see)
        note that qmake doesn't work  for me but I saw it in a ton of documentation and it threw me on the wrong path.  qtdeploy can work when qmake doesnt.

- Proceed with the remaining install steps from therecipe/qt

I did a bunch of fuckery around this point, cant remember it all.  Good luck.

- Run qtsetup Takes a while, should work after this
- Now run qtdeploy build desktop 

All of this stuff was done in the source directory.  I didn't do this command mentioned above: cd $(go env GOPATH)/src/github.com/nokka/slashdiablo-launcher
Instead I was just working within c:/www/slashdiablo-launcher running all of the related install and deployment commands.

To build an msi there are scripts to do it better, but one solution is:
Use qtdeploy build desktop, then use this program to make an msi
https://www.advancedinstaller.com/user-guide/tutorial-create-simple-msi-installer.html 



Then for setting up the patch server the key thing to know is about the file manifests.  You need to update the files you want, then each folder needs a new manifest.json.  You'll need to have a script that generates the crc value with the expected hash so it matches the local machine.  If you don't get these values right you'll get stuck with file sync errors.  This is a basic script from chat gpt that got the job done, but is certainly not an efficient method at all.

Also consider that we need to set the flag ignore_crc in the manifest for any config files.  Specifically d2gl.ini and d2fps.ini are packaged in to help users avoid setup steps, but we don't want to sync back over their changed config.


# Remove previous type if already defined
if ([System.Management.Automation.PSTypeName]'CRC32'.Type) {
    Remove-TypeData -TypeName CRC32 -ErrorAction SilentlyContinue
    Remove-TypeData -TypeName CRC32 -ErrorAction SilentlyContinue
}

# Add the CRC32 class with corrected scope and accessibility
Add-Type -TypeDefinition @"
using System;
using System.IO;

public class CRC32 {
    uint[] table;

    public uint ComputeChecksum(byte[] bytes) {
        uint crc = 0xffffffff;
        for (int i = 0; i < bytes.Length; i++) {
            byte index = (byte)((crc & 0xff) ^ bytes[i]);
            crc = (crc >> 8) ^ table[index];
        }
        return ~crc;
    }

    public CRC32() {
        uint poly = 0xedb88320;
        table = new uint[256];
        for (int i = 0; i < 256; i++) {
            uint crc = (uint)i;
            for (int j = 8; j > 0; j--)
                crc = (crc & 1) == 1 ? (crc >> 1) ^ poly : crc >> 1;
            table[i] = crc;
        }
    }

    public string ComputeChecksumString(byte[] bytes) {
        return ComputeChecksum(bytes).ToString("x8");  // Output CRC in lowercase
    }
}
"@

# Compute CRC for each file and prepare the manifest
$files = Get-ChildItem -File | ForEach-Object {
    $crcInstance = New-Object CRC32
    $fileData = [System.IO.File]::ReadAllBytes($_.FullName)
    $crc = $crcInstance.ComputeChecksumString($fileData)
    @{
        name = $_.Name
        crc = $crc
        last_modified = $_.LastWriteTime.ToString("yyyy-MM-ddTHH:mm:ss.fffffffK")
        content_length = $_.Length
        ignore_crc = $false
        deprecated = $false
    }
}

# Convert file information to JSON and save to 'manifest.json'
$manifest = @{
    files = $files
}
$manifest | ConvertTo-Json -Depth 100 | Set-Content -Path 'manifest.json'


