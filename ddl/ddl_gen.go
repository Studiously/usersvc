// Code generated by go-bindata.
// sources:
// postgres/1_init.sql
// tmpl/consent.html
// tmpl/error.html
// tmpl/login.html
// tmpl/register.html
// DO NOT EDIT!

package ddl

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

var _postgres1_initSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x91\x4f\x4f\x02\x31\x10\xc5\xef\xfd\x14\x2f\x9c\x76\xa3\x24\x46\x8f\x9c\x4a\x3b\xe0\xc6\xda\x62\x69\x8d\x9c\x48\xc3\x36\xa6\x09\xff\xb2\xbb\xca\xd7\x37\x5d\x24\xa2\x06\xe6\xd6\xe9\xbc\xce\xfb\xbd\x0e\x87\xb8\xd9\xa4\xf7\x26\x74\x11\x7e\xcf\x98\xb0\xc4\x1d\xc1\xf1\xb1\x22\x7c\xb4\xb1\x69\x51\x30\x20\xd5\xc8\xe5\x7d\x25\xf1\xbf\xb4\x71\xd0\x5e\x29\xcc\x6c\xf5\xcc\xed\x02\x4f\xb4\xb8\x65\xc0\x36\x6c\x22\x00\x47\x6f\xee\x8a\x2a\x4f\xc6\x4d\x48\x6b\x40\x3c\x72\xcb\x85\x23\x8b\x57\x6e\x17\x95\x9e\x16\x0f\xf7\x77\xe5\xaf\xc9\xb0\xea\xd2\x67\xc4\xd8\x18\x75\xcd\x89\xa4\x09\xf7\xca\xc1\x59\x4f\x59\xe6\x75\xf5\xe2\x09\x45\xbf\xa8\x64\xe5\xe8\x0f\xea\x7a\xb7\x0a\xeb\x65\xaa\xe3\xb6\x4b\x5d\x8a\x47\xea\xcc\xbf\xcc\xe8\x3d\xf7\x25\xca\x89\xb1\x54\x4d\x75\x3e\xa2\x18\x7c\x6b\x06\x25\x2c\x4d\xc8\x92\x16\x34\x3f\x05\x99\xea\x12\x46\x43\x92\x22\x47\x10\x7c\x2e\xb8\xa4\xdc\xf1\x33\xc9\x7f\x3a\xf9\xd1\x7d\x68\xdb\xc3\xae\xa9\x8f\xe1\x9d\x56\xf7\xbe\xcf\xbf\x4c\xee\x0e\x5b\xc6\xa4\x35\xb3\x0b\x1c\xa3\xf3\xcb\xde\xc6\xe8\x2b\x00\x00\xff\xff\xba\xb3\xab\xb4\xf3\x01\x00\x00")

func postgres1_initSqlBytes() ([]byte, error) {
	return bindataRead(
		_postgres1_initSql,
		"postgres/1_init.sql",
	)
}

