// Code generated by go-bindata.
// sources:
// data/config_schema_v3.0.json
// data/config_schema_v3.1.json
// data/config_schema_v3.2.json
// DO NOT EDIT!

package schema

import (
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

var _dataConfig_schema_v30Json = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5a\x4b\x8f\xdb\x38\x12\xbe\xfb\x57\x08\x4a\x6e\x71\x77\x07\xd8\x60\x81\xcd\x6d\x8f\x7b\xda\x39\x4f\x43\x11\x68\xa9\x6c\x33\x4d\x91\x4c\x91\x72\xda\x09\xfc\xdf\x07\xd4\xcb\x14\x4d\x8a\xb2\xad\x3c\x30\x98\x53\xb7\xc5\xaa\x62\xbd\xf8\x55\xb1\xa4\xef\xab\x24\x49\xdf\xaa\x62\x0f\x15\x49\x3f\x26\xe9\x5e\x6b\xf9\xf1\xe9\xe9\xb3\x12\xfc\xa1\x7d\xfa\x28\x70\xf7\x54\x22\xd9\xea\x87\xf7\x1f\x9e\xda\x67\x6f\xd2\xb5\xe1\xa3\xa5\x61\x29\x04\xdf\xd2\x5d\xde\xae\xe4\x87\x7f\x3d\xbe\x7f\x34\xec\x2d\x89\x3e\x4a\x30\x44\x62\xf3\x19\x0a\xdd\x3e\x43\xf8\x52\x53\x04\xc3\xfc\x9c\x1e\x00\x15\x15\x3c\xcd\xd6\x2b\xb3\x26\x51\x48\x40\x4d\x41\xa5\x1f\x13\xa3\x5c\x92\x0c\x24\xfd\x03\x4b\xac\xd2\x48\xf9\x2e\x6d\x1e\x9f\x1a\x09\x49\x92\x2a\xc0\x03\x2d\x2c\x09\x83\xaa\x6f\x9e\xce\xf2\x9f\x06\xb2\xb5\x2b\xd5\x52\xb6\x79\x2e\x89\xd6\x80\xfc\x8f\x4b\xdd\x9a\xe5\x4f\xcf\xe4\xe1\xdb\x7f\x1f\xfe\x7c\xff\xf0\x9f\xc7\xfc\x21\x7b\xf7\x76\xb4\x6c\xfc\x8b\xb0\x6d\xb7\x2f\x61\x4b\x39\xd5\x54\xf0\x61\xff\x74\xa0\x3c\x75\xff\x9d\x86\x8d\x49\x59\x36\xc4\x84\x8d\xf6\xde\x12\xa6\x60\x6c\x33\x07\xfd\x55\xe0\x4b\xcc\xe6\x81\xec\x17\xd9\xdc\xed\xef\xb1\x79\x6c\xce\x41\xb0\xba\x8a\x46\xb0\xa7\xfa\x45\xc6\xb4\xdb\xdf\x17\xbf\x55\x6f\xf4\x24\x6d\x4b\x61\xed\xdd\x28\x38\xca\x76\x9f\xab\x7c\xd9\x16\xf6\xd5\xe0\xac\x80\x97\x4a\x90\x4c\x1c\xcd\xb3\x80\x3f\x5a\x82\x0a\xb8\x4e\x07\x17\x24\x49\xba\xa9\x29\x2b\x5d\x8f\x0a\x0e\xff\x37\x22\x9e\xad\x87\x49\xf2\xdd\x3d\xd8\x96\x9c\x66\x7d\xf4\x2b\x1c\xf0\x61\x3d\x60\xcb\xb0\x5e\x08\xae\xe1\x55\x37\x46\x4d\x6f\xdd\xba\x40\x14\x2f\x80\x5b\xca\x60\x2e\x07\xc1\x9d\x9a\x70\x19\xa3\x4a\xe7\x02\xf3\x92\x16\x3a\x3d\x39\xec\x17\xf2\xe2\xf9\x34\xb0\x5a\xbf\xb2\x95\x47\x60\x5a\x10\x99\x93\xb2\x1c\xd9\x41\x10\xc9\x31\x5d\x27\x29\xd5\x50\x29\xbf\x89\x49\x5a\x73\xfa\xa5\x86\xff\x75\x24\x1a\x6b\x70\xe5\x96\x28\xe4\xf2\x82\x77\x28\x6a\x99\x4b\x82\x26\xc1\xa6\xdd\x9f\x16\xa2\xaa\x08\x5f\x2a\xeb\xae\xb1\x63\x86\xe7\x05\xd7\x84\x72\xc0\x9c\x93\x2a\x96\x48\xe6\xd4\x01\x2f\x55\xde\xd6\xbf\xc9\x34\xda\xe6\x2d\xbf\x72\x04\x0c\xc5\x70\xd1\x78\x94\x7c\x2a\xb1\x5b\x31\x26\xb5\x8d\x6e\xa9\xc3\x98\x2b\x20\x58\xec\x6f\xe4\x17\x15\xa1\x7c\x8e\xef\x80\x6b\x3c\x4a\x41\xdb\x7c\xf9\xed\x12\x01\xf8\x21\x1f\xb0\xe4\x6a\x37\x00\x3f\x50\x14\xbc\xea\x4f\xc3\x1c\x80\x19\x40\xde\xf0\xbf\x4a\xa1\xc0\x75\x8c\x63\xa0\xbd\x34\x98\x3a\xf2\x49\xcf\xf1\xdc\x1b\xbe\x4e\x52\x5e\x57\x1b\x40\xd3\xd2\x8d\x28\xb7\x02\x2b\x62\x94\xed\xf7\xb6\x96\x47\x9e\xf6\x64\x9e\xed\x40\xdb\x06\x53\xd6\x09\xcb\x19\xe5\x2f\xcb\xa7\x38\xbc\x6a\x24\xf9\x5e\x28\x3d\x1f\xc3\x2d\xf6\x3d\x10\xa6\xf7\xc5\x1e\x8a\x97\x09\x76\x9b\x6a\xc4\x2d\x94\x9e\x93\xe4\xb4\x22\xbb\x38\x91\x2c\x62\x24\x8c\x6c\x80\xdd\x64\xe7\xa2\xce\xb7\xc4\x8a\xdd\xce\x90\x86\x32\xee\xa2\x73\xe9\x96\x63\x35\xbf\x44\x7a\x00\x9c\x5b\xc0\x85\x3c\x37\x5c\xee\x62\xbc\x01\x49\xe2\xdd\xe7\x88\xf4\xd3\x63\xdb\x7c\x4e\x9c\xaa\xe6\x3f\xc6\xd2\xcc\x6d\x17\x12\xa7\xee\xfb\x9e\x38\x16\xce\x6b\x28\x46\x51\xa9\x48\x61\xfa\x06\x04\x15\x88\xeb\x99\xb4\x6b\xf6\xf3\x4a\x94\xa1\x04\xbd\x20\x76\x7d\x13\x44\xea\xab\x0b\x61\x72\x53\xff\x38\x2b\x74\xd1\x0b\x44\xc4\x9a\x90\x7a\x73\xd5\x3c\xab\x1b\x4f\xb1\x86\x8e\x30\x4a\x14\xc4\x0f\x7b\xd0\x91\x23\x69\x54\x1e\x3e\xcc\xcc\x09\x1f\xef\xbf\x27\x79\x03\xac\x41\x99\xf3\x7b\xe4\x88\xa8\xb3\x2a\xcd\x71\xf3\x29\x92\x45\x4e\xdb\x0f\x6e\xe1\x25\x2d\xc3\x58\xd1\x20\x84\x7d\xc0\xa4\x40\x7d\x71\xba\x7e\x4e\xb9\x6f\xb7\xbe\xbb\xda\x4b\xa4\x07\xca\x60\x07\xe3\x5b\xcb\x46\x08\x06\x84\x8f\xa0\x07\x81\x94\xb9\xe0\xec\x38\x83\x52\x69\x82\xd1\x0b\x85\x82\xa2\x46\xaa\x8f\xb9\x90\x7a\xf1\x3e\x43\xed\xab\x5c\xd1\x6f\x30\x8e\xe6\x19\xef\x3b\x41\xd9\x88\xe7\xa8\x0a\x7d\x5b\xbd\x56\xba\xa4\x3c\x17\x12\x78\xd4\x3b\x4a\x0b\x99\xef\x90\x14\x90\x4b\x40\x2a\x4a\x9f\x81\x6b\x3b\xd6\x65\x8d\xc4\xec\x7f\x29\x46\xd1\x1d\x27\x2c\xe6\x68\x5d\xc9\xed\x8d\x17\x0b\xad\xe3\xe1\xae\x19\xad\x68\xf8\x1c\x78\x00\x76\x46\x0d\x68\xf1\xdf\x0f\xfb\x13\x90\x7f\xd6\x94\x72\x0d\x3b\x40\x1f\x52\x4e\x74\x1d\xd3\x4d\xc7\x8c\x6e\x63\x4f\x70\x1c\xd0\x09\x3d\x1a\x06\x25\xb6\xda\xcf\xe0\xeb\x45\xbc\x7a\x8d\x86\xbf\x8d\xbc\x75\xa7\x48\xe6\xa5\xbf\x0a\xce\x5d\x35\xb2\x20\xa2\x9e\xbc\x88\x5a\xab\x68\x63\xd8\xd0\x70\x35\xd5\xd4\x0c\xa4\xd6\x14\x73\x51\xbc\x30\x8d\x92\x39\x04\x25\xf5\x6b\xbb\x72\x2c\xbb\x62\x8e\xec\xdc\x59\x7a\x01\xbe\x89\xa2\x4d\x1a\x9d\xc0\x4e\x4f\x37\x3b\xa2\xe0\xe4\x91\x2a\xb2\x71\x66\x6e\xbe\xc3\x6d\xb2\x11\x0f\x71\x8c\x41\xd0\x48\x9d\xb8\x74\x68\x3b\xc2\x13\x50\xbf\xe7\xe0\x40\xd3\x0a\x44\xed\xaf\x59\x2b\x3b\xbf\x3b\xa6\xd4\x9a\xcc\x46\x82\x6a\x51\xba\x31\x7d\x1e\x82\xda\xf7\x17\xd1\xc0\xcd\x39\x24\x08\x92\xd1\x82\xa8\x18\x10\xdd\x71\x41\xad\x65\x49\x34\xe4\xed\x8b\xaa\xab\xa0\x7f\x02\xf3\x25\x41\xc2\x18\x30\xaa\xaa\x39\x18\x9a\x96\xc0\xc8\xf1\xa6\xf2\xd9\xb0\x6f\x09\x65\x35\x42\x4e\x0a\xdd\xbd\x0b\x8b\xe4\x5c\x5a\x09\x4e\xb5\xf0\x22\xc4\xbc\x2d\x2b\xf2\x9a\xf7\xdb\x36\x24\xde\x03\x13\x6c\xeb\xe6\xde\x2d\xad\x4c\x50\xa2\xc6\xe2\xc2\xd9\x37\x87\xe8\x5c\xeb\x03\x19\xd3\xef\x78\x61\x3a\x82\x32\x48\x32\x5c\xfd\xa3\xfc\xd1\xd2\xd2\xf5\x99\xb9\x14\x8c\x16\xc7\xa5\x2c\x2c\x04\x6f\x9d\x3c\x27\x21\xee\xcc\x40\x93\x0e\xa6\x15\xaa\xa4\x8e\x1e\xd6\x86\xe1\x2b\xe5\xa5\xf8\x7a\xc5\x86\xcb\xa5\x92\x64\xa4\x00\x07\xef\xee\x75\xb4\xd2\x48\x28\xd7\x57\x97\xf3\x7b\xcd\xba\xa3\x9a\x0f\xf9\x19\x41\xfd\x81\x2e\xfe\x26\x35\x80\xf4\x85\xac\xa3\xf3\xa0\x0a\x2a\x81\xde\x04\x5c\xe0\xcd\x77\xcc\xc4\x9e\x6c\x81\xaa\x36\x6b\x80\xd8\x51\x99\xfb\xe2\xe2\xb7\x8d\xf8\x90\x30\x8b\x03\x12\x95\xa4\x5a\xea\x74\xcc\x1e\xa9\xa6\xde\x1a\x9c\x4c\x8f\x22\x92\xf0\x38\x22\xa6\x75\x5c\xf7\x8e\x42\xd5\x1b\x0e\x93\x1d\x95\xe5\x4f\xdf\x7b\xde\xf9\xd7\x94\x53\xf8\x52\x72\x1f\xe8\xf5\x6f\x43\x02\x51\x7d\x1e\x7a\xe6\xf5\xe0\xab\x6c\x76\x88\x83\xaf\x22\x96\xd3\xbf\x69\xdf\xdd\x11\x81\xaf\xcf\xbf\xb2\x13\xbc\x03\x5c\xba\x4f\x3c\x22\xd8\xd2\x51\xfd\x03\x2d\x7f\x93\x44\xfc\x79\xf9\xe5\x4c\xb3\xac\x3c\xbb\xbc\x68\x4e\xa5\xc4\xec\x31\x7e\xc7\x91\x8d\xd5\x70\xc9\x3c\x5f\xda\x8d\x61\x79\x6a\x7a\xd1\x93\x04\xc6\xba\xce\xa6\x9d\x13\xa7\x2d\x5f\x30\xc3\x1f\xdf\x4d\x14\x9f\xa9\xd7\x6d\x3f\x08\xb5\x17\x98\x0c\xf9\x63\xea\x74\xac\xbd\x77\x2f\x3f\x17\x0b\x80\x9a\xc5\x7f\xf1\xf1\x98\xb1\x93\x1f\x2f\x06\x21\xdf\xc7\xd3\xbd\xf6\xc3\xaf\x6c\xe4\x1f\x87\xa4\x7d\x79\x6d\x41\x4a\x66\x37\xf1\xa1\x30\x7a\x3f\x29\x73\x67\x8b\xfd\xa7\x5d\x99\x1f\xae\x56\xf6\xdf\xe6\x33\xbc\xd5\x69\xf5\x57\x00\x00\x00\xff\xff\x78\x30\xec\x51\x0e\x2b\x00\x00")

func dataConfig_schema_v30JsonBytes() ([]byte, error) {
	return bindataRead(
		_dataConfig_schema_v30Json,
		"data/config_schema_v3.0.json",
	)
}

func dataConfig_schema_v30Json() (*asset, error) {
	bytes, err := dataConfig_schema_v30JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/config_schema_v3.0.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataConfig_schema_v31Json = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x1a\xcb\x8e\xdb\x36\xf0\xee\xaf\x10\x94\xdc\xe2\xdd\x4d\xd1\xa0\x40\x73\xeb\xb1\xa7\xf6\xdc\x85\x23\xd0\xd2\x58\x66\x96\x22\x19\x92\x72\xd6\x09\xfc\xef\x05\xf5\x32\x45\x91\x22\x6d\x2b\xd9\x45\xd1\xd3\xae\xc5\x99\xe1\xbc\x67\x38\xe4\xf7\x55\x92\xa4\x6f\x65\xbe\x87\x0a\xa5\x1f\x93\x74\xaf\x14\xff\xf8\xf0\xf0\x59\x32\x7a\xd7\x7e\xbd\x67\xa2\x7c\x28\x04\xda\xa9\xbb\xf7\x1f\x1e\xda\x6f\x6f\xd2\xb5\xc6\xc3\x85\x46\xc9\x19\xdd\xe1\x32\x6b\x57\xb2\xc3\xaf\xf7\xbf\xdc\x6b\xf4\x16\x44\x1d\x39\x68\x20\xb6\xfd\x0c\xb9\x6a\xbf\x09\xf8\x52\x63\x01\x1a\xf9\x31\x3d\x80\x90\x98\xd1\x74\xb3\x5e\xe9\x35\x2e\x18\x07\xa1\x30\xc8\xf4\x63\xa2\x99\x4b\x92\x01\xa4\xff\x60\x90\x95\x4a\x60\x5a\xa6\xcd\xe7\x53\x43\x21\x49\x52\x09\xe2\x80\x73\x83\xc2\xc0\xea\x9b\x87\x33\xfd\x87\x01\x6c\x6d\x53\x35\x98\x6d\xbe\x73\xa4\x14\x08\xfa\xf7\x94\xb7\x66\xf9\xd3\x23\xba\xfb\xf6\xc7\xdd\x3f\xef\xef\x7e\xbf\xcf\xee\x36\xef\xde\x8e\x96\xb5\x7e\x05\xec\xda\xed\x0b\xd8\x61\x8a\x15\x66\x74\xd8\x3f\x1d\x20\x4f\xdd\x7f\xa7\x61\x63\x54\x14\x0d\x30\x22\xa3\xbd\x77\x88\x48\x18\xcb\x4c\x41\x7d\x65\xe2\x29\x24\xf3\x00\xf6\x42\x32\x77\xfb\x3b\x64\x1e\x8b\x73\x60\xa4\xae\x82\x16\xec\xa1\x5e\x48\x98\x76\xfb\x65\xec\x27\x21\x17\xa0\xc2\x2e\xdb\x42\xbd\x98\xc7\xea\xed\x6f\x13\x78\xd5\x0b\x3d\x0b\xdb\x42\x18\x7b\x37\x0c\x8e\xc2\xdb\xa5\x2a\x57\x78\xf9\x75\x35\x28\xcb\xa3\xa5\x02\x38\x61\x47\xfd\xcd\xa3\x8f\x16\xa0\x02\xaa\xd2\x41\x05\x49\x92\x6e\x6b\x4c\x0a\x5b\xa3\x8c\xc2\x5f\x9a\xc4\xa3\xf1\x31\x49\xbe\xdb\x99\xcc\xa0\xd3\xac\x8f\x7e\xf9\x0d\x3e\xac\x7b\x64\x19\xd6\x73\x46\x15\x3c\xab\x46\xa8\xf9\xad\x5b\x15\xb0\xfc\x09\xc4\x0e\x13\x88\xc5\x40\xa2\x94\x33\x2a\x23\x58\xaa\x8c\x89\xac\xc0\xb9\x4a\x4f\x16\xfa\x84\x5e\xd8\x9f\x06\x54\xe3\xd7\x66\xe5\x20\x98\xe6\x88\x67\xa8\x28\x46\x72\x20\x21\xd0\x31\x5d\x27\x29\x56\x50\x49\xb7\x88\x49\x5a\x53\xfc\xa5\x86\x3f\x3b\x10\x25\x6a\xb0\xe9\x16\x82\xf1\xe5\x09\x97\x82\xd5\x3c\xe3\x48\x68\x07\x9b\x57\x7f\x9a\xb3\xaa\x42\x74\x29\xaf\xbb\x44\x8e\x08\xcd\x33\xaa\x10\xa6\x20\x32\x8a\xaa\x90\x23\xe9\xa8\x03\x5a\xc8\xac\x2d\xf8\xb3\x6e\xb4\xcb\x5a\x7c\x69\x11\x18\xaa\xff\xa2\xf6\x28\xe8\x9c\x63\xb7\x64\xb4\x6b\x6b\xde\x52\x0b\x31\x93\x80\x44\xbe\xbf\x12\x9f\x55\x08\xd3\x18\xdd\x01\x55\xe2\xc8\x19\x6e\xfd\xe5\xd5\x39\x02\xd0\x43\x36\xe4\x92\x8b\xd5\x00\xf4\x80\x05\xa3\x55\x1f\x0d\x31\x09\x66\x48\xf2\x1a\xff\x99\x33\x09\xb6\x62\x2c\x01\xcd\xa5\x41\xd4\x91\x4e\x7a\x8c\xc7\x5e\xf0\x75\x92\xd2\xba\xda\x82\xd0\x3d\xec\x08\x72\xc7\x44\x85\x34\xb3\xfd\xde\xc6\xf2\x48\xd3\x0e\xcf\x33\x15\x68\xca\xa0\xcb\x3a\x22\x19\xc1\xf4\x69\x79\x17\x87\x67\x25\x50\xb6\x67\x52\xc5\xe7\x70\x03\x7d\x0f\x88\xa8\x7d\xbe\x87\xfc\x69\x06\xdd\x84\x1a\x61\x33\xa9\x62\x9c\x1c\x57\xa8\x0c\x03\xf1\x3c\x04\x42\xd0\x16\xc8\x55\x72\x2e\xaa\x7c\x83\x2c\x2b\x4b\x0d\xea\xf3\xb8\x49\xe7\xd2\x2d\x87\x6a\x7e\x21\xf0\x01\x44\x6c\x01\x67\xfc\xdc\x70\xd9\x8b\xe1\x06\x24\x09\x77\x9f\x23\xd0\x4f\xf7\x6d\xf3\x39\x13\x55\xcd\x7f\x84\xa4\x1b\xbb\x5d\x48\xac\xba\xef\xfa\x62\x49\x18\xd7\x50\x8c\xac\x52\xa1\x5c\xf7\x0d\x02\xa4\xc7\xae\x67\xd0\xee\x74\x93\x55\xac\xf0\x39\xe8\x04\xd8\xd6\x8d\x37\x53\x5f\x5c\x08\x93\xab\xfa\xc7\x28\xd3\x05\x0f\x10\x01\x69\x7c\xec\xc5\xb2\x79\x66\x37\xec\x62\x0d\x1c\x22\x18\x49\x08\x07\xbb\x57\x91\x23\x6a\x98\x1f\x3e\x44\xfa\x84\x0b\xf7\xb7\x59\x5c\x0f\xaa\x97\x66\x7c\x8f\x1c\x20\x75\x66\xa5\x09\x37\x17\x23\x9b\x40\xb4\xfd\xe0\x16\x9e\xe3\xc2\x9f\x2b\x9a\x0c\x61\x06\x18\x67\x42\x4d\xa2\xeb\xe7\x94\xfb\x76\xeb\x9b\xab\x3d\x17\xf8\x80\x09\x94\x30\x3e\xb5\x6c\x19\x23\x80\xe8\x28\xf5\x08\x40\x45\xc6\x28\x39\x46\x40\x4a\x85\x44\xf0\x40\x21\x21\xaf\x05\x56\xc7\x8c\x71\xb5\x78\x9f\x21\xf7\x55\x26\xf1\x37\x18\x5b\xf3\x9c\xef\x3b\x42\x1b\x8b\x21\x6b\x42\x72\xa5\x41\x7d\x29\x29\x1c\xc6\x8e\x44\x18\x4c\x54\xe1\x14\x95\x4a\x56\x8b\x3c\xf6\x80\xad\xf7\x44\xa2\x84\xd8\x23\xbc\x76\xb7\x71\xd8\xcc\x03\x97\x97\x00\x4f\x0a\x5d\x67\xc2\x50\x55\xb6\x7f\x9b\x79\xe5\xe4\x0c\x7d\x79\x94\xb9\xba\xae\x5b\x93\xaa\xc0\x34\x63\x1c\x68\x30\x36\xa4\x62\x3c\x2b\x05\xca\x21\xe3\x20\x30\x73\xaa\x62\x6d\x46\x7a\x51\x0b\xa4\xf7\x9f\x92\x91\xb8\xa4\x88\x84\xc2\x4c\x55\x7c\x77\xe5\xb1\x52\xa9\x70\xb0\xd7\x04\x57\xd8\x1f\x34\x0e\xaf\x8d\xe8\x00\xda\xea\xef\x2e\xfa\x33\x05\xff\xcc\x29\xa6\x0a\x4a\xed\x26\x53\xa7\x9a\xe9\x39\xe7\x5b\xce\x88\x5e\x73\x8f\xc4\xd8\xa0\x33\x7c\x24\x6d\x60\xee\x94\x1b\xc1\xd5\x89\x3a\xf9\x1a\xdd\x75\x34\xf4\xd6\x1d\x23\x1b\x27\xfc\x45\xc5\xdc\x66\x63\xe3\xad\xa7\xee\xa0\xaa\x65\xf0\x58\xd0\xc0\x50\x39\xd7\xd2\x0e\xa0\xc6\xd0\x7e\xd1\x6a\xa1\xdb\x64\x1d\x04\x05\x76\x73\xbb\xb2\x24\xbb\x60\xec\x6e\x9d\x58\x7b\x02\xae\x79\xb2\x09\x1a\x9c\xbf\xcf\xcf\xb6\x3b\x20\xef\xdc\x19\x4b\xb4\xb5\x26\xae\xae\xe0\xd6\xde\x28\x0e\xe1\x1c\x23\x40\x09\x6c\xd9\xa5\x4f\xd4\x66\x3e\x01\xf9\x3a\xc7\x46\x0a\x57\xc0\x6a\x77\xc1\x5b\x99\xfe\xdd\x21\xa5\xc6\x5c\x3e\x60\x54\x03\xd2\xb6\xe9\xe3\x60\xd4\xbe\xbb\x0c\x1a\x2e\x26\x48\x04\x70\x82\x73\x24\x43\x89\xe8\x86\xf1\x44\xcd\x0b\xa4\x20\x6b\xef\x65\x2f\x4a\xfd\x33\x39\x9f\x23\x81\x08\x01\x82\x65\x15\x93\x43\xd3\x02\x08\x3a\x5e\x55\x3e\x1b\xf4\x1d\xc2\xa4\x16\x90\xa1\x5c\x75\x57\xbf\x01\x9f\x4b\x2b\x46\xb1\x62\xce\x0c\x11\xb7\x65\x85\x9e\xb3\x7e\xdb\x06\x24\xd4\xd9\x8c\x9b\xfa\xd8\xc9\x82\xe1\x09\x6d\xe3\x77\x59\x75\x9e\x31\xd1\xb9\xd6\x7b\x3c\xa6\xdf\x71\x22\xba\x00\xa9\x33\xc9\x30\xf8\x09\xe2\x07\x4b\x4b\x77\xca\xc8\x38\x23\x38\x3f\x2e\x25\x61\xce\x68\xab\xe4\x18\x87\xb8\xd1\x03\xb5\x3b\xe8\x56\xa8\xe2\x2a\x18\xac\x0d\xc2\x57\x4c\x0b\xf6\xf5\x82\x0d\x97\x73\x25\x4e\x50\x0e\x56\xbe\xbb\x55\xd1\x52\x09\x84\xa9\xba\xb8\x9c\xdf\x2a\xd6\x0d\xd5\x7c\xf0\xcf\x40\xd6\x1f\xe0\xc2\xf7\xe8\x9e\x4c\x9f\xf3\x3a\x38\x0d\xac\xa0\x62\xc2\xe9\x80\x0b\x3c\xf4\x08\x89\xd8\x83\x2d\x50\xd5\xa2\xc6\xc7\x1d\x54\xc6\xf8\xf2\xa7\x8d\xf0\x88\x78\x13\x4e\x48\x98\xa3\x6a\xa9\xe8\x88\x1e\xa8\xa7\xce\x1a\x9c\xcc\xcf\x2d\x12\xff\xec\x22\xc4\x75\x98\xf7\x0e\x42\xd6\x5b\xea\x19\x21\x4c\x4f\x19\xae\x5b\xfe\xf8\x63\xca\xc9\x7f\x28\xb9\x2d\xe9\xf5\x77\x61\x1e\xab\x3e\x0e\x3d\xf3\x7a\xd0\xd5\x26\xda\xc4\xde\x8b\xa8\xe5\xf8\x6f\xda\x77\x7b\x44\xe0\xea\xf3\x2f\xec\x04\x6f\x48\x2e\xdd\x8b\xa6\x40\x6e\xe9\xa0\xfe\x4f\x2d\xff\x11\x47\xfc\x79\xfe\xd5\x3d\x20\x0b\xbe\xdc\x6a\xa0\xae\x2e\xce\x11\xcf\x95\x5e\x81\xcd\x5e\xda\x14\xe3\xc1\xa2\x61\x92\xe9\x99\x7f\x4e\x93\xd1\xf7\x69\x1d\xc6\x66\xcc\x86\x0d\xe6\x78\xe3\x3b\xae\x90\x73\x83\xa4\x1e\xc4\x73\xbf\x62\x6d\xda\x29\x71\x5e\xf2\x05\x93\xcd\xfd\xbb\x99\x3e\x60\xee\xde\xfb\x07\x15\xd0\x05\x86\x74\x6e\x9b\x5a\x87\x87\x5e\xbb\xd3\x77\x9b\x9e\xf8\x37\xf0\x27\xaf\x38\xb5\x9c\xf4\x38\x99\x49\x7d\x1f\x0f\x5a\xdb\x17\x98\x9b\x91\x7e\x2c\x90\xf6\x15\x89\x91\xdd\x37\xe6\x79\xca\x67\x46\xe7\xdb\x4e\x7b\xcc\xdb\xbf\xb1\xf4\xdc\x6a\xac\xcc\xbf\xcd\x7b\xd8\xd5\x69\xf5\x6f\x00\x00\x00\xff\xff\xfc\xf3\x11\x6a\x88\x2f\x00\x00")

func dataConfig_schema_v31JsonBytes() ([]byte, error) {
	return bindataRead(
		_dataConfig_schema_v31Json,
		"data/config_schema_v3.1.json",
	)
}

func dataConfig_schema_v31Json() (*asset, error) {
	bytes, err := dataConfig_schema_v31JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/config_schema_v3.1.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataConfig_schema_v32Json = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5b\xcd\x73\xdc\x28\x16\xbf\xf7\x5f\xa1\x52\x72\x8b\x3f\x52\xbb\xa9\xad\xda\xdc\xf6\xb8\xa7\x99\xf3\xb8\x3a\x2a\x1a\xbd\x56\x13\x4b\x40\x00\xb5\xdd\x49\xf9\x7f\x9f\xd2\x67\x03\x02\x81\xba\xe5\x38\x33\x35\x27\xdb\xe2\xf7\x80\xf7\xfd\x1e\xe0\x1f\x9b\x24\x49\xdf\x4b\x7c\x80\x0a\xa5\x9f\x93\xf4\xa0\x14\xff\x7c\x7f\xff\x55\x32\x7a\xdb\x7d\xbd\x63\xa2\xb8\xcf\x05\xda\xab\xdb\x8f\x9f\xee\xbb\x6f\xef\xd2\x9b\x86\x8e\xe4\x0d\x09\x66\x74\x4f\x8a\xac\x1b\xc9\x8e\xff\xbe\xfb\xd7\x5d\x43\xde\x41\xd4\x89\x43\x03\x62\xbb\xaf\x80\x55\xf7\x4d\xc0\xb7\x9a\x08\x68\x88\x1f\xd2\x23\x08\x49\x18\x4d\xb7\x37\x9b\x66\x8c\x0b\xc6\x41\x28\x02\x32\xfd\x9c\x34\x9b\x4b\x92\x11\x32\x7c\xd0\xa6\x95\x4a\x10\x5a\xa4\xed\xe7\x97\x76\x86\x24\x49\x25\x88\x23\xc1\xda\x0c\xe3\x56\xdf\xdd\x9f\xe7\xbf\x1f\x61\x37\xf6\xac\xda\x66\xdb\xef\x1c\x29\x05\x82\xfe\x3e\xdd\x5b\x3b\xfc\xe5\x01\xdd\x7e\xff\xdf\xed\x1f\x1f\x6f\xff\x7b\x97\xdd\x6e\x3f\xbc\x37\x86\x1b\xf9\x0a\xd8\x77\xcb\xe7\xb0\x27\x94\x28\xc2\xe8\xb8\x7e\x3a\x22\x5f\xfa\xdf\x5e\xc6\x85\x51\x9e\xb7\x60\x54\x1a\x6b\xef\x51\x29\xc1\xe4\x99\x82\x7a\x62\xe2\x31\xc4\xf3\x08\x7b\x23\x9e\xfb\xf5\x1d\x3c\x9b\xec\x1c\x59\x59\x57\x41\x0d\x0e\xa8\x37\x62\xa6\x5b\x7e\x1d\xfd\x49\xc0\x02\x54\xd8\x64\x3b\xd4\x9b\x59\x6c\xb3\xfc\x75\x0c\x6f\x06\xa6\x67\xb1\x1d\x42\x5b\xbb\xdd\xa0\xe1\xde\x2e\x51\xb9\xdc\xcb\x2f\xab\x51\x58\x1e\x29\xe5\xc0\x4b\x76\x6a\xbe\x79\xe4\xd1\x01\x2a\xa0\x2a\x1d\x45\x90\x24\xe9\xae\x26\x65\x6e\x4b\x94\x51\xf8\xad\x99\xe2\x41\xfb\x98\x24\x3f\xec\x48\xa6\xcd\xd3\x8e\x1b\x7f\xf9\x15\x3e\x8e\x7b\x78\x19\xc7\x31\xa3\x0a\x9e\x55\xcb\xd4\xfc\xd2\x9d\x08\x18\x7e\x04\xb1\x27\x25\xc4\x52\x20\x51\xc8\x19\x91\x95\x44\xaa\x8c\x89\x2c\x27\x58\x39\xe9\x31\xc2\x07\xc8\xf6\x82\x55\xc1\x59\xf6\x59\xb7\x0f\x99\xbe\x58\xf3\x4c\x26\x0e\x1b\xe6\x48\xaa\xfd\xb5\xdd\x38\x26\x4c\x31\xe2\x19\xca\x73\x43\x20\x48\x08\x74\x4a\x6f\x92\x94\x28\xa8\xa4\x5b\x56\x49\x5a\x53\xf2\xad\x86\xff\xf7\x10\x25\x6a\xb0\xe7\xcd\x05\xe3\xeb\x4f\x5c\x08\x56\xf3\x8c\x23\xd1\x58\xea\xbc\x1e\x53\xcc\xaa\x0a\xd1\xb5\xcc\x77\x09\x1f\x11\x92\x67\x54\x21\x42\x41\x64\x14\x55\x21\x8b\x6c\xdc\x17\x68\x2e\xb3\xae\x72\x88\xb5\x24\x63\x82\xb1\x8c\x58\x55\x1f\x39\x9d\xf3\x90\x6e\x9a\xc6\x47\x9a\xbd\xa5\x16\x61\x26\x01\x09\x7c\xb8\x90\x9e\x55\x88\xd0\x18\xd9\x01\x55\xe2\xc4\x19\xe9\xec\xe5\x97\x33\x04\xa0\xc7\x6c\x0c\x4a\x8b\xc5\x00\xf4\x48\x04\xa3\xd5\xe0\x0d\x71\x91\x4a\xa3\x7f\xe6\x4c\x82\x2d\x18\x8b\x41\x7d\x68\x64\xd5\x90\xc9\x40\xf1\x30\x30\x7e\x93\xa4\xb4\xae\x76\x20\x9a\x62\xd8\x40\xee\x99\xa8\x50\xb3\xd9\x61\x6d\x6d\xd8\x90\xb4\xc3\xf2\x74\x01\xea\x3c\x34\xf5\x01\x2a\xb3\x92\xd0\xc7\xf5\x4d\x1c\x9e\x95\x40\xd9\x81\x49\x75\x49\x32\x48\x0f\x80\x4a\x75\xc0\x07\xc0\x8f\x33\xe4\x3a\xca\xa0\x66\x52\xc5\x18\x39\xa9\x50\x11\x06\x71\x1c\x82\x94\x68\x07\xe5\x45\x7c\xae\x2a\x7c\x6d\x5a\x56\x14\x0d\xd4\x67\x71\x93\x12\xa8\x1f\x0e\x15\x0f\xb9\x20\x47\x10\xb1\x95\x00\xe3\xe7\xca\xcd\x1e\x0c\x57\x32\x49\xb8\x8c\x35\xa0\x5f\xee\xba\x2a\x76\xc6\xab\xda\xdf\xca\x32\xdd\xda\xe5\x42\x62\xe5\x7d\xd7\x17\x8b\xc3\xb8\x82\xc2\xd0\x4a\x85\x70\x53\x37\x08\x90\x1e\xbd\x9e\xa1\x7d\x9b\x94\x55\x2c\xf7\x19\xe8\x04\x6c\xcb\xc6\x1b\xa9\x17\x27\xc2\xe4\xa2\x42\x34\x4a\x75\xc1\x4e\x24\xc0\x8d\x6f\x7b\xb1\xdb\x3c\x6f\x37\x6c\x62\x2d\x0e\x95\x04\x49\x08\x3b\xbb\x57\x90\xc6\x6c\x84\x1f\x3f\x45\xda\x84\x8b\xf6\x3f\xb3\xb4\x1e\x52\xef\x9c\xf1\x35\x72\x60\xaa\xf3\x56\x5a\x77\x73\x6d\x64\x1b\xf0\xb6\x57\x2e\xe1\x39\xc9\xfd\xb1\xa2\x8d\x10\xba\x83\x71\x26\xd4\xc4\xbb\x96\xa7\x7b\x9f\x05\xeb\xe2\x1a\xe2\xd4\x39\xe1\x77\x8b\x4f\xa4\x31\x51\x77\x14\xd1\xd4\xff\x82\xfe\x11\xf6\x8c\x74\x26\x4a\x39\xd0\x0a\x89\x02\xcc\x36\x84\x50\x05\x05\x08\x0f\x01\xaf\x77\x25\x91\x07\xc8\x97\xd0\x08\xa6\x18\x66\x65\x9c\x63\x38\xfb\xd8\x78\x67\x30\x27\xdc\x5e\x5d\x9b\x71\x41\x8e\xa4\x84\xc2\xe2\x78\xc7\x58\x09\x88\x1a\x89\x42\x00\xca\x33\x46\xcb\x53\x04\x52\x2a\x24\x82\xed\x9f\x04\x5c\x0b\xa2\x4e\x19\xe3\x6a\xf5\xaa\x50\x1e\xaa\x4c\x92\xef\x60\xfa\xde\xd9\xea\xfb\x89\xb6\xd6\x86\xac\x83\xb1\xe4\xb5\xdc\xcf\x67\xb6\xaf\xe4\x36\x92\xd5\x02\x5f\xe7\x38\xb3\xf8\xda\x0c\x72\xf3\xe0\x62\x09\x78\xe2\xf0\xbd\x0a\x43\x35\xd4\xac\xab\x38\x03\xb5\x3c\x49\xac\x2e\xab\xad\xa5\xca\x09\xcd\x18\x07\x1a\xf4\x0d\xa9\x18\xcf\x0a\x81\x30\x64\x1c\x04\x61\x4e\x51\x18\x01\x36\xaf\x05\x6a\xd6\x9f\x4e\x23\x49\x41\x91\x3b\xee\x68\x50\x55\xf1\xfd\x85\x87\x00\x4a\x85\x9d\xbd\x2e\x49\x45\xfc\x4e\xe3\xb0\xda\x88\x7a\xad\xab\xd5\xdc\x25\xda\x4c\x79\x16\x15\xb2\x67\x3a\x84\xf9\x06\x21\xa2\x33\x38\x20\xb1\x20\x75\xb4\x8e\xb9\xf7\xe4\x27\x57\xdf\xe0\xdc\x97\x71\xc5\xd5\xce\x77\xd3\x6f\x64\xeb\xc4\x2f\x2a\xbd\xec\x6d\x6c\xbd\xd5\x8f\xdb\xa9\x6a\x19\x6c\xe2\x5a\x0c\x95\x73\x0d\xc8\x08\x9d\xde\xd5\x24\x7f\x89\x08\x6d\xe8\xa8\x85\x3b\x74\x13\x11\xc7\xfb\x95\x22\x63\xe7\x6b\x47\xfd\xe8\x8a\x40\xa3\xd9\x91\xc9\x81\xef\x12\x49\xc6\xc9\x69\x44\xa1\xa2\x0b\x9d\xd1\x3d\x4b\xbc\xdb\xf5\x37\x72\x3f\x85\x15\xca\x30\xe3\x1e\x29\xc7\xb3\xb1\x34\x63\x5a\xa7\x10\x33\x25\xa5\xcf\xfb\x9f\x98\x78\x6c\x72\x4b\x4e\xdc\x41\x60\x63\x91\x2c\xb8\xc4\xb4\x8e\xed\x86\x09\x5c\xb7\x73\x3a\x34\x78\x9b\x39\x7f\x53\xd8\x83\xbc\xb7\x78\x44\xa2\x9d\x75\x7f\xe5\xca\x99\x4d\x90\x17\xc7\x70\xea\x16\xa0\x04\xb1\x6e\x05\x86\xfa\x47\x4f\xd3\x20\x7f\xcd\xb3\x73\x45\x2a\x60\xb5\x3b\xa2\x6c\x74\xc3\xe9\x89\x52\xed\x96\x33\xa0\x54\x0d\x69\xeb\xf4\x61\x54\xea\xd0\x62\x07\x15\x17\x93\x7b\x80\xe6\xed\x2d\x45\x54\xa2\x12\xc0\x4b\x82\x91\x0c\x15\x03\x57\x1c\xe8\xd6\x3c\x47\x0a\xb2\xee\x49\xcc\xa2\xf2\x6b\xa6\xee\xe2\x48\xa0\xb2\x84\x92\xc8\x2a\xa6\x8e\x49\x73\x28\xd1\xe9\xa2\x12\xb6\x25\xdf\x23\x52\xd6\x02\x32\x84\xbd\x61\xda\xa2\xa8\x18\x25\x8a\x39\xc3\x49\xdc\x92\x15\x7a\xce\x86\x65\x5b\x48\xa8\xbb\x30\x1b\xeb\xd8\xb3\x58\xcd\x12\xba\x34\xbc\xac\x42\x9e\x51\xd1\xb9\xde\xf6\x58\xcc\xb0\xe2\x84\x75\x01\xb2\x09\x3b\xe3\x51\x79\x90\x3e\x18\xe0\xfb\x4e\x3f\xe3\xac\x24\xf8\xb4\x16\x87\x98\xd1\x4e\xc8\x31\x06\x71\xa5\x05\x36\xe6\xd0\xb4\x23\x15\x57\x41\x67\x6d\x09\x9e\x08\xcd\xd9\xd3\x82\x05\xd7\x33\x25\x5e\x22\x0c\x56\x70\xbc\x56\xd0\x52\x09\x44\xa8\x5a\x7c\x33\x74\x2d\x5b\x57\xa4\xfe\xd1\x3e\x03\x29\x62\xc4\x05\x93\xbe\x2f\x2d\x60\x5e\x07\xef\x4f\x2a\xa8\x98\x70\x1a\xe0\x0a\x6f\xec\x42\x2c\x0e\xb0\x15\x52\x60\xd4\x85\x5b\x8f\xca\x18\x5f\xbf\xe3\x0f\x5f\xaa\x6d\xc3\x01\x89\x70\x54\xad\xe5\x1d\xd1\x57\x90\xa9\x33\x07\x27\xf3\x9d\x69\xe2\xef\x4e\x43\xbb\x0e\xef\xbd\x47\xc8\x7a\x47\x3d\x0d\xdd\xb4\x19\x58\xf3\x60\x7a\xc5\xa0\x37\xbc\x1e\xf0\x68\xf5\x61\x2c\xb0\x6f\x46\x59\x6d\xa3\x55\xec\xbd\xba\x5f\x6f\xff\x6d\xad\x6f\x1f\xd3\xb9\x9a\x02\xa4\x14\xc2\x87\xa8\xfe\x61\x61\xd1\x78\x45\x1c\x9a\x74\xb9\xce\x30\xd4\xa3\xfe\x89\x42\x7f\x13\x9b\xfd\x79\xf6\xd5\x3f\xf3\x0d\xbe\xaf\x6d\x51\x17\xe7\xf1\x88\x47\xa5\xbf\x80\xce\xde\x5a\x15\xe6\x3d\x80\xa6\x92\xe9\x59\xc2\x9c\x24\x97\x3e\xa4\xdd\x9a\xdb\xb0\x61\x8e\xff\xc4\x30\x93\xe9\xdc\x2d\xe1\x00\xf1\x9c\x5d\x59\x8b\xf6\x42\x9c\xe7\x7c\xc5\x60\x73\xf7\x61\xa6\x64\x98\x7b\x54\xf4\x4a\xb9\x76\x85\x1b\x58\xb7\x4e\xad\x3e\x63\x90\xee\xf4\x75\xbd\xc7\xff\x35\xfa\xc9\x5b\xfb\x86\x4f\x7a\x9a\x9c\x75\xfd\x30\xcf\xdc\xbb\x77\xf2\x5b\x43\x3e\x16\xa4\x7b\xa2\xa7\x45\xf7\xad\xde\x7a\xf9\xd4\xe8\x7c\x81\x6f\x9f\xf8\x0f\x2f\xe1\x3d\x97\x90\x1b\xfd\x67\xfb\x5f\x0b\x9b\x97\xcd\x9f\x01\x00\x00\xff\xff\x57\x1e\x50\xbb\x2e\x35\x00\x00")

func dataConfig_schema_v32JsonBytes() ([]byte, error) {
	return bindataRead(
		_dataConfig_schema_v32Json,
		"data/config_schema_v3.2.json",
	)
}

func dataConfig_schema_v32Json() (*asset, error) {
	bytes, err := dataConfig_schema_v32JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/config_schema_v3.2.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"data/config_schema_v3.0.json": dataConfig_schema_v30Json,
	"data/config_schema_v3.1.json": dataConfig_schema_v31Json,
	"data/config_schema_v3.2.json": dataConfig_schema_v32Json,
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
	"data": &bintree{nil, map[string]*bintree{
		"config_schema_v3.0.json": &bintree{dataConfig_schema_v30Json, map[string]*bintree{}},
		"config_schema_v3.1.json": &bintree{dataConfig_schema_v31Json, map[string]*bintree{}},
		"config_schema_v3.2.json": &bintree{dataConfig_schema_v32Json, map[string]*bintree{}},
	}},
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

