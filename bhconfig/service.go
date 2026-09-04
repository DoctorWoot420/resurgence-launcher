package bhconfig

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/DoctorWoot420/resurgence-launcher/clients/bhconfig"
	"github.com/DoctorWoot420/resurgence-launcher/config"
	"github.com/DoctorWoot420/resurgence-launcher/log"
	"github.com/DoctorWoot420/resurgence-launcher/storage"
)

type bhSyncCache struct {
	ParamsHash string `json:"params_hash"`
	SourceSHA  string `json:"source_sha"`
}

type ProgressFunc func(message string, progress float32)

type Service interface {
	SetMaphackTextData(gameIndex int) error
	SyncMaphackConfigsIfChanged(status ProgressFunc) (updated bool, err error)
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

	if gameIndex < 0 || gameIndex >= len(conf.Games) {
		return fmt.Errorf("invalid game index: %d", gameIndex)
	}

	_, err = s.applyGameMaphackConfig(conf.Games[gameIndex], true, s.currentFilterSourceSHA(), nil)
	return err
}

func (s *service) SyncMaphackConfigsIfChanged(status ProgressFunc) (bool, error) {
	s.logger.Debug("bhconfig/service.go start SyncMaphackConfigsIfChanged")

	games, err := s.eligibleMaphackGames()
	if err != nil {
		return false, err
	}
	if len(games) == 0 {
		return false, nil
	}

	sourceSHA := s.currentFilterSourceSHA()

	var updated bool
	var syncErr error
	for i, g := range games {
		changed, err := s.applyGameMaphackConfig(g, false, sourceSHA, func() {
			if status != nil {
				progress := 0.2 + (float32(i) / float32(len(games)) * 0.7)
				status("Downloading new BH.cfg...", progress)
			}
		})
		if err != nil {
			s.logger.Debug(fmt.Sprintf("Failed to sync BH.cfg for game %s: %v", g.ID, err))
			if syncErr == nil {
				syncErr = err
			}
			continue
		}
		if changed {
			updated = true
			if status != nil {
				status("Updating BH.cfg...", 0.9)
			}
		}
	}

	if updated && status != nil {
		status("Updated BH.cfg", 1)
	}

	return updated, syncErr
}

func (s *service) eligibleMaphackGames() ([]storage.Game, error) {
	conf, err := s.configService.Read()
	if err != nil {
		s.logger.Debug("Failed to read data from configService")
		return nil, err
	}

	var games []storage.Game
	for _, g := range conf.Games {
		if g.Location == "" || g.OverrideBHCfg {
			continue
		}
		if g.MaphackVersion == "" || g.MaphackVersion == config.ModVersionNone {
			continue
		}
		games = append(games, g)
	}

	return games, nil
}

func (s *service) currentFilterSourceSHA() string {
	sha, err := s.client.GetFilterSourceSHA()
	if err != nil {
		s.logger.Debug(fmt.Sprintf("Failed to get filter source sha: %v", err))
		return ""
	}
	return sha
}

func (s *service) applyGameMaphackConfig(g storage.Game, forceDownload bool, sourceSHA string, onDownload func()) (bool, error) {
	s.logger.Debug(fmt.Sprintf("Game settings: %+v", g))

	itemNameOption := g.MaphackItemNameOption
	if itemNameOption == "" {
		itemNameOption = "Default"
	}

	maphackConfigParams := bhconfig.Payload{
		DefaultGs:       g.MaphackDefaultGs,
		DefaultGameName: g.MaphackDefaultGameName,
		DefaultPassword: g.MaphackDefaultPassword,
		RuneDesign:      g.MaphackRuneDesign,
		ItemNameOption:  itemNameOption,
		FilterBlocks:    g.MaphackFilterBlocks,
	}

	s.logger.Debug(fmt.Sprintf("Maphack config params: %+v", maphackConfigParams))

	paramsHash := bhCfgParamsHash(maphackConfigParams)
	cache := readSyncCache(g.Location)
	paramsMatch := cache.ParamsHash == paramsHash
	sourceMatch := sourceSHA != "" && cache.SourceSHA == sourceSHA
	if !forceDownload && localBHExists(g.Location) && paramsMatch && (sourceMatch || sourceSHA == "") {
		s.logger.Debug("BH.cfg source and params unchanged, skipping download")
		settingsUpdated, err := s.updateBHSettings(g.Location, maphackConfigParams)
		return settingsUpdated, err
	}

	if onDownload != nil {
		onDownload()
	}

	contents, err := s.client.GetMaphackTextFromParams(maphackConfigParams)
	if err != nil {
		s.logger.Debug("Failed to get maphack text from params")
		return false, err
	}
	defer contents.Close()

	bytes, err := ioutil.ReadAll(contents)
	if err != nil {
		s.logger.Debug("Failed to read response contents")
		return false, err
	}

	bhUpdated, err := s.updateBHConfig(g.Location, bytes)
	if err != nil {
		s.logger.Debug("Failed to update BH.cfg")
		return false, err
	}

	if sourceSHA == "" {
		sourceSHA = cache.SourceSHA
	}
	writeSyncCache(g.Location, paramsHash, sourceSHA)

	settingsUpdated, err := s.updateBHSettings(g.Location, maphackConfigParams)
	if err != nil {
		s.logger.Debug("Failed to update BH_settings.cfg")
		return bhUpdated, err
	}

	return bhUpdated || settingsUpdated, nil
}

