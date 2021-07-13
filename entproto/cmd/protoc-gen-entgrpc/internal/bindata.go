// Package internal Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// template/enums.tmpl
// template/method_delete.tmpl
// template/method_get.tmpl
// template/method_mutate.tmpl
// template/service.tmpl
// template/to_ent.tmpl
// template/to_proto.tmpl
package internal

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templateEnumsTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x92\x4f\x6b\xdc\x30\x10\xc5\xef\xfe\x14\x83\xe8\xc1\x86\x22\x4a\x7b\x6b\xe9\xad\x9b\x25\x90\x96\x85\xa6\xa7\x50\x82\x76\x3d\x4e\x45\x6c\x49\x95\xc7\x2e\x41\x9d\xef\x5e\x24\x7b\xd7\xf6\xfe\x29\xdb\xea\x64\x34\xf3\xde\x6f\xf4\xc6\x21\x40\x89\x95\x36\x08\x02\x4d\xd7\xb4\x02\x98\x33\x00\x80\x10\xe0\x95\xb7\x96\xe0\xfd\x47\x90\xb3\x4b\xaf\xcc\x13\x82\xbc\xd1\x58\x97\x9f\x95\x93\xab\xa8\xda\xd7\xf7\xc2\x68\x75\xff\xe2\x30\x89\x37\xdb\xd4\xfc\x09\xdb\x9d\xd7\x8e\xac\x97\x6b\xa4\xd5\xbe\xe3\x8c\xf2\x8b\x6a\x92\xd2\x79\x6d\x68\x98\x42\xae\x0c\xc5\x76\x99\x6a\xe2\x51\x4c\x8c\xe8\x96\x6e\x8f\x9c\xdc\x36\x32\x6e\x4b\x34\xe9\x0d\x83\xcd\x8d\xae\x51\xae\xed\x6d\xe3\xac\xa7\x8d\xa2\x1f\x72\x68\x98\xb8\x70\x3a\x11\xdd\xed\x54\x9b\x26\x6a\x8d\x7a\xc6\x73\x13\x9d\x6a\x16\x70\x34\x74\xe0\x8c\x6e\x72\xb3\xfd\x4a\xbe\xdb\x51\x0a\x67\x6e\x50\x75\x66\x07\x64\x37\xde\x92\x3d\x7a\x88\x5c\xdb\x11\x07\x39\x46\x92\x3e\xb8\x4e\x3c\x60\x2e\x66\xb5\x79\x0e\xcc\x10\x0e\x9c\x78\x74\x05\xfd\x6b\xb0\xcf\x71\xc8\x8b\xac\xc7\x5e\xd5\x1d\x3e\x84\x00\x3f\x3b\x55\xeb\xea\x05\x44\x4b\x5e\x9b\xa7\x56\x80\xb8\xb7\xdf\x9c\x43\x1f\xff\x9b\x7c\xb8\xcd\xb1\x28\xbe\x7f\x88\xa6\x4b\x56\x3c\x1e\xa9\xf3\xe6\x64\x3f\xbf\xc7\x61\x99\xf3\xbe\x58\x88\x38\xfb\x37\xf9\x9b\x49\xce\xd9\x71\xa4\x2b\x43\x17\x1f\xb9\xc8\x73\x99\x59\xf1\x97\xa4\xff\x2b\x4f\xa3\x1a\x7c\xd0\x86\xde\xbd\xcd\xf1\x8a\xa8\xce\xa1\x99\xf3\x4b\xfb\xb8\xb3\xbf\xc6\x7d\xf4\xc5\x15\x61\x0a\x91\x2d\xeb\x21\x00\x9a\x92\x39\x1b\x3e\x80\xf9\x4f\x00\x00\x00\xff\xff\xeb\x69\x69\x55\x25\x04\x00\x00")

func templateEnumsTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateEnumsTmpl,
		"template/enums.tmpl",
	)
}

