// Code generated by go-bindata.
// sources:
// assets/add.html
// assets/footer.html
// assets/header.html
// assets/index.html
// assets/view.html
// assets/views.go
// DO NOT EDIT!

package assets

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

var _assetsAddHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x92\x41\x8b\xd6\x40\x0c\x86\xef\xfe\x8a\x90\x7b\xed\x1f\x98\x0e\xb8\x28\x22\x78\x52\x3c\x4b\xfa\x25\xb5\x83\xed\x64\x48\xd3\x75\x97\xa1\xff\x5d\xda\xed\x7e\xbb\x2a\x0a\xb2\xb7\xd0\xa4\xcf\xfb\xc0\xbc\xb5\x02\xcb\x90\xb2\x00\x12\x33\xc2\xb6\xbd\x02\xa8\x15\x5c\xe6\x32\x91\x0b\xe0\x28\xc4\x62\xe7\x26\x94\x18\x08\x46\x93\xa1\xc3\x5a\xe1\xf5\x0d\x2d\xf2\xe5\xd3\x47\xd8\x36\x8c\x21\xc1\x65\xa2\x65\xe9\x70\x20\x18\xa8\x21\x33\xfd\xd1\x4c\x32\x38\x02\x59\xa2\x66\x4c\xcc\x92\x3b\x74\x5b\x05\x63\x68\x53\x84\xf7\x0a\x3d\x5d\xbe\x83\x2b\xa4\xcc\x72\x17\x5a\x8a\xa1\x2d\x71\x8f\x1a\xd4\x66\x98\xc5\x47\xe5\x0e\x8b\x2e\x8e\xfb\x67\x80\xc0\xe9\xf6\x9a\xa4\x36\x37\xdf\x4c\xd7\x72\x2e\x01\xc2\x44\xbd\x4c\x30\xa8\x75\x98\x18\xe3\x8d\x92\x31\x7c\x78\x1b\xda\x63\x71\x3d\x4b\xb9\xac\x0e\x7e\x5f\xa4\x43\x97\x3b\xc7\x5f\x98\x17\xcd\x6e\x3a\x21\x24\x3e\x30\x90\x69\x96\x87\xa9\x4c\x74\x91\x51\x27\x16\xeb\xf0\x91\xfe\xe8\xd6\x72\xba\xfd\x6f\xcd\xc5\xc9\xfc\x2b\x93\x0b\xc6\xcf\xfb\x0c\xfb\xfc\x2f\xe1\xe3\xf6\xef\xc2\xcf\x80\xa7\xf8\xf3\x88\x17\x98\x4a\xe6\x13\xf2\x2e\xf3\x4b\x2d\xaf\xb0\xd3\xf1\x09\xfe\x87\x61\xbf\xba\x6b\x3e\xb1\xcb\xda\xcf\xe9\xe9\xbd\x7a\xcf\xd0\x7b\x6e\x8a\xa5\x99\xec\x1e\xe3\x1b\xe6\xd0\x3e\xfc\x71\x14\xa9\xdd\xa3\xe3\xef\xbd\x1e\x54\xfd\xec\x75\xad\x20\x99\x61\xdb\x7e\x06\x00\x00\xff\xff\xc6\xfe\x50\xc4\x0e\x03\x00\x00")

func assetsAddHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsAddHtml,
		"assets/add.html",
	)
}

func assetsAddHtml() (*asset, error) {
	bytes, err := assetsAddHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/add.html", size: 782, mode: os.FileMode(420), modTime: time.Unix(1480513461, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsFooterHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\xcf\xc1\x6a\x84\x30\x10\x06\xe0\xbb\x4f\x31\xe4\xae\x69\x0b\x6d\xb7\x25\xe6\xd0\x87\xe8\xb5\x8c\x66\xb2\xa6\x8d\x89\x24\x23\x61\x11\xdf\xbd\xa8\xcb\x2e\xec\x5c\xe6\xf0\xf3\x7f\xf0\x2f\x0b\x18\xb2\x2e\x10\x08\x1b\x23\x53\x12\xb0\xae\x15\x00\x80\xca\x7d\x72\x13\x43\x4e\x7d\x2b\x06\xe6\x29\x7f\x4a\x39\x67\x6a\x6c\x0c\x8c\x85\x72\x1c\xa9\xe9\xe3\x28\x4f\xcf\x2f\x6f\xaf\xef\x27\xf3\xf1\xd4\xfc\x66\xa1\x95\x3c\x8a\xfa\x50\x26\xfd\x75\x01\x85\x30\x24\xb2\x77\xa7\x94\xd2\xe4\x12\x93\xe9\x88\x71\x57\x04\x30\xa6\x33\x71\x2b\x7e\x3a\x8f\xe1\x4f\xe8\x5b\xae\x24\x6a\xa8\xe1\xf0\x1e\xa5\xb3\xe3\x61\xee\x76\xe2\x56\x90\x9c\xc8\xfb\x58\x77\x73\x0a\x26\x96\x70\xb5\xef\xf4\x4e\x6d\xf7\xed\xa8\x40\x0c\xa0\x1c\xf4\x1e\x73\x6e\x85\x45\xb0\x58\x1f\x6a\x8d\x9e\xb7\x41\xee\xba\x45\xa2\x56\x72\xd2\x95\x92\x5d\x34\x97\xed\x0f\x3c\x7a\x5d\x2d\x0b\x50\x30\xb0\xae\xff\x01\x00\x00\xff\xff\x73\x31\x31\x8b\x4e\x01\x00\x00")

func assetsFooterHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsFooterHtml,
		"assets/footer.html",
	)
}

