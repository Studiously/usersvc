// Code generated by go-bindata.
// sources:
// postgres/1_init.sql
// tmpl/consent.html
// tmpl/error.html
// tmpl/login.html
// tmpl/me.html
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

var _tmplConsentHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x91\x41\x6f\xa3\x30\x10\x85\xef\xfc\x8a\x59\x9f\x76\x0f\x8b\xaf\x55\x64\xa8\xa2\xb4\xb9\x26\x6a\xe8\xa1\x47\x63\x26\x60\xd5\xd8\xd4\x1e\xda\x46\x96\xff\x7b\x45\x80\x28\xe5\x32\xd8\x6f\xe6\x7b\xb6\x9f\xf8\xf3\x74\xd8\x55\x6f\xc7\x67\xe8\xa8\x37\x65\x26\xa6\x02\x46\xda\xb6\x60\x68\xd9\xb4\x81\xb2\x29\x33\x00\x00\xd1\x23\x49\x50\x9d\xf4\x01\xa9\x60\xaf\xd5\xfe\xff\x03\x5b\x24\xd2\x64\xb0\x3c\x1a\x94\x01\xa1\xd5\x9f\x08\x17\x37\x7a\x50\xce\x06\xb4\x24\xf8\xac\x67\x82\xcf\x38\x51\xbb\xe6\x52\x66\x62\x98\xc7\xb7\x16\xe4\x30\x18\xad\x24\x69\x67\xe1\xaf\x6e\x36\x10\x63\xbe\x1d\x1b\x8d\x56\x61\x4a\xff\xc0\xe3\xc7\x88\x81\xb0\x59\x99\x40\x0e\xa4\x52\x18\x02\x78\x0c\x6e\xf4\x0a\x03\x38\x3b\xfb\xd6\xd8\x49\x73\xce\xa1\xea\xf0\x17\xf9\x4b\x5a\x0a\x57\xcb\x65\x94\xdc\x26\x13\x7c\x28\x33\x71\x76\xbe\x07\xa9\xa6\xb6\x82\xf1\xc5\xe4\x51\x75\xd2\x18\xb4\x2d\x16\x31\xe6\xbb\x75\x91\x12\x83\x1e\xa9\x73\x4d\xc1\x8e\x87\x53\xb5\xbe\xc2\x68\xe6\x9f\xe9\x8b\xd1\x4b\xdb\x22\xe4\x2f\xeb\xc9\x4f\xca\x0d\x18\x52\xba\xb5\x08\xa3\x4b\xa1\xed\x30\x12\xd0\x65\xc0\x82\xa9\x0e\xd5\x7b\xed\xbe\x19\x58\xd9\x63\xc1\x62\xcc\x53\x62\xe5\xb5\x08\x6e\xf4\x3d\x1d\x6d\xb3\xa0\x04\x5f\x7d\x63\xcc\x55\xf0\xe7\xbd\x46\x73\x13\xef\xf9\x61\xac\x7b\x4d\x53\xaa\x7c\xba\xee\x54\x97\x20\xf8\x35\xfe\x9f\x00\x00\x00\xff\xff\x23\xff\x5c\x28\x0e\x02\x00\x00")

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

	info := bindataFileInfo{name: "tmpl/consent.html", size: 526, mode: os.FileMode(420), modTime: time.Unix(1495035516, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplErrorHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x3c\x8f\x41\x6b\x03\x21\x10\x85\xef\xfe\x8a\x17\xef\xa9\xd7\x52\x26\x5e\x9a\xf4\xda\x42\xb7\x94\x1e\x4d\x76\xa8\x82\x51\x51\x97\xb2\x2c\xfb\xdf\x8b\xeb\x92\x93\x32\xdf\xbc\xf7\x31\x74\x38\xbf\xbf\x0e\x3f\x1f\x17\xd8\x7a\xf7\x5a\x50\x7b\xe0\x4d\xf8\x3d\x49\x0e\xb2\x0d\xd8\x8c\x5a\x00\x00\xdd\xb9\x1a\xdc\xac\xc9\x85\xeb\x49\x7e\x0d\x6f\xc7\x67\xb9\xa3\xea\xaa\x67\x7d\xc9\x39\x66\x1c\xf1\x59\xa7\xd1\xc5\xa9\xf8\x99\x54\x27\x82\x54\x2f\xa2\x6b\x1c\x67\x2d\x28\xf5\xe0\xb7\x8d\x31\x95\x03\x06\xcb\x99\xf1\x67\x0a\x4c\x00\xb7\x9a\x17\x2c\x0b\x9e\xb6\x2f\xd6\xb5\x5b\xd4\x35\xf7\xd8\x03\x9d\xb9\xdc\xb2\x4b\xd5\xc5\xd0\xb6\x48\xa5\xa6\xda\x1d\x6a\xbb\xe9\x3f\x00\x00\xff\xff\xcb\x78\xe2\x5c\xe3\x00\x00\x00")

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

	info := bindataFileInfo{name: "tmpl/error.html", size: 227, mode: os.FileMode(420), modTime: time.Unix(1495035516, 0)}
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

var _tmplMeHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x55\xdf\x57\xe2\x3a\x10\x7e\xe7\xaf\xc8\xe6\x3e\x6a\x1b\x10\x10\xbc\xa7\xe5\x1c\x14\xd0\x55\x57\x45\x56\x59\x7d\x0b\xc9\xd0\xa6\x4d\x93\x9a\xa4\x85\xb2\x77\xff\xf7\x7b\xac\x3f\xef\xdd\xb3\x7b\xf1\x5c\x5e\x86\x99\xf9\x9a\x99\xaf\x99\x6f\x1a\x7c\x1a\x5d\x1e\x7d\xbd\xbb\x1a\xa3\xd8\x65\x72\xd0\x08\x1e\x0d\x92\x54\x45\x21\x06\x85\x1f\x03\x40\xf9\xa0\x81\x10\x42\xc1\x27\xcf\x43\xd7\xf0\x50\x08\x03\x1c\x65\xe0\x28\x72\x34\xb2\xc8\xf3\x9e\xf3\x75\x88\xc5\xd4\x58\x70\x21\x2e\xdc\xd2\xeb\xe3\xf7\x29\x45\x33\x08\x71\x29\x60\x95\x6b\xe3\x30\x62\x5a\x39\x50\x2e\xc4\x2b\xc1\x5d\x1c\x72\x28\x05\x03\xaf\x76\x76\x91\x50\xc2\x09\x2a\x3d\xcb\xa8\x84\xb0\xb5\x8b\x6c\x6c\x84\x4a\x3d\xa7\xbd\xa5\x70\xa1\xd2\x78\xd0\x78\x6b\xeb\x50\x6b\x67\x9d\xa1\x39\x3a\x9a\xcd\xde\x3a\x92\x42\xa5\xc8\x80\x0c\xb1\x75\x95\x04\x1b\x03\x38\x8c\x62\x03\xcb\x10\xc7\xce\xe5\xf6\x4f\x42\x32\xba\x66\x5c\xf9\x8b\x97\x13\x1e\x1d\xa6\x33\xf2\x1a\x20\x1d\xbf\xe9\x37\x3d\x2a\xf3\x98\xfa\xfb\x84\x59\xfb\x96\xf3\x33\xa1\x7c\x66\x2d\xae\xeb\x3d\xfd\x84\x72\x10\x19\xe1\xaa\x10\xdb\x98\xb6\xfb\x1d\xcf\xac\xf4\xe7\x6b\xb0\xc9\xcd\x5e\xc5\xda\x9b\xfe\xf1\x2d\xb9\xb8\x82\xfb\xf9\xb0\xec\xee\x9b\x59\x76\x2e\xf9\x51\xfb\x9a\x0c\xef\x37\xc7\xd7\xea\x78\x3d\x9d\x9e\xa9\xb3\x54\x4f\x6e\xe3\xc9\x34\xbe\xb8\x59\x8d\xab\x53\x8c\x98\xd1\xd6\x6a\x23\x22\xa1\x42\x4c\x95\x56\x55\xa6\x0b\x8b\x3f\x40\x93\x71\x95\x58\x9f\x49\x5d\xf0\xa5\xa4\x06\x6a\x8e\x34\xa1\x6b\x22\xc5\xc2\x92\xa5\x56\xce\xa3\x2b\xb0\x3a\x03\xd2\xf1\x7b\x7e\xb3\x26\xfa\x3e\xfc\xca\x75\xd0\x08\xc8\xd3\x58\x04\x0b\xcd\xab\x41\x23\x50\xb4\x44\x4c\x52\x6b\x43\xac\x68\xb9\xa0\x06\x3d\x19\x4f\xa8\x12\x8c\x05\xb4\x88\x3c\x5b\x30\x06\xf6\xb5\x67\xfa\xcf\x07\xbc\x85\xa1\x8a\xbf\x34\xfd\x07\x1e\xcc\x5c\xc1\x85\x2e\xac\xac\xd0\x5f\xe8\x4b\x85\x86\x8c\xe9\x42\xb9\x80\xd0\xc7\xf2\x8a\x96\x83\x46\xc0\xc5\x6b\x59\x9b\x7a\xac\x58\x80\x17\x19\xc1\x5f\x4a\xfc\x9c\x46\xcf\xb6\x85\x07\x01\xe1\xa2\xfc\x4f\xe0\xde\xb6\xc0\xf6\xb6\xc0\xce\xb6\xc0\xee\xb6\xc0\xfd\x6d\x81\xbd\x6d\x81\xfd\x6d\x81\x07\xaf\xc0\x17\x93\x0f\x4e\x40\x4a\xbd\x8b\xbe\x7f\xf7\x0b\x0b\xc6\xbf\xa0\x19\xfc\xf8\xf1\x29\x20\x79\x9d\x0d\xe8\xf3\x15\x13\xa9\x23\x5d\x38\x3c\x98\x89\x48\xa1\xcb\xa2\xbe\xd9\x1a\xd5\xa8\x15\x9d\x4c\x0b\x30\x15\x5a\x0a\x63\xdd\x2e\x72\x31\x28\xf4\x15\x5c\x0c\xe6\xd9\x79\x53\xfc\xe9\xcc\xaf\x15\x1f\x58\x66\x44\xee\x90\x35\xec\xdd\xd8\x6b\x0e\x7e\xf2\xf0\x78\x56\x3d\xf1\x4f\x7f\xbd\xb6\xdf\xf2\x5b\xbe\x95\x22\xab\xa7\x3a\x79\x27\xe0\x9f\xe5\x3b\xec\x4d\xee\x93\x5e\xb9\xc3\x89\xe5\xd9\x97\x87\x9c\xa8\xcb\xe9\x4a\x8a\xf3\xf2\xc6\x9e\x2e\x47\x27\xf3\x9d\xf4\xe0\x32\x8b\x08\x25\xe3\x18\x86\x3c\x72\x9b\x0b\xdb\x8e\xf3\x25\x8d\xf6\xc7\xfc\xa0\xdb\x54\x6f\x67\xff\x4a\xc6\x01\x79\xea\xfd\x57\x24\x7e\xaf\x5d\x57\xbf\x17\xd2\xf2\x3b\x7e\x93\x24\x2f\xfe\x16\xc4\x46\x1b\xc7\x87\x57\x87\xf3\xab\xeb\x6f\xb3\x21\x69\xc3\xdd\x78\x7c\x33\x37\xf3\xa3\xaa\x77\xdc\x3d\x9b\x2c\xa0\xbf\x9c\x24\x69\xf7\x74\xf8\x79\x7d\x73\x77\x72\x96\x8e\xd6\xfb\x53\xa1\x5a\xa3\x74\xbe\xee\xb6\x16\x87\x66\xf1\xbf\x89\x7d\x74\xf7\x26\xff\x5e\xbd\xbf\xe7\x57\x1e\xce\xe7\x1b\x79\x7f\xda\x07\x7a\x40\x8f\xbe\x75\xf2\xf1\xbc\x6d\x6e\x4f\x92\x28\x71\xbd\x4d\x9e\x5e\xe4\xf7\xe9\x4e\x73\x6f\x74\x90\xc7\x9b\x0a\x6e\xd3\xf1\x4e\xa2\x9b\x02\x8e\xc5\xe6\xe1\xea\x7c\xa2\xcd\xc7\x2e\x8e\x3c\xef\x42\x52\x7f\x49\xff\x0e\x00\x00\xff\xff\x79\xc0\x2b\xed\x59\x07\x00\x00")

func tmplMeHtmlBytes() ([]byte, error) {
	return bindataRead(
		_tmplMeHtml,
		"tmpl/me.html",
	)
}

func tmplMeHtml() (*asset, error) {
	bytes, err := tmplMeHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/me.html", size: 1881, mode: os.FileMode(420), modTime: time.Unix(1495035516, 0)}
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
	"tmpl/me.html": tmplMeHtml,
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
		"me.html": &bintree{tmplMeHtml, map[string]*bintree{}},
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