func templateEnumsTmpl() (*asset, error) {
	bytes, err := templateEnumsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/enums.tmpl", size: 1061, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateMethod_deleteTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x52\xc1\x6e\xdb\x30\x0c\xbd\xf7\x2b\x08\xa2\x1b\x92\xa2\x91\xef\x01\x72\x4b\x5a\xf8\xb0\xae\xc0\x86\x5d\x03\x45\xa2\x5d\xa1\x0a\xe5\xca\x74\xda\xc0\xd3\xbf\x0f\x92\x93\x74\xc5\x36\x60\x04\x0c\x3c\x8b\x8f\xef\x51\x14\xc7\x71\x01\xd5\x4d\x1b\xe4\xd8\xd1\x12\x88\xa5\x0d\xca\x85\xca\x04\x96\xe8\x76\x15\xb1\x74\x31\x48\xa8\xcc\xde\x56\x05\x99\x45\x4b\xbc\xc8\xc4\xd8\x19\xb5\x27\x79\x0a\xb6\xe6\x6e\x90\x9b\x0a\x16\x29\x5d\x8d\x23\x58\x6a\x1c\x13\xe0\x94\xdc\x5a\xf2\x24\x84\x90\xd2\x15\x00\x40\x76\xbc\x86\xe5\x0a\xd4\x7d\x29\xb8\x9c\x39\x7b\xe7\xc8\xdb\x9c\xba\x56\x05\x7e\xd1\x9d\xaa\xd7\x1f\x59\x07\x1d\x1f\xf4\x9e\x0a\xeb\x54\xa1\x36\x2c\x13\x28\x99\x33\xfd\xa0\x23\x50\x2c\x5f\x88\x17\x01\xa1\x7d\xe7\xb5\x10\x60\x93\x4b\xb6\x12\xb6\xc4\x82\x60\x9d\x11\xc0\x22\x83\xef\xbd\xe0\x8f\xc9\x0e\xff\x65\x86\xb5\x2d\xe5\xb3\x2e\x3a\x16\xc0\x48\x2f\xea\x9e\xe4\x37\xfe\xe3\xee\x9b\xc4\xc1\xc8\x49\x70\x36\xc7\xf9\x79\x14\xb9\xbb\x15\xf4\x07\xa3\x8c\x77\xc4\xa2\xc6\x11\xae\xb3\xc1\xf7\x63\x47\x93\x7e\x4a\x6a\x5d\xe6\xf7\x95\xa9\x5e\xcf\x32\xe1\x3c\x81\x94\xe6\x6a\xf3\x46\x66\x66\xe4\x6d\x5e\xf4\xfa\x57\x27\xe6\x09\xc6\xf2\x93\xc3\xe8\x9e\x26\x97\x15\xb0\xf3\xcb\x4b\x22\x47\x24\x19\x22\xc3\xe7\x71\x84\x97\x41\x7b\xd7\x1c\x01\xdb\x10\x5a\x4f\xaa\x0d\x5e\x73\xab\x42\x6c\xa7\x47\xdf\x0d\x4d\x95\x37\xa4\xaf\x9e\x39\xbc\x72\x45\xfb\x4e\x8e\xdd\x0e\x01\x37\x19\xe5\xb7\x1d\xd3\x6d\xb6\xf8\x68\x7d\xbe\xcf\xa3\x36\xcf\xba\x25\x55\x86\x05\x58\xf7\x0f\x41\xee\xc2\xc0\x16\xe1\x27\xb8\x72\x98\xd2\x8c\x62\x9c\xff\xb5\x45\x76\xfe\x36\x6b\xf5\xa2\x65\xe8\x37\x31\x36\x80\xef\x0a\xc8\x41\xa0\xc9\x78\x09\x9f\x7a\x04\xa4\x18\xf1\x34\xe1\x1c\x96\x1a\x3d\x78\xf9\x6f\xe5\x9a\x85\x22\x6b\x8f\x80\xee\x04\xa7\x1d\xfa\x53\xbe\xac\x3b\xb1\x85\x94\x7e\x05\x00\x00\xff\xff\xfe\xff\xf2\xdb\x49\x03\x00\x00")

func templateMethod_deleteTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateMethod_deleteTmpl,
		"template/method_delete.tmpl",
	)
}

