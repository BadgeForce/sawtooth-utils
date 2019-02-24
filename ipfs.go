package utils

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	shell "github.com/ipfs/go-ipfs-api"
)

const (
	add         = "/api/v0/add"
	nocopy      = "nocopy=true"
	fileNameFmt = "%s.%s.bf"
)

// IPFSClient ...
type IPFSClient struct {
	Host  string
	Shell *shell.Shell
}

// Response ...
type Response struct {
	Name  string `json:"Name"`
	Hash  string `json:"Hash"`
	Bytes int64  `json:"Bytes"`
	Size  string `json:"Size"`
}

// File ...
type File struct {
	Data         string `json:"data"`
	LastModified int64  `json:"last_modified"`
	Name         string `json:"name"`
	FileType     string `json:"file_type"`
}

// GetFileName ...
func (ipfs *IPFSClient) GetFileName(name, filetype string) string {
	return fmt.Sprintf(fileNameFmt, name, filetype)
}

// NewFile ...
func (ipfs *IPFSClient) NewFile(name string, data []byte) (*File, []byte) {
	lastModified := time.Now().Unix()
	encoded := hex.EncodeToString(data)
	file := &File{
		Data:         encoded,
		LastModified: lastModified,
		Name:         name,
	}

	b, _ := json.Marshal(file)

	return file, b
}

// AddFile ...
func (ipfs *IPFSClient) AddFile(filename, filetype string, data []byte) (string, string, error) {
	fileName := ipfs.GetFileName(filename, filetype)
	_, fileBytes := ipfs.NewFile(fileName, data)
	hash, err := ipfs.Shell.Add(bytes.NewReader(fileBytes))
	return fileName, hash, err
}

// NewIPFSHTTPClient ...
func NewIPFSHTTPClient(host string) *IPFSClient {
	return &IPFSClient{host, shell.NewShell(host)}
}