func assetsFooterHtml() (*asset, error) {
	bytes, err := assetsFooterHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/footer.html", size: 334, mode: os.FileMode(420), modTime: time.Unix(1480463928, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsHeaderHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x44\x90\x4d\xcf\xd2\x40\x14\x85\xf7\xfd\x15\xd7\xd9\x9a\xb7\x13\x2d\x11\x62\xda\x26\xa0\x28\xca\x02\x02\x48\xc2\x72\x3a\x73\xe9\x8c\xcc\x07\x99\x3b\x15\x6b\xd3\xff\x6e\x80\xe8\xbb\xba\xb9\xcf\x39\x9b\xf3\x0c\x03\x28\x3c\x1b\x8f\xc0\x34\x0a\x85\x91\xc1\x38\x66\xe5\x9b\xcf\x9b\x4f\x87\xd3\x76\x09\x3a\x39\x5b\x67\xe5\xbf\x83\x42\xd5\x19\x40\x99\x4c\xb2\x58\x1f\x22\x5a\x1b\x60\xd1\x45\xaf\xc2\xcd\x97\xfc\x89\xef\x05\x87\x49\x80\xd4\x22\x12\xa6\x8a\x75\xe9\xfc\x32\x63\xc0\x1f\x91\x35\xfe\x02\x11\x6d\xc5\x28\xf5\x16\x49\x23\x26\x06\x3a\xe2\xb9\x62\x3a\xa5\x2b\x7d\xe4\xdc\x89\xdf\x52\xf9\xbc\x09\x21\x51\x8a\xe2\x7a\x7f\x64\x70\xfc\x3f\xe0\x45\x5e\xe4\x53\x2e\x89\x5e\x59\xee\x8c\xcf\x25\x11\xcb\x00\x00\x8c\x4f\xd8\x46\x93\xfa\x8a\x91\x16\xc5\x6c\xf2\xb2\x38\x9e\x8c\xd9\x7f\xfb\x82\xeb\x77\xea\xab\xfb\xbe\x9b\x5f\x7a\xd9\xad\xe6\xab\x5d\x5b\xbc\xdf\xb8\x1f\xf2\x76\x9b\x06\x5f\xec\x4e\xaa\x9d\x1c\xc5\xdb\xad\xdb\x1f\xe8\x0f\x5f\x7f\x98\xfd\x6a\xd4\xf2\xa7\x9e\x74\x0c\x64\x0c\x44\x21\x9a\xd6\xf8\x8a\x09\x1f\x7c\xef\x42\x47\xac\xce\x4a\xfe\x54\x53\x36\x41\xf5\x20\xad\x20\xaa\x98\x0c\x3e\x09\xe3\x31\xb2\xc7\xee\x26\xde\x05\x0c\x03\xa0\x57\x30\x8e\x7f\x03\x00\x00\xff\xff\x59\xe7\xf9\x64\x7a\x01\x00\x00")

func assetsHeaderHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsHeaderHtml,
		"assets/header.html",
	)
}

func assetsHeaderHtml() (*asset, error) {
	bytes, err := assetsHeaderHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/header.html", size: 378, mode: os.FileMode(420), modTime: time.Unix(1480463928, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x53\xdb\x8a\xdc\x30\x0c\x7d\xef\x57\x08\xd3\xd7\x8c\xa7\xfb\xd0\x27\x8f\x61\x6f\x85\x42\x29\xa5\x97\x0f\xd0\x8c\x94\xc6\xe0\xd8\xc1\xd6\xb6\x85\xe0\x7f\x2f\x4e\xe6\xb6\xb3\xd9\xa1\x7d\x09\x96\x8f\x8e\x14\x9d\x63\x8d\x23\x10\xb7\x2e\x30\x28\x17\x88\xff\x28\x28\xe5\x0d\xc0\x38\x82\x70\x3f\x78\x14\x06\xd5\x31\x12\xa7\x3d\x62\x06\x6b\x10\xba\xc4\xed\x46\x8d\x23\xac\xee\x30\xf3\x8f\xaf\x9f\xa0\x14\x24\x52\xd6\x38\xd8\x79\xcc\x79\xa3\x5a\x84\x16\x9b\xc1\x3f\x65\x65\x8d\x76\x16\x6e\x89\x40\x12\x7b\x1f\x61\x1b\x31\x91\xd1\x68\x8d\x1e\x6c\x2d\x2a\xb8\xf5\x7c\x60\xce\xc1\xf4\x6d\xb6\x31\x11\x27\xa6\x7d\x98\x25\xb9\x81\x49\x55\x52\xa5\xa5\xf9\x50\x8f\x9d\x35\x5a\xba\xf3\xf8\xae\x76\x81\x80\x3d\x5f\x22\xdf\x04\x93\x00\xa1\xbc\x40\x1e\x03\x2d\xde\xdf\x63\xa2\x6c\xb6\x09\xb4\xbd\x8f\xfd\xe0\x59\x98\x40\xc3\xf7\x28\xe8\x2f\x73\xbf\x44\x17\xe4\x5f\x93\x6f\x77\xe2\x62\xc8\xa7\x6b\xa3\x0f\x53\x8d\x23\x24\x0c\x3f\x19\x56\xd3\x24\x79\xd2\xff\x62\xee\x1a\xd0\x33\x47\xde\x9e\x59\xf2\xcb\xf1\x6f\x5d\x5d\xfa\xf8\x00\xa5\xbc\x74\xc7\xbb\xc0\xcd\xae\xc3\x24\x0a\x30\x39\x6c\x3a\x47\xc4\x61\xa3\x24\x3d\xf1\x6c\xdb\xec\x92\xd0\xf3\x86\xb5\xe6\x67\xec\x19\x4a\x59\x06\x1f\x50\x78\x52\x79\xf5\x21\xa6\x1e\x05\xd4\xfa\xa6\x59\xbf\x6b\x6e\xd6\xeb\xf7\xea\x2a\xeb\x31\xd0\x7f\x71\x26\x67\x4e\x32\x97\x02\x1a\x8e\xf7\xaf\xb2\x66\x8f\x16\x68\x33\xb0\xc8\x3b\x06\x00\xaf\x09\x4e\x5c\xcb\x5d\x93\x5c\x12\xe6\xee\xaa\xda\xa7\x9e\x67\x7f\x30\xbd\x0a\x80\xc3\xbb\xe0\x40\xfb\x75\xd4\xd3\x62\xd8\xcb\x9d\x6d\x63\x94\xfd\xce\x1e\xd3\xff\x06\x00\x00\xff\xff\x1f\x4b\x74\x81\xec\x03\x00\x00")

func assetsIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsIndexHtml,
		"assets/index.html",
	)
}

func assetsIndexHtml() (*asset, error) {
	bytes, err := assetsIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/index.html", size: 1004, mode: os.FileMode(420), modTime: time.Unix(1480513578, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsViewHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x55\x4d\x6f\xe3\x36\x10\xbd\xef\xaf\x18\xe8\x12\x07\x2b\x4b\x8e\xdb\xb4\x88\x63\x19\xe8\x26\xdb\xa2\x40\x0f\x41\x3f\x4e\x41\x0e\x63\x71\x64\x31\xa5\x49\x81\x1c\xd9\x0e\x04\xfd\xf7\x82\xb4\x25\x4b\xd9\x64\xb1\x05\xf6\x60\x80\x9a\x79\x7c\xf3\xe6\x71\x4c\x36\x0d\x08\x2a\xa4\x26\x88\x76\x92\xf6\x11\xb4\xed\x07\x80\xa6\x01\xa6\x6d\xa5\x90\x09\xa2\x92\x50\x90\x3d\x65\x96\xd5\x6a\x89\x50\x5a\x2a\xb2\xa8\x69\x20\xf9\x84\x8e\xfe\xf9\xf3\x0f\x68\xdb\x68\xb5\x94\x90\x2b\x74\x2e\x8b\x0a\x84\x02\xa7\x68\xad\xd9\x4f\x15\x15\x1c\x01\x5a\x89\xd3\x52\x0a\x41\x3a\x8b\xd8\xd6\x14\xad\x96\xa9\x5c\xc1\x6f\x06\xd6\x98\xff\x0b\x6c\x40\x6a\x41\x87\x65\x8a\xab\x65\x5a\xad\x7c\x29\x21\x77\xb0\x97\x82\xcb\x2c\x9a\xcf\x66\xd5\x21\x82\x92\xe4\xa6\xe4\xee\xd3\x83\x00\x96\x39\xea\x1d\x3a\x90\x22\x8b\xd6\xb5\xd5\xc2\xec\xf5\x5d\x89\x96\x7d\x85\x63\x2e\xb0\xa5\x42\xee\xc2\xc2\xe5\x56\x56\x0c\xce\xe6\x59\x54\x32\x57\x6e\x91\xa6\xb9\xd0\xcf\x2e\xc9\x95\xa9\x45\xa1\xd0\x52\x92\x9b\x6d\x8a\xcf\x78\x48\x95\x5c\xbb\x34\xf0\x25\xcf\x2e\x9d\x27\x3f\x26\xb3\xd3\xe7\xba\xd6\x42\x51\xb2\x95\x3a\x79\x76\xbe\xd8\x91\x78\x50\xe3\x28\x70\x87\x16\x72\x3e\x40\x06\xc2\xe4\xf5\x96\x34\x27\x1b\xe2\xcf\x8a\xfc\xf2\xd3\xcb\xef\x62\xf2\x4a\xf7\xe5\x6d\xbf\x8f\x0d\xa3\x7a\x30\x52\xb3\x83\x0c\x82\xe1\x06\xad\x48\x4e\xa1\xb6\x3d\x43\xab\x10\xfa\xab\x22\xcd\x0f\x64\xef\xf1\x05\x32\x78\x0c\x59\x7f\x9e\x53\xb0\xa8\x37\xd4\xed\xbf\x43\x2b\x1e\xac\xd9\x58\x72\x0e\xa6\xe1\x64\x3b\x24\x9c\xc9\xe3\xc1\x76\xd2\xa2\x07\x3e\x9d\xab\x0a\x7c\x71\x6f\x17\xba\x47\xa6\x31\x77\x18\x98\x5f\x8d\xdd\x22\x83\x3f\xc2\x9f\xa6\xb3\xab\xe9\x6c\xee\x27\x2b\xfa\x86\x52\x23\x93\x20\x03\x4d\x7b\x08\xeb\x49\xce\x87\x18\x9a\x13\x03\xbf\x54\xb4\x80\x0b\x25\x35\x5d\x74\xac\x02\x19\x17\x3d\x02\x40\xe1\x9a\x94\x5b\x04\xf5\x71\x1f\xf5\x28\x47\xec\x16\x7d\x3b\x41\xd1\x60\x0d\x50\x48\xa5\x16\xe0\x07\x38\x1e\xc5\xfd\x10\x6f\xac\xa9\xb5\xb8\x33\xca\xd8\x05\x44\x76\xb3\xc6\xc9\xcf\xd7\xf1\xd5\xcd\x3c\xfc\x66\xc9\xfc\x32\x1a\x6f\x0a\x32\x16\x70\xf1\xd9\xb1\xdc\x22\x93\xb8\x18\xe7\x8f\xb2\xfd\xb0\x74\x80\x7b\x64\x9c\x5c\x0e\x51\x6d\xfc\xbd\xb4\xce\xaf\xaf\x63\xb8\xb9\x89\xe1\xea\x87\x79\x0c\x5f\x91\xfb\x4b\xce\x35\xaa\xf7\xb4\x1e\xb3\x6f\x08\xed\xd7\x4f\x1f\x46\xd2\xdb\xd3\xbc\x17\xb5\xce\x59\x1a\xfd\x46\xc3\x7d\x63\xd4\xc5\xfd\xcc\x9d\x66\x03\xa0\xea\xe6\x7d\xf8\x77\x49\xc3\xe9\x26\x8a\xf4\x86\xcb\x0e\xe9\x6f\xb5\xbf\x3d\x68\x0c\xee\xd2\x85\xb1\x30\x91\x90\xc1\xec\x16\x24\x2c\x47\x0c\x20\x3f\x7e\xbc\x1c\x38\xdc\x4b\x79\x94\x4f\x9e\xad\x67\x9e\x9e\xf4\xdc\xf6\xd0\x51\xd5\x77\x71\x9d\x41\x96\xb8\xb6\xfa\xcc\x7f\xcc\xb7\x5f\x78\x34\x34\xba\xd7\x85\x21\x38\x72\xe7\x7f\xf6\xfc\xc5\x3d\xf2\x9e\x01\xc7\x52\x6f\x74\xff\x9a\xe0\x51\x3e\x7d\x83\x17\x5f\xd9\xf5\xca\x99\x63\xe1\xb3\x2d\xc3\x9b\x77\xf4\x72\x15\xc6\xf0\xe9\xe5\x6a\x9a\x70\xad\xb4\xed\x7f\x01\x00\x00\xff\xff\x27\x42\x5f\xfd\xf1\x06\x00\x00")

func assetsViewHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsViewHtml,
		"assets/view.html",
	)
}

func assetsViewHtml() (*asset, error) {
	bytes, err := assetsViewHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/view.html", size: 1777, mode: os.FileMode(420), modTime: time.Unix(1480513473, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsViewsGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x99\xdb\x6f\x1b\xc7\x15\xc6\x9f\xb9\x7f\xc5\x46\x40\x02\xb2\x50\xa5\xbd\x5f\x04\xf8\xa1\xb9\x21\x7e\x68\x02\xa4\xe9\x53\xa7\x08\x66\x77\x67\x14\xa2\x12\xa9\x92\x54\x32\xb6\xe1\xff\xbd\xf8\xcd\x39\x2b\x52\xb2\xe2\xc4\x86\x81\x3e\xd0\x12\x57\x3b\x33\xe7\xfa\x7d\xdf\x19\x5f\x5e\xa6\x5f\x6d\x27\x97\x5e\xbb\x8d\xdb\xd9\x83\x9b\xd2\xe1\x55\x7a\xbd\xfd\xeb\xb0\xde\x4c\xf6\x60\x2f\x92\xcb\xcb\x74\xbf\xbd\xdf\x8d\x6e\x7f\xc5\xef\x76\xbf\x77\x87\xfd\xa5\x9d\xa6\x8b\x5f\x0e\xb7\x37\x27\x8f\xfc\x76\x7b\x70\xbb\xa7\x4f\x7f\x71\x76\x7a\xf7\xe9\x7a\x33\xb9\xf0\xf4\xe1\xaf\x6b\xf7\xdb\x73\xcf\xf6\x17\xd7\x5b\x1e\x7d\xfd\x43\xfa\xfd\x0f\x3f\xa5\xdf\x7c\xfd\xf2\xa7\xcf\x92\xe4\xce\x8e\xff\xb1\xd7\x4e\x5f\x4c\x92\xf5\xed\xdd\x76\x77\x48\x97\xc9\xe2\x6c\x78\x75\x70\xfb\xb3\x64\x71\x36\x6e\x6f\xef\x76\x6e\xbf\xbf\xbc\x7e\xbd\xbe\xe3\x81\xbf\x3d\xf0\x63\xbd\x95\x7f\x2f\xd7\xdb\xfb\xc3\xfa\x86\x2f\xdb\xb8\xe0\xce\x1e\x7e\xb9\xf4\xeb\x1b\xc7\x2f\x3c\xd8\x1f\x76\xeb\xcd\x75\xfc\xdb\x61\x7d\xeb\xce\x92\x55\x92\xf8\xfb\xcd\x98\x6a\x7c\x7e\x74\x76\x5a\xf2\x4b\xfa\xaf\x7f\x73\xec\x79\xba\xb1\xb7\x2e\x95\x65\xab\x74\x39\x3f\x75\xbb\xdd\x76\xb7\x4a\xdf\x24\x8b\xeb\xd7\xf1\x5b\x7a\xf5\x22\xc5\xaa\x8b\xef\xdd\x6f\x3f\xc6\x18\x2d\xa3\xd9\x7c\xff\xf2\xde\x7b\xb7\x8b\xdb\xae\x56\xc9\x62\xed\xe3\x82\xcf\x5e\xa4\x9b\xf5\x0d\x5b\x2c\x76\xee\x70\xbf\xdb\xf0\xf5\x3c\xf5\xb7\x87\x8b\x6f\xd8\xdd\x2f\xcf\xd8\x28\xfd\xfc\xbf\x57\xe9\xe7\xbf\x9e\x89\x25\xf1\xac\x55\xb2\x78\x9b\x24\x8b\x5f\xed\x2e\x1d\xee\x7d\x2a\xe7\xc8\x21\xc9\xe2\x67\x31\xe7\x45\xba\xde\x5e\x7c\xb5\xbd\x7b\xb5\xfc\x62\xb8\xf7\xe7\xe9\xf5\xeb\x55\xb2\x18\x6f\xbe\x99\x2d\xbd\xf8\xea\x66\xbb\x77\xcb\x55\xf2\xa9\xec\x61\x1b\xd9\xff\x77\x36\x72\xbb\x9d\xd8\xad\x0f\x87\x7b\x7f\xf1\x25\xa6\x2f\x57\xe7\xbc\x91\xbc\x4d\x92\xc3\xab\x3b\xad\x00\x42\x7e\x3f\x1e\xd8\x25\xfa\xa7\xf9\x48\x16\xeb\x8d\xdf\xa6\xe9\x76\x7f\xf1\xed\xfa\xc6\xbd\xdc\xf8\xed\xc3\x3a\x4d\xe1\xfc\xfc\x64\x87\x98\xc3\x34\xd5\x34\x26\x8b\xfd\xfa\x75\xfc\xbe\xde\x1c\x9a\x2a\x59\xdc\xd2\x30\xe9\xc3\xa6\x7f\xdf\x4e\x2e\x3e\xfc\x69\x7d\xeb\x52\xca\xe4\x82\xdf\x38\x27\x96\xca\xd2\xaf\x9f\x9e\xb5\x4a\xbf\xb7\xb7\x6e\xb9\xd2\x13\x38\x53\xbd\xf4\xeb\x0b\x4e\x4f\xde\xbe\x67\xed\x3f\xd6\xaf\x59\x1b\xad\x79\xbc\x14\x43\xdf\xbb\x14\x5b\x97\xab\x53\xcb\x1f\x6f\x80\x6b\x7f\xb4\x01\xce\x2d\x57\x47\x47\xdf\xd9\x41\xbd\xff\xfd\x4d\x5e\xee\xbf\x5e\xef\x96\xab\x74\xd8\x6e\x6f\x4e\x57\xdb\x9b\xfd\x1f\x78\xfe\x6a\x2f\x8e\xbb\x9d\xb7\xa3\x7b\xf3\xf6\x64\xb5\x96\x04\x55\xfe\xb3\x80\xc2\xdf\xa6\xe9\xbb\xc3\xed\x4d\xfa\x42\x8b\x61\x79\x66\x42\xee\x4d\xe8\x06\x13\xb2\xce\x84\x2c\xd3\x4f\x6f\x42\xe3\x4c\xe8\xf4\x99\xf7\x26\xd8\xca\x84\xbe\x30\xa1\xca\xe5\xfd\xa9\x31\xa1\xe2\xdd\xd1\x84\xae\x31\xc1\x79\x13\x3c\x6b\xac\x09\x7d\x66\x42\x3b\x98\xe0\x26\xd9\xbf\x67\x1f\x67\xc2\xd0\x99\x50\xf0\x29\x4c\x68\x3b\x13\xea\xc2\x84\x72\x34\xa1\x1a\x4c\xf0\xd6\x84\xa2\x36\x61\xa8\x4d\xe8\x4a\x59\xdb\x54\x26\x54\x9d\x09\x53\x69\x42\x5b\x9b\xd0\xb7\x26\xd8\x5c\xec\xa9\x27\x13\x26\x2b\xef\xb5\xec\x3d\x98\x50\x58\x13\x32\x6b\xc2\x50\x98\x30\xb4\x26\x4c\x99\xd8\x3d\x62\xdb\x60\xc2\x98\x99\x30\x8c\x72\x46\x56\x98\x30\x0e\x62\x2b\xef\xe3\x67\x8e\x3d\xa5\x09\x23\xeb\x1b\x13\x86\x49\xde\xb3\x9d\x09\x79\x6d\x42\x3d\x9a\xe0\x1a\x13\xca\xc2\x84\x3e\x37\x21\xc3\xc7\x4c\x7c\x1a\x2b\x13\x9a\xc2\x04\xd7\x9a\x50\x34\x26\xf4\x95\x09\x79\x27\x71\xad\xf8\x5e\x8a\xed\x63\x69\x42\x8d\xdd\xf8\x81\x1d\x93\x09\xc5\x64\x82\x67\x6d\x2d\xbe\x56\xd8\xde\x99\x50\x36\x26\x74\xa3\x09\x05\xeb\x72\x13\x9a\xda\x04\x5b\xc8\x4f\xd7\x9b\xd0\x72\x76\x26\xe7\x60\x23\xef\x61\xbf\x27\x36\xb9\x09\xd5\x28\xb6\x96\x9d\xf8\x51\xf7\xba\xbe\x91\xbf\x8d\xa3\xe4\xb4\x1c\x4c\x68\x2b\x13\x6a\xea\xa0\x36\xa1\x29\x4d\x68\xc8\x4f\x69\x42\x57\x99\xe0\x5b\x89\x6b\x39\x49\xdc\x07\x27\x39\x2a\x06\x8d\xef\x68\x42\x5b\x98\x90\xb7\xe2\x1b\x35\x80\xaf\xf8\xd5\x52\x2f\xde\x84\xdc\x9a\x30\x55\x72\x36\xf5\x30\xd6\xe2\x27\x3e\x53\x1b\xd4\x54\xc1\x4f\x27\xb9\x6a\x5a\x13\xba\x4c\x72\x86\xaf\x9e\x75\x56\xce\x23\x2e\x25\x31\xf2\xe2\xc7\xd4\xca\xf9\xac\xcf\x72\xc9\x5f\x55\x49\xfe\x62\x0c\x32\x89\x4f\xac\xa1\x4e\xe2\xe5\xf0\x6d\x52\xff\x73\x13\xda\x51\x6a\x32\x1f\xa4\xae\x88\x01\xb6\xe3\x33\xf5\x39\x10\xbb\x51\x6c\xa5\xde\x6a\xed\x89\xaa\x94\xba\x24\x86\x63\x2b\xcf\xe3\x19\xad\x09\xe3\x24\x7d\x54\x5a\xc9\x4d\x81\xed\x4e\xec\xa1\xee\x9a\xde\x84\x9e\x1c\x53\xd7\xbd\xd8\x4a\x1e\xa8\xad\x9a\xf7\x5b\x13\x72\x7a\x8b\xfe\xc8\xe4\x39\x3d\xe6\x3a\xe9\x3f\x7c\x1e\xac\xe4\xbb\xf1\x72\x1e\x71\x1d\x89\xd7\x28\xf9\xa1\xee\xa8\xd3\xb1\x39\xf6\x01\x3d\xcb\xcf\xf8\x8e\x97\x5a\x24\x1e\x5e\xfb\x98\xf8\xf1\x2e\xf1\xb7\xad\x09\xbe\x33\xc1\x97\x82\x07\xf8\x85\x7f\x95\x95\x5e\xc8\x4b\xa9\x5d\x62\xcf\x3b\xc4\x8a\xbc\x5b\x6f\xc2\x90\x49\x3d\xfb\x5c\x30\x25\xe2\x43\x6b\x42\x93\x9b\x30\x78\xb1\xbd\x19\x24\x26\x43\x2e\xfd\x48\xec\x39\x9b\x9c\xf3\x21\x87\xad\x3e\x9b\x14\x57\x1a\xc5\x19\x4b\x3e\x79\x6f\x34\x21\x77\xe2\x03\xf9\xc3\x2e\xde\x65\x5f\x7c\x24\x8f\x79\x25\x31\x9e\xb0\xad\x94\x77\xf1\x95\xfd\x59\x5b\x57\x12\x47\xf6\xa2\x4e\xec\x24\xbd\xc5\xfe\xd8\x3b\x0d\x92\xf7\xac\x39\xe2\x24\x38\xc4\x87\xd8\xe2\x5b\x9d\x09\x0e\x50\x23\x59\x39\xbf\x77\x36\x4b\xa4\x47\x20\xac\xcc\xfd\x9c\x22\x9a\xf9\xfd\x44\x51\x25\x8b\xc5\x63\x0c\x3f\x4f\x16\x8b\xb3\x27\xda\xf3\xec\x3c\x59\xac\x1e\x58\xf6\xd1\xeb\x1c\xf4\x97\xf8\xe4\xf4\xa0\x28\x0c\x1e\xd4\xd7\x73\xe6\xfd\x91\xb8\x79\xd0\x24\x51\x55\x5c\xbd\x78\xca\x50\x6f\xe0\xee\xab\xf4\x5d\x4b\x53\x98\xf9\x2a\x6d\xbb\xe2\x3c\x85\x63\xaf\x4e\x29\x78\x59\x15\xd9\x2a\x3e\x87\x39\xaf\x84\x59\xff\xb9\x59\x87\x65\x5e\x75\x59\x9d\x97\x55\x93\x9f\xa7\xd9\xea\x6d\xb2\xb0\x1c\xfa\x45\xdc\xfe\x4d\x74\xe7\x2a\x55\xaf\xb0\xe8\x2a\xfe\xfb\xf6\x21\xa4\xf6\xfc\x19\x56\xfc\x36\xea\xf4\x8f\x26\x46\x48\x81\xe2\x8c\x00\x6d\x05\x30\x69\xf0\x3c\x93\x62\x81\x20\x20\xa7\xca\x0b\xd0\x38\x0a\xd1\x49\xf3\xd3\x98\xcd\x24\x64\x05\xf9\xcd\x85\x4b\x93\xd0\xe4\x91\x0c\x47\x01\x4d\x48\xca\x36\x02\x5a\x5d\x2f\x60\x52\x94\x52\x9c\x79\x2e\x60\x48\x31\x03\x74\xa3\x82\x29\xc5\x3c\x93\x16\x00\x42\x83\xb6\x5e\x7e\xa7\xf1\x39\x1f\x30\x1c\xb4\x81\xa3\xcd\x0a\x82\xec\x0d\x01\x40\x8a\x34\x32\x36\x43\x82\xf8\x1d\x81\xd9\x9a\xd0\x4e\x0a\xfc\xa5\x00\x61\xe5\xe4\x19\xc0\x13\x7d\x6f\x4c\x28\x7a\x39\x13\xc0\x28\x7b\x01\x76\xe2\xd4\x8c\x2a\x1e\xf8\xd4\xba\xcf\x28\x4d\x1a\x9b\xbf\x14\x62\x25\x6e\xc4\x17\x7b\x01\x39\x80\x85\xc6\x05\x1c\xbd\x82\x0b\xc4\xe2\x35\x4e\x10\x4f\x5f\x8b\xb0\x00\x24\x10\x16\x34\x26\xa4\x4c\x83\xd3\xdc\x90\x04\x67\x92\x27\xe2\x88\xff\x6d\x2b\x60\x07\x71\x4f\x85\xe4\x09\xdf\x01\x50\xec\x21\xe7\x00\x49\x0d\x19\x6a\x8e\xc9\x47\xa9\x44\x51\x28\x40\x00\xf6\x9d\x82\x2a\xb6\x93\x47\x88\x95\xf8\x71\x16\xbe\x93\xa3\x28\x00\x32\x01\x1f\x80\x6c\x50\x50\x22\x9f\x8e\x38\x12\x3f\x48\xbb\x11\xb2\xe8\xa9\xb3\x4e\x80\xbb\xee\xc4\xde\xb6\x14\x62\x8e\x82\xa3\x11\x41\x40\xcd\x44\x31\x56\x49\x3e\xa8\x2f\x08\xdc\x69\x6d\xcc\xa2\xcd\x66\x12\x6f\xde\xc3\x06\xf6\x8a\xb5\x0d\x31\xd7\x92\x73\xce\xa1\x0f\x62\x5d\x43\x96\x4e\x6a\xb5\x52\x1b\x01\xf0\x4a\x45\x49\xad\x84\x44\xec\x88\x3f\xc4\x8a\x58\x28\x2b\x21\x48\xce\xcf\xbc\xe4\x05\x40\xaf\x55\xf8\x50\x83\xe4\x87\x78\xce\x75\x46\x5f\x91\xa3\xa7\x40\x8b\x8d\xf4\x50\xa9\x22\xb4\x72\xa7\xef\x3d\x01\xda\x63\x5f\x7f\x2c\xd6\x1e\x77\x38\x85\xdb\x93\xb9\xfe\x39\xc4\x3d\x2e\xfa\xf3\xa0\xfb\x8e\xa9\x9f\x18\x77\x1f\x99\xac\xd0\x5b\x96\xd5\x07\x43\x6f\xd5\x94\x7d\xd1\x7d\x3a\xe8\xfd\x2e\x0e\xfa\x1f\x0d\xbd\x68\x3c\x5a\xa3\x9a\x54\x1f\x14\x52\xde\x70\x3d\x65\x4c\xd9\xd3\xf2\xc0\x16\x1a\x71\xea\x45\x43\x52\xc2\x40\x16\xe5\x07\x74\xa2\xdb\x27\x6d\x23\x3b\xeb\x79\x2b\x9a\x99\x0f\x33\x08\x9a\x88\xf2\xa6\xbd\xdb\x19\x12\x54\x43\x67\xad\xe8\x04\x34\x20\x67\xa1\x69\xa2\xee\xf1\x6a\x73\xa6\x70\x3e\x48\xcb\xa0\x25\xb1\x17\x38\xec\x07\x81\x31\x5a\x32\x53\xc8\xa3\x45\x80\x60\x60\x04\xfd\x4b\x0b\xd1\xe6\xf8\xd4\xeb\x2c\x80\xa6\x07\xee\xd0\xcf\xec\xc1\x7e\xbd\xd2\x0c\xf4\x11\x67\xa6\x46\xe2\x56\x2a\xf4\xd6\xaa\xab\x59\x83\xfe\x42\xc3\x57\xf8\x5e\x9b\x90\xf7\x12\xbb\x5e\x75\x28\x90\x38\xb7\x3f\xf3\x1a\xba\x1e\x9b\x9a\x4c\x67\x8b\x5a\xa0\x38\x6a\xc5\x49\x5a\x1c\x08\x86\x9a\x80\x1f\xe6\x08\xec\xa8\x7a\xf1\x1f\x88\x66\x9f\x7c\xa6\x31\x2b\x50\xec\x54\xab\x32\xa7\xa0\xb9\xf1\x37\xce\x8c\xf8\x58\x8b\xae\x22\x07\x31\x4f\xaa\x6f\x89\x11\xb3\x13\xda\x98\xfd\xa1\x2a\x72\x07\xe5\xe0\x2b\xb0\x49\x8c\x1b\xcd\x17\xf0\x1a\xe1\x79\x12\x58\x9f\xd4\x4e\x28\x93\x19\xd4\xf7\x32\x17\x12\xab\x42\xb5\x77\xa7\xfb\x00\x55\xd0\x17\xb3\x28\xf0\x1a\x6d\xf5\x42\xeb\xc4\xa0\x76\xb2\x27\x54\x19\xe7\x97\x5e\x9e\x31\x6b\x01\x91\x9d\xca\x03\xe8\x3d\x57\x8d\x3b\xd7\x34\x7f\x23\x5f\xd3\xc9\x8c\xe8\x75\x7e\xc2\x7f\xfc\xc3\x5f\xa8\x12\x8d\x3d\xe8\x2c\xd7\xeb\xbe\x93\x52\x2b\x94\xd0\x15\x32\x27\x40\x0b\x0e\x9f\x07\x79\xce\xac\x46\x5d\x30\x97\xe0\x0b\x6b\xd0\xb6\x50\x33\xef\x00\xc3\xd4\x05\xfe\x47\x09\xa1\x39\xe7\x3c\xea\xa7\x57\x2a\x67\x3d\x72\x02\xc8\xb5\xd8\x37\x09\x7d\x60\x1f\xda\x98\x3d\xf9\x99\x2b\xe5\x01\xf5\xf1\x4c\x9d\x87\xe2\x8c\xa2\x33\x20\x7b\x47\xaa\x75\x32\xef\x50\xfb\x8d\xfe\xa4\xa7\xc9\x01\x7d\xca\x8c\xc3\xec\x41\x3c\xc8\x0d\x7b\x33\xa7\x51\x5f\xd4\x2d\xf3\x23\x33\xd9\xe8\x44\x5e\x78\xd5\xf1\xe4\x82\x19\xb1\xd2\x19\x1b\x1d\x3f\xeb\xf9\x68\x7b\x27\x67\x31\x1f\x64\x2a\x35\xa0\x14\xe2\xcb\x0c\x47\xde\xa8\xa7\xc9\x49\x1d\xcf\xbd\x09\x36\x40\xff\x50\x15\x73\x2a\xbe\x1d\x75\xfe\x91\xa6\x62\xee\x5b\xa9\xab\x58\x37\xf6\x3d\x34\x75\xc4\xc0\x8f\xa5\xa9\xe3\x0e\xa7\x34\x75\x72\xd1\xfc\x1c\x4d\x1d\x17\xfd\x79\x9a\x7a\xc7\xd4\x4f\x4c\x53\x8f\x4c\x9e\x69\xaa\xed\xfe\xef\x34\xf5\x72\x33\xb9\xf0\xd1\x2c\xd5\x6b\x45\xd2\x1d\x54\xf2\xa4\x37\x12\x54\x55\x44\x24\xaf\xa2\x52\x6f\xba\x60\x2a\xba\x3b\x4e\xde\x83\x0c\x04\x88\x5d\xd8\x00\x71\x88\x08\x8e\x22\xad\x10\x71\x6d\xf5\x66\x2c\xd3\x09\x99\xb5\x9c\x39\xea\xf0\x01\xba\xc0\x18\x74\xd3\xd0\xc8\x5a\x9e\xb7\x2a\xaa\xe9\x68\xa7\x7f\x43\x84\x46\x84\xc8\xc5\xb6\x6c\xbe\xa1\xf0\x52\xf1\x30\x2b\x9d\xdf\xe8\xed\x09\xc8\xcb\xd0\x10\x07\x18\xbd\xd9\x28\xf4\x96\x80\x18\xe0\x7f\xa1\xb7\x7e\x4e\x6f\xbb\x22\xa3\x75\x82\x58\x20\x2a\x68\x1a\x6f\x5e\x2a\x41\x1a\xd8\x88\x6e\xcc\x15\x25\xca\x49\x90\x9c\xbf\xc1\xac\xf1\x3c\x2b\x93\x36\xa2\xb2\x6a\x8e\xb6\x80\x06\x4e\xed\x88\x37\x16\x3a\x30\xf4\x5e\x05\x6f\x25\x22\x1c\xd4\x27\x1e\xd8\x41\x7c\xda\x5e\x6f\xb1\x7a\xb9\x1d\xa8\x75\x90\xcb\x15\x01\x23\xe3\x7a\xb9\x79\xe3\xc3\x79\xa0\x23\x7f\x6f\x14\xf5\x2b\xbd\x41\x6c\x15\x05\xe3\x60\x36\x88\x2f\xa0\x39\xac\xd9\x74\xb2\x16\xc1\x0d\xaa\xa3\x40\x0a\x2b\xa8\x0b\x52\x81\x6e\x30\x2c\x88\xc6\x99\x88\xf3\x38\x20\x36\x1a\x13\xbd\x15\x82\x41\xc9\x0f\x48\x16\x6f\x4d\x61\xd9\x5c\xd8\xb0\xae\xc5\x3f\xea\x82\x0f\x28\x84\xc0\x26\x16\xe4\x0e\x56\xc5\xbf\x78\x73\x54\x0a\xc2\x5a\x1d\x38\xd8\x03\xb6\x8e\x37\x7d\xca\xfc\xb0\x02\x03\x4f\xaf\x48\x47\x6d\xc0\x16\x95\x0e\x2f\x99\x2a\x29\xd0\x92\x77\xa9\x8d\x78\xcb\x38\x8a\xdd\xd4\xd0\xa0\xb7\x71\xd8\x10\x6f\x59\x32\x41\xc9\x5e\x07\x92\xb1\x14\xb5\x16\xeb\xa4\x93\xba\x44\x65\x0d\xaa\xf2\xa8\xa7\xc8\xb2\x5e\xec\xa0\x07\x22\xd3\x15\x52\x1b\x0c\x46\xd4\x0a\xf5\x41\xfe\x22\xfb\x9c\xd4\x68\xa1\xec\x42\x0e\x51\x3a\x0c\x2b\x93\x3e\x27\xe7\x28\xc7\xa8\xd6\x0a\xa9\x9f\x78\x5b\x99\xc9\x70\x17\x6f\x63\x9d\xb0\x15\x36\x7b\xbd\x39\xf6\xda\x8b\xb5\xde\x8a\xc2\xdc\x5e\xfb\x3e\x0e\x74\x7a\x9b\x07\xa3\x0c\x7a\xfb\x8c\x52\x22\xd6\xe3\x6c\x93\x95\x5a\x8d\x83\xf1\x28\x6a\x85\xf3\xe2\x6d\x72\x2e\xbe\xc4\x5b\xc3\x49\xcf\x1b\x05\x47\x9c\xde\xdc\x4e\x73\x9d\x37\x3a\xf8\x35\xaa\xf0\x46\x55\x74\x56\x6c\xa6\x77\xa8\x49\xea\x60\xbe\x79\xc7\xef\xc8\x8e\xb9\xd4\x53\xae\x6c\x17\x19\xda\xca\x33\xd6\x92\x83\x46\x6f\xdd\xa6\xf9\x36\x52\x55\x06\xb5\xeb\xb4\xc6\x27\xc5\x94\x38\x0c\xab\x1d\x60\x4b\xbc\x2d\x6f\x55\xc5\x59\xbd\xe1\x6d\xf5\x86\x50\x6f\xa3\xe3\x0d\x71\x23\xfd\x41\xcc\x19\x04\xc9\x69\x39\x1c\x6f\xdc\x78\x9f\x5a\x62\xcf\x78\xe9\xa2\xea\x85\x7a\xe0\x3b\xbd\x87\xff\x4e\x15\x0a\x8a\x7e\x56\x19\xd8\x82\x6d\x71\xc8\x9c\x54\x4d\x6b\xdd\x0d\x7a\x8b\x4f\x1d\x82\xb7\xad\xc6\x80\xfa\x88\x93\x81\x5e\xa8\x44\xdc\x53\xec\xeb\xf5\xb6\x90\x3e\x25\x6e\xb3\xda\x7f\xee\x76\x90\xb8\x57\x7a\xc3\x4e\x7f\xb9\xf1\x3d\xb7\x83\x0f\x54\xf3\xb1\x62\xe0\x61\x83\x53\x2d\x70\xfc\xef\xe5\xe7\xa4\xc0\xc3\x92\x3f\xaf\x04\x9e\x9a\xf9\x89\x85\xc0\xa9\xbd\xaa\x03\xf2\x2c\xfb\xf0\x79\xb5\xce\xcb\xba\xfd\x30\x21\xf0\xbf\x00\x00\x00\xff\xff\xfa\xd6\x74\xbf\x00\x20\x00\x00")

func assetsViewsGoBytes() ([]byte, error) {
	return bindataRead(
		_assetsViewsGo,
		"assets/views.go",
	)
}

func assetsViewsGo() (*asset, error) {
	bytes, err := assetsViewsGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/views.go", size: 20480, mode: os.FileMode(420), modTime: time.Unix(1480513627, 0)}
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
	"assets/add.html": assetsAddHtml,
	"assets/footer.html": assetsFooterHtml,
	"assets/header.html": assetsHeaderHtml,
	"assets/index.html": assetsIndexHtml,
	"assets/view.html": assetsViewHtml,
	"assets/views.go": assetsViewsGo,
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
	"assets": &bintree{nil, map[string]*bintree{
		"add.html": &bintree{assetsAddHtml, map[string]*bintree{}},
		"footer.html": &bintree{assetsFooterHtml, map[string]*bintree{}},
		"header.html": &bintree{assetsHeaderHtml, map[string]*bintree{}},
		"index.html": &bintree{assetsIndexHtml, map[string]*bintree{}},
		"view.html": &bintree{assetsViewHtml, map[string]*bintree{}},
		"views.go": &bintree{assetsViewsGo, map[string]*bintree{}},
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