func templateMethod_deleteTmpl() (*asset, error) {
	bytes, err := templateMethod_deleteTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/method_delete.tmpl", size: 841, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateMethod_getTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x54\x4d\x6f\xf2\x46\x10\xbe\xe7\x57\x8c\x56\x54\xb2\xd1\xcb\x72\x47\xca\x21\x0d\x0e\xf5\x21\x29\x15\x69\x72\x44\x1b\x7b\x6c\x56\x31\xbb\x66\x3d\x86\x22\xd7\xff\xbd\xda\xf5\x07\x36\xa1\x6a\xdf\x58\x02\x86\xd9\xf9\x78\xe6\xf1\x33\x5b\x55\x33\x98\x4f\x53\x4d\xe7\x1c\x17\x80\x8a\x52\xcd\xa5\x9e\x47\x5a\x91\x91\x1f\x73\x54\x94\x1b\x4d\x7a\x1e\xed\xe3\xb9\xb3\xa2\x59\x8a\x6a\x66\x03\x4d\x1e\xf1\x3d\xd2\x4e\xc7\xa1\xca\x4b\x9a\xce\x61\x56\xd7\x77\x55\x05\x31\x26\x52\x21\xb0\xe6\x70\x9b\x22\x31\xa8\xeb\x3b\x00\x00\xdb\x6e\x02\x8b\x7b\xe0\x2b\x17\xdd\xfb\x64\xfc\x24\x31\x8b\xed\xd1\x84\x3b\xf3\x59\xe4\x3c\x5c\x8e\xa3\x8e\xc2\xbc\x88\x3d\xba\xa8\x36\x83\x07\x8a\x1a\xc3\x9d\x8c\x8b\x5a\x5c\x5d\x02\x7f\x76\x70\xb8\x03\xcb\x57\x3a\x8c\x51\xd9\xdf\x51\xda\x51\x18\xf0\x9c\x65\x1f\x34\xc6\x7e\xb4\xe9\x3d\x29\x12\x4c\xab\x0a\x26\xb6\xed\x5a\x44\x9f\x22\x45\xee\x2a\x35\xae\xd7\x73\x8e\x0d\x90\xbf\x41\x3a\x77\x5b\xd8\xef\x51\x11\xee\xf3\x4c\x10\x02\x4b\x2c\xec\x2d\xe9\x2d\x2a\x62\x10\xcb\x88\x80\xb9\x51\xd8\x85\x0f\xf6\xd6\x8c\xcc\xfe\x6d\x60\xe6\xba\x33\xf0\x72\x23\x15\x01\x33\x78\xe0\x2b\xcb\x78\x1f\xbf\xfe\xd8\x90\x29\x23\x6a\x0b\x7a\x3e\xf3\x3b\x54\xc5\x49\x52\xb4\x83\x36\xe7\x4d\xe2\xc9\xf3\xa1\xea\xa7\x8d\x44\x81\x60\xa7\xbd\x10\x59\xd7\xdb\xb7\x30\x78\xdf\xfe\xf9\xb2\x59\x07\x8f\xe1\x53\x18\x2c\x7f\x7c\x0d\xf9\xf5\x61\x13\x3e\x2e\xfa\x3a\x2d\x73\x3f\x1c\xa1\xf7\x50\x1c\x23\x1e\x65\xd2\xd2\xdf\x51\x79\xe1\xad\xae\x2d\x14\x2f\xa2\xbf\x9a\xba\xdd\x2b\xaf\x6b\xff\x3f\x70\xbd\x87\xaf\xbf\x6d\x83\xe5\x2a\xd8\x86\xcb\xcd\xb7\x9b\xff\x51\xa2\x39\x7b\x3e\x1f\xe5\xbf\xef\xd0\xa0\x57\x55\x70\x28\x45\x26\x93\x73\xc7\xb6\x57\xaa\x43\xa9\x09\xc7\x7a\xd8\x90\x91\x2a\xf5\x81\xcd\xd9\xa0\x41\x7b\xea\x03\x0b\x97\x76\x1f\xbc\xab\xe9\xae\x5a\x56\x15\x18\xa1\x52\x1c\xae\x43\x10\xa7\x58\x74\xef\x6e\x1c\x3c\x83\x09\x92\x93\x79\xa0\xc8\xc6\x71\xdb\xb4\xd7\xf5\x68\x18\x49\xbb\xaa\xba\x04\x0e\xd5\x51\xd7\x5e\x52\xaa\xc8\x3b\x58\x16\x60\xda\xb1\x84\xd4\xf1\xe3\xe8\x19\x6a\x64\xf8\xb8\x2c\xbe\xc1\x0c\x23\xb2\xe3\x7d\x83\x2e\xa4\x8e\x28\xf0\xdd\xbf\x70\xc9\x1f\xb5\x2a\x48\x28\x82\x91\x08\xba\xa7\xfe\x4a\x1c\xaa\xf8\x9a\xa5\xdf\x55\x76\xb6\xa2\xba\xe4\xc7\x98\x88\x32\xa3\xb1\x50\x0c\x52\x69\x14\x28\x99\x39\xf5\x15\x24\xa8\x2c\x02\x63\x80\x85\xea\x28\x32\x19\x3f\x98\xb4\xdc\xbb\x85\x63\xb2\xf1\x80\x68\x5d\x0b\x28\xd5\xa7\xd2\x27\x05\x47\x89\x27\xd6\x02\x18\x2d\xda\xd5\x6a\x39\x45\xde\xdb\x6e\x37\x51\x90\x5e\xdb\x1b\xf7\x96\x4c\xbd\x14\xe9\xc6\x3e\x7c\xbd\x95\x58\x58\xbc\x68\x7a\xd2\xa5\x8a\xd9\xe0\x4e\xf2\xd0\x18\xff\xff\x8d\x9e\x00\xbb\x54\x60\x4a\x13\x24\xd6\x5e\xc0\x2f\x05\x03\x86\xc6\xb0\x21\xd7\x3f\xc5\x6a\x62\x69\x25\x34\x4a\x64\x8e\xcf\xc6\x6c\xae\xdc\x1b\xf5\x9b\xef\x61\x31\x25\xb3\xbb\xfe\x7d\xff\x13\x00\x00\xff\xff\x48\x39\xea\x99\xcb\x06\x00\x00")

func templateMethod_getTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateMethod_getTmpl,
		"template/method_get.tmpl",
	)
}

