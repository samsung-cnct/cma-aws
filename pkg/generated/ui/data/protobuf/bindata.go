// Code generated by go-bindata.
// sources:
// api/api.proto
// DO NOT EDIT!

package protobuf

import (
	"github.com/elazarl/go-bindata-assetfs"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _apiProto = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x59\x6d\x6f\xdb\x38\xf2\x7f\xef\x4f\x31\xf0\x9b\x7f\xfa\x47\x62\x35\xe9\xf6\x6e\x11\x6f\x0e\xe7\x75\xba\xad\xd1\xe6\x01\x75\xb6\x41\x5f\x19\x63\x6a\x2c\xf3\x42\x91\x2a\x49\xd9\x71\x8b\x7c\xf7\x03\x1f\x64\x4b\xb2\xdc\x6d\xb7\x7d\x71\x06\xda\x52\xe4\xcc\x68\xe6\x37\x8f\x54\x93\x04\xc6\xaa\xd8\x68\x9e\x2d\x2d\x9c\x3d\x3f\xfd\x15\xa6\x98\x9b\x52\x66\x30\xbd\x9c\xc2\x58\xa8\x32\x85\x6b\xb4\x7c\x45\x30\x56\x79\x51\x5a\x2e\x33\xb8\x23\xcc\x01\x4b\xbb\x54\xda\x0c\x7a\x49\xd2\x4b\x12\x78\xc7\x19\x49\x43\x29\x94\x32\x25\x0d\x76\x49\x30\x2a\x90\x2d\xa9\x3a\x39\x86\x0f\xa4\x0d\x57\x12\xce\x06\xcf\xe1\xc8\x11\xf4\xe3\x51\xff\xd9\xd0\x89\xd8\xa8\x12\x72\xdc\x80\x54\x16\x4a\x43\x60\x97\xdc\xc0\x82\x0b\x02\x7a\x64\x54\x58\xe0\x12\x98\xca\x0b\xc1\x51\x32\x82\x35\xb7\x4b\xff\x9e\x28\xc5\x69\x02\x1f\xa3\x0c\x35\xb7\xc8\x25\x20\x30\x55\x6c\x40\x2d\xea\x84\x80\x36\x2a\xed\x7e\x4b\x6b\x8b\xf3\x24\x59\xaf\xd7\x03\xf4\x0a\x0f\x94\xce\x12\x11\x48\x4d\xf2\x6e\x32\x7e\x75\x3d\x7d\x75\x72\x36\x78\x1e\x99\xfe\x94\x82\x8c\x01\x4d\x9f\x4a\xae\x29\x85\xf9\x06\xb0\x28\x04\x67\x38\x17\x04\x02\xd7\xa0\x34\x60\xa6\x89\x52\xb0\xca\x29\xbd\xd6\xdc\xe1\x76\x0c\x46\x2d\xec\x1a\x35\x39\x31\x29\x37\x56\xf3\x79\x69\x1b\x98\x55\x2a\x72\xd3\x20\x50\x12\x50\x42\x7f\x34\x85\xc9\xb4\x0f\xbf\x8f\xa6\x93\xe9\xb1\x13\x72\x3f\xb9\x7b\x73\xf3\xe7\x1d\xdc\x8f\xde\xbf\x1f\x5d\xdf\x4d\x5e\x4d\xe1\xe6\x3d\x8c\x6f\xae\x2f\x27\x77\x93\x9b\xeb\x29\xdc\xfc\x01\xa3\xeb\x8f\xf0\x76\x72\x7d\x79\x0c\xc4\xed\x92\x34\xd0\x63\xa1\x9d\x05\x4a\x03\x77\x68\x52\xea\xa1\x9b\x12\x35\x54\x58\xa8\xa0\x92\x29\x88\xf1\x05\x67\x20\x50\x66\x25\x66\x04\x99\x5a\x91\x96\x2e\x12\x0a\xd2\x39\x37\xce\xab\x06\x50\xa6\x4e\x8c\xe0\x39\xb7\x68\xfd\xd6\x9e\x5d\x83\x9e\x23\xb9\xe2\x6c\x89\x24\xe0\x03\x49\xfa\xcc\x11\x7e\xcb\x57\x61\xf5\xef\x2c\x47\x2e\x06\x4c\xe5\xff\x72\x74\x23\xc1\x1f\x10\xde\xa1\x36\x24\xe1\x37\x74\x4f\x03\xe1\x9f\xea\x84\x3d\xb3\x91\x16\x1f\xe1\x02\xfa\x85\x56\x56\xbd\xe8\x0f\x7b\xbd\x02\xd9\x83\x53\x95\xe5\x88\x6b\x33\xec\xf5\x78\x5e\x28\x6d\xa1\x9f\x29\x95\x09\x4a\xb0\xe0\x09\x4a\xa9\xa2\xa6\x03\xcf\xd9\x1f\x6e\xc9\xfc\x33\x3b\xc9\x48\x9e\x98\x35\x66\x19\xe9\x44\x15\x9e\xb4\x93\xad\xd7\x0b\xa7\x70\x94\xe9\x82\x0d\x32\xb4\xb4\xc6\x4d\x38\x66\xb3\x8c\xe4\x2c\x4a\x19\x44\x29\x03\x55\x90\xc4\x82\xaf\xce\xaa\x93\x67\x70\x01\x5f\x7a\x00\x5c\x2e\xd4\xb9\x5f\x01\x58\x6e\x05\x9d\x43\x7f\x2c\x4a\x63\x49\xc3\x15\x4a\xcc\x48\xc3\xe8\x7e\x0a\x6f\x48\x14\x6e\x79\x3b\xe9\x0f\x3d\xf1\x2a\xa4\xd7\x39\xf4\x57\xcf\x07\xa7\x83\xe7\x71\x9b\x29\x69\x91\xd9\x4a\xa4\xfb\x49\xcc\x9d\xd4\x96\x1b\x22\xbd\xfb\x95\x5a\x9c\x43\xdf\x65\x86\x39\x4f\x92\x8c\x5b\x81\x73\x07\x76\x52\x39\x2a\x61\x39\x9e\xe0\xda\xd4\x78\xc8\x79\xe4\x1c\xfa\xfb\xbe\x8c\x44\x4f\xee\x1f\xff\x17\x3d\x5a\xd2\x12\xc5\x2c\x55\xcc\x54\x8a\x7d\xef\x3b\x53\x32\x4c\x73\x8f\xa6\xb3\x45\x69\x02\x9c\xab\xd2\xc2\x37\x80\xf5\xd4\x03\x30\x6c\x49\x39\x99\x73\x78\x73\x77\x77\x3b\x1d\xb6\x77\xdc\x06\x53\xd2\x94\x7e\xa7\x1f\x33\xdc\xbd\x2d\xf9\x8f\x51\xd2\x8b\x29\xb4\x4a\x4b\x76\xe8\xfc\x69\xd8\xeb\x19\xd2\x2b\xce\x68\xab\x53\x30\xd5\x25\x2e\x17\xc2\xf1\xaf\xb8\x2f\x89\x08\x2c\x50\xf8\x73\x5d\x30\x18\x6b\x42\x4b\x15\xdf\x51\xe3\xf1\xca\x64\xcf\x40\x93\x2d\xb5\x34\xad\xa3\xf7\x54\x88\xcd\xb3\x9a\xaf\xb7\x71\xe9\xe3\x7e\x80\x05\x1f\x38\x8c\xab\x68\xdb\xfd\x8a\xd2\xc2\x39\xf4\x7d\x66\xac\x4e\x93\xa8\x4f\xbf\x41\x33\x57\xe9\xc6\x11\xfd\xff\x6e\xfb\x29\x3a\xb7\x61\x98\x26\xab\x39\xad\x42\x3d\x31\x16\x6d\x69\x5c\x0d\xde\x5a\xe9\x6a\x05\x70\x6b\xe0\xa1\x9c\x13\x53\x72\xc1\x33\x5f\x6e\x98\x92\x92\x98\xe5\x2b\x6e\x37\x5b\x24\x5e\x93\xdd\xc2\xb0\x5b\x37\x31\xd8\xed\xff\x7d\x00\x32\xfa\x3a\x00\x9d\x96\xa6\x24\xc8\x52\x87\xff\x2e\xfd\xc1\x56\xf1\xc6\x63\x53\xf7\xc6\xd1\xdf\x57\x3f\x6a\xf2\xdd\x16\x6c\x7d\x85\x20\xb8\xb1\xce\x4f\x91\xd1\x74\xb8\xe0\x9d\x23\x39\x6a\x3e\x1f\x72\x85\x3b\xfb\xd9\xee\x48\x9c\x8e\x7f\x6d\x51\xa9\x65\x55\x0f\x7d\x41\xd5\xb9\x4f\xcd\x58\x21\xb0\xe0\xe0\x32\xb3\xe6\xae\xd7\x64\xe3\x78\x32\xa9\x91\x1f\xed\xb6\xf7\x8c\x8c\xfb\x3f\xcd\xc0\xa8\x6e\x87\x6d\x4f\xbd\x5e\x4e\xc6\xb8\x76\xd6\x2e\x03\xbb\x82\x72\x8d\x39\x55\x73\x4e\x95\x65\x56\xc1\x9c\x76\x55\x86\x52\x4f\xec\xa6\x0a\x99\xf9\x26\x00\x17\x70\x3a\xac\x24\xdc\x2d\x23\xad\xeb\xd9\x55\xd3\xf7\x38\x78\x8a\xc6\xab\x6f\x23\xdd\xb4\x20\xb6\x63\xba\x80\xb3\xe1\x41\x6d\x3d\x50\xb5\x02\xb8\x24\x3f\x8c\x28\xed\xe7\xbd\xba\xda\x6b\x34\x75\xa5\xdd\x80\xe5\x47\x41\x37\x71\x91\xb1\xbd\x50\x89\x94\x00\xf5\xb0\x67\x40\x4a\x16\xb9\x30\x6d\x24\x22\x2b\x68\x32\x85\x92\x86\x82\x45\xe1\x70\x62\x29\xdf\x12\xb6\x4d\x68\x14\x9c\x6f\x41\x5b\x28\xf5\xe0\x26\xba\xe2\xab\x58\x8f\x35\xa5\x24\x2d\x47\x61\x1c\xdf\xa7\x92\xf4\x66\x3b\x70\xd5\x4b\xc9\xe8\x7e\x5a\xa7\x65\xb5\xf5\x61\x55\x5b\x50\x4f\x4c\x43\x4f\x2e\x43\x59\xde\x18\x4b\xf9\x3e\x98\x75\x68\x2e\x3d\x9a\x5f\x05\xa8\x5d\xd8\xea\x1e\x46\xeb\xe6\xd8\xda\xbb\xff\xcf\x04\x28\xac\x72\xed\xdb\x6a\xb5\xf9\x1e\x94\x62\x85\xfb\x01\x88\xf6\x4b\xed\x4e\xdd\xb1\x2a\x45\xda\x00\x6a\x4e\x95\x96\x31\x73\xba\x82\x6e\xba\xed\x6e\x8e\xb5\x1e\xa2\xd1\xaa\xd8\xfe\x0e\x7b\x2b\x96\xd0\x9a\x26\x3f\x25\x38\x4e\xbf\xf2\xba\x1f\x0a\x90\xc8\xf4\xae\xb3\x59\x50\xe1\x52\x3e\xed\xca\xad\x7d\x0c\xea\x44\x3b\x65\x2e\x5b\x89\x55\x07\x93\xa7\x0d\x1d\x3a\xd2\xb0\x23\xa0\xce\x86\x5d\x21\x69\x1a\x8e\xeb\xe0\xde\x3a\xee\x45\x97\xd2\xb5\xd4\xf8\xdf\x56\xbd\x83\xbf\x36\x75\x59\x55\x0d\x5d\x6e\x79\x40\x5c\x8d\xfe\x02\x7e\x39\x5c\xe2\x1b\x5d\xa1\xb3\x0e\x6c\x5b\xc5\x09\xb0\x52\x6b\x92\x56\xc4\xe2\xce\x0d\xe0\xda\xdf\x49\x73\x44\xf3\x97\x8d\xaa\x6a\xee\x6a\x01\x6f\xcb\x39\x69\x49\x96\x1a\x5c\x0f\xbf\x9a\x59\x45\xe4\x71\xf4\x87\x4a\x92\x5a\x6c\xb5\x98\xd5\x47\x83\x5d\x73\x8e\xaf\x70\xf7\x86\xfd\x36\xb8\xd7\x0a\x47\xf7\x53\x6f\xaf\xd3\x7e\x0b\xf8\x53\xef\x1b\xfa\x1c\x37\xf0\x66\xb4\xcb\xaf\x25\xcf\x96\x33\x5c\x21\x17\x38\xe7\x82\xdb\x4d\x00\xbb\xa6\xd1\x02\xe7\x9a\xb3\xd8\x68\x4a\xd3\xea\xe7\x64\xd7\x4a\x3f\xcc\x22\xd1\x05\xbc\x1c\xf6\x9c\xa7\x22\x2f\x6b\x16\x95\x32\x5e\xf0\x99\x33\xc5\xb1\xd7\x5d\x5f\x39\xb7\x55\x5f\xbe\xd4\x75\x99\x12\xd3\x64\xdf\xd2\x66\x92\x7a\x41\xa3\xdb\x09\x8c\x18\x23\xd3\xf0\x82\xf1\x54\xb3\x07\xda\xcc\x5a\xf1\xbf\x93\x11\xb8\xde\xd2\x66\x2b\x07\x0f\xc9\x09\x07\x4e\x5c\x23\x37\x9c\xac\xf7\x94\x39\x2f\x1e\x16\xa1\x03\xc1\x5e\x3a\x77\xf9\xb2\x61\xa9\x8b\x83\x4b\xb4\x08\x63\x92\xb5\xca\xeb\xb6\xc2\x0e\xa4\x68\x71\xc6\xc2\xfa\x2b\xfd\x6b\x5e\xf2\x66\x87\xf9\xa6\xf6\x15\x65\xfd\xfe\xf1\x06\xb8\xa5\xdc\x54\x4c\xb7\x3a\xe6\x64\xa9\x29\x75\x75\xc8\x0d\x48\x46\x95\x9a\x51\x33\xf3\x27\xd2\x58\xff\xc1\x2c\xd3\xaa\x2c\x5a\x75\x7a\x74\x3f\xad\xce\x5f\xbb\x63\xe0\xf1\x69\x16\xa8\x43\x0c\xee\x82\x99\xb3\xe5\x1e\x18\xb5\x70\xa9\x81\xd2\xc8\xa7\xc0\x18\x5d\x70\x54\x9a\x13\x42\x63\x4f\x4e\x8f\x81\x2c\x1b\x3c\xdb\x52\xb6\x5d\x75\x3a\xdc\x17\xd2\xc8\x90\xcf\x4a\x92\xa9\x09\x9c\x1f\x43\xb5\x3e\x63\x7e\xbd\x26\xb7\x4e\xdb\x6f\xda\x02\x10\x5f\x59\x97\x3a\x0b\x52\xb7\xe8\xef\x92\xf9\x0f\xa5\x61\xbd\x24\x09\x46\xe5\xfe\xfb\xa4\xcc\x0c\xa0\x26\x40\xa1\x09\xd3\x4d\x48\xa7\x98\x97\x35\x58\x3a\x7c\xb5\x57\x6e\x3e\xdc\x8e\x81\xa7\xc7\x30\x17\x28\x1f\x7c\x1c\xbb\x3f\xfd\x20\xd1\x15\x2e\xff\xbc\x51\x65\xff\x18\x16\x5c\x08\x4a\x81\x2f\xfc\x37\x53\xa7\x80\x0b\x8f\x0f\xb7\xe3\x36\x92\xab\x82\x35\x12\xaf\x1a\x5a\x88\x95\xda\xe1\xe7\x9d\xdc\x66\x32\xf1\x34\x84\x40\xe0\x3f\x1b\xb6\xf5\x9d\x8c\xae\x40\x2b\x41\xed\xc9\x04\x8e\x50\xcb\x3d\x97\x72\xcc\x67\x8e\x7a\x86\x5a\x76\x95\xc9\x56\x90\x42\x4a\x0b\x2e\xdd\x9d\xd4\x6e\x0a\xf2\x9f\x0b\x64\x99\xcf\x5d\x1d\x5d\x6c\x43\xd4\xb4\x61\x6e\x46\x72\x03\xe1\xad\x7c\x2f\xef\x28\x7f\x39\x10\xa8\x33\x3a\x10\x80\x9e\xa8\x0d\xda\x15\x97\x3c\x2f\xf3\x2e\x45\xe0\x28\xa5\x05\x96\xc2\xfa\x3c\xff\x4c\x5a\xed\x44\x72\x69\x5f\x9c\x41\xce\xe5\xec\x53\x89\xd2\x86\xba\xde\x84\xf3\x0a\x1f\x7f\x40\x32\x3e\xd6\x25\xbf\xa8\xdd\x1a\x93\xc4\xcd\x7d\xf5\x5e\xe9\x4a\xe3\x34\xdc\x7c\x6b\x93\xe1\xee\x8a\x0b\x5f\x22\x5f\x98\x10\x9d\x6b\x2b\xee\x6a\xb4\xdd\xe7\x6b\x4f\x93\x0b\x50\x05\xe9\xd0\x55\xdd\x5d\xee\xe6\xed\x81\x5b\x46\x25\xaa\xe3\xe6\xbd\x97\x1e\x16\x33\x50\x61\x30\xcd\xb8\xbb\xc8\x15\xca\x70\xab\xf4\xa6\xed\xbb\x8c\xdb\x5a\xe3\x3f\xdd\x8b\xdb\x25\x9a\x65\x35\x3a\x39\x49\x4c\xe5\x39\xb7\x5d\x52\xc2\xc9\x9e\xb7\x3a\x3a\xba\xd5\x44\xde\x54\x26\x08\x65\x28\x11\xae\xda\x77\x8a\x75\xc4\x33\x37\xa1\xd1\xce\x5d\x51\xf4\xa5\x4f\xf5\x45\xe8\x14\x6d\x5e\xbf\x39\x4b\x03\xdf\x2f\x0d\xbe\x0f\x3b\x0f\x67\xbe\xb7\xa7\x61\xb0\xcb\x0b\x2e\x68\x4f\x07\x55\xc3\xe7\x65\x43\xce\x38\x70\xe8\xdd\x70\x51\xe3\x63\xd5\xe1\x05\xfc\xa3\xc1\x75\x2b\xd0\x3a\xcf\x01\xb7\x01\x84\x40\x18\xe6\x82\x04\x74\x29\xfd\xff\x15\xd4\x06\xa8\x28\xb1\xa8\x18\x2f\xe0\x9f\xed\x82\x50\x99\x54\x0b\x0a\x7f\xd4\x11\x2b\xd1\x9a\xc6\x34\x57\x5d\x35\x7a\xff\x0d\x00\x00\xff\xff\xc0\xcd\xf2\xe1\xe1\x1a\x00\x00")

func apiProtoBytes() ([]byte, error) {
	return bindataRead(
		_apiProto,
		"api.proto",
	)
}

func apiProto() (*asset, error) {
	bytes, err := apiProtoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "api.proto", size: 6881, mode: os.FileMode(420), modTime: time.Unix(1536710176, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"api.proto": apiProto,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"api.proto": &bintree{apiProto, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}


func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}
