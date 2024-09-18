package bhconfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/DoctorWoot420/resurgence-launcher/clients/bhconfig"
	"github.com/DoctorWoot420/resurgence-launcher/config"
	"github.com/DoctorWoot420/resurgence-launcher/log"
)

type Service interface {
	SetMaphackTextData(gameIndex int) error
}

type service struct {
	client        bhconfig.Client
	bhconfigModel *Model
	configService config.Service
	logger        log.Logger
}

func NewService(
	client bhconfig.Client,
	bhconfigModel *Model,
	configService config.Service,
	logger log.Logger,
) Service {
	return &service{
		client:        client,
		configService: configService,
		logger:        logger,
	}
}

func (s *service) SetMaphackTextData(gameIndex int) error {
	s.logger.Debug("bhconfig/service.go start SetMaphackTextData")

	conf, err := s.configService.Read()
	if err != nil {
		s.logger.Debug("Failed to read data from configService")
		return err
	}

	g := conf.Games[gameIndex]
	s.logger.Debug(fmt.Sprintf("Game settings: %+v", g))

	maphackConfigParams := bhconfig.Payload{
		DefaultGs:       g.MaphackDefaultGs,
		DefaultGameName: g.MaphackDefaultGameName,
		DefaultPassword: g.MaphackDefaultPassword,
		RuneDesign:      g.MaphackRuneDesign,
		FilterBlocks:    g.MaphackFilterBlocks,
	}

	s.logger.Debug(fmt.Sprintf("Maphack config params: %+v", maphackConfigParams))

	contents, err := s.client.GetMaphackTextFromParams(maphackConfigParams)
	if err != nil {
		s.logger.Debug("Failed to get maphack text from params")
		return err
	}

	bytes, err := ioutil.ReadAll(contents)
	if err != nil {
		s.logger.Debug("Failed to read response contents")
		return err
	}

	if err := s.updateBHConfig(g.Location, bytes); err != nil {
		s.logger.Debug("Failed to update BH.cfg")
		return err
	}

	if err := s.updateBHSettings(g.Location, maphackConfigParams); err != nil {
		s.logger.Debug("Failed to update BH_settings.cfg")
		return err
	}

	return nil
}

func (s *service) updateBHConfig(location string, bytes []byte) error {
	s.logger.Debug("Start updateBHConfig")

	location = cleanLocation(location)
	filePath := filepath.Join(location, "bh.cfg")

	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		s.logger.Debug(fmt.Sprintf("Failed to create directories: %v", err))
		return err
	}

	var jsonData map[string]string
	if err := json.Unmarshal(bytes, &jsonData); err != nil {
		s.logger.Debug(fmt.Sprintf("Failed to unmarshal JSON: %v", err))
		return err
	}

	configContent, ok := jsonData["config"]
	if !ok {
		return fmt.Errorf("config field not found in the JSON data")
	}

	configContent = strings.ReplaceAll(configContent, "\\n", "\n")

	existingContent, err := ioutil.ReadFile(filePath)
	if err == nil {
		customLines, _ := extractCustomBuildLines(string(existingContent))
		configContent = mergeContent(configContent, customLines)
	}

	if err := ioutil.WriteFile(filePath, []byte(configContent), 0644); err != nil {
		s.logger.Debug(fmt.Sprintf("Failed to write to file %s: %v", filePath, err))
		return err
	}

	s.logger.Debug(fmt.Sprintf("File %s updated successfully", filePath))
	return nil
}

func (s *service) updateBHSettings(location string, params bhconfig.Payload) error {
	s.logger.Debug("Start updateBHSettings")

	location = cleanLocation(location)
	filePath := filepath.Join(location, "BH_settings.cfg")

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		s.logger.Debug(fmt.Sprintf("Failed to read settings file: %v", err))
		return err
	}

	contentStr := string(content)
	contentStr = updateConfigLine(contentStr, "Default Game Name", params.DefaultGameName)
	contentStr = updateConfigLine(contentStr, "Default Password", params.DefaultPassword)
	gsIndex := transformGs(params.DefaultGs)
	contentStr = updateConfigLine(contentStr, "Default Gs", fmt.Sprintf("%d", gsIndex))

	if err := ioutil.WriteFile(filePath, []byte(contentStr), 0644); err != nil {
		s.logger.Debug(fmt.Sprintf("Failed to write updated settings file: %v", err))
		return err
	}

	s.logger.Debug("BH_settings.cfg updated successfully")
	return nil
}

func cleanLocation(location string) string {
	location = strings.TrimPrefix(location, "\\")
	location = strings.TrimPrefix(location, "/")
	return filepath.FromSlash(location)
}

func extractCustomBuildLines(content string) (string, error) {
	startMarker := "// START CUSTOM BUILD LINES - DO NOT EDIT THIS SEPARATOR TEXT"
	endMarker := "// END CUSTOM BUILD LINES - DO NOT EDIT THIS SEPARATOR TEXT"

	startIndex := strings.Index(content, startMarker)
	endIndex := strings.Index(content, endMarker)

	if startIndex == -1 || endIndex == -1 {
		return "", nil // No custom build lines found
	}

	endIndex += len(endMarker)
	return content[startIndex:endIndex], nil
}

func mergeContent(newContent, customLines string) string {
	if customLines == "" {
		return newContent
	}

	startMarker := "// START CUSTOM BUILD LINES - DO NOT EDIT THIS SEPARATOR TEXT"
	endMarker := "// END CUSTOM BUILD LINES - DO NOT EDIT THIS SEPARATOR TEXT"

	startIndex := strings.Index(newContent, startMarker)
	endIndex := strings.Index(newContent, endMarker)

	if startIndex == -1 || endIndex == -1 {
		return newContent + "\n\n" + customLines
	}

	return newContent[:startIndex] + customLines + newContent[endIndex+len(endMarker):]
}

func updateConfigLine(content, key, value string) string {
	lines := strings.Split(content, "\n")
	keyPattern := regexp.MustCompile(`^` + regexp.QuoteMeta(key) + `\s*:`)
	found := false

	for i, line := range lines {
		if keyPattern.MatchString(line) {
			if value == "" {
				lines[i] = key + ":"
			} else {
				lines[i] = key + ": " + value
			}
			found = true
			break
		}
	}

	if !found && value != "" {
		newLine := key + ": " + value
		lines = append(lines, newLine)
	}

	return strings.Join(lines, "\n")
}

func transformGs(gs string) int {
	switch gs {
	case "GS1 - New York":
		return 0
	case "GS2 - Los Angeles":
		return 1
	case "GS3 - Amsterdam":
		return 2
	default:
		return -1
	}
}
