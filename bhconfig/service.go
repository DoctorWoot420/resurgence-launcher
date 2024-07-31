package bhconfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/DoctorWoot420/resurgence-launcher/log"

	"github.com/DoctorWoot420/resurgence-launcher/clients/bhconfig"
	"github.com/DoctorWoot420/resurgence-launcher/config"
)

// Service is responsible for all things related to the maphack config.
type Service interface {
	SetMaphackTextData(gameIndex int) error
}

type service struct {
	client        bhconfig.Client
	bhconfigModel *Model
	configService config.Service
	logger        log.Logger
}

// JSONItem represents a bhconfig item from JSON, before it's turned into a model item.
type JSONItem struct {
	ConfigText []string `json:"config_text"`
}

// SetMaphackTextData will fetch the new BH.cfg text from the server given our configuration parameters - also updates BH_settings.cfg with game/pass/gs
func (s *service) SetMaphackTextData(gameIndex int) error {
	s.logger.Debug("bhconfig/service.go start SetMaphackTextData")
	// Get the config data from game settings
	conf, err := s.configService.Read()
	if err != nil {
		s.logger.Debug("bridge/bhconfig.go broke reading data from configService")
		return err
	}
	// Doesn't feel like a good way to get the right game settings, hardcoded gameIndex passed in from GameSettings.qml right now
	g := conf.Games[gameIndex]
	s.logger.Debug("bridge/bhconfig.go location: " + g.Location)

	// Prepare our maphack config params
	maphackConfigParams := bhconfig.Payload{
		DefaultGs:       g.MaphackDefaultGs,
		DefaultGameName: g.MaphackDefaultGameName,
		DefaultPassword: g.MaphackDefaultPassword,
		RuneDesign:      g.MaphackRuneDesign,
		FilterBlocks:    g.MaphackFilterBlocks,
	}

	s.logger.Debug(fmt.Sprintf("g: %+v", g))
	s.logger.Debug(fmt.Sprintf("maphackConfigParams: %+v", maphackConfigParams))

	// Send params to server and get bfg.cfg text in return
	contents, err := s.client.GetMaphackTextFromParams(maphackConfigParams)
	if err != nil {
		s.logger.Debug("bridge/bhconfig.go broke getting GetMaphackTextFromParams from client")
		return err
	}
	bytes, err := ioutil.ReadAll(contents)
	if err != nil {
		s.logger.Debug("bridge/bhconfig.go broke preparing the response for writing")
		return err
	}

	// Write the contents to the file
	err = s.writeToFile(g.Location, bytes)
	if err != nil {
		s.logger.Debug("bridge/bhconfig.go broke writing the file to disk")
		return err
	}

	//Ugly add-on to update BH_settings.cfg - maybe get motivated and clean this up later
	s.logger.Debug("bhconfig/service.go about to call updateBhSettings")
	s.updateBhSettings(g.Location, maphackConfigParams)

	return nil
}

// newItem will create a new QObject item that we can pass to the model.
func newItem(item JSONItem) *Item {
	i := NewItem(nil)
	i.ConfigText = item.ConfigText

	return i
}

// NewService returns a service with all the dependencies.
func NewService(
	client bhconfig.Client,
	bhconfigModel *Model,
	configService config.Service,
	logger log.Logger, // Ensure logger is passed here
) Service {
	return &service{
		client:        client,
		configService: configService,
		logger:        logger, // Initialize logger
	}
}

