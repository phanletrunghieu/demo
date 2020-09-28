// Code generated by go-bindata.
// sources:
// schema/model/auth_payload.graphql
// schema/model/link.graphql
// schema/model/user.graphql
// schema/model/vote.graphql
// schema/mutation.graphql
// schema/query.graphql
// schema/schema.graphql
// schema/subscription.graphql
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

var _schemaModelAuth_payloadGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x70\x2c\x2d\xc9\x08\x48\xac\xcc\xc9\x4f\x4c\x51\xa8\xe6\x52\x50\x28\xc9\xcf\x4e\xcd\xb3\x52\x08\x2e\x29\xca\xcc\x4b\x57\xe4\x52\x50\x28\x2d\x4e\x2d\xb2\x52\x08\x2d\x4e\x2d\x52\xe4\xaa\x05\x04\x00\x00\xff\xff\x4c\x33\xf3\x47\x33\x00\x00\x00")

func schemaModelAuth_payloadGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_schemaModelAuth_payloadGraphql,
		"schema/model/auth_payload.graphql",
	)
}

func schemaModelAuth_payloadGraphql() (*asset, error) {
	bytes, err := schemaModelAuth_payloadGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/model/auth_payload.graphql", size: 51, mode: os.FileMode(420), modTime: time.Unix(1548436725, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaModelLinkGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\xf0\xc9\xcc\xcb\x56\xa8\xe6\x52\x50\x50\x50\xc8\x4c\xb1\x52\xf0\x74\x51\x04\xb3\x53\x52\x8b\x93\x8b\x32\x0b\x4a\x32\xf3\xf3\xac\x14\x82\x4b\x8a\x32\xf3\xd2\x21\x12\xa5\x45\x39\xa8\x02\x05\xf9\xc5\x25\xa9\x29\x4e\x95\x56\x0a\xa1\xc5\xa9\x45\x5c\xb5\x80\x00\x00\x00\xff\xff\x48\xc5\x5a\x5e\x56\x00\x00\x00")

func schemaModelLinkGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_schemaModelLinkGraphql,
		"schema/model/link.graphql",
	)
}

func schemaModelLinkGraphql() (*asset, error) {
	bytes, err := schemaModelLinkGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/model/link.graphql", size: 86, mode: os.FileMode(420), modTime: time.Unix(1548436725, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaModelUserGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x08\x2d\x4e\x2d\x52\xa8\xe6\x52\x50\x50\x50\xc8\x4c\xb1\x52\xf0\x74\x51\x04\xb3\xf3\x12\x73\x53\xad\x14\x82\x4b\x8a\x32\xf3\xd2\x21\x22\xa9\xb9\x89\x99\x39\xa8\x42\x39\x99\x79\xd9\xc5\x56\x0a\xd1\x3e\x99\x79\xd9\x8a\xb1\x8a\x5c\xb5\x80\x00\x00\x00\xff\xff\xda\x7f\x66\xd8\x52\x00\x00\x00")

func schemaModelUserGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_schemaModelUserGraphql,
		"schema/model/user.graphql",
	)
}

func schemaModelUserGraphql() (*asset, error) {
	bytes, err := schemaModelUserGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/model/user.graphql", size: 82, mode: os.FileMode(420), modTime: time.Unix(1548436725, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaModelVoteGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x08\xcb\x2f\x49\x55\xa8\xe6\x52\x50\x50\x50\xc8\x4c\xb1\x52\xf0\x74\x51\x04\xb3\x73\x32\xf3\xb2\xad\x14\x7c\x32\xf3\xb2\x21\xfc\xd2\xe2\xd4\x22\x2b\x85\xd0\xe2\xd4\x22\x45\xae\x5a\x40\x00\x00\x00\xff\xff\x36\xeb\x70\x81\x39\x00\x00\x00")

func schemaModelVoteGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_schemaModelVoteGraphql,
		"schema/model/vote.graphql",
	)
}

func schemaModelVoteGraphql() (*asset, error) {
	bytes, err := schemaModelVoteGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/model/vote.graphql", size: 57, mode: os.FileMode(420), modTime: time.Unix(1548436725, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaMutationGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\xcc\x31\x0a\x02\x31\x10\x85\xe1\xde\x53\xbc\xed\x56\xf0\x04\xdb\xd9\x2b\x08\x9e\x60\x30\x21\x0e\x66\x27\x43\x66\x82\x2c\xe2\xdd\x45\x05\x09\x36\xdb\x3e\xde\xff\xf9\xa2\x11\xc7\xe6\xe4\x5c\x04\x8f\x0d\x00\x68\x31\x1f\x5b\xcd\x13\xce\x5e\x59\xd2\xb0\x43\x88\x76\xa9\xac\xef\xd3\x6f\xdd\x4e\x38\xb0\xdc\x86\x4f\x63\x9c\xa4\xe9\x18\x67\xe2\xbe\x53\x32\xbb\x97\x1a\xba\x49\x68\x8e\xbd\xb1\x6f\x7e\x3d\xd1\x92\x0b\x85\x2f\x95\x4b\x62\x59\x97\xfe\xd3\xe7\x2b\x00\x00\xff\xff\x16\x02\xd9\x6b\xcb\x00\x00\x00")

func schemaMutationGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_schemaMutationGraphql,
		"schema/mutation.graphql",
	)
}

func schemaMutationGraphql() (*asset, error) {
	bytes, err := schemaMutationGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/mutation.graphql", size: 203, mode: os.FileMode(420), modTime: time.Unix(1548436725, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaQueryGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x08\x2c\x4d\x2d\xaa\x54\xa8\xe6\x52\x50\x50\x50\xc8\xcc\x4b\xcb\xb7\x52\x08\x2e\x29\xca\xcc\x4b\x57\x04\x8b\xa4\xa5\xa6\xa6\x68\xa4\x65\xe6\x94\xa4\x16\xc1\x24\x74\x14\xf2\xd3\xd2\x8a\x53\x4b\xac\x14\x3c\xf3\x4a\x74\x14\x72\x32\x73\x33\x21\x6c\x4d\x2b\x85\x68\x9f\xcc\xbc\x6c\xc5\x58\x45\xae\x5a\x40\x00\x00\x00\xff\xff\x69\xd8\xe2\x96\x5c\x00\x00\x00")

func schemaQueryGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_schemaQueryGraphql,
		"schema/query.graphql",
	)
}

func schemaQueryGraphql() (*asset, error) {
	bytes, err := schemaQueryGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/query.graphql", size: 92, mode: os.FileMode(420), modTime: time.Unix(1548436725, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaSchemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x4e\xce\x48\xcd\x4d\x54\xa8\xe6\x52\x50\x50\x50\x28\x2c\x4d\x2d\xaa\xb4\x52\x08\x04\x51\x60\x81\xdc\xd2\x92\xc4\x92\xcc\xfc\x3c\x2b\x05\x5f\x28\x8b\xab\x16\x10\x00\x00\xff\xff\x8e\x43\x79\x00\x32\x00\x00\x00")

func schemaSchemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_schemaSchemaGraphql,
		"schema/schema.graphql",
	)
}

func schemaSchemaGraphql() (*asset, error) {
	bytes, err := schemaSchemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/schema.graphql", size: 50, mode: os.FileMode(420), modTime: time.Unix(1548436725, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaSubscriptionGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x08\x2e\x4d\x2a\x4e\x2e\xca\x2c\x28\xc9\xcc\xcf\x53\xa8\xe6\x52\x50\x50\x50\xc8\x4b\x2d\xf7\xc9\xcc\xcb\xb6\x52\x00\x91\x30\x91\xb0\xfc\x92\x54\x2b\x05\x10\xc9\x55\x0b\x08\x00\x00\xff\xff\xb6\x47\xd6\x1e\x39\x00\x00\x00")

func schemaSubscriptionGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_schemaSubscriptionGraphql,
		"schema/subscription.graphql",
	)
}

func schemaSubscriptionGraphql() (*asset, error) {
	bytes, err := schemaSubscriptionGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/subscription.graphql", size: 57, mode: os.FileMode(420), modTime: time.Unix(1548436725, 0)}
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
	"schema/model/auth_payload.graphql": schemaModelAuth_payloadGraphql,
	"schema/model/link.graphql": schemaModelLinkGraphql,
	"schema/model/user.graphql": schemaModelUserGraphql,
	"schema/model/vote.graphql": schemaModelVoteGraphql,
	"schema/mutation.graphql": schemaMutationGraphql,
	"schema/query.graphql": schemaQueryGraphql,
	"schema/schema.graphql": schemaSchemaGraphql,
	"schema/subscription.graphql": schemaSubscriptionGraphql,
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
	"schema": &bintree{nil, map[string]*bintree{
		"model": &bintree{nil, map[string]*bintree{
			"auth_payload.graphql": &bintree{schemaModelAuth_payloadGraphql, map[string]*bintree{}},
			"link.graphql": &bintree{schemaModelLinkGraphql, map[string]*bintree{}},
			"user.graphql": &bintree{schemaModelUserGraphql, map[string]*bintree{}},
			"vote.graphql": &bintree{schemaModelVoteGraphql, map[string]*bintree{}},
		}},
		"mutation.graphql": &bintree{schemaMutationGraphql, map[string]*bintree{}},
		"query.graphql": &bintree{schemaQueryGraphql, map[string]*bintree{}},
		"schema.graphql": &bintree{schemaSchemaGraphql, map[string]*bintree{}},
		"subscription.graphql": &bintree{schemaSubscriptionGraphql, map[string]*bintree{}},
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