func templateMethod_getTmpl() (*asset, error) {
	bytes, err := templateMethod_getTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/method_get.tmpl", size: 1739, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateMethod_mutateTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\x4d\x6f\xdb\x38\x10\xbd\xf7\x57\xcc\x12\x59\xc0\x2e\x12\xfa\x6e\x20\x87\xa0\xf1\x06\x3a\x74\x1b\x20\xdb\x5e\x0d\x46\x1a\x29\x44\x25\x4a\x26\x47\xd9\x06\x5a\xfd\xf7\xc5\x50\x8a\x4d\x7d\xb8\x8e\x77\x1b\x01\x86\x29\xce\xcc\xe3\x9b\x99\x47\x8a\x4d\x73\x05\xab\x8f\x59\x49\x2f\x15\xae\x01\x0d\x65\xa5\xd4\xe5\x2a\x2e\x0d\x59\xfd\xb8\x42\x43\x95\x2d\xa9\x5c\xc5\x45\xb2\xf2\xa3\xf8\x2a\x43\x73\xc5\x8e\xb6\x8a\x65\x81\xf4\x54\x26\x91\xa9\x6a\xfa\xb8\x82\xab\xb6\xfd\xd0\x34\x90\x60\xaa\x0d\x82\xe8\x8c\xdb\xa2\x26\x45\x28\xa0\x6d\x3f\x00\x00\xf0\x8a\x17\xb0\xbe\x06\x79\xe7\x03\xf6\x73\x3a\xf9\x43\x63\x9e\xb0\xe9\x42\xfa\xe1\x67\x55\xc9\xe8\x76\xe8\xf5\xac\xec\x9f\xaa\x40\xef\xd5\x47\xc8\x8d\xa1\x6e\xe0\x2d\x43\x50\xa6\xf6\x1a\x20\x3f\x7b\x46\xd2\xf3\x95\x77\x65\x94\xa0\xe1\xff\x69\x58\x47\x7d\x1c\x37\xe7\x69\x71\xf7\x4d\x59\xf6\x8a\x55\x81\x39\x5c\x30\x9b\xbf\x5e\x2a\x1c\x93\xd9\xbb\xb6\x2d\x7b\x5b\xdc\xc9\x3b\x24\x9e\x1f\x46\xb4\xed\x62\xb9\x87\xd7\x29\xe0\x6e\xbc\xbe\xf8\x64\x31\xac\x28\x3f\x05\x63\xba\xe7\x58\xc6\xb9\xe6\xa4\xe6\x70\x65\x17\x17\xc0\x63\xee\x30\x44\x19\x57\xb8\x4b\x69\x51\x59\x6d\x68\xcf\x5f\x6c\xc5\x91\xd2\x2f\xf7\xe9\x06\x4d\x65\x9c\x11\x00\x27\x1e\x60\xdc\x3f\x3e\x90\xad\xe3\x0e\x08\xc4\x62\x09\x62\x02\x44\x58\x54\xb9\x22\x04\x91\xb2\xd7\x96\xca\x2d\x1a\x12\x90\xe8\x98\x40\xf8\xc8\x03\x22\x88\x6f\x5d\x0e\xe2\x90\x8d\xf0\xcd\xf6\x3e\x67\xd7\xed\x6b\x95\x28\xc2\x2f\x06\xa3\xdb\x05\x3b\xbc\x62\xb6\x6d\x50\x4a\x93\x84\x0a\xb7\xca\x64\x18\xea\xd8\x0f\xdc\xa4\xd8\xee\xbb\xae\xa2\x82\xb7\xc8\x63\xee\x4b\xae\x4c\x02\x0b\x6e\x7a\xa8\x41\xd1\x31\x10\xb0\x84\x43\xc5\x0f\x61\x93\xaa\x33\x2a\x83\x95\x16\x64\xe4\xa2\xdb\xae\x2a\xa3\xc5\xc6\x51\x3a\x05\x53\x52\x1f\x1c\xd8\xde\x2c\x0c\xaf\x0c\xf8\x99\x24\xde\x22\x8b\x93\x72\x08\xf8\x1e\xd6\xfa\x52\x91\x2e\x8d\xca\xc7\xc4\xf9\xd1\xa9\xdf\x7e\xbe\xf1\xf0\xdb\x35\x18\x9d\x43\x33\x81\x0b\x1a\x78\xb6\xf0\xe4\x19\x8a\xf3\xaa\x93\x0f\x7e\xeb\x1f\xf8\x87\x29\xb7\xed\xac\xca\xfe\x53\xee\xd3\x84\x46\x79\xce\x68\xf7\x94\x94\x37\x49\x86\x13\x25\xf7\x8c\xd8\x26\xbf\x1a\xbd\xab\xf1\x7f\x6a\xc8\x23\xfd\x2a\x09\xf5\xf3\x8c\x19\xdd\x4e\xac\xf3\xfa\x7a\xf7\xce\xfb\x0c\x87\x8d\x3f\x72\xc0\xc0\x91\xf3\x9a\x9f\xb4\xb4\xb0\xbd\x04\x4d\xe8\xcf\xb2\xae\x5b\x83\xcf\x4d\xff\x9d\x19\x55\x85\xbf\x33\xa3\x6d\x30\xe9\xd2\xa1\x4d\x47\xf8\xce\x87\x87\xad\x11\x4c\xec\x44\xf9\x8f\xe0\xbc\x4b\x03\xba\x26\xdc\x24\x49\xd3\x80\xd3\x26\xab\x73\x65\x8f\x77\xc3\xfd\x74\x27\xbe\x75\x17\x59\x74\x97\x80\xd6\xdf\x12\x0a\xf9\xa0\x9e\x71\x11\xd3\x8f\x0e\xcb\xfd\xad\x29\x7e\x0a\x1a\x11\x2b\x87\xde\xf9\xda\x9f\x54\xeb\xc1\x8a\xfe\x0a\xb6\xc7\xa2\xf2\x9e\xdf\x67\xef\x10\x16\xdd\x90\x2c\x5f\x25\xac\x9d\x3f\xff\x3a\x92\x54\x5b\xc3\xc6\x4b\x96\x8f\x23\x45\xb5\xdb\x58\x9b\x82\x88\x0c\xa1\x35\x2a\x17\x20\x74\x3f\x64\xac\xd2\xae\xe1\x77\x27\x40\xa0\xb5\x62\x5c\xe7\xe1\x5b\x0f\xde\xd3\x37\x3a\x1f\xa6\xdb\x34\xb0\xab\x55\xae\xd3\x17\x10\xfb\x3b\x28\x1a\x5a\x25\x5a\xe5\x18\xd3\xca\xed\x72\xfe\x65\x56\x55\x4f\x02\x44\xe4\xba\x23\xe6\x53\x69\x1c\x59\xa5\x0d\x6d\x98\x0e\x93\x58\xa0\xb5\xcb\xf5\xdc\xe2\x73\x99\xdd\xe4\x16\x55\xf2\xb2\xf9\xa1\x1d\x71\x26\xaa\x7b\x07\xf4\x13\x41\x7a\x41\x76\xaf\x8c\x7d\xcd\xef\x55\xfc\x5d\x65\x28\xbd\xe0\x98\xd8\x84\xd2\x3f\xa0\xbd\xed\x5c\x6a\x91\x79\x56\xb9\x4e\x6e\x6c\x56\x17\x5e\xcc\x42\x77\x33\xa0\xfa\xa9\x79\x7a\x09\xa6\xaa\xce\xe9\x8c\x75\x4e\x37\xb7\x87\xf7\xb7\xfb\x4e\xd7\xff\x06\x00\x00\xff\xff\xe4\x37\x23\x6a\x38\x0c\x00\x00")

