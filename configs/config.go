// Code generated by go-bindata. DO NOT EDIT.
// sources:
// configs/config.yaml
// 只需要执行 b, _ := configs.Asset("configs/config.yaml") 就可以读取对应内容了
package configs

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

var _configsConfigYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x52\x5f\x6f\x9b\x4e\x10\x7c\xe7\x53\xac\xf4\x7b\x0e\xe6\xcf\xcf\x90\xdc\x53\x12\xc7\x51\x5d\xd9\x2d\x2a\x44\x7e\x8c\xd6\xb0\x3e\x63\x1d\xdc\xf9\xee\xb0\x71\x3f\x7d\x75\x50\x87\x54\xed\xdb\x6a\x76\x98\x9b\xd9\x21\x27\x7d\x26\xcd\x3c\x80\x1f\x5d\xbb\x91\x15\x31\xa8\x68\xd7\x71\x0f\xe0\x8b\xb5\x2a\x93\xda\x32\xb8\x0f\x82\xc0\x31\x08\xab\xa2\x6e\x48\x76\x96\x41\xe2\x90\xad\xae\x2d\x7d\x86\x9e\x94\x72\x5a\x2f\xb4\xc7\x4e\xd8\x0c\x39\xe5\xf5\x4f\x62\x10\x3a\xf6\x06\xfb\xcf\x88\x83\xd6\x92\xe7\x78\xa6\x0c\xed\x81\x81\xb1\x52\x23\xa7\x99\x90\xdc\x8c\xbb\xd7\x5a\xd0\x37\x6c\x88\x01\x2a\x35\x41\xcb\xde\x32\xf0\x85\x74\x2e\xdf\xd4\x5a\x62\xf5\xb7\x48\xa7\x84\xc4\xca\x0c\x0c\x37\x8d\x41\xdf\xb4\x60\x70\xb0\x56\xb1\xd9\x2c\x8c\x52\x3f\xf0\x03\x3f\x64\x2e\xdf\xcc\x58\xb4\x75\xf9\xc1\x5f\x35\xc8\x69\x83\xfd\xe8\x76\x0e\xff\xc1\xe6\xf9\xcf\xe5\x93\x10\xf2\xb2\xec\xad\x71\x89\x01\xee\xc0\x3f\xaa\xd3\x34\x12\xbf\xcd\xaa\xe5\xde\x0b\x5a\xdc\xa1\xa1\xe1\x3a\xcf\xc5\x55\x11\x83\xe6\x6a\x4e\xc2\x69\x1a\xd2\xed\x90\x52\x4b\x69\x3d\x80\x0c\x8d\xb9\x48\x5d\x31\x28\x9b\x63\xf8\xf0\x90\x86\x51\x94\xba\x46\xa4\xb1\x0c\x26\xdf\x71\x1c\x24\x83\xde\x78\xa3\x9d\x90\xfc\xdd\x90\x3e\xd7\x25\x79\x00\x05\xee\x04\x65\x9a\xf6\x75\xff\x7b\xe7\x01\x2c\x0e\xa8\x0d\x59\x06\x9d\xdd\xdf\x0f\x4f\x69\x33\x34\xc8\xa0\xd0\x1d\x8d\x2d\xad\x2a\x41\x0b\xd9\xb6\x64\xa6\xe6\xbe\x2b\x6a\x6f\x58\x1c\x78\x5f\xb7\x85\x8b\x92\x53\xa9\x9d\x1a\x55\xd5\xb5\x3c\x5e\x3d\x80\x95\x31\x1d\xe9\xf1\xc1\xbb\xc9\xcc\xb2\x57\xb5\x26\x06\x69\x14\x04\xde\xb2\xc1\x5a\xb0\x8f\x40\xa6\xb1\xca\x3f\x9d\xfc\x52\x36\xce\xd1\xf0\xcb\xfd\x9f\xcc\x07\xb1\x3c\x5f\x33\xb0\xa3\x33\x77\xa7\x31\x69\x9c\xc4\x69\x12\xa5\xd1\xfc\x71\xfa\x0c\x8d\xd9\xfe\xe3\x66\xaf\x5a\x36\x0c\xb6\xd8\xf2\xc5\x81\xda\x8c\x5a\x57\x4b\x21\x6f\x9d\x5d\x4a\xf5\xee\xd8\x8f\x61\x12\x3b\xa5\x5f\x01\x00\x00\xff\xff\x6c\x6c\xf1\xed\x10\x03\x00\x00")

func configsConfigYamlBytes() ([]byte, error) {
	return bindataRead(
		_configsConfigYaml,
		"configs/config.yaml",
	)
}

func configsConfigYaml() (*asset, error) {
	bytes, err := configsConfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "configs/config.yaml", size: 784, mode: os.FileMode(420), modTime: time.Unix(1636881090, 0)}
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
	"configs/config.yaml": configsConfigYaml,
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
	"configs": &bintree{nil, map[string]*bintree{
		"config.yaml": &bintree{configsConfigYaml, map[string]*bintree{}},
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

