package bhconfig

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/DoctorWoot420/resurgence-launcher/log"
)

// Client encapsulates the details of the Configuration Builder API.
type Client struct {
	address string
	logger  log.Logger
}

type Payload struct {
	DefaultGs       string   `json:"default_gs"`
	DefaultGameName string   `json:"default_game_name"`
	DefaultPassword string   `json:"default_password"`
	RuneDesign      string   `json:"rune_design"`
	ItemNameOption  string   `json:"item_name_option"`
	FilterBlocks    []string `json:"filter_blocks"`
}

// GetMaphackTextFromParams will fetch the new bh.cfg text from the server given our configuration parameters
func (c *Client) GetMaphackTextFromParams(maphackConfigParams Payload) (io.ReadCloser, error) {
	c.logger.Debug("clients/bhconfig/client.go start GetMaphackTextFromParams")

	// Marshal the JSON payload
	jsonData, err := json.Marshal(maphackConfigParams)
	if err != nil {
		return nil, fmt.Errorf("clients/bhconfig/client.go failed to marshal JSON: %w", err)
	}

	// Log the JSON payload in a readable format
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, jsonData, "", "  "); err != nil {
		return nil, fmt.Errorf("failed to indent JSON: %w", err)
	}
	c.logger.Debug("clients/bhconfig/client.go JSON Payload:\n" + prettyJSON.String())

	// Make the HTTP POST request with JSON content
	resp, err := http.Post(fmt.Sprintf("%s/generate-config", c.address), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	c.logger.Debug("clients/bhconfig/client.go Successful response from server")

	return resp.Body, nil
}

const filterSourceSHAURL = "https://api.github.com/repos/DoctorWoot420/cosmic-resurgence-bh/commits/main"

// GetFilterSourceSHA returns the current main commit SHA of the BH filter source repo.
func (c *Client) GetFilterSourceSHA() (string, error) {
	c.logger.Debug("clients/bhconfig/client.go start GetFilterSourceSHA")

	req, err := http.NewRequest(http.MethodGet, filterSourceSHAURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "application/vnd.github.sha")
	req.Header.Set("User-Agent", "resurgence-launcher")

	httpClient := &http.Client{Timeout: 5 * time.Second}
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	sha := strings.TrimSpace(string(body))
	if strings.HasPrefix(sha, "{") {
		var payload struct {
			SHA string `json:"sha"`
		}
		if err := json.Unmarshal(body, &payload); err == nil {
			sha = strings.TrimSpace(payload.SHA)
		}
	}
	if resp.StatusCode != http.StatusOK || len(sha) < 7 {
		return "", fmt.Errorf("unexpected filter source sha response (%d): %s", resp.StatusCode, sha)
	}

	c.logger.Debug("clients/bhconfig/client.go filter source sha=" + sha)
	return sha, nil
}

// NewClient returns a new client with all dependencies setup.
func NewClient(
	logger log.Logger,
) Client {
	return Client{
		address: "https://resurgence-bh-generator-812678563239.us-west2.run.app",
		logger:  logger,
	}
}