func templateMethod_mutateTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateMethod_mutateTmpl,
		"template/method_mutate.tmpl",
	)
}

func templateMethod_mutateTmpl() (*asset, error) {
	bytes, err := templateMethod_mutateTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/method_mutate.tmpl", size: 3128, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateServiceTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\x41\x6f\xe2\x3c\x10\xbd\xe7\x57\x8c\xa2\xea\x53\xa8\xc0\xb9\x23\x7d\x27\x68\x11\x87\xc2\x4a\xdb\x9e\x91\x9b\x0c\x59\x6b\x13\x3b\x38\x13\x5a\x94\xcd\x7f\x5f\x79\x9c\x00\xd9\x86\x6d\xf7\x84\x3d\xf3\xe6\xcd\x9b\xf1\x23\x4d\x33\x83\xf8\x3e\x33\x74\x2a\x71\x0e\xa8\x29\x33\x42\x99\x38\x31\x9a\xac\x7a\x8d\x51\x53\x69\x0d\x99\x38\x29\xd2\x98\x4f\xc9\x2c\x43\x3d\x73\x40\x5b\x26\xa2\x42\x7b\x54\x09\xae\x50\xa3\x95\x64\xec\x7d\x0c\xb3\xb6\x0d\x1c\xeb\x1d\xcc\xff\x07\xc1\xd7\x38\x86\x85\x49\x11\x32\x0f\xc3\x14\x5e\x4f\x30\xc2\x06\xcb\x2d\x6c\xb6\xcf\xf0\xb0\x5c\x3f\x8b\xa0\x94\xc9\x4f\x99\x21\x34\x0d\x88\x47\x95\xa3\x58\x99\x6f\x3e\xb4\x91\x05\x42\xdb\x06\x8e\xd8\x65\xbf\x7b\x15\x62\x65\xba\x0c\xa8\xa2\xcc\xb1\x40\x4d\xd5\x38\xc0\x05\xd0\x06\x6e\xea\x1b\x0c\x15\xd9\x3a\x21\x68\x02\x00\x80\x24\x57\xa8\x09\xee\x1d\xf4\x41\x53\x27\x43\xac\x53\x17\x0d\x17\x9c\x0d\xe1\x17\x28\x0e\xb4\x2d\x17\xbd\xe8\xb3\x0a\x4c\xff\xa6\xc2\x0f\xb2\xc1\xb7\x71\x25\x16\xa9\xb6\xba\x02\x09\x1a\xdf\xc6\xd5\x06\xfb\x5a\x27\x37\x19\xa2\x7f\x95\x3f\xf1\xd0\x8f\x52\xfc\x36\xbc\x20\xf8\x6f\x14\xe3\x21\x97\xa5\xcd\xbb\xdf\x29\xc7\x5b\x37\x6c\xd3\x00\x61\x51\xe6\x92\x10\x42\xd4\x75\x51\x85\x20\xf8\x3d\x07\x19\x32\x3b\xf6\xc8\xce\x0d\x77\x85\xb0\x52\x67\x78\xe9\xfc\x84\xf4\xc3\xa4\x55\xbf\x75\xb6\x9e\x4a\x1f\x15\xe6\xa9\x73\xe0\x9d\xe0\xe3\x93\x2c\xc5\x7a\xc9\x6e\x3c\xa3\x8e\xd2\xb2\x68\x87\xea\x2a\xdc\x7a\xfc\x81\x33\x03\x78\xc1\x8d\xfa\x8a\x7e\xe2\x01\x44\xe9\xb2\xa6\x33\x62\xed\x6e\x62\x65\x78\xcf\xd7\x78\x2e\xe8\xbc\x7b\xcb\xb3\x77\xb7\xec\x22\x06\x55\x4c\xc5\xaf\x1f\x55\xc7\x84\xdf\x6d\xa4\x72\x32\x6c\x15\x25\xf4\xee\x22\x87\x5a\xe6\x6a\x7f\x82\xd0\xfd\xdd\xf1\x9d\x42\x08\x17\xfd\xa9\x6d\xa7\x60\xf1\xc0\x8c\xde\x17\xc3\x79\x98\x35\xba\xca\x6e\x6b\x1a\xa6\xa7\x80\xd6\x1a\x3b\x81\x8b\x25\xdc\x92\xd4\x1e\xf0\x30\xd8\x66\xb8\x42\x6e\x78\x86\x79\xe8\x95\x15\x3c\x78\x97\x39\x5c\xe4\x2f\x20\x26\xd7\x25\x8e\x19\xf3\x0a\xc7\xe8\x97\x98\x23\xe1\x97\x3a\xa4\x1d\xf4\xf3\x26\xc6\x42\xf4\x67\xa3\x85\x45\x49\x18\x4e\x3e\x66\x5e\xca\xd4\x67\xbe\xa0\xa1\xa8\x49\x7e\xa2\x41\xa7\x7d\xc8\x7d\x6b\xfb\xfb\xef\x00\x00\x00\xff\xff\xc6\xfe\x90\x34\xcc\x05\x00\x00")

func templateServiceTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateServiceTmpl,
		"template/service.tmpl",
	)
}

func templateServiceTmpl() (*asset, error) {
	bytes, err := templateServiceTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/service.tmpl", size: 1484, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateTo_entTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x94\x4f\x6f\xd3\x40\x10\xc5\xef\xfd\x14\x83\x05\x28\xad\xc8\xfa\x1e\x94\x03\x54\xa5\xea\x01\x2e\xfc\xb9\x56\xcb\x7a\x1c\x46\x72\x66\xa2\xd9\x89\x51\xb5\xda\xef\x8e\x76\x53\x05\xa7\xd0\x9a\x03\x87\xfa\x34\xde\x99\xf7\x9e\xe6\xe7\x3f\x29\x2d\xa1\xbd\xd8\x88\xdd\xed\x70\x05\xc8\xb6\x11\x47\xd2\x06\x61\x53\xfa\xde\x22\xdb\x4e\xc5\xa4\x0d\xdb\xae\xad\x55\x58\x6e\x90\x97\x65\x50\x77\xc1\x45\xd4\x91\x02\x5e\x23\xa3\x7a\x13\xbd\x68\x61\x99\xf3\x59\x4a\xd0\x61\x4f\x8c\xd0\xf4\x84\x43\x77\x6b\x72\x8b\x6c\x0d\xe4\x7c\x06\x00\x50\x52\x5f\x52\x07\xab\x35\xb8\x9b\x0e\xd9\xaa\xea\xd8\x09\xc2\x63\xe9\x31\xfe\xbc\x14\x1e\x51\x0d\x15\xdc\x87\xe2\x74\x32\x48\xfd\x61\xd6\x7d\x91\x2b\xb6\x8f\xd2\x51\x4f\xa8\xc7\x91\x69\xd2\x1a\x76\x4a\x6c\xb5\x7e\x4a\x53\xe6\x91\x9f\x8c\xf1\x1a\x7f\xf8\x61\x40\xbd\x14\x8e\xa6\xfb\x60\xa2\xee\x5a\x3e\xf9\x2d\xc2\x24\x79\xf4\x0a\x29\x81\xfb\xe6\xf5\xbe\x55\x6e\xa9\x6e\x3b\x67\x37\xb1\xa1\x1e\x50\xb5\xd0\x58\xbc\x3e\xb5\x3b\x77\x5f\x79\x7b\x50\xbf\x27\xf6\x7a\xb7\x28\x01\x65\xc1\x9c\xcf\xdf\x56\xd5\x8b\x35\x30\x0d\x90\x8e\x6e\xe5\x52\xb4\xbd\x72\x69\xbc\x29\x82\x68\xde\xf6\xf1\x4a\xb5\x87\xe6\x86\x47\x3f\x50\xf7\x4e\x37\xfb\x6d\x7d\x5e\x0d\x1d\x4e\xc0\xdf\x1f\xad\xe0\x55\x6c\xa0\x41\xd5\x66\xba\xec\x04\xde\x10\xf1\x01\xb1\xcf\xc1\x33\xcf\xe1\x7a\x80\x6a\xb5\xfe\x2b\xad\x3f\xad\x20\xe7\xf4\x6f\xb4\x8a\x76\x82\x08\x9e\x15\xa3\xff\x01\xe7\x94\xca\xe2\xf7\xcb\x30\x13\x3c\xa2\x46\x12\x9e\x0d\x7c\x4c\xf5\x68\xd2\xac\x61\x15\x9d\x7c\x77\xf5\xe7\x71\x2c\x7f\x05\x00\x00\xff\xff\xda\x56\x49\xac\x9e\x04\x00\x00")

