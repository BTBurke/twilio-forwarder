// Code generated by go-bindata.
// sources:
// templates/voicemail.html
// templates/voicemail.mjml
// DO NOT EDIT!

package main

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

var _templatesVoicemailHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x59\x6d\x6f\xdb\x36\x10\xfe\x9e\x5f\x71\x51\x30\xa0\x1d\x42\x4b\x89\xfb\x16\x59\xf2\x9a\x26\xed\x56\xa0\xdb\x8a\x2e\x01\x36\x14\xc1\x40\x53\x67\x9b\x0d\x45\x6a\x24\x65\xcb\xf3\xf2\xdf\x07\x52\xb2\xad\xd8\x4e\x1a\xb4\x0b\xb0\x02\xf1\x07\x2b\x7c\xbb\x7b\xf8\xdc\xf1\x39\x99\x49\x76\x33\xc5\xec\xac\x40\x18\xdb\x5c\xf4\x77\x12\xf7\x80\x2a\x17\xd2\xa4\xc1\xd8\xda\x22\x0e\xc3\xe9\x74\xda\x99\x76\x3b\x4a\x8f\xc2\x83\xa3\xa3\xa3\xb0\x72\x73\x82\x7a\x52\x3c\x49\x83\x52\xcb\xd8\xb0\x31\xe6\xd4\x90\x9c\x33\xad\x8c\x1a\x5a\xc2\x54\x1e\x4f\x56\xf3\xd4\x6d\xf3\xd4\x70\xc8\x19\x36\x8f\xc0\xa1\x40\x9a\xf5\x77\x00\x12\xcb\xad\xc0\x7e\x12\xd6\x4f\xd7\xb3\x4b\xc8\x47\x3e\x84\xdd\xdc\xa8\x8b\xbe\x6b\x01\x21\x7e\x20\x47\x4b\xc1\x41\x26\xf8\x57\xc9\x27\x69\xf0\x3b\x39\x3f\x26\x27\x2a\x2f\xa8\xe5\x03\x81\x01\x30\x25\x2d\x4a\x9b\x06\x6f\x5f\xa7\x98\x8d\x9c\xa7\xda\x60\xb2\xfb\x11\x65\xc6\x87\x17\xce\xd4\xa6\xa1\x93\x7a\x1d\x39\x9b\x15\x6d\x2b\x16\x2b\x1b\x3a\x2e\x7a\xc0\xc6\x54\x1b\xb4\xe9\xf9\xd9\x1b\xf2\xc2\x6d\xc0\xd8\x99\x40\x70\xc4\x36\xf3\x98\x31\xde\xdd\x9e\x2a\xad\x50\xea\x12\x28\xcc\xa1\xa0\x59\xc6\xe5\x28\x86\xa8\x07\x57\x3b\x00\x9d\x0f\x48\xb3\x9f\xcd\xe8\x95\xca\x66\x30\x87\x29\xcf\xec\x38\x86\x83\x28\xfa\xae\x19\x7f\x5d\x59\xd4\x92\x8a\x13\x41\x8d\xb9\xc3\x8c\xef\x61\x0e\x82\x4b\x24\x63\xe4\xa3\xb1\x8d\x57\x13\x07\xb5\x8b\x9c\xea\x11\x97\x1e\x40\x1b\x0c\x99\xe2\xe0\x92\x5b\xe2\xa0\x13\xc3\xff\x46\x42\xb3\x4f\xa5\xb1\x0b\x57\x24\x37\x37\x8e\x39\xeb\x96\x0e\x04\xee\x83\xcd\x60\x0e\x03\xa5\x33\xd4\x84\x29\x21\x68\x61\x30\x5e\xfc\xd1\x83\xdc\x28\xe2\x67\x12\x61\x0a\xca\x30\x86\xa8\xb0\xed\x6e\xdd\xee\x76\x76\x79\x3e\x5a\x5a\xf4\x40\x9b\x7d\x01\x2d\xad\xea\x5d\xdb\x6a\x83\xc6\xd1\xcd\x25\xc6\x20\x95\xc4\x1e\x78\xd0\x19\x32\xa5\xa9\xe5\x4a\x2e\xba\xdd\x7e\xb8\xb4\xa8\x0b\x25\xfc\x00\xc9\x55\x86\x31\x0c\x38\x2b\x07\x9c\xd5\xce\x0b\x98\x43\xc6\x4d\x21\xe8\x2c\x86\x81\x50\xec\xb2\xb7\xe4\xef\xa0\x5b\x54\x75\x14\x93\xd0\xc7\xbe\xbf\xb3\x99\xa9\xb7\xe5\xc5\xcb\x1c\x33\x4e\x41\x49\x31\x03\xc3\x34\xa2\x04\x2a\x33\x78\x94\xd3\x8a\xd4\x51\x7e\xf2\x22\x2a\xaa\xc7\x30\xdf\x01\x00\x78\xe9\x10\x4f\x38\x4e\x0b\xa5\xed\x32\x11\xba\x87\x51\x51\xd5\x60\x01\x5e\xde\x36\xbc\x86\xf3\xfa\x01\x68\x80\x7b\xdc\x3b\x49\xe5\x74\x01\x20\x51\xf1\xaf\xfe\x7c\x9e\x2a\x56\xe6\x28\xed\x6f\x68\x2d\x97\x23\xd3\xf7\xde\x12\x15\x1f\x0b\xa1\xa6\xef\x7f\xf9\x31\x5c\xf6\xbc\xe7\x15\x0a\xf3\x1e\xf5\x5b\xc9\xc6\xfd\xa3\x67\x49\xb8\xde\xe7\x0c\x87\x37\x5b\x4e\x42\xef\x7d\x2b\x3c\x61\xd1\x41\x84\x83\x83\x8b\xdb\x88\xed\x34\x07\x8e\x8c\xb4\x2a\x0b\x32\xe4\x55\x43\x61\xcd\x89\xcb\x12\xd8\xe5\xb9\x23\x8a\x4a\xdb\x5b\xe7\xa6\xe5\x78\x6b\x44\xfd\x4e\x05\x97\x97\x30\xd6\x38\xac\x35\xd3\xc4\x61\x38\x54\xd2\x9a\xce\x48\xa9\x91\x40\x5a\x70\xd3\x61\x2a\x77\x98\x7e\x18\xd2\x9c\x8b\x59\x7a\x3e\x28\xa5\x2d\xe3\x6e\x14\xed\x3f\x89\xa2\xfd\xa7\x51\xb4\xff\x3c\x8a\x02\xd0\x28\xd2\xc0\x7b\x37\x63\x44\x1b\x6c\xdb\x12\xc0\x4d\xdb\xf5\x83\x3e\xf8\xf5\x8e\xa0\xd4\xe2\xd1\xd7\x41\x7a\xdc\xab\xad\x2e\x39\xd9\x94\xcc\x2f\xc8\x6a\x2e\xb7\x65\x75\x27\xff\xe4\x74\xa2\xcc\x25\x29\x50\x93\x83\x28\x5a\xe6\xae\x8b\x53\x2b\x4c\x1b\x49\x1c\xd6\x25\x23\x71\xba\xe6\x31\x66\x7c\xd2\x6f\x65\x32\xfc\x03\x6f\x5f\x5f\xf4\x1b\x82\x12\xaf\x2f\xa0\x95\xc0\x34\x28\x34\x1a\x94\xd6\x9f\xfb\xa0\x11\x97\x34\x88\x02\x60\x28\x44\xa3\x8a\xcb\xb6\x53\xa4\x45\xdb\x23\x4b\x83\x67\x2e\x6e\x54\xf0\x91\x4c\x03\x86\x4e\x45\x02\xf0\xb0\xd2\xa0\xc6\xfe\x2c\x72\xe7\x2e\xe8\x2f\xa3\x93\x58\xbd\x6a\xb8\x66\xb6\x58\xd0\x56\x30\xb7\xc8\xc5\xcc\x8b\xac\x6f\x39\x65\x6c\xcd\x20\xba\x14\x18\x63\x45\x99\x15\xb3\x95\xfd\x6b\xa1\xc9\xf8\x64\x61\xbb\xd1\xaa\xa8\xa8\x6a\xb9\x5c\x49\xcb\x02\xe0\x2d\xb4\x7c\x8e\x8b\xc6\xc7\x75\xbc\xab\xd0\xf5\x36\x18\x5a\xf1\xdc\x4f\xac\x8f\x9a\x63\xa5\x45\x85\x97\x6a\xbf\x28\x16\x38\xb4\xbd\x09\x6a\xcb\x19\x15\x4d\x9f\x55\x45\x2f\xe3\x1a\x99\xd7\x71\x61\xf5\x1a\x57\x8b\x72\xe6\x34\x0f\x9a\xed\xdd\x6b\x3a\xac\xe3\xdf\x82\x77\x6b\x3a\x6c\x84\x8b\xb9\xd2\x9d\x06\x9b\x67\x61\x43\xc9\x82\x5b\x7c\x2d\x2a\x15\x97\x3e\x63\xea\x7a\x75\x13\x61\xae\x7c\xf5\xd6\x09\x6f\x47\xef\x6b\x52\xa3\x39\x26\xce\xd0\x1d\xc2\x3e\x55\x3a\x23\x03\x8d\xf4\x32\xf6\xdf\xc4\x75\xdc\x10\xdb\x03\x17\xdb\xc3\xa7\x8e\xcd\x45\x7a\x39\xe4\xc1\x35\x1e\x97\x24\xb1\x52\x1b\xa5\x63\x9f\xfb\x4c\x09\xa5\xe3\xbd\xee\x49\xf7\xb4\x7b\x5a\x5b\xaf\x55\x30\xae\x55\x70\x1f\x7e\x42\x31\x41\x47\xea\x3e\x1c\x6b\x4e\xc5\x3e\x18\x2a\x0d\x31\xa8\xf9\xb0\xcd\xdc\x61\x51\xf5\xda\xa7\xf6\xf0\x70\x0b\x95\x41\xfb\xb8\xaf\x3e\x7f\xa8\x12\xc6\x74\x82\xa0\x91\x21\x9f\x60\x06\x14\x26\x8a\x33\xcc\x29\x17\x30\xd4\x2a\x87\xf9\xbc\xf3\x46\xab\xfc\xea\xaa\x03\x70\x36\x46\xb0\x9a\x4a\xc3\x34\x2f\x1c\xf7\xc0\x0d\x0c\x50\xa8\x69\xbc\x61\x3e\x09\xbd\x00\x86\x36\x73\x5f\xda\x7d\xd5\x6c\x87\x3e\x90\xfd\x66\xfc\xc6\x13\xd1\x5e\xe8\x17\x6c\xcb\xd6\x7b\xb5\xbe\xec\x7a\xd0\xf0\xbb\x6b\xf8\x80\xb2\x4b\x27\x0f\x32\x8b\xf7\x4e\x9f\x9f\x1e\x9f\xbe\xba\x5f\x59\xdf\xe6\xef\x41\xe9\x1f\x94\xfe\x3e\x94\xfe\x1b\x15\xf9\xf9\xbc\x73\xd6\x16\xed\x33\xac\xec\xd5\x95\x1b\x79\x10\xed\x07\xd1\xfe\x7f\xbf\x78\xd7\x6b\x1e\x04\xf9\x41\x90\xef\xf0\xea\xdd\x24\xd8\x7f\x90\xcc\xeb\x57\x86\x06\x0b\xaa\xa9\xc5\x2f\x48\xe6\xe6\xae\xd0\x5f\xf3\x35\x66\x35\xcd\x78\x69\x62\xc7\x7b\x53\x21\x86\xfe\xd3\x6b\xd7\x8e\xcf\xef\x13\x26\x4d\x3b\xe7\x59\x26\x30\x80\xc1\xc8\x9b\x4b\x83\xbd\x43\xd6\xc5\xa7\x0e\x12\x6d\x6e\x85\xe6\xf3\xce\x07\x64\x4a\x3b\x93\xe7\x1f\xde\x5d\x5d\x05\xd7\x0e\x5b\xeb\x42\xd2\x03\xdd\xb8\xb2\x6d\xbf\x66\xd5\xc6\xd7\xa0\x7f\x4d\x71\xeb\x2e\xe4\x6e\x5a\x7b\x94\x4a\xe7\x54\xd4\x39\xe9\x7f\x6e\x0c\x95\xce\x6b\x60\x2b\x25\xeb\x05\x60\xa9\x1e\xa1\x4d\x83\x3f\x07\x82\xca\xcb\x1b\xaa\xdf\x3b\x6e\x2c\x4a\xb0\x6a\xf5\xc3\x66\x4b\xd9\xa3\xb7\x96\xa5\x6f\xb7\x1c\xd6\x36\x76\x92\xb0\xbe\x96\x4a\xfc\xbf\x0b\xfa\xff\x06\x00\x00\xff\xff\x6c\xe1\x6d\x41\x75\x19\x00\x00")

func templatesVoicemailHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesVoicemailHtml,
		"templates/voicemail.html",
	)
}

func templatesVoicemailHtml() (*asset, error) {
	bytes, err := templatesVoicemailHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/voicemail.html", size: 6517, mode: os.FileMode(420), modTime: time.Unix(1488247115, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesVoicemailMjml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x93\x4f\x4f\x83\x40\x10\xc5\xef\x7e\x8a\xc9\x7a\x86\xfe\x21\xc6\xc4\x94\x26\x2a\xf1\xd4\x53\x53\x0f\x1e\x97\x65\x68\xb7\xee\xee\x34\xcb\x50\xab\x84\xef\x6e\xa0\x6d\xb0\x16\x13\x35\xe9\x9c\xe0\x0d\x33\x79\xef\x17\x66\x62\xd7\xd6\x4c\xaf\x00\x00\x26\x76\x1d\xa4\x94\xbd\xef\xdf\x8e\x8a\x22\xc7\x52\x3b\xf4\xa7\x72\x81\x8a\x35\x39\x60\xdc\x71\x20\x8d\x5e\xba\x58\x18\xcc\x59\x74\x9f\x75\x1b\x4c\x69\xdd\xa9\x7e\xec\x35\xe3\x90\x93\xe3\xa0\xd0\x1f\x18\x8b\xd1\x58\x80\x22\x43\x3e\x16\xd7\xd1\x63\x94\x44\x89\x38\x1f\x6c\xea\x85\x4a\x58\xc9\x2d\x82\x47\x85\x7a\x8b\x19\x48\xd8\x92\x56\x68\xa5\x36\x90\x7b\xb2\x50\x55\xe1\x93\x27\x5b\xd7\x21\xc0\x62\x85\xc0\x5e\xba\x42\x79\xbd\x69\x8d\xeb\x02\x52\x34\xf4\x76\x77\xee\x6b\x70\x30\xf6\x2d\xca\xa0\x27\x4b\x2b\x1e\x58\xf4\x03\x4a\xa5\x7a\x5d\x7a\x2a\x5d\x16\x1c\x83\x25\xb7\xc9\x7d\xf2\x20\x2e\xc1\x6e\x23\xb3\x4c\xbb\x65\x2c\x46\xc3\x5f\x82\xac\xaa\x70\xf1\x15\xcc\x02\x77\x5c\xd7\x4d\xe7\xb2\x60\xfe\x96\x35\x2d\x99\xfb\x61\x8e\x55\x84\x37\x43\x01\x2b\x8f\x79\x2c\xaa\x2a\x9c\xa3\x22\xdf\x40\x78\x9e\xcf\xea\xba\xa3\x90\xb7\xf5\x03\x85\x99\x2e\x18\x1d\x30\x75\x3f\x51\x7f\xfc\xbd\x91\xff\x03\x18\x9c\x9f\xd4\x7e\x6f\x7b\x78\xcd\xa3\x35\xd3\xcf\x00\x00\x00\xff\xff\xa1\xf9\xc0\xf9\x95\x03\x00\x00")

func templatesVoicemailMjmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesVoicemailMjml,
		"templates/voicemail.mjml",
	)
}

func templatesVoicemailMjml() (*asset, error) {
	bytes, err := templatesVoicemailMjmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/voicemail.mjml", size: 917, mode: os.FileMode(420), modTime: time.Unix(1488247105, 0)}
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
	"templates/voicemail.html": templatesVoicemailHtml,
	"templates/voicemail.mjml": templatesVoicemailMjml,
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
	"templates": &bintree{nil, map[string]*bintree{
		"voicemail.html": &bintree{templatesVoicemailHtml, map[string]*bintree{}},
		"voicemail.mjml": &bintree{templatesVoicemailMjml, map[string]*bintree{}},
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