func (s *service) updateBhSettings(location string, maphackConfigParams bhconfig.Payload) {
	s.logger.Debug("bhconfig/service.go start updateBhSettings")

	// Remove leading slash if present
	if strings.HasPrefix(location, "\\") || strings.HasPrefix(location, "/") {
		location = location[1:]
		s.logger.Debug(fmt.Sprintf("Removed leading slash: %s", location))
	}
	settingsFilePath := filepath.Join(location, "BH_settings.cfg")
	s.logger.Debug(fmt.Sprintf("Settings file path: %s", settingsFilePath))

	// Read the existing file content
	fileBytes, err := ioutil.ReadFile(settingsFilePath)
	if err != nil {
		s.logger.Debug(fmt.Sprintf("Failed to read settings file: %v", err))
		return
	}

	content := string(fileBytes)

	// Update Default Game Name
	content = updateConfigLine(content, "Default Game Name", maphackConfigParams.DefaultGameName)

	// Update Default Password
	content = updateConfigLine(content, "Default Password", maphackConfigParams.DefaultPassword)

	// Update Default Gs with index transformation
	gsIndex := transformGs(maphackConfigParams.DefaultGs)
	content = updateConfigLine(content, "Default Gs", fmt.Sprintf("%d", gsIndex))

	// Write the updated content back to the file
	err = ioutil.WriteFile(settingsFilePath, []byte(content), 0644)
	if err != nil {
		s.logger.Debug(fmt.Sprintf("Failed to write updated settings file: %v", err))
		return
	}

	s.logger.Debug("bhconfig/service.go stop updateBhSettings")
}

func updateConfigLine(content, key, value string) string {
	lines := strings.Split(content, "\n")
	keyPattern := regexp.MustCompile(`^` + key + `\s*=`)
	found := false

	for i, line := range lines {
		if keyPattern.MatchString(line) {
			if value == "" {
				lines[i] = key + " = "
			} else {
				lines[i] = key + " = " + value
			}
			found = true
			break
		}
	}

	if !found && value != "" {
		newLine := key + " = " + value
		lines = append(lines, newLine)
	}

	return strings.Join(lines, "\n")
}

// Helper function to transform the DefaultGs string to an index
func transformGs(gs string) int {
	switch gs {
	case "GS1 - New York":
		return 0
	case "GS2 - Los Angeles":
		return 1
	case "GS3 - Amsterdam":
		return 2
	default:
		return -1 // handle invalid GS values
	}
}

func (s *service) writeToFile(location string, bytes []byte) error {
	// Log the initial location
	s.logger.Debug(fmt.Sprintf("Initial location: %s", location))

	// Remove leading slash if present
	if strings.HasPrefix(location, "\\") || strings.HasPrefix(location, "/") {
		location = location[1:]
		s.logger.Debug(fmt.Sprintf("Removed leading slash: %s", location))
	}

	// Ensure the location is correctly formatted
	location = filepath.FromSlash(location)
	s.logger.Debug(fmt.Sprintf("Formatted location: %s", location))

	// Check if the location is an absolute path
	if !filepath.IsAbs(location) {
		s.logger.Debug(fmt.Sprintf("Location is not an absolute path: %s", location))
		return fmt.Errorf("location must be an absolute path: %s", location)
	}

	// Define the file path
	filePath := filepath.Join(location, "bh.cfg")
	s.logger.Debug(fmt.Sprintf("File path: %s", filePath))

	// Ensure the directory exists
	dir := filepath.Dir(filePath)
	s.logger.Debug(fmt.Sprintf("Directory path: %s", dir))
	if err := os.MkdirAll(dir, 0755); err != nil {
		s.logger.Debug(fmt.Sprintf("failed to create directories: %v", err))
		return err
	}

	// Parse the JSON to extract the config field
	var jsonData map[string]string
	if err := json.Unmarshal(bytes, &jsonData); err != nil {
		s.logger.Debug(fmt.Sprintf("Failed to unmarshal JSON: %v", err))
		return err
	}

	// Extract the "config" field from the JSON object
	configContent, ok := jsonData["config"]
	if !ok {
		return fmt.Errorf("config field not found in the JSON data")
	}

	// Replace \n with actual newline characters in the config content
	configContent = strings.ReplaceAll(configContent, "\\n", "\n")

	// Write the extracted config content to the file
	err := ioutil.WriteFile(filePath, []byte(configContent), 0644)
	if err != nil {
		s.logger.Debug(fmt.Sprintf("Failed to write to file %s: %v", filePath, err))
		return err
	}
	s.logger.Debug(fmt.Sprintf("File %s updated successfully", filePath))

	return nil
}
