package resurgence

import (
	"fmt"
	"io"
	"net/http"
)

// Client encapsulates the details of the Resurgence API.
type Client struct {
	address string
}

// GetFile will the file by the given path in the repository set on the service.
func (c *Client) GetFile(filePath string) (io.ReadCloser, error) {
	resp, err := http.Get(fmt.Sprintf("%s/resurgence-patches/%s", c.address, filePath))
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

// GetNews will fetch the remote news source.
func (c *Client) GetNews() (io.ReadCloser, error) {
	resp, err := http.Get(fmt.Sprintf("%s/news.json", c.address))
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

// GetAvailableMaphackOptions will fetch the remote news source.
func (c *Client) GetAvailableMaphackOptions() (io.ReadCloser, error) {
	resp, err := http.Get(fmt.Sprintf("%s/available_maphack_options.json", c.address))
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

// GetAvailableMods will fetch the remote available mods source.
func (c *Client) GetAvailableMods() (io.ReadCloser, error) {
	resp, err := http.Get(fmt.Sprintf("%s/available_mods_1.1.0.json", c.address))
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

// NewClient returns a new client with all dependencies setup.
func NewClient() Client {
	return Client{
		address: "https://d2resurgence.github.io/patch-d2lod/files",
	}
}