func postgres1_initSql() (*asset, error) {
	bytes, err := postgres1_initSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "postgres/1_init.sql", size: 499, mode: os.FileMode(420), modTime: time.Unix(1495035516, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplConsentHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x56\xdd\x6e\xdb\x38\x13\xbd\xcf\x53\xcc\xa7\xa0\x40\x52\xd8\x92\xec\x38\x6d\xa0\xc8\xee\x17\xa4\x1b\x60\xaf\x1a\xb4\xd9\x8b\xbd\xa4\xa8\x91\x44\x84\xe2\x68\x49\xca\x3f\x31\xfc\xee\x0b\x4a\x72\x62\xcb\x76\xbb\xdd\xad\x9c\x44\x31\x67\x74\xc8\x39\x67\x66\x34\xf1\xff\x3e\x7f\xb9\x7f\xfa\xf3\xf1\x37\x28\x6c\x29\x67\x67\xb1\xbb\x81\x64\x2a\x9f\x7a\xa8\x3c\xb7\x80\x2c\x9d\x9d\x01\x00\xc4\x25\x5a\x06\xbc\x60\xda\xa0\x9d\x7a\x7f\x3c\x3d\x0c\x6f\xbc\xce\x64\x85\x95\x38\x7b\x94\xc8\x0c\x42\x2e\xe6\x08\x2b\xaa\x35\x70\x52\x06\x95\x8d\x83\xd6\xde\xfa\x1a\xbb\xda\xfe\xef\xae\xff\x8b\xb2\x22\x6d\xa1\xd6\xf2\xa2\xb0\xb6\x32\x51\x10\x64\xa4\xac\xf1\x73\xa2\x5c\x22\xab\x84\xf1\x39\x95\x01\x37\xe6\x53\xc6\x4a\x21\x57\xd3\xaf\x94\x90\xa5\xe8\x2a\x0c\x2f\x6f\xcf\x5e\x91\x12\x4a\x57\xb0\x7e\xfd\xda\x2c\x31\xfe\x9c\x6b\xaa\x55\x1a\xc1\xf9\xc7\x0f\xc9\xcd\xf5\xf8\x16\x82\xf7\x90\x31\x29\x9d\x0d\x32\xd2\x40\x32\x85\x44\xd3\xc2\xa0\x36\xf0\x3e\x38\x09\x30\x5c\x60\xf2\x2c\xec\x50\x0a\x85\x4c\x0f\x73\xcd\x52\x81\xca\x5e\x68\x91\x17\x76\xb0\xc5\x1f\xc0\xf9\xcd\xe7\xfb\xf1\x87\x87\xcb\xdb\xd3\x48\x25\xbd\xfc\x0a\x18\xfa\x05\x20\x7d\x04\x4b\x20\x31\xfb\x31\x86\xd3\x68\xd8\xea\x11\x81\xd7\x2a\xe2\x0d\xc0\x30\x65\x86\x06\xb5\xc8\xf6\xdd\x69\x8e\x3a\x93\xb4\x88\xa0\x10\x69\x8a\x6a\xdf\xba\xa5\xb6\x01\x35\x25\x91\x2d\x84\xca\x23\x60\xca\x0a\x26\x05\x33\x98\xf6\x1e\x70\x0c\x92\x59\x1e\x3c\x91\x6b\xb6\x32\x9c\x49\x7c\xf3\xdf\xbc\xa5\x88\xbf\xd0\xac\xaa\x50\xf7\xd2\x64\x21\x52\x5b\x44\x70\xf5\x21\xac\x96\xfb\xfb\x54\x2c\x4d\x1b\xdc\x9b\x77\x10\x42\xb8\x6f\x2c\x99\xce\x85\x8a\x80\xd5\x96\x8e\x6f\x57\x31\x85\xb2\xb7\x59\x45\x46\x58\x41\x2a\x02\x8d\x92\x59\x31\xc7\x7d\xd4\x97\xa1\x50\x29\x2e\x23\x18\x9d\x16\xed\xfc\xa1\xb9\xfa\xc7\x59\x0e\x4f\x47\xb2\x3d\x6c\xd8\x1c\x17\x46\xe1\xe9\x58\x27\xd7\x7d\x53\x42\xcb\xa1\x29\x58\xea\xf4\x0b\x21\x84\x71\x58\x2d\x21\x04\x9d\x27\xec\x22\x1c\x40\xf7\xe3\x8f\x2f\x07\x10\xc2\x75\xb5\x6c\x7e\x8f\xd8\x27\x97\x47\x79\x62\x3d\x8a\x38\x49\xd2\x11\x9c\x4f\xee\xef\x1e\xae\x7b\xa4\x5b\x5c\xda\x61\x8a\x9c\x34\x6b\x59\x54\xa4\x8e\x8b\x9d\x91\x2e\x41\xa8\xaa\xb6\x3d\xf8\x9f\x4d\xdd\xda\xba\x22\x89\xfa\xf2\xef\x09\x92\x8d\xdd\xe7\xf6\x58\x5a\x8d\xc2\xf0\x5d\x9f\x4e\x9d\xa2\x3e\x00\x7c\x93\x28\x84\xd1\x81\x06\xaf\xf2\x1c\x9a\x1a\x79\xc4\x4b\x63\x6d\xb1\x87\x09\x2d\x8f\xd4\xab\x11\x2f\x18\xc1\x68\xb2\x0b\xd0\x67\x2c\xa9\xad\x25\xf5\xdf\x28\x6b\x44\xb2\x9a\x29\xe3\x20\x23\xa8\x5d\xd1\x71\x66\xf0\x5f\x30\x7b\x2c\x09\x7e\x96\xd9\xef\x70\xb7\xcd\xb5\x63\x15\x75\x92\x33\xd8\xe9\x59\x4d\x98\x5d\x45\x33\x29\x21\xf4\xaf\x00\x0f\x42\xfd\x67\x5e\xbc\xd6\xc6\x9d\xa6\x22\xa1\x2c\xea\x1f\x89\x14\x15\xae\xad\x0e\xc0\xdf\x5d\x63\xdc\xf5\x94\xde\x62\x46\xbc\x36\xdf\x7b\x39\x4e\xae\xee\xc2\xc9\xc7\xd3\x1b\xfa\x25\x1a\xc3\x72\xec\x61\x6c\x53\x76\xd4\x16\x7c\x78\x9c\xdb\xe4\xca\x7d\x4e\x73\x3b\xde\xcf\x47\xf7\x37\x0e\xba\x39\x21\x0e\xda\x11\x24\x76\xaf\xf7\xd9\x59\x9c\x8a\x39\x70\xc9\x8c\x99\x7a\x5d\x33\xdf\x0e\x21\x3b\x96\xa6\xef\x7a\x6f\x43\x46\xdc\xc4\xe0\x98\x21\x35\xf5\x82\x6e\x2e\xf9\xc4\x0b\x26\x25\xaa\x1c\xa7\xeb\xb5\x7f\xbf\xfd\xb2\xd9\x78\x50\xa2\x2d\x28\x9d\x7a\x8f\x5f\xbe\x3d\xed\xe0\x34\x58\xc5\x08\x98\x14\xb9\x9a\x7a\xee\x45\xe9\xcd\xee\xb7\x53\x4e\x31\xea\x79\x56\xfb\x8e\x7b\x46\x77\xdd\x29\x60\x55\x25\x05\x6f\x5a\x19\x5c\x88\x34\x82\xf5\xda\xbf\xab\xdd\xab\x98\xe3\x66\x73\x09\x1a\xff\xaa\xd1\x58\x4c\xb7\xc3\x14\x58\x02\xc6\x39\x1a\x03\x1a\x0d\xd5\x9a\xa3\x01\x52\xed\xc0\x95\x60\xc1\x64\xe6\xc3\x53\x81\xbb\xc8\x07\x1b\x2f\x98\xb2\xe6\x60\xb5\x83\xb5\x14\xed\x87\x11\x54\xbd\xb8\x6a\x79\x18\xcb\x7a\xad\x99\xca\x11\xfc\xaf\xdb\x13\x7f\xe3\x54\xa1\xd9\x6c\x0e\x5c\x63\x29\x66\x71\xdb\x98\xed\xaa\xc2\xa9\xc7\x0b\xe4\xcf\x09\x2d\x3d\x50\xac\xc4\xa9\xb7\x5e\xfb\x9b\x8d\x37\x6b\x6e\x71\x20\xc5\xb1\xdd\x50\xa5\x3d\xe8\x38\xe8\x9f\x6b\xbd\xf6\xb9\xd1\xd9\x83\x40\x79\xe0\xdc\xb5\xb9\xf6\x00\xa6\x4e\x4a\x61\xbd\xd9\x5d\x55\x69\x9a\x63\x1c\xb4\xd6\x03\x39\xbb\xf4\xea\x6a\xc1\x9b\xfd\x9e\x39\xde\x21\x25\x50\x64\x41\x23\xa7\x5c\x89\x17\x04\x5b\x08\xb3\x2b\xc0\xa0\x71\xe3\x4c\x81\x61\x19\xca\x15\xe0\x52\xd8\xd6\xab\x62\x39\xfa\x7b\x14\xc7\x81\xcb\xd6\x2e\xa9\x83\x54\xcc\x5d\x11\x74\xb7\xae\x08\x82\x66\x5c\xff\x3b\x00\x00\xff\xff\xb5\x4c\x66\xb6\xbe\x0b\x00\x00")

func tmplConsentHtmlBytes() ([]byte, error) {
	return bindataRead(
		_tmplConsentHtml,
		"tmpl/consent.html",
	)
}

func tmplConsentHtml() (*asset, error) {
	bytes, err := tmplConsentHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/consent.html", size: 3006, mode: os.FileMode(420), modTime: time.Unix(1495322613, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplErrorHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x55\x4f\x6f\xdb\xb8\x13\xbd\xe7\x53\x4c\x14\xfc\xf0\x8b\x0b\xcb\x92\xd3\x24\x1b\x28\xb2\x76\x17\x69\x73\xdd\x3f\x4d\x0f\x3d\x8e\xc5\x91\xc8\x0d\xc5\x11\xc8\x91\x6d\x75\xb1\xdf\x7d\x21\xd9\x4e\x6a\x37\xc6\x5e\x4a\xd8\xa0\xc4\x37\x7c\x1c\xbe\x19\x3c\xe5\xe7\x1f\x7e\x7b\x78\xfa\xf2\xfb\x47\xd0\xd2\xd8\xe2\x2c\x1f\x26\xb0\xe8\xea\x45\x44\x2e\x1a\x16\x08\x55\x71\x06\x00\x90\x37\x24\x08\xa5\x46\x1f\x48\x16\xd1\xe7\xa7\xc7\xf8\x2e\xda\x41\x62\xc4\x52\xf1\xd1\x7b\xf6\x10\xc3\x27\xe9\x94\xe1\x2e\xd8\x3e\x4f\xb6\xc8\x36\x2a\x48\xbf\x7f\x1e\xc6\x2f\xa6\x69\xd9\x0b\x74\xde\x5e\x6a\x91\x36\x64\x49\x52\xb1\x93\x30\xab\x99\x6b\x4b\xd8\x9a\x30\x2b\xb9\x49\xca\x10\x7e\xae\xb0\x31\xb6\x5f\xfc\xc9\x4b\x16\xce\xde\xa7\xe9\xe4\xfe\xec\x85\x69\xc9\xaa\x87\xbf\x5f\x5e\xc7\x25\x2c\x9f\x6b\xcf\x9d\x53\x19\x5c\xfc\x74\xbb\xbc\xbb\xb9\xba\x87\xe4\x1d\x54\x68\xed\x80\x41\xc5\x1e\xd8\x2a\x58\x7a\x5e\x07\xf2\x01\xde\x25\x27\x09\xe2\x35\x2d\x9f\x8d\xc4\xd6\x38\x42\x1f\xd7\x1e\x95\x21\x27\x97\xde\xd4\x5a\xa6\x7b\xfe\x29\x5c\xdc\x7d\x78\xb8\xba\x7d\x9c\xdc\x9f\x66\x6a\xf8\xeb\x8f\xa0\xe1\x1f\x40\x72\xcc\x20\x0c\x96\xaa\xff\xe6\x18\x6a\x14\x6f\xeb\x91\x41\xb4\xad\x48\x34\x85\x80\x2e\xc4\x81\xbc\xa9\x0e\xc3\x79\x45\xbe\xb2\xbc\xce\x40\x1b\xa5\xc8\x1d\xa2\x7b\x69\x47\xd2\xd0\x30\x8b\x36\xae\xce\x00\x9d\x18\xb4\x06\x03\xa9\xa3\x0d\x83\x82\x1c\x36\xdf\xed\xa8\x3d\xf6\xa1\x44\x4b\xaf\xf1\xff\xbc\xb6\xc8\x6c\xed\xb1\x6d\xc9\x1f\xb5\xc9\xda\x28\xd1\x19\xbc\xbf\x4d\xdb\xcd\xe1\x39\x2d\x2a\x35\xf2\xde\xfd\x0f\x52\x48\x0f\xc1\x06\x7d\x6d\x5c\x06\xd8\x09\xbf\x7d\x5c\x8b\x8e\xec\xd1\x61\x2d\x07\x23\x86\x5d\x06\x9e\x2c\x8a\x59\xd1\x21\xeb\xd7\xd8\x38\x45\x9b\x0c\xe6\xa7\x8b\x76\xf1\x38\x8e\xe3\x74\x36\xf1\xe9\x9b\xec\x93\x4d\xc7\x74\x61\x9e\x9e\xbe\xeb\xf5\xcd\x31\xb4\xe4\x4d\x1c\x34\xaa\xa1\x7e\x29\xa4\x70\x95\xb6\x1b\x48\xc1\xd7\x4b\xbc\x4c\xa7\xb0\xfb\xcd\xae\x26\x53\x48\xe1\xa6\xdd\x8c\xff\x37\xf0\xeb\xc9\x9b\x3a\xe1\x91\x44\x25\x5b\xf6\x19\x5c\x5c\x3f\xfc\xfa\x78\x73\x24\xba\xd0\x46\x62\x45\x25\x7b\xdc\xaa\xe8\xd8\x7d\x5f\xec\x3c\xd9\xd9\x4c\x9e\x6c\xbd\x2b\x1f\xdc\xa1\x38\xcb\x95\x59\x41\x69\x31\x84\x45\xb4\xeb\x85\xbd\x7b\x7d\x83\x8c\x65\x8b\x5e\x3d\x2a\xd7\xf3\xe2\xb3\x06\xd6\xe7\x79\xa2\xe7\xdf\xac\xb7\xc5\x17\xee\xfe\xbf\xa2\xc1\x3f\x9e\xc9\x81\x68\x02\xe3\x84\xbc\x23\x01\x74\x0a\xd0\xd5\xe4\x49\x8d\x80\x35\x22\x96\xa0\x31\xaa\x26\x09\x20\x1a\x05\x7c\xe7\xe0\x49\x13\x7c\xea\x83\x50\x33\xcb\x93\xf6\x80\xfd\x8f\xce\x94\xcf\x53\xa8\x19\x72\x04\xed\xa9\x5a\x44\x7b\x8f\x0c\x2f\xf6\x3a\x2b\x39\x2a\x34\x37\x94\x27\x58\x00\xfb\xd7\xd8\xbf\x70\x85\xa1\xf4\xa6\x95\x4c\x9b\x20\xec\xfb\x59\xcd\x97\xf1\x7c\x12\x15\x43\x37\x8d\xf1\x4b\xaa\xd8\xd3\xa1\xc4\x9a\x7a\xa8\x49\xa0\xe7\xee\xfc\x25\xa5\x3c\x51\x66\x35\x08\xba\x9b\x76\x82\x26\xe3\x37\xe3\xdf\x00\x00\x00\xff\xff\xda\x5e\xab\x27\x43\x06\x00\x00")

func tmplErrorHtmlBytes() ([]byte, error) {
	return bindataRead(
		_tmplErrorHtml,
		"tmpl/error.html",
	)
}

func tmplErrorHtml() (*asset, error) {
	bytes, err := tmplErrorHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/error.html", size: 1603, mode: os.FileMode(420), modTime: time.Unix(1495322055, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplLoginHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x57\x7b\x77\xda\xb8\x12\xff\x3f\x9f\x62\xea\x9e\x3d\xa7\xdd\xe2\x07\x81\x3c\x4a\x6d\x7a\xd3\x3c\xb6\xb7\xdb\xdb\xa6\xaf\xcd\xb6\xff\x09\x6b\x6c\x0b\x64\xc9\x91\x64\x02\xe4\xe6\xbb\xdf\x23\x1b\x93\x60\xa0\xb7\xbb\x5d\x48\x80\x99\x91\x7f\x9a\xf9\xcd\x68\x3c\x0e\x1f\x9d\xbd\x3f\xfd\xfc\xf5\xf2\x1c\x32\x93\xf3\xe1\x5e\x68\xbf\x80\x13\x91\x46\x0e\x0a\xc7\x2a\x90\xd0\xe1\x1e\x00\x40\xf8\xc8\x75\xe1\x23\x5e\x97\x4c\x21\x85\x1c\x0d\x01\x43\x52\x0d\xae\xbb\xb4\x57\xaa\x38\x23\x4a\xa3\x89\x9c\xd2\x24\xee\xb1\xf3\xd0\x24\x48\x8e\x91\x33\x65\x78\x53\x48\x65\x1c\x88\xa5\x30\x28\x4c\xe4\xdc\x30\x6a\xb2\x88\xe2\x94\xc5\xe8\x56\x42\x07\x98\x60\x86\x11\xee\xea\x98\x70\x8c\xba\x1d\xd0\x99\x62\x62\xe2\x1a\xe9\x26\xcc\x44\x42\x36\xd0\x86\x19\x8e\xc3\xb7\x32\x65\x02\xfe\x0b\x9f\x4c\x49\x99\x2c\x35\x9f\x87\x7e\x6d\xa9\x57\x69\x33\x6f\x7e\xdb\xd7\xbf\x58\x6e\x7d\x80\x52\xf1\x27\x99\x31\x85\x1e\xf8\x7e\x22\x85\xd1\x5e\x2a\x65\xca\x91\x14\x4c\x7b\xb1\xcc\xfd\x58\xeb\x97\x09\xc9\x19\x9f\x47\x1f\xe5\x48\x1a\x39\xe8\x05\xc1\xd3\x17\x7b\x2b\xa4\x91\xa4\x73\xb8\x5d\x89\x95\x8a\xc4\x93\x54\xc9\x52\xd0\x01\x3c\x3e\x3a\x1c\x1d\x1f\xec\xbf\x00\xff\x57\x48\x08\xe7\xd6\x06\x89\x54\x20\x39\x85\x91\x92\x37\x1a\x95\x86\x5f\xfd\x9d\x00\xee\x0d\x8e\x26\xcc\xb8\x9c\x09\x24\xca\x4d\x15\xa1\x0c\x85\x79\xa2\x58\x9a\x99\x4e\x83\xdf\x81\xc7\xc7\x67\xa7\xfb\x87\x17\x4f\x5f\xec\x46\xca\xe5\xe2\x9f\x80\x91\xff\x00\x48\x1b\xc1\x48\xe0\x98\xfc\x7f\x0c\x9b\x23\xb7\xce\xc7\x00\x9c\x3a\x23\x4e\x07\x34\x11\xda\xd5\xa8\x58\xb2\xbe\x5c\x4e\x51\x25\x5c\xde\x0c\x20\x63\x94\xa2\x58\xb7\x36\xd4\x56\xa0\x3a\x97\xd2\x64\x4c\xa4\x03\x20\xc2\x56\x1e\x23\x1a\x69\xeb\x02\xcb\xa0\xd4\xb3\x8d\x2b\x52\x45\xe6\x55\xa1\xde\xaf\xbf\xbb\x2f\x11\xef\x46\x91\xa2\x40\xd5\x2a\x93\xaa\xd0\x07\xd0\x3b\x0c\x8a\xd9\xfa\x3e\x05\xa1\xb4\xc2\x3d\xfe\x05\x02\x08\xd6\x8d\x39\x51\x29\x13\x03\x20\xa5\x91\xdb\xb7\x2b\x88\x40\xde\xda\xac\x90\x9a\x19\x26\xc5\x00\x14\x72\x62\xd8\x14\xd7\x51\x17\x2e\x13\x14\x67\x03\xe8\xee\x4e\xda\xe3\x8b\xea\xd5\x76\x67\xe6\xee\x8e\xa4\x71\x36\xa8\xdc\x85\x6e\xb0\x3b\xd6\xfe\x41\xdb\x64\x70\x66\x5c\xc2\x59\x2a\x06\x10\xa3\x30\xa8\x5a\xbe\xc9\x99\xab\x33\x42\x6d\x7e\x03\x08\x60\x3f\x28\x66\x10\x80\x4a\x47\xe4\x49\xd0\x81\xe5\x9f\xb7\xff\xb4\x03\x01\x1c\x14\xb3\xea\x7f\x8b\xbd\xff\x74\x2b\x8f\x89\x54\x39\x30\x51\x94\xa6\xc5\xe5\x5f\x2d\xc2\xd2\xd8\x72\x1f\xb4\x13\xb9\x46\x6d\xb2\x6f\xdf\x2f\xb6\x15\x48\x37\x08\x7e\x69\x07\xae\x28\xaa\x0d\xc0\x7b\xb2\x03\xe8\x6e\xb0\xb9\x22\x7a\xd3\x54\x11\xc9\x16\x95\xb5\xc6\x76\x47\x72\xb6\xe5\xe4\x69\xb6\xc0\x01\x74\xfb\x0f\x01\xda\x8c\x8d\x4a\x63\xa4\xf8\x39\xca\xaa\xcc\x1b\x45\x84\xb6\x90\x03\x28\xed\xf1\x89\x89\xc6\xbf\xc1\x6c\xff\xf4\xe4\xe2\x20\xf8\x39\x66\xbf\xc3\x5d\x2c\xb9\x54\x3b\xce\xc6\x4e\xce\xe0\x41\xf7\xa9\xc2\x5c\x9e\x4d\xc2\x39\x04\x5e\x0f\x70\x23\xd4\x1f\x5b\x15\x97\x4a\x5b\x6f\x0a\xc9\xd6\x8f\xcb\xf6\x24\x0d\x32\xdb\x20\x3b\xe0\x3d\xd4\x91\xd8\x76\x87\x96\x32\x91\x71\xa9\xbf\x77\x9b\xeb\xf7\x4e\x82\xfe\xd1\xee\x0d\xbd\x1c\xb5\x26\x29\xb6\x30\x9a\x92\xed\xd6\x47\x33\xd8\xce\xed\xa8\x67\xdf\xbb\xb9\xdd\xff\x5e\x3d\xae\x76\x26\xad\xbd\x1b\xf4\x6d\x05\x52\x15\x20\xc5\x58\x2a\x52\x73\x2e\xa4\xd8\xd1\xdc\x51\x29\xd9\x6e\xed\x0d\x74\x92\x04\x41\x10\xc0\xa3\x7a\xda\x20\xc2\x3c\x84\xb0\x9f\xa1\xbf\x1c\x4c\x42\xbf\x1e\xb4\x42\x3b\x4f\x0c\xf7\x42\xca\xa6\x10\x73\xa2\x75\xe4\x2c\xef\x1e\xcd\xbc\xf3\xc0\x52\x35\x7a\xe7\x7e\xaa\x09\xeb\x96\x45\x23\x87\xdb\x71\xc8\x01\x9b\x4b\x29\x22\xc7\xaf\xe4\x97\x71\x46\x38\x47\x91\x62\x74\x7b\x0b\xde\x4a\x82\xbb\x3b\xc7\xce\x73\x99\xa4\x91\x73\xf9\xfe\xd3\xe7\x07\x90\x15\x6c\xd6\x85\xaa\x0d\x47\x8e\xbd\x49\x3b\xf5\xb0\x15\xfa\x59\x77\x7d\xdd\xed\x2d\xb0\xa4\x21\xe4\xee\x6e\x1d\xa3\x58\x83\x68\x22\xa8\xd6\x3a\x43\xeb\x4e\x73\x59\xe8\x17\x1b\xb0\x28\xe8\x06\x60\xdd\x99\xeb\x99\x12\x73\xc2\xb8\x03\x66\x5e\xdc\x0b\x05\x27\x31\x66\x92\x53\x54\x2b\xdd\x94\xf0\x12\x23\xa7\xda\xce\x6a\x6c\xe4\xfe\x70\x37\x6e\x41\xb4\xbe\x91\x8a\x36\xd0\xf7\xf2\x1a\xfa\x4a\xed\x6f\x78\xee\xc5\x5a\x25\x17\x0c\xf9\x66\x00\xcb\x46\x59\x23\xeb\x72\x94\x33\xe3\x0c\x79\x4d\x6d\x6d\x1b\xb6\x29\x5c\xb2\xb6\xac\x68\x67\xf8\x4e\x1a\x50\x98\x32\x6d\x50\x21\x7d\x09\x21\x81\x4c\x61\x12\x39\x7e\xa3\x5d\xcb\xf9\x7d\xca\xef\xee\x9c\xe1\xa9\x42\x62\x10\x88\x00\x12\xc7\xb2\x14\x26\xf4\xc9\x70\x8d\xfd\xd0\xb7\x05\xb5\xac\x3b\x9f\xb2\xa9\xad\xd3\xfa\xcb\x3e\x0b\x8c\x3f\x94\xa8\xe6\x90\x30\xa5\x4d\x07\x4c\x86\x02\x3e\xa3\xc9\x6c\x47\xa9\x84\x57\x52\x1a\x6d\x14\x29\xe0\xcd\x27\xaf\x7a\x4c\x08\x75\xac\x58\x61\x40\xab\x38\x72\x9a\xb9\x3b\x96\x14\xbd\xf1\xb5\xc5\xaa\x46\xee\xfa\xa7\xdb\xf3\xba\x5e\xd7\xd3\x9c\xe5\x5e\xce\x84\x37\xd6\xce\xca\x2f\xdb\xdc\x52\xc5\xcc\x3c\x72\x74\x46\x7a\xc7\x7d\xf7\xe4\xe8\xe2\xdb\xf8\x68\xfa\x8c\xfa\x9a\xe6\xff\xb9\x2e\x7c\xf1\xfe\xc3\x0d\x67\x6f\xa7\x5f\xf4\x9b\xe4\xec\xf5\xd5\xb3\xc9\xf3\xf7\x79\xea\x13\xff\x3c\xc3\x13\x9a\x9a\xc5\x3b\xdd\xcb\x8a\x84\xa4\x87\xe7\xf4\xf9\x41\x20\xee\xb1\x63\x25\xb5\x96\x8a\xa5\x4c\x44\x0e\x11\x52\xcc\x73\x59\x6a\x67\x18\xfa\xb5\xef\xbb\x82\xa0\x62\xac\xbd\x98\xcb\x92\x26\x9c\x28\xac\x22\x21\x63\x32\xf3\x39\x1b\x69\xdf\x54\xbc\xf8\x5d\xaf\xef\x05\xfe\xb8\x91\x7f\x20\xb0\xb3\x85\xa1\x27\x97\xaf\xae\x2e\x3f\xfe\xf9\xe9\xc4\xef\xe1\xd7\xf3\xf3\x2f\x57\xea\xea\x74\x7e\xf4\xdb\xc1\xef\x17\x23\x3c\x4e\x2e\xc6\x93\x83\x37\x27\xff\x9e\x7d\xf9\xfa\xfa\xf7\xc9\xd9\xec\xf0\x03\x13\xdd\xb3\xc9\xd5\xec\xa0\x3b\x7a\xa5\x46\x3f\x1d\x58\x4e\x66\x31\x15\xde\xa8\xc9\xa5\x15\x6c\x6c\x2b\x85\xdf\xf7\x02\x2f\x70\x09\x2f\x32\xe2\x1d\xda\xe0\x56\xa6\x1f\x88\x6f\xfa\xea\xea\x6a\xc1\xbf\xbd\x39\x46\xf2\x9c\x9c\xfe\xd9\x2f\xce\xaf\x7a\xea\x8f\xd7\xe3\x74\x6c\x8e\x16\xc5\xe4\x5d\xf1\x6d\xf2\x2c\xd8\x3f\x7b\x5e\x64\x8b\x39\xfe\x31\x39\x7f\x36\x96\x01\xc3\xdf\xd8\xe2\xfa\xf2\xed\x85\x54\x7f\x29\x71\x7b\xa1\xbf\xec\xae\x7e\xf5\xb4\xfb\xbf\x00\x00\x00\xff\xff\x09\xa6\xc8\x62\xfd\x0e\x00\x00")

func tmplLoginHtmlBytes() ([]byte, error) {
	return bindataRead(
		_tmplLoginHtml,
		"tmpl/login.html",
	)
}

func tmplLoginHtml() (*asset, error) {
	bytes, err := tmplLoginHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/login.html", size: 3837, mode: os.FileMode(420), modTime: time.Unix(1495035516, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplRegisterHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x57\x5b\x6f\xdb\x38\x13\x7d\xcf\xaf\x98\xb2\x28\xd0\x14\x96\x45\xe7\xd2\x06\x8a\xe4\x7e\x45\xfa\xe5\x69\x81\x16\x6d\x5f\xf6\x91\x16\x47\x12\x11\x8a\xd4\x92\x54\x6c\xc7\x9b\xff\xbe\xa0\x2e\x89\x25\x5b\xd9\x4b\x2b\x25\x52\x38\x43\x1e\x92\x67\x0e\x47\x93\xf8\xd5\xe7\x2f\x37\x3f\x7e\xff\xfa\x7f\x28\x5c\x29\x97\x27\xb1\x7f\x81\x64\x2a\x4f\x08\x2a\xe2\x0d\xc8\xf8\xf2\x04\x00\x20\x7e\x15\x04\xf0\x0d\xff\xa8\x85\x41\x0e\x25\x3a\x06\x8e\xe5\x16\x82\xa0\xf3\x37\xa6\xb4\x60\xc6\xa2\x4b\x48\xed\xb2\xe0\x8a\xec\xbb\x14\x2b\x31\x21\xf7\x02\xd7\x95\x36\x8e\x40\xaa\x95\x43\xe5\x12\xb2\x16\xdc\x15\x09\xc7\x7b\x91\x62\xd0\x34\x66\x20\x94\x70\x82\xc9\xc0\xa6\x4c\x62\xb2\x98\x81\x2d\x8c\x50\x77\x81\xd3\x41\x26\x5c\xa2\x74\x0f\xed\x84\x93\xb8\xfc\x2e\x72\x05\x75\x05\x7f\xc2\x77\x57\x73\xa1\x6b\x2b\xb7\x71\xd8\xfa\xda\x7e\xd6\x6d\xfb\xbf\xfd\xf5\x3f\x51\xfa\x55\x40\x6d\xe4\xdb\xc2\xb9\xca\x46\x61\x98\x69\xe5\xec\x3c\xd7\x3a\x97\xc8\x2a\x61\xe7\xa9\x2e\xc3\xd4\xda\x8f\x19\x2b\x85\xdc\x26\xdf\xf4\x4a\x3b\x1d\x9d\x53\x7a\x7a\x7d\xf2\x84\xb4\xd2\x7c\x0b\xbb\xa7\x66\x63\x62\xe9\x5d\x6e\x74\xad\x78\x04\xaf\x3f\xbc\x5f\x5d\x5d\x9e\x5d\x43\xf8\x0e\x32\x26\xa5\xf7\x41\xa6\x0d\x68\xc9\x61\x65\xf4\xda\xa2\xb1\xf0\x2e\x9c\x04\x08\xd6\xb8\xba\x13\x2e\x90\x42\x21\x33\x41\x6e\x18\x17\xa8\xdc\x5b\x23\xf2\xc2\xcd\x7a\xfc\x19\xbc\xbe\xfa\x7c\x73\xf6\xfe\xf6\xf4\x7a\x1a\xa9\xd4\x0f\xbf\x02\x46\xff\x02\x90\x31\x82\xd3\x20\x31\xfb\x7b\x0c\x1f\xa3\xa0\x8d\x47\x04\xa4\x8d\x08\x99\x81\x65\xca\x06\x16\x8d\xc8\x86\xdd\xf5\x3d\x9a\x4c\xea\x75\x04\x85\xe0\x1c\xd5\xd0\xdb\x53\xdb\x80\xda\x52\x6b\x57\x08\x95\x47\xc0\x94\xd7\x9e\x60\x16\xf9\x68\x80\x67\x50\xdb\xcd\xc1\x88\xdc\xb0\x6d\x23\xd5\xe7\xfe\x8f\xcf\x12\x99\xaf\x0d\xab\x2a\x34\x23\x99\x34\x52\x8f\xe0\xfc\x3d\xad\x36\xc3\x79\x2a\xc6\x79\x83\x7b\xf5\x06\x28\xd0\xa1\xb3\x64\x26\x17\x2a\x02\x56\x3b\x7d\x7c\xba\x8a\x29\x94\xa3\xc9\x2a\x6d\x85\x13\x5a\x45\x60\x50\x32\x27\xee\x71\x88\xfa\x10\x08\xc5\x71\x13\xc1\x62\x3a\x68\xaf\x6f\x9b\x6b\xbc\x9c\x4d\x30\xbd\x93\x7e\xb1\xb4\x59\x2e\x2c\xe8\xf4\x5e\x2f\x2e\xc7\x2e\x87\x1b\x17\x30\x29\x72\x15\x41\x8a\xca\xa1\x19\xad\x4d\x6f\x02\x5b\x30\xee\xe3\x4b\x81\xc2\x19\xad\x36\x40\xc1\xe4\x2b\xf6\x96\xce\xa0\xfb\x99\x9f\x9d\xce\x80\xc2\x65\xb5\x69\x7e\x8f\xf8\x2f\x4e\x8f\xf2\x98\x69\x53\x82\x50\x55\xed\x46\x5c\xfe\x5b\x11\xd6\xce\xcb\x3d\x1a\x07\x72\x40\x6d\x76\xe6\xef\xeb\x63\x02\x59\x50\xfa\x66\xbc\x71\xc3\xd1\x1c\x00\x3e\x93\x4d\x61\x71\xc0\xe6\x13\xd1\x87\xae\x86\x48\xf1\xd0\x78\x5b\xec\x60\xa5\x37\x47\x4e\x9e\x15\x0f\x18\xc1\xe2\x62\x1f\x60\xcc\xd8\xaa\x76\x4e\xab\x9f\xa3\xac\x89\xbc\x33\x4c\x59\x0f\x19\x41\xed\x8f\x4f\xca\x2c\xfe\x07\x66\x2f\x6e\x3e\xdd\x5e\xd2\x9f\x63\xf6\x05\xee\x52\x2d\xb5\x99\x38\x1b\x93\x9c\xc1\x5e\xf6\x69\xb6\xd9\x9d\x4d\x26\x25\xd0\xf9\x39\xe0\xc1\x56\xff\x59\xaf\xb4\x36\xd6\xaf\xa6\xd2\x62\x78\x5c\x8e\x07\x29\x2a\x7c\x82\x9c\xc1\x7c\xdf\xc6\x52\x9f\x1d\x46\xc6\x4c\xa7\xb5\x7d\xe9\x33\x77\x71\xfe\x89\x5e\x7c\x98\x9e\x70\x5e\xa2\xb5\x2c\xc7\x11\x46\x2f\xd9\x45\x7b\x34\xe9\x71\x6e\x57\xe7\xfe\x9e\xe6\xf6\xec\x25\x3d\x3e\xcd\xcc\x46\x73\xf7\xe8\xc7\x04\xd2\x08\x90\x63\xaa\x0d\x6b\x39\x57\x5a\x4d\x24\x77\x34\x46\x8f\x53\x7b\x0f\x9d\x65\x94\x52\x0a\xaf\xda\x6a\x83\x29\xb7\x0f\xe1\x9f\x71\xd8\x15\x26\x71\xd8\x96\x5a\xb1\xaf\x27\x96\x27\x31\x17\xf7\x90\x4a\x66\x6d\x42\xba\xaf\x47\x5f\xf1\xec\x79\x9a\x44\x4f\x9e\xab\x9a\xb8\x4d\x59\x3c\x21\x06\x73\x61\x1d\x1a\x02\x3e\x9c\x5a\x25\x24\xec\x4d\x1f\xd3\x82\x49\x89\x2a\xc7\x64\xb7\x9b\x3f\x35\x1e\x1f\x89\xaf\xea\x0a\xcd\x13\xf2\xf5\xcb\xf7\x1f\x7b\xb0\x0d\x74\xb1\x80\x26\x15\x27\xc4\x7f\xa8\xc9\xf2\xc6\x20\x73\x08\x4c\x01\x4b\x53\x5d\x2b\x17\x87\xc5\x62\x38\x66\xb7\x03\x91\xf5\x04\x3d\x3e\x0e\xf1\xaa\x01\x5c\xbf\xa3\xa6\x2f\x59\xee\x76\xcf\xc3\xe2\xb0\x3a\x80\x45\xc5\x0f\x00\xdb\x4c\xdd\x56\x99\xfe\x49\xc0\x6d\x2b\x4c\x88\x0f\x25\x81\x4a\xb2\x14\x0b\x2d\x39\x9a\xce\x1d\x2e\xa7\xc7\x57\xcc\xda\xb5\x36\xbc\xc7\x78\x6e\x0f\x70\x9e\xcc\x2f\x61\x61\xc9\x84\xec\x81\xba\xc6\x00\xa5\xb1\x01\xe3\xdc\xa0\xb5\x63\x28\xcf\x44\x6a\x4d\x76\x2b\x50\x1e\xee\xb9\xcb\xb5\x2d\xb6\xad\x57\xa5\x70\x64\x99\x36\x91\x89\xc3\xd6\xb9\x1c\xd3\xde\x31\xdd\x9d\x0a\xb2\xfc\x24\x0d\x32\xbe\x85\x5e\x1f\xc8\x3f\x42\xcc\xa0\x30\x98\x25\x24\x94\x3a\x17\x6a\x5a\x32\xcb\xdf\xbc\x3f\x0e\xd9\x72\x10\xa6\x38\xf4\x4a\xec\x04\x1b\x72\x71\xef\x05\xde\xbc\x4e\xe2\xb0\x53\x78\xd8\xfc\xcf\xf1\x57\x00\x00\x00\xff\xff\x90\x0b\xac\xd5\x83\x0c\x00\x00")

func tmplRegisterHtmlBytes() ([]byte, error) {
	return bindataRead(
		_tmplRegisterHtml,
		"tmpl/register.html",
	)
}

func tmplRegisterHtml() (*asset, error) {
	bytes, err := tmplRegisterHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/register.html", size: 3203, mode: os.FileMode(420), modTime: time.Unix(1495035516, 0)}
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
	"postgres/1_init.sql": postgres1_initSql,
	"tmpl/consent.html": tmplConsentHtml,
	"tmpl/error.html": tmplErrorHtml,
	"tmpl/login.html": tmplLoginHtml,
	"tmpl/register.html": tmplRegisterHtml,
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
	"postgres": &bintree{nil, map[string]*bintree{
		"1_init.sql": &bintree{postgres1_initSql, map[string]*bintree{}},
	}},
	"tmpl": &bintree{nil, map[string]*bintree{
		"consent.html": &bintree{tmplConsentHtml, map[string]*bintree{}},
		"error.html": &bintree{tmplErrorHtml, map[string]*bintree{}},
		"login.html": &bintree{tmplLoginHtml, map[string]*bintree{}},
		"register.html": &bintree{tmplRegisterHtml, map[string]*bintree{}},
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