func templateTo_entTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateTo_entTmpl,
		"template/to_ent.tmpl",
	)
}

func templateTo_entTmpl() (*asset, error) {
	bytes, err := templateTo_entTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/to_ent.tmpl", size: 1182, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateTo_protoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x56\x4d\x6f\xdb\x30\x0c\xbd\xe7\x57\xb0\x42\x31\x24\x41\x22\xdf\x33\xe4\xb2\xb6\x2b\x7a\x68\x51\x60\x5d\xaf\x81\x6a\xd1\xae\x50\x47\x72\x65\xc5\x45\xe0\xe9\xbf\x0f\x92\xed\xc4\x5f\x69\x13\x60\xc5\x36\x9f\x64\x99\x7c\x24\xdf\xa3\x28\x17\xc5\x1c\x82\x69\xac\xcc\x36\xc5\x05\xa0\x34\xb1\xa2\x42\x05\xa1\x92\x46\x8b\xa7\x00\xa5\x49\xb5\x32\x2a\x08\xd7\x3c\xf0\xab\x70\x1e\xa3\x9c\x3b\x43\x9d\x86\x34\x43\x9d\x8b\x10\xaf\x51\xa2\x66\x46\xe9\x69\x00\x73\x6b\x47\x45\x01\x1c\x23\x21\x11\x88\x51\x2b\xef\xb7\x8a\x36\x32\x24\x60\xed\x08\x00\x20\x08\xc0\xa8\x7b\xb7\x5f\x14\x40\xaf\xa4\x79\xd8\xa6\x48\xef\xd8\x1a\xc1\x5a\x30\x9a\xc9\x2c\x52\x7a\x9d\x81\x79\x46\x97\x15\xb8\xfc\xc0\x28\xff\x9e\x3e\xf9\x57\x0f\xe4\x50\xdf\x81\x1a\x23\x4c\xab\xed\x7b\x16\xbe\xb0\x18\xe9\x0d\x77\x78\x6d\xc3\x5f\x20\xfc\xae\xb5\x13\x18\x4f\x07\x70\x66\x80\x5a\x2b\x3d\x81\xc2\x47\x75\x4f\x0e\x8b\x25\x7c\x19\xb0\x2d\xec\xce\xc6\xb1\xab\x99\x8c\x11\xe8\x77\x81\x09\xbf\x65\x69\xb9\xc8\x6a\x22\x9a\x96\xe7\x39\xd3\x1e\x63\xb1\x84\x90\xad\x31\xf1\xd0\xde\x9e\xfe\x30\x7a\x13\x96\x6b\xcf\x70\xcf\x57\x70\xe7\x96\x6a\x21\x0d\x10\xa4\xe4\x68\x5f\x30\xb8\x4e\x13\x66\x10\x48\xe4\x4c\x56\xb5\x62\x04\xb8\x08\x0d\x10\xef\x48\x80\x02\x79\x2c\xf3\x23\xfb\x4c\x89\x67\x93\xf8\xf0\x1d\xe4\x9c\x3a\x6e\xee\x9f\x9a\xd1\xad\x85\xa5\x0b\xb9\xf3\xb7\x6d\xae\x50\xf2\xee\x56\x97\xbe\x2b\x1e\xe3\x09\xec\x39\x73\xea\xd5\xb9\xb9\x3c\x95\x44\x1e\x57\x34\x9e\x06\x22\xab\x24\x76\x9e\x1f\x79\x88\x68\x6f\xfb\x53\x8a\xd7\x0d\x76\xeb\x73\x8f\x88\x00\x79\xec\x80\xb1\x24\xc1\xf1\x5b\x06\xb3\xf6\xab\xff\x76\xb6\x04\x29\x92\x46\x8f\x76\x23\x7d\x92\xd4\x2d\xc9\xeb\x94\x60\x7f\x3a\xf6\xfc\xd5\x47\x64\xd0\xbd\xea\x47\x5f\xdd\xcd\x65\xb7\x73\x16\x9d\xc6\x99\x0d\x62\xf4\x13\xeb\xf3\x8d\x49\x36\xc8\x70\xa4\x34\xac\x66\x35\xcb\x65\xe3\x0d\x70\xfd\xef\xf0\xcb\xd2\x14\x25\x1f\xb7\xb6\x67\x7f\x85\xf5\xc9\x31\xb4\xf7\x0f\x77\x67\x4b\xa3\xd9\x68\x09\xf9\xcc\xb5\xf1\xa8\x44\x29\x8a\xda\xaa\x79\xab\x74\xe9\xad\x40\x1a\x87\xb8\x9a\xf3\xf3\xe6\x97\x50\x49\x3f\xb5\x25\xbe\x5d\x28\x99\xa3\x36\xa8\xab\xd1\xd2\x32\x14\x51\x69\x4b\x1f\xca\x8b\xa5\x34\xce\x84\x92\xdd\x02\x5c\xb0\x7a\x60\x1c\x72\x21\xe3\x52\x54\x32\x21\xad\x28\x8d\xe2\x3b\x41\xaf\xa4\xb9\x65\x3a\x7b\x66\x49\x82\xfa\x42\xc9\xcc\x4b\xa2\x34\xbd\x56\xfd\xa9\x09\xf4\x71\xaf\x8e\xbb\xa6\x5c\x85\x4e\x34\xdf\x47\xb4\x02\xfa\x26\x24\xd3\xdb\xf1\x5e\x27\x37\x4e\xb4\x1e\x1e\x19\x95\x0e\x52\x24\x1e\x70\xd4\xd6\x74\x77\x84\x44\x04\x4c\xf2\x76\xdd\x8f\x2c\xd9\xa0\xee\x71\x71\x74\x05\xde\xbf\x51\xc6\xb9\xe0\xd6\x52\xbf\xfb\x47\x92\xef\x47\x74\x27\x84\xcf\x40\xbd\x54\xbc\x75\x93\xa1\x63\x47\xe6\x40\x8d\xd6\xb6\x12\x3a\x53\x2f\xef\x64\x52\x14\xf0\xba\x61\x89\x88\xb6\x40\xfc\xaf\x44\x46\x80\xdc\xe1\x9b\x6b\xdd\x31\x09\x59\x66\x84\x8c\x21\x77\xc8\xee\x37\xe7\x60\x48\x32\x79\xbf\x20\x6b\xab\x32\xca\x1f\x9a\x43\x3a\xb8\xa8\x2d\x2f\xcf\xc2\xa4\xa7\xef\x70\x12\x47\x6a\x57\xb7\xe0\x67\xaa\xf7\xff\x0a\x77\x90\xe6\x53\x8e\xcb\x29\x62\x97\x6a\x74\x34\xfe\x08\xb9\x71\x1b\x35\x66\x56\x63\xf9\x3b\x00\x00\xff\xff\x03\x59\x9f\x2c\x41\x0c\x00\x00")