func bhCfgParamsHash(params bhconfig.Payload) string {
	blocks := append([]string(nil), params.FilterBlocks...)
	sort.Strings(blocks)
	raw := strings.Join([]string{params.RuneDesign, params.ItemNameOption, strings.Join(blocks, ",")}, "\n")
	sum := sha256.Sum256([]byte(raw))
	return hex.EncodeToString(sum[:])
}

func syncCachePath(location string) string {
	return filepath.Join(cleanLocation(location), "bh.cfg.sync")
}

func localBHExists(location string) bool {
	_, err := os.Stat(filepath.Join(cleanLocation(location), "bh.cfg"))
	return err == nil
}

func readSyncCache(location string) bhSyncCache {
	var cache bhSyncCache
	data, err := ioutil.ReadFile(syncCachePath(location))
	if err != nil {
		return cache
	}
	_ = json.Unmarshal(data, &cache)
	return cache
}

func writeSyncCache(location string, paramsHash string, sourceSHA string) {
	data, err := json.Marshal(bhSyncCache{ParamsHash: paramsHash, SourceSHA: sourceSHA})
	if err != nil {
		return
	}
	_ = ioutil.WriteFile(syncCachePath(location), data, 0644)
}

func (s *service) updateBHConfig(location string, bytes []byte) (bool, error) {
	s.logger.Debug("Start updateBHConfig")

	location = cleanLocation(location)
	filePath := filepath.Join(location, "bh.cfg")

	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		s.logger.Debug(fmt.Sprintf("Failed to create directories: %v", err))
		return false, err
	}

	var jsonData map[string]string
	if err := json.Unmarshal(bytes, &jsonData); err != nil {
		s.logger.Debug(fmt.Sprintf("Failed to unmarshal JSON: %v", err))
		s.logger.Debug(fmt.Sprintf("Raw JSON data: %s", string(bytes))) // Output the raw JSON
		return false, err
	}

	configContent, ok := jsonData["config"]
	if !ok {
		return false, fmt.Errorf("config field not found in the JSON data: %s", string(bytes)) // Output the full JSON
	}

	configContent = strings.ReplaceAll(configContent, "\\n", "\n")

	existingContent, err := ioutil.ReadFile(filePath)
	if err == nil {
		customLines, _ := extractCustomBuildLines(string(existingContent))
		configContent = mergeContent(configContent, customLines)
		if normalizeBHContent(string(existingContent)) == normalizeBHContent(configContent) {
			s.logger.Debug(fmt.Sprintf("File %s is unchanged, skipping write", filePath))
			return false, nil
		}
	}

	if err := ioutil.WriteFile(filePath, []byte(configContent), 0644); err != nil {
		s.logger.Debug(fmt.Sprintf("Failed to write to file %s: %v", filePath, err))
		return false, err
	}

	s.logger.Debug(fmt.Sprintf("File %s updated successfully", filePath))
	return true, nil
}

func (s *service) updateBHSettings(location string, params bhconfig.Payload) (bool, error) {
	s.logger.Debug("Start updateBHSettings")

	location = cleanLocation(location)
	filePath := filepath.Join(location, "BH_settings.cfg")

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		s.logger.Debug(fmt.Sprintf("Failed to read settings file: %v", err))
		return false, err
	}

	contentStr := string(content)
	contentStr = updateConfigLine(contentStr, "Default Game Name", params.DefaultGameName)
	contentStr = updateConfigLine(contentStr, "Default Password", params.DefaultPassword)
	gsIndex := transformGs(params.DefaultGs)
	contentStr = updateConfigLine(contentStr, "Default Gs", fmt.Sprintf("%d", gsIndex))

	if contentStr == string(content) {
		s.logger.Debug("BH_settings.cfg is unchanged, skipping write")
		return false, nil
	}

	if err := ioutil.WriteFile(filePath, []byte(contentStr), 0644); err != nil {
		s.logger.Debug(fmt.Sprintf("Failed to write updated settings file: %v", err))
		return false, err
	}

	s.logger.Debug("BH_settings.cfg updated successfully")
	return true, nil
}

func normalizeBHContent(content string) string {
	content = strings.ReplaceAll(content, "\r\n", "\n")
	return strings.TrimRight(content, "\n")
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