func templateTo_protoTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateTo_protoTmpl,
		"template/to_proto.tmpl",
	)
}

func templateTo_protoTmpl() (*asset, error) {
	bytes, err := templateTo_protoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/to_proto.tmpl", size: 3137, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
	"template/enums.tmpl":         templateEnumsTmpl,
	"template/method_delete.tmpl": templateMethod_deleteTmpl,
	"template/method_get.tmpl":    templateMethod_getTmpl,
	"template/method_mutate.tmpl": templateMethod_mutateTmpl,
	"template/service.tmpl":       templateServiceTmpl,
	"template/to_ent.tmpl":        templateTo_entTmpl,
	"template/to_proto.tmpl":      templateTo_protoTmpl,
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
	"template": &bintree{nil, map[string]*bintree{
		"enums.tmpl":         &bintree{templateEnumsTmpl, map[string]*bintree{}},
		"method_delete.tmpl": &bintree{templateMethod_deleteTmpl, map[string]*bintree{}},
		"method_get.tmpl":    &bintree{templateMethod_getTmpl, map[string]*bintree{}},
		"method_mutate.tmpl": &bintree{templateMethod_mutateTmpl, map[string]*bintree{}},
		"service.tmpl":       &bintree{templateServiceTmpl, map[string]*bintree{}},
		"to_ent.tmpl":        &bintree{templateTo_entTmpl, map[string]*bintree{}},
		"to_proto.tmpl":      &bintree{templateTo_protoTmpl, map[string]*bintree{}},
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
