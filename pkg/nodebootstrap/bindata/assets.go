// Code generated by go-bindata. DO NOT EDIT.
// sources:
// bindata/assets/10-eksctl.al2.conf (1.025kB)
// bindata/assets/bootstrap.al2.sh (773B)
// bindata/assets/bootstrap.helper.sh (1.302kB)
// bindata/assets/bootstrap.legacy.al2.sh (1.286kB)
// bindata/assets/bootstrap.legacy.ubuntu.sh (2.275kB)
// bindata/assets/bootstrap.ubuntu.sh (597B)
// bindata/assets/efa.al2.sh (351B)
// bindata/assets/efa.managed.boothook (484B)
// bindata/assets/install-ssm.al2.sh (159B)
// bindata/assets/kubelet.yaml (480B)

package bindata

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _bindataAssets10EksctlAl2Conf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x52\xdf\x6b\xdb\x30\x10\x7e\xf7\x5f\x71\x90\x3e\x6c\x10\xd9\xac\x7d\x2b\xf8\xc1\x8b\xdd\x12\x96\x3a\x25\x76\xb7\xc1\x36\x8c\x2c\x5d\xd2\x23\xb2\x64\x24\x39\x69\x57\xf2\xbf\x0f\xc7\xf1\x96\xb2\x32\xf6\x26\xdd\x77\xf7\x7d\xdf\xfd\x98\x00\x6e\x9d\xf0\x8a\xb9\x16\x05\xad\x49\x80\x7b\x76\x1e\x1b\x09\xd2\x9a\x96\x91\x86\x4e\x93\x87\xb5\xb1\xb0\xed\x6a\x54\xe8\xa7\xc7\x4f\xd2\xf0\x9f\x46\xc3\x82\x74\xf7\x04\x97\xf0\x2e\x59\x5c\xbe\x0f\x26\x50\x2e\xd3\x25\xa4\xd8\x5a\x14\xdc\xa3\x9c\xc2\x9e\x94\x82\x1a\xc1\x62\x63\x76\x28\xc1\x19\xa3\x83\xe0\x5b\x81\x76\x47\x02\x7f\x04\x13\x58\x18\xc1\x15\x34\xe8\xb9\xe4\x9e\x43\xcb\x2d\x6f\xd0\xa3\x75\xd7\xb0\xca\x6e\xe7\xcb\x7c\x0a\xc9\x97\xa2\x4a\xb3\x9b\xe4\x61\x51\x56\x43\x2c\xc8\xf4\x8e\xac\xd1\x0d\x6a\x7f\x43\x0a\xe3\x08\xbd\x88\x86\x56\xa2\x91\x2b\x44\xbd\x0b\x26\x70\xab\x4c\xcd\x15\x70\x2d\xc1\x79\xee\x49\xbc\xd2\x98\x2d\x1e\x8a\x32\x5b\x55\x69\x5e\x4c\x21\x5f\xa6\x59\xb5\x48\x3e\x66\x8b\xf1\x53\x26\xf3\xbc\x2c\xfe\x29\x77\x9a\xcb\x49\x6d\x68\x47\x1b\xcd\xde\x10\x3b\x52\xce\xef\xa7\x30\xcf\x8b\x32\xc9\x67\x59\x35\x4f\xff\x8b\x5b\xf5\xac\x47\x85\x20\x7b\x42\x51\x78\x6e\x7d\x7c\xf6\x8c\x3a\x67\xa3\x9a\xf4\x58\x00\xdf\x03\x00\xc6\xb4\x91\xc8\xa8\x8d\x2f\x5e\x4e\xca\x87\x73\x40\xf1\x1a\x95\x1b\xc1\xa1\xed\xc3\x94\xab\xf6\x91\x87\x83\x7e\x48\x26\x22\xed\x3c\xd7\x02\x19\xc9\xf8\xe2\xe5\xcc\xf8\xc8\xd5\xf0\x27\xd6\x1a\xd9\x13\xdd\x25\x5f\xab\xfb\x65\x5a\x8c\x90\xc5\x0d\x39\x8f\xf6\xa8\x17\x7b\xdb\xe1\x79\x70\x4f\xfe\x91\x79\x4e\xda\xff\x36\x31\x8c\x7b\x2c\x17\xca\x74\x92\xb5\xd6\xec\x48\xa2\x8d\xf9\xde\x8d\x80\xd1\x7d\x1d\x5a\x66\x3b\xed\xa9\xc1\x58\x1a\xb1\x45\x3b\x76\x87\x7e\x6f\xec\x96\xb5\xaa\xdb\x90\x8e\x85\xa6\xb1\x4e\x13\xab\x49\x33\x49\x36\x8e\x4c\xeb\x23\xa1\xa9\x1f\xdb\x19\x2c\x8c\x5e\x0f\x78\xbf\x86\x1e\xd7\xe8\x43\x79\xca\x68\x8d\x64\xa4\xd7\x96\x9f\x59\xa0\x86\x6f\x30\xbe\x78\xe9\xaf\x34\xfb\x54\x54\xd9\x6c\x55\x25\xb3\xd9\xf2\x21\x2f\x0f\xa1\xdc\xda\x10\x85\x0d\x07\xf8\xf5\x11\x1f\x4e\xd1\x22\x5b\x7d\x9e\xcf\xb2\xa2\x4a\x97\x77\xc9\x3c\x3f\xf4\xcb\x8f\x5a\xde\x39\xbc\xbe\x0a\xaf\x18\x6e\x5d\xdd\x91\x92\xe1\x87\x93\x89\x7e\xc7\xbd\x4d\xda\xfc\x75\x2b\x43\x38\x7c\xe6\x8d\xfa\x33\xaa\xb7\x12\xfb\xa3\xea\xb3\x82\x5f\x01\x00\x00\xff\xff\x04\xfc\xe9\x45\x01\x04\x00\x00")

func bindataAssets10EksctlAl2ConfBytes() ([]byte, error) {
	return bindataRead(
		_bindataAssets10EksctlAl2Conf,
		"bindata/assets/10-eksctl.al2.conf",
	)
}

func bindataAssets10EksctlAl2Conf() (*asset, error) {
	bytes, err := bindataAssets10EksctlAl2ConfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bindata/assets/10-eksctl.al2.conf", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xb0, 0x2b, 0x5f, 0x5a, 0xc4, 0x94, 0x74, 0x90, 0x34, 0x26, 0x3b, 0xc9, 0x7, 0x3f, 0x89, 0x8f, 0xe9, 0x6f, 0xe5, 0x33, 0x58, 0x2b, 0x82, 0x48, 0xd, 0xca, 0x2b, 0x91, 0x5f, 0x6f, 0x4a, 0xab}}
	return a, nil
}

var _bindataAssetsBootstrapAl2Sh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x91\xd1\x6b\xdb\x30\x10\xc6\xdf\xf5\x57\xdc\xb4\x42\x60\x20\x7b\x83\xd2\x87\x3d\x0c\xd2\xcc\x2b\x61\x69\x56\x9c\x64\x14\xba\x62\x64\xf9\x92\xa8\xb1\x25\xef\x4e\x0e\x1d\x25\xff\xfb\x10\x4b\xbc\x2c\xcd\x1e\x75\xf7\xfb\xbe\xbb\xef\xf4\xf6\x4d\x5a\x5a\x97\x96\x9a\xd7\x42\x30\x06\x50\x1e\x90\x08\x9f\x6d\x38\x3c\x5b\xdb\xe2\x52\xdb\xfa\xf0\x76\xbe\x73\x8c\x41\x08\xf6\x1d\x19\x84\x74\xab\x29\xad\x6d\x99\x9a\xda\x77\x55\xca\x86\x6c\x1b\x38\xc5\x0d\x9b\x50\xa7\xa5\xf7\x81\x03\xe9\x36\x59\x63\xdd\x22\x25\x71\x10\x9a\xb5\x07\xf9\x87\xf8\x08\xd4\x39\x67\xdd\x0a\x52\x0c\x26\xca\xfe\x6a\xa4\x78\x5d\x4b\x78\x0d\xf2\xe2\x65\x34\x59\xcc\xe6\x59\x5e\x4c\x87\xb7\xd9\x4e\xc2\x0f\x01\xa0\x94\x6e\x2d\x23\x6d\x91\x14\xba\xaa\xf5\xd6\x85\x88\x0e\xef\xc6\xc5\x2c\xcb\xbf\x67\x79\xb1\xc8\x27\x3d\x5c\x5e\x5d\x2a\x53\x77\x1c\x90\x94\xd1\x11\xbc\xbe\xba\x2c\x0e\xbe\xa3\x61\x0f\x56\x8e\x7b\xd0\xb6\xc7\xc3\x3f\x4f\x67\x3d\xb5\xe9\x4a\xac\x31\x28\x7c\x0e\xa4\x95\xa6\x15\x47\xf2\xeb\xe2\x3a\x9b\x64\xf3\x22\xbb\x9f\xe7\xc3\x62\x98\xdf\xcc\x76\xf2\x34\x7f\x83\xb4\x8a\xf9\x3b\x46\x02\xdf\x06\xeb\x1d\x83\x75\xc1\xc3\xc1\xd3\x78\xb7\xb4\xab\xe4\x89\xbd\x93\x22\x1e\x01\x06\xd4\x80\x5a\xc2\xc5\xcb\xfc\xf6\xae\x88\x43\x8a\xd1\xb7\xe9\x97\xdd\x00\xb2\xfb\xf1\x5c\x3c\xfd\x04\xc5\x30\x48\x1e\xde\x3f\xc2\x3b\x48\x1e\x3e\x3c\x0e\x8e\x97\x89\xe8\xf8\x66\x27\x5f\x2f\xd8\x77\x3e\xc5\xde\xbf\xde\x52\x34\xdb\x33\xd5\x73\xc6\x42\xf0\x2f\x0e\xd8\x98\x50\x43\xa5\xb1\xf1\x4e\x11\xd6\x5e\x57\xa7\x3f\x8f\x1c\x34\x85\x18\xbe\xf2\x66\x83\xb4\xa7\xe5\x91\x7e\xcf\xec\x81\xff\x1b\xf4\xf7\xdf\xf0\x39\xf9\xbe\x7d\xa2\xaf\xbc\x43\x29\x7e\x07\x00\x00\xff\xff\x82\xe9\xcc\x95\x05\x03\x00\x00")

func bindataAssetsBootstrapAl2ShBytes() ([]byte, error) {
	return bindataRead(
		_bindataAssetsBootstrapAl2Sh,
		"bindata/assets/bootstrap.al2.sh",
	)
}

func bindataAssetsBootstrapAl2Sh() (*asset, error) {
	bytes, err := bindataAssetsBootstrapAl2ShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bindata/assets/bootstrap.al2.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd9, 0x95, 0x58, 0x9f, 0x70, 0xdb, 0x6, 0x93, 0x8d, 0x96, 0x89, 0x18, 0xb9, 0xa0, 0x83, 0xaa, 0x1c, 0x99, 0x6d, 0x43, 0x6f, 0xdc, 0x9b, 0x71, 0x25, 0x73, 0x4d, 0x91, 0xe6, 0x76, 0x2d, 0x4f}}
	return a, nil
}

var _bindataAssetsBootstrapHelperSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x94\x61\x6b\xe3\x46\x10\x86\xbf\xef\xaf\x98\x6e\xcc\x25\xa6\x59\xeb\x2e\x5c\x03\x0d\x18\xaa\xd8\xba\xab\x38\x45\x31\xb6\x5c\x52\x42\x10\x6b\x69\x1c\x6f\x23\xef\x8a\xdd\xd1\x39\x47\xd0\x7f\x2f\x92\xa3\x58\xbe\xd2\x7c\x92\x66\x77\x9e\x77\xdf\x19\x86\x39\xf9\xc5\x5b\x29\xed\xad\xa4\xdb\x30\xe6\x90\x40\x18\x40\x6b\xf1\x59\x51\x17\x96\xaa\xc4\xb5\x54\x45\x17\x6b\x53\x69\x87\xc4\x98\x33\x95\xcd\x10\x3c\xa4\xcc\xc3\x27\x97\x51\xe1\x3d\x55\x2b\x2c\x90\x46\xa8\xbf\xc3\x09\xac\x55\x81\xb0\xb3\x8a\x08\x35\xac\x7e\xc0\xca\x18\x72\x64\x65\x59\xa2\x65\xec\x04\x96\x0e\x21\xbc\x99\x2e\xbe\x5f\x00\x19\x78\x44\x82\x2d\x92\xcc\x25\x49\x96\xdc\x7e\x0b\xe2\x31\x1f\x9c\x65\x95\x2d\x40\x08\xa7\x0a\xd4\x04\xe2\x0e\x66\xcb\x04\xc4\x9f\xc0\xef\x84\xdc\x39\x81\xd9\x85\xe8\x20\x41\xe6\x09\xb5\x20\x2a\x84\xc3\xcc\xe8\xdc\x5d\xc1\xe5\xc7\x8f\x1c\x36\x44\xe5\x95\xe7\x7d\xba\xfc\x7d\x74\xf1\xdb\xe7\xd1\xeb\xd7\x2b\x24\xa1\x23\x4f\x96\xca\x6b\xc9\x21\x67\xeb\x4a\x67\xa4\x8c\x6e\xcc\xa4\x9d\xee\xd9\x10\x5e\x18\xc0\x4f\x4e\xde\xb1\x70\x05\x83\xd6\x3f\x07\xfe\xfe\xd3\x0d\x26\x1a\xce\x1b\x7c\xe2\xac\x66\xcc\x9f\x85\xe9\x22\x98\xff\x15\xcc\xd3\xe5\x3c\x1a\xf3\xc1\xcb\xf1\x49\xcd\xd9\xf5\xe5\xe7\x74\x12\x2d\x17\x49\x30\x4f\x27\x7e\x93\x72\x7c\x52\x73\x16\xc6\x8b\xc4\x8f\x27\x41\x1a\x4e\x9b\x16\xf6\x6b\x01\xa5\x1d\x49\x9d\xa1\x50\xf9\xb0\x97\x19\x85\x5f\x82\xc9\xdf\x93\x28\xf8\x7f\xa0\x50\x6b\x14\xd9\x8f\xac\xc0\x21\x67\xdd\x7b\xd3\x78\xd1\x58\xe8\x85\x57\xa2\xe6\x2c\xbe\x9d\x06\x69\xe2\x87\x71\xd2\x5e\xf7\xc2\xf6\xfa\xc6\xbf\x4b\x67\xb7\xd3\xf6\xae\xfb\x3f\x70\x91\x7f\x1d\x44\x07\x6e\x1f\xd6\xe7\xda\xe4\x7b\x13\xad\x87\xf1\xe0\xe5\xbf\xe6\xeb\x73\x59\x94\x1b\x39\xda\x4f\xe3\x48\x19\xaf\x57\x6e\x9f\x08\xa7\x35\x67\xec\xdb\xf2\x3a\x88\x82\x24\xf5\xe7\x5f\x17\xe3\x33\x2e\xc4\xfe\x09\xb9\xc2\xc2\x8d\x8f\x5f\xe7\x43\x76\x7f\x0f\x42\xc3\x71\x35\x35\x87\x87\x07\xf8\xf0\x01\xfa\x52\xbf\xb6\x5a\x16\x1f\x95\x23\xb4\x62\xa7\x68\x23\x48\x2a\x4d\x6f\xa2\x1d\x3c\x64\x27\x20\xc4\x56\x3e\x8b\xd2\xe4\x0e\xa4\x03\x09\x93\x28\x04\x69\x1f\xab\x6d\x33\x65\xca\x41\x8e\xa5\xc5\x4c\x12\xe6\xe7\x40\x1b\xe5\x9a\x33\x09\x3b\x63\x9f\xa4\x35\x95\xce\xa1\xd2\xa4\x0a\xd8\xe1\x21\x13\x5c\x55\x96\xc6\x12\xac\x8d\x85\xad\x7c\x9e\x99\xdc\xcd\xd0\xc6\x26\xc7\x43\x15\x5d\xdf\xdf\x29\xa1\x33\x36\xee\x67\x0f\xdf\xda\x16\xdc\x25\x73\x7f\xdf\x3c\x3e\x78\xe9\xe3\xf7\x7f\x3c\x34\xfd\xed\xa6\x22\xf6\x6f\x82\xfe\x94\x34\x71\xcd\xdf\x74\x26\xb7\xf1\x97\xf0\xeb\xf8\xb4\x5d\x24\xcd\x06\xb1\x1a\x09\x5d\xb7\x4c\xba\xaf\xc8\x8c\x5e\xab\xc7\xd1\x3f\xce\xe8\xd3\x9f\x4c\x1c\x49\x1c\xef\x22\x81\xcf\x64\xe5\x2b\x95\xdc\xcc\xd2\x86\x6c\x81\xf1\xa9\x47\xdb\xf2\x48\xfe\x35\xed\xdf\x00\x00\x00\xff\xff\xa3\xa1\xa2\xa6\x16\x05\x00\x00")

func bindataAssetsBootstrapHelperShBytes() ([]byte, error) {
	return bindataRead(
		_bindataAssetsBootstrapHelperSh,
		"bindata/assets/bootstrap.helper.sh",
	)
}

func bindataAssetsBootstrapHelperSh() (*asset, error) {
	bytes, err := bindataAssetsBootstrapHelperShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bindata/assets/bootstrap.helper.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc1, 0x5f, 0xa5, 0xe5, 0xfa, 0xb3, 0x61, 0x47, 0xca, 0x28, 0xf9, 0xb7, 0x13, 0xed, 0x4, 0x40, 0x34, 0x22, 0x7e, 0xc2, 0x9, 0x87, 0x4e, 0x67, 0x56, 0x30, 0x55, 0x2b, 0x62, 0x44, 0x4c, 0x95}}
	return a, nil
}

var _bindataAssetsBootstrapLegacyAl2Sh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x93\x61\x4f\x02\x47\x10\x86\xbf\xef\xaf\x98\x1e\x17\x23\xa9\xcb\xa9\xb1\x26\xa2\x34\xa1\xdc\x99\x5e\x8a\x40\x0a\xb6\x1a\x63\x2f\xcb\xde\x50\x36\x2e\xbb\x97\xdb\x01\x35\xe4\xfa\xdb\x9b\xc5\xc3\x82\x5a\x3f\xc1\xcc\xbc\xb3\xf3\xee\x33\xb7\x8d\x1f\xa2\xa9\x32\xd1\x54\xb8\x39\x6b\xc0\x64\x18\x0f\x21\xc6\xa2\x44\x29\x08\xf3\x23\x78\x56\x5a\xc3\x14\xa1\xc4\x85\x5d\x61\x0e\xce\x5a\xc3\x98\x43\x02\x6e\x01\xcb\x12\x5f\x14\x6d\xc3\x42\x15\x38\x13\x4a\x6f\x63\x63\x97\xc6\x21\x31\x36\x5b\x1a\x49\xca\x1a\xf8\x1b\x29\x5b\x88\x97\xac\xb0\xb9\x3b\x6c\xc2\x9a\x01\x3c\xcf\x95\xf6\xc7\x8b\x1c\x94\x71\x24\x8c\xc4\x8c\x5e\x0b\x04\xaf\xb9\x84\xdc\x32\x00\x00\x35\x03\x78\x78\x80\x20\x5c\xef\x89\xaa\x00\x3a\x1d\x9f\x3d\xa9\x02\x78\x7c\x84\x83\x83\x5a\xe5\x9b\x7d\xf1\x1f\xf8\xeb\xe1\x98\x5f\x3c\xfe\x18\xfa\xf2\x25\xd0\x1c\xcd\xe6\x40\x00\x94\x73\x0b\xb5\xb2\x4e\x95\x48\xcb\xf2\xad\x3e\x53\x0c\x20\xb7\x06\xe1\x0a\x22\x24\x19\xe1\x93\x93\xa4\xa3\xad\xfb\xd6\x42\x14\xac\x62\xac\x01\xb7\x0e\x21\xbd\x89\xc7\xab\x53\x20\xeb\x6f\x08\x0b\x24\x91\x0b\x12\x6c\x32\xfc\x2d\x19\x74\x82\xf0\x50\x2e\x4b\x0d\x9c\x3b\xa5\xd1\x10\xf0\x3b\x18\xdd\x4e\x80\xff\x0a\xc1\x1d\x17\xcf\x8e\xa3\x3c\xe5\xdb\x26\x4e\xf6\x09\x0d\x27\xd2\xdc\xa1\xb4\x26\x77\x6d\x38\x3f\x3e\x0e\x60\x4e\x54\xb4\xa3\xe8\xe4\xfc\xa2\x75\xfa\xd3\x59\xab\xfe\x8d\xb4\x20\x74\x14\x89\x42\x45\x9b\xce\x66\xf0\x01\x77\x7d\x6e\x8d\xfb\x83\x93\x6f\x2c\xb4\x21\xdc\xf8\x0f\x20\xf8\x7e\xb4\x6f\xe3\xbe\x2f\x0a\x4f\x02\xcf\x64\x30\x8c\x93\x2c\x1d\xf9\x8b\xef\x3a\x00\x6d\xa5\xd0\x5c\x15\xab\xb3\x66\xc0\xd2\xc1\x78\xd2\x1d\xf4\x92\x2c\x8d\x3f\x09\xb7\x3b\xe6\x2a\xdf\x55\x4e\xee\x47\xc9\xff\x6b\xfd\xf7\xd0\x0c\x58\xf7\xcf\x71\x36\x4e\x7e\xff\x23\xed\x25\xe3\x2c\x1e\xde\x74\xd3\xc1\xa7\x1e\x87\xe5\x4a\x49\x74\x51\x6e\x17\x42\x79\x64\x8c\x39\xbb\x2c\x25\xee\xed\xfa\x69\x39\x45\x8d\xd4\x42\xb3\x82\x06\xd0\x5c\x39\x90\xc2\x80\x5d\x61\x59\xaa\x1c\xe1\xa6\x7b\x97\x8d\x86\xf1\x98\xfd\x67\xb1\x9f\x5e\x27\xbd\xfb\x5e\xff\x1b\x9f\x5a\xcd\x90\xcb\x57\xa9\xbd\xdb\x0d\xaa\x7e\xf7\x97\xa4\x3f\xee\x04\xe1\x7a\x27\xac\x8e\x8c\xcd\xdf\xd4\x1b\x71\x27\x5c\x7f\x9e\x52\x79\xe7\x52\x10\xfc\xfc\xa5\xf1\x0d\xf0\x8d\xfd\xab\xab\x64\x78\xfd\xbe\x98\x7a\x50\x3a\xaa\xf6\xd6\xb0\x33\x21\x8d\xab\x0f\xdc\x77\x8a\x3e\xae\xbe\x04\x1d\xae\xbf\xc8\x56\x6c\x0b\xaa\x13\xae\xb7\x7f\xdb\xbc\xe6\x53\xbf\x27\xff\x68\xf7\x07\x04\xcd\x6a\x0f\xcf\x3e\x1d\xe6\xef\xc3\xdc\xab\x23\x5c\x48\xd2\x90\x0b\x5c\x58\xc3\x4b\xd4\x56\xe4\x3b\x79\x34\x62\xaa\x11\x6a\x22\x3b\x05\x47\xa2\xa4\xf7\xfc\xbf\x01\x00\x00\xff\xff\x46\x01\xeb\xc2\x06\x05\x00\x00")

func bindataAssetsBootstrapLegacyAl2ShBytes() ([]byte, error) {
	return bindataRead(
		_bindataAssetsBootstrapLegacyAl2Sh,
		"bindata/assets/bootstrap.legacy.al2.sh",
	)
}

func bindataAssetsBootstrapLegacyAl2Sh() (*asset, error) {
	bytes, err := bindataAssetsBootstrapLegacyAl2ShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bindata/assets/bootstrap.legacy.al2.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x35, 0xb1, 0x58, 0xf5, 0xe6, 0x40, 0x92, 0xbe, 0xa6, 0x27, 0xc2, 0xfc, 0xbd, 0xc2, 0xe8, 0x54, 0x5e, 0x40, 0x37, 0x77, 0x73, 0x87, 0x9, 0xf, 0x97, 0xc4, 0xbc, 0x10, 0x5, 0x1b, 0x44, 0xf8}}
	return a, nil
}

var _bindataAssetsBootstrapLegacyUbuntuSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x55\xf1\x6f\x1a\xb9\x12\xfe\xdd\x7f\xc5\xbc\x0d\xea\x0b\x7a\x31\xdb\xa4\x7d\x95\x9a\x96\xd3\x71\x81\xde\xa1\xa6\x10\x15\x72\xd7\x2a\xca\x21\x63\x0f\x59\x0b\xaf\xbd\xb2\x67\xa1\x51\xc4\xfd\xed\x27\x2f\xbb\x84\xd0\x36\x3f\xb1\xf6\xf7\xcd\xf8\xf3\xf8\x9b\xe1\xe8\x3f\xe9\x5c\xdb\x74\x2e\x42\xc6\x8e\x60\x3a\xee\x8f\xa1\x8f\x85\x47\x29\x08\xd5\x09\xac\xb5\x31\x30\x47\xf0\x98\xbb\x15\x2a\x08\xce\x59\xc6\x02\x12\x70\x07\xe8\x3d\x7e\xd3\xd4\x2c\x0b\x5d\xe0\x42\x68\xd3\xac\xad\x2b\x6d\x40\x62\x6c\x51\x5a\x49\xda\x59\xb8\x43\x9a\xe5\xe2\xdb\xac\x70\x2a\x1c\xb7\xe1\x81\x01\xac\x33\x6d\x62\x7a\xa1\x40\xdb\x40\xc2\x4a\x9c\xd1\x7d\x81\x10\x39\xef\x40\x39\x06\x00\xa0\x17\x00\x37\x37\x90\xb4\x1e\x9e\x90\x36\x09\x74\xbb\x71\xf7\x74\x93\xc0\xed\x2d\xbc\x78\x51\xb3\x62\x70\x04\xff\x81\xbf\x6f\x5e\xf2\xb7\xb7\xff\x6b\x45\xf8\x1d\x50\x86\xb6\x4a\x08\x80\x32\x73\x50\x33\xdf\xd5\x7b\x1e\xa9\xf4\x5b\xc2\x42\x33\x00\xe5\x2c\xc2\x7b\x48\x91\x64\x8a\xcb\x20\xc9\xa4\x8d\xfc\x4e\x2e\x0a\xb6\x61\xec\x08\xae\x03\xc2\xf0\x53\x7f\xb2\x3a\x03\x72\xf1\x8a\x90\x23\x09\x25\x48\xb0\xe9\xf8\xe3\x60\xd4\x4d\x5a\xc7\xb2\xf4\x06\x38\x0f\xda\xa0\x25\xe0\x5f\xe0\xea\x7a\x0a\xfc\x0f\x48\xbe\x70\xb1\x0e\x1c\xe5\x19\x6f\x82\x38\xb9\x25\x5a\x4e\x64\x78\x40\xe9\xac\x0a\xe7\xf0\xe6\xe5\xcb\x04\x32\xa2\xe2\x3c\x4d\x4f\xdf\xbc\xed\x9c\xfd\xff\x75\xa7\xfe\x4d\x8d\x20\x0c\x94\x8a\x42\xa7\x55\x64\x3b\x39\xa8\x77\x9d\xb7\xae\xf7\x81\x92\x67\x24\x9c\x43\xab\xd2\x9f\x40\xf2\xfc\xd1\x31\x8c\xc7\xb8\xb4\x75\x9a\xc4\x9a\x8c\xc6\xfd\xc1\x6c\x78\x15\x2f\xbe\xaf\x00\x8c\x93\xc2\x70\x5d\xac\x5e\xb7\x13\x36\x1c\x4d\xa6\xbd\xd1\xc5\x60\x36\xec\x7f\x47\x6c\x1e\x99\x6b\xb5\xcf\x9c\x7e\xbd\x1a\xfc\x9c\x1b\x0d\xd1\x4e\x58\xef\xaf\xc9\x6c\x32\xf8\xfc\xe7\xf0\x62\x30\x99\xf5\xc7\x9f\x7a\xc3\xd1\x77\x31\x01\xfd\x4a\x4b\x0c\xa9\x72\xb9\xd0\xb1\x64\x2c\xb8\xd2\x4b\x7c\xf2\xd4\xcb\x72\x8e\x06\xa9\x83\x76\x05\x47\x40\x99\x0e\x20\x85\x05\xb7\x42\xef\xb5\x42\xf8\xd4\xfb\x32\xbb\x1a\xf7\x27\x8c\x3d\x4a\xbc\x1c\x7e\x18\x5c\x7c\xbd\xb8\x7c\x46\xa7\xd1\x0b\xe4\xf2\x5e\x9a\xa8\x96\x49\x41\xf0\xcb\x0f\x8f\xad\xaa\x55\x1d\xfe\xfe\xfd\x60\xfc\x61\x57\xd5\xd6\x43\xfd\xb5\x79\x52\xc3\xd6\xc3\xde\x6a\x73\x50\xb4\x3d\x30\xae\x37\xac\xd1\xde\x6d\x3d\x34\x9f\xe7\xbc\x56\x5c\x3b\x3c\xf6\xd1\xd3\xa8\xa4\xbd\x61\x51\x09\x0b\x56\x14\x20\x8c\x16\x01\x6a\xb5\x1c\x97\xa1\x53\x7f\x37\x7b\x87\x34\x49\x66\x47\x93\x64\x9a\xbd\x2d\x2d\x90\x2b\xf6\x93\xb1\x70\x1f\x08\xf3\xc8\xf3\x18\x90\x78\x9c\x2c\xa8\x18\x3b\x66\x00\xdb\x41\x75\x1e\xdb\x39\x20\x84\xcc\x95\x46\xc5\x29\x65\x9c\x5b\xa2\x02\x41\x80\x2b\xf4\xf7\x40\x3a\xc7\x26\x29\x04\x12\x9e\x02\x94\xc5\x49\x95\x61\x9d\x69\x99\x81\x0e\xb0\xce\x04\xc1\x1a\x41\x39\xd0\x16\x7a\x97\x67\x70\xbc\xc3\xe6\x22\xa0\x02\x67\xa1\x30\x42\x5b\xd8\x6a\x52\xdb\x04\xc2\x2a\xc8\x51\x58\x8a\x6d\x3f\x8f\x03\xcb\x93\x98\x1b\x8c\xcb\xdc\x05\x6a\xd8\xa0\x74\x20\xef\x42\xfb\x04\xe6\x25\x81\xa6\xff\x86\x2a\xde\x3a\x02\x69\x50\x78\xc8\xdc\x3a\x06\x19\x27\x54\x7d\xa5\x85\x77\xf9\xa3\xf0\x58\x9f\xb5\xa6\xcc\x95\x04\x99\x58\x69\x7b\x57\x25\x20\x07\xb2\x0c\xe4\x72\x1d\x30\xc6\x6d\x89\x9a\x02\x9a\x05\x03\x78\xc6\xd1\x3b\x6b\x3d\x4f\xfb\x29\xa1\x31\x75\xc5\x60\x00\x0b\x23\xee\x42\xf7\xb8\x1a\x9c\x89\x75\x0a\xb9\x2e\xf6\x7c\x9a\x6c\x81\x5c\x7c\xe3\xd1\x58\x7b\x9e\x6b\xa0\x2a\xc6\x88\x39\x9a\xd0\xc4\x5d\xf6\x7e\x1b\x5c\x4e\x36\x27\xc2\x14\x99\xe8\x6c\x0f\xee\x68\x97\xee\xcd\x86\x03\xcf\x9f\x6c\xb3\xe8\x05\x56\xdd\xb5\x8f\xee\xda\xb2\x39\xb0\x70\x8a\x6b\xbb\xf0\x82\x4b\x67\x49\x68\x8b\x9e\xeb\x5c\xdc\xc5\xa8\x38\x41\x06\x1f\x27\xb3\xc1\xc5\xe7\x59\xef\xe2\x62\x7c\x3d\x9a\x6e\x3a\x6a\xe9\x3b\x28\x7d\x67\x0b\xf7\x07\x1f\x7a\xd7\x97\xd3\xd9\xe7\xc1\xef\xc3\xf1\x68\x53\xef\x1e\x8c\x9d\x4d\x2c\x57\x5a\x88\x32\xe0\xf9\xab\xce\xab\xe8\xea\x79\xa9\x8d\xea\x9c\xd6\x22\xa4\x71\xa5\xe2\x85\x77\x2b\xad\xd0\x77\xc5\x3a\x34\x80\xd5\x7c\xae\x2d\x57\xda\x77\x53\x57\x50\x2a\xad\x8e\x7f\xd3\x7b\xb0\x74\x76\xb1\xc5\xe3\xbb\x44\xdc\x22\x75\x54\xc3\xd8\x5d\xca\x97\x36\x76\x41\x57\x39\xb9\x44\xdf\x94\x1b\x69\xed\xfc\x92\x17\xa6\xbc\xd3\xb6\x2b\xad\xae\x01\x8f\x77\x3a\x10\x7a\x1e\x4b\xd9\x25\x5f\xe2\x21\x10\x7d\xc8\x63\x6e\xda\xbd\xd4\xb4\x37\x1c\x4d\x77\x4f\x59\x75\xb5\xb3\x0b\x7d\xd7\x3d\xf4\xd4\x76\xbb\x73\x2f\x72\xf3\xa8\xf3\x47\xc4\x68\xbe\x86\xd5\x8e\x06\xdb\x8e\x88\xc7\xd1\x12\x6b\x19\xe7\x53\x65\xbc\x9b\x5f\x6f\x37\x09\x6b\xb3\x66\x90\x08\xff\x84\xc7\xfe\x0d\x00\x00\xff\xff\x76\x63\xcb\x48\xe3\x08\x00\x00")

func bindataAssetsBootstrapLegacyUbuntuShBytes() ([]byte, error) {
	return bindataRead(
		_bindataAssetsBootstrapLegacyUbuntuSh,
		"bindata/assets/bootstrap.legacy.ubuntu.sh",
	)
}

func bindataAssetsBootstrapLegacyUbuntuSh() (*asset, error) {
	bytes, err := bindataAssetsBootstrapLegacyUbuntuShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bindata/assets/bootstrap.legacy.ubuntu.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x38, 0xdc, 0xce, 0x7f, 0x6d, 0x30, 0x36, 0x5b, 0x8b, 0xc2, 0x38, 0x8d, 0x9e, 0xe4, 0xf8, 0x27, 0xd5, 0x77, 0x77, 0x8b, 0xaf, 0x59, 0xb3, 0x1a, 0xcd, 0x96, 0xe1, 0xcd, 0xc6, 0x2b, 0x19, 0x6b}}
	return a, nil
}

var _bindataAssetsBootstrapUbuntuSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x91\xdf\x6b\xdb\x30\x10\xc7\xdf\xf5\x57\xdc\xb4\x82\x61\x20\x6b\x7b\xdd\xc3\x20\xeb\xbc\x52\xd6\x66\x23\x71\xa1\x90\x05\x23\x2b\x17\x5b\x89\x23\x69\x3a\x39\x04\x82\xff\xf7\x21\x86\xb3\x34\xc9\xe3\xdd\xf7\xc7\x7d\x6c\xbd\x7f\x27\x6b\x63\x65\xad\xa8\x65\x8c\x30\x82\x70\x80\x21\xe0\xc1\xc4\x71\xf4\xc6\xe3\x5a\x99\x6e\x9c\xad\xeb\x2d\x61\x64\x8c\x5c\x1f\x34\x82\xdc\xab\x20\x3b\x53\x4b\xdd\xb9\x7e\x25\x49\x07\xe3\x23\x49\xdc\x92\x8e\x9d\xac\x9d\x8b\x14\x83\xf2\x79\x8b\x9d\xc7\x90\xa7\x43\xa8\x5b\x07\xfc\x9f\xe3\x33\x84\xde\x5a\x63\x1b\x90\x18\x75\x8a\xfd\xcf\x70\x76\xbd\xcb\xa9\x05\x7e\x77\xbc\x7f\x7a\x99\x97\xc5\xac\x9a\x4e\x9e\x8b\x81\xc3\x6f\x06\x20\xc4\xca\x92\xd0\x5d\x4f\x11\x83\x30\xfe\xdc\xf6\x6d\x3a\x3f\xb9\xb6\x7d\x8d\x1d\x46\x81\x87\x18\x94\x50\xa1\xa1\xe4\xfc\xf1\xf2\xb5\x78\x2a\xca\xaa\x78\x2d\x67\x93\x6a\x32\x7b\x98\x0f\xfc\x92\x74\x87\xa1\x49\xa4\x3d\x61\x00\xe7\xa3\x71\x96\xc0\xd8\xe8\x60\xec\xd4\xce\xae\x4d\x93\x6f\xc8\x59\xce\x12\x2e\x64\x61\x07\x62\x0d\x77\xc7\xf2\xf9\x57\x95\x8e\x54\xf7\x3f\xa7\xdf\x87\x0c\x8a\xd7\xc7\x92\x6d\xfe\x80\x20\xc8\xf2\xc5\xc7\x25\x7c\x80\x7c\xf1\x69\x99\x9d\xc3\x24\xeb\xe3\xc3\xc0\xaf\x01\x4f\xca\x97\xa4\xbd\xed\xe6\x6c\xb7\xbf\xb1\xbd\x55\x7c\xf5\x16\x48\x51\x85\x98\x3e\xf2\xf4\x9b\xb6\xc4\x19\x59\xe5\x47\xf1\x5c\xb9\x88\xaf\x9c\x45\xce\xfe\x06\x00\x00\xff\xff\xad\x08\x49\xa5\x55\x02\x00\x00")

func bindataAssetsBootstrapUbuntuShBytes() ([]byte, error) {
	return bindataRead(
		_bindataAssetsBootstrapUbuntuSh,
		"bindata/assets/bootstrap.ubuntu.sh",
	)
}

func bindataAssetsBootstrapUbuntuSh() (*asset, error) {
	bytes, err := bindataAssetsBootstrapUbuntuShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bindata/assets/bootstrap.ubuntu.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x6d, 0x94, 0xc2, 0xe1, 0x4f, 0x4c, 0xb5, 0x4d, 0x57, 0x4a, 0xa2, 0x51, 0x56, 0xf7, 0x8a, 0x83, 0xfe, 0xe7, 0x9c, 0xf3, 0x3f, 0x7b, 0xb6, 0x17, 0xc1, 0x29, 0x5b, 0x9c, 0x33, 0x76, 0xed, 0x28}}
	return a, nil
}

var _bindataAssetsEfaAl2Sh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\xd0\x4d\x6a\xc3\x40\x0c\x05\xe0\xbd\x4e\xa1\xd2\xb5\x46\x25\xdd\x15\xba\xea\x01\x7a\x84\xa0\xa4\x1a\x7b\xc0\xf3\xd3\x91\x8c\x93\x9c\xbe\x38\x4d\xe8\xc2\x34\x9b\x81\x37\x7c\x12\x3c\x3d\x3f\xf1\x21\x15\x3e\x88\x8d\x00\xa6\x8e\x54\x51\x7b\xd7\x53\xf2\x7b\x6c\xa9\x69\x94\x34\xdd\x73\xa9\x73\x31\x75\x80\xf3\x9c\x31\x15\x73\x99\x26\xa4\x33\x2e\x83\x3a\xac\x0f\xd2\x37\x12\x79\xca\x5a\x67\x7f\xdf\xbd\xe0\xe8\xde\xec\x8d\xd9\x5e\x69\x36\x5a\xd4\x9c\x76\x41\xb2\x5c\x6a\x91\xc5\xc2\xb1\x66\x96\xc5\x48\xa3\xd0\x6d\x9f\xf6\xed\x0f\x4d\xe2\x6a\x1e\x5c\x7a\x18\x2e\x48\x9f\xc8\x9e\xdb\xd6\xdd\x00\xb8\x74\xa4\x53\x7c\xac\x90\x3e\xae\x00\x8e\x5f\xff\x40\x08\xac\x51\xf6\x7f\x83\x36\xae\x6d\x69\x00\xae\xcd\xf9\xb7\xc6\x4a\xae\x87\x8c\x69\x9f\x4a\xac\x48\x0d\x35\x0a\xfc\x04\x00\x00\xff\xff\x8f\x52\xee\x9a\x5f\x01\x00\x00")

func bindataAssetsEfaAl2ShBytes() ([]byte, error) {
	return bindataRead(
		_bindataAssetsEfaAl2Sh,
		"bindata/assets/efa.al2.sh",
	)
}

func bindataAssetsEfaAl2Sh() (*asset, error) {
	bytes, err := bindataAssetsEfaAl2ShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bindata/assets/efa.al2.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x8f, 0xd9, 0x1d, 0x49, 0x21, 0x82, 0x43, 0xb0, 0xeb, 0x4d, 0xa2, 0x59, 0x42, 0xc6, 0xba, 0xc9, 0xb5, 0x60, 0xcd, 0x1d, 0x96, 0x18, 0x48, 0x80, 0x1b, 0x66, 0xda, 0x3e, 0x7, 0x93, 0xdb, 0x9e}}
	return a, nil
}

var _bindataAssetsEfaManagedBoothook = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\x4d\x4e\xc4\x30\x0c\x85\xf7\x39\x85\x2f\xe0\x04\x0d\x3b\x24\x56\x1c\x80\x23\x54\xa6\xe3\xb4\x91\x9a\x1f\x62\x47\x65\xe6\xf4\x28\xed\x88\x0d\xad\x60\x67\xf9\x3d\x7f\xcf\x7a\xe3\x92\xdb\x15\x43\x0a\x8a\x85\x2b\xe4\x34\x32\xdc\x5a\x1c\xd6\x89\xb5\x0f\x10\x92\x28\x2d\x0b\xe0\x0d\xfa\xce\x1c\x1d\x74\x61\x60\x4f\xdb\x00\xf8\x09\x88\x1a\x22\xe7\xa6\xaf\x97\x27\x98\x55\x8b\xbc\x38\x27\xcf\xd8\x04\x57\x16\xc5\x8b\xa5\x48\xf7\x9c\x68\x15\x3b\xe6\xe8\x68\x15\x64\x4f\xf8\x08\xe3\xfa\x7b\x83\x0b\x29\x8b\x5a\xa5\x6a\xa7\x3b\xe0\x3b\x38\x8d\xe5\x2f\x9f\x39\xfc\x57\xa9\x6e\xef\x2a\x55\xc0\x2f\xff\x2f\x12\xe0\xdb\xe6\x33\xa5\xc9\x7c\x3d\x39\x39\x4c\x7b\xa8\x5b\xa2\x75\xec\x69\xf8\xf1\x5b\x99\x7b\xb1\x38\x99\x92\xcb\x19\xf3\x10\xba\x63\x7c\x06\x97\x8b\xba\xbd\xcd\xce\x76\x1f\x21\x39\x1f\x76\x0d\x4b\xf7\x99\xef\x00\x00\x00\xff\xff\xd7\xd4\x31\xc2\xe4\x01\x00\x00")

func bindataAssetsEfaManagedBoothookBytes() ([]byte, error) {
	return bindataRead(
		_bindataAssetsEfaManagedBoothook,
		"bindata/assets/efa.managed.boothook",
	)
}

func bindataAssetsEfaManagedBoothook() (*asset, error) {
	bytes, err := bindataAssetsEfaManagedBoothookBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bindata/assets/efa.managed.boothook", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa6, 0x40, 0xf1, 0xae, 0x67, 0xa7, 0xe, 0xc, 0x31, 0x61, 0x86, 0x41, 0x8b, 0xd1, 0x55, 0x3f, 0x84, 0xbd, 0x4c, 0xdd, 0x84, 0xc9, 0xee, 0x9d, 0x36, 0xda, 0x8f, 0x6c, 0xbc, 0x6a, 0xaf, 0xed}}
	return a, nil
}

var _bindataAssetsInstallSsmAl2Sh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\xca\x41\x0a\x02\x31\x0c\x05\xd0\x7d\x4f\x11\x71\x5d\xe6\x4c\xa9\x44\x0d\xa4\xe9\xd0\xff\x07\xac\xa7\x77\x35\x2b\x61\x96\x0f\xde\xfd\xb6\x35\xcf\xad\x29\xde\xa5\xc0\x28\x75\x88\xcd\x69\x1f\xe7\xc9\xdd\x77\x7b\xaa\xc7\xe9\x1c\x47\xc2\x58\xca\x3a\xba\x78\x82\x1a\x21\x75\x89\x76\xfd\x8e\xac\x40\xaf\xfa\xb2\x64\xc1\x02\xad\x3f\x18\x62\xa9\x2d\xec\x6a\x80\x3a\xf9\x1f\x7e\x01\x00\x00\xff\xff\x93\x2c\xf6\x43\x9f\x00\x00\x00")

func bindataAssetsInstallSsmAl2ShBytes() ([]byte, error) {
	return bindataRead(
		_bindataAssetsInstallSsmAl2Sh,
		"bindata/assets/install-ssm.al2.sh",
	)
}

func bindataAssetsInstallSsmAl2Sh() (*asset, error) {
	bytes, err := bindataAssetsInstallSsmAl2ShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bindata/assets/install-ssm.al2.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x18, 0xf9, 0xf8, 0x5e, 0xb9, 0xcf, 0xfb, 0x94, 0xb4, 0x85, 0xa3, 0x62, 0xf0, 0x3b, 0x88, 0x44, 0xe3, 0x84, 0xfa, 0x85, 0x39, 0xdb, 0xed, 0xa2, 0x6a, 0x1, 0x7b, 0xe5, 0x49, 0x21, 0xef, 0x3d}}
	return a, nil
}

var _bindataAssetsKubeletYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x91\xc1\xca\xdb\x30\x10\x84\xef\x7a\x8a\x85\x5e\x8b\xed\xb4\x04\x5a\xdd\xd2\x84\xf6\xd0\x40\xa0\x71\xdb\xf3\x5a\x1a\x37\xc2\xb2\x36\x48\xeb\xa4\xed\xd3\x97\xd8\xee\x0f\x81\x1f\x9d\x86\x99\x61\x3e\x49\x6f\xa8\x3d\x1d\x4e\x74\xc0\x35\xc3\xb1\xc2\xbf\xa5\x7b\x88\x91\x3a\x50\xc6\x28\x37\x78\x2a\x22\xc9\x0c\x21\x79\x4b\x5f\xa7\x0e\x11\xba\x97\xd4\x87\x5f\x53\x66\x0d\x92\x0c\x5f\xc3\x0f\xe4\x12\x24\x59\x1a\x96\x40\xe5\xe6\x44\x35\x7c\x28\x55\x90\xfa\xb6\xe9\xa0\xbc\x31\x86\xbd\xcf\x28\xc5\x52\x53\xcd\xc7\xb8\x38\x15\x45\x3e\xc8\xc8\x21\x59\x5a\x65\x15\xc5\x71\x34\x86\x27\xbd\x20\x69\x70\xf3\x90\x35\x44\x9c\x24\xfd\x19\x65\x2a\x0f\x41\x84\xc4\x5d\x84\xb7\xd4\x73\x2c\x30\x44\x77\x74\x17\x91\x61\x71\x1d\xbb\x0b\xda\xf6\x68\xe9\xdd\xd8\x94\xe7\x82\xe6\xe9\x91\xff\xbd\x6d\x3e\xae\xe1\x18\x90\x74\xbf\xfb\x1c\x22\x2c\xd5\x50\x57\x63\x28\x4e\x63\xed\xb8\x72\x59\x17\x1a\xc9\xe1\xef\x0b\xcc\x28\x1e\x96\x7e\x2e\x93\xaf\x8e\xef\xd6\x0a\xfc\x8c\xb1\xfd\x8f\x31\x9b\xdf\x13\x3f\xdb\xef\x9b\x62\x4c\x41\xbe\x21\xb7\xc7\xf3\x27\x11\x2d\x9a\xf9\xba\xc2\x9a\x1e\xac\x53\xc6\x17\x56\xcc\xd7\xff\x26\xca\x8a\xf5\x4b\xce\x73\x6d\x8f\xac\xa1\x7f\xbc\x17\xd6\xd6\xbf\x00\x00\x00\xff\xff\x46\x42\xbb\xf2\xe0\x01\x00\x00")

func bindataAssetsKubeletYamlBytes() ([]byte, error) {
	return bindataRead(
		_bindataAssetsKubeletYaml,
		"bindata/assets/kubelet.yaml",
	)
}

func bindataAssetsKubeletYaml() (*asset, error) {
	bytes, err := bindataAssetsKubeletYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bindata/assets/kubelet.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x90, 0x31, 0x7a, 0x5b, 0xa3, 0xbf, 0x93, 0x44, 0x7, 0x84, 0xfb, 0x4e, 0x10, 0x44, 0x47, 0xaa, 0xd1, 0x9, 0x71, 0x33, 0xdc, 0x75, 0x6b, 0x7c, 0x1c, 0x25, 0x25, 0xfa, 0x82, 0x6d, 0xc4, 0xf}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"bindata/assets/10-eksctl.al2.conf":         bindataAssets10EksctlAl2Conf,
	"bindata/assets/bootstrap.al2.sh":           bindataAssetsBootstrapAl2Sh,
	"bindata/assets/bootstrap.helper.sh":        bindataAssetsBootstrapHelperSh,
	"bindata/assets/bootstrap.legacy.al2.sh":    bindataAssetsBootstrapLegacyAl2Sh,
	"bindata/assets/bootstrap.legacy.ubuntu.sh": bindataAssetsBootstrapLegacyUbuntuSh,
	"bindata/assets/bootstrap.ubuntu.sh":        bindataAssetsBootstrapUbuntuSh,
	"bindata/assets/efa.al2.sh":                 bindataAssetsEfaAl2Sh,
	"bindata/assets/efa.managed.boothook":       bindataAssetsEfaManagedBoothook,
	"bindata/assets/install-ssm.al2.sh":         bindataAssetsInstallSsmAl2Sh,
	"bindata/assets/kubelet.yaml":               bindataAssetsKubeletYaml,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"bindata": {nil, map[string]*bintree{
		"assets": {nil, map[string]*bintree{
			"10-eksctl.al2.conf": {bindataAssets10EksctlAl2Conf, map[string]*bintree{}},
			"bootstrap.al2.sh": {bindataAssetsBootstrapAl2Sh, map[string]*bintree{}},
			"bootstrap.helper.sh": {bindataAssetsBootstrapHelperSh, map[string]*bintree{}},
			"bootstrap.legacy.al2.sh": {bindataAssetsBootstrapLegacyAl2Sh, map[string]*bintree{}},
			"bootstrap.legacy.ubuntu.sh": {bindataAssetsBootstrapLegacyUbuntuSh, map[string]*bintree{}},
			"bootstrap.ubuntu.sh": {bindataAssetsBootstrapUbuntuSh, map[string]*bintree{}},
			"efa.al2.sh": {bindataAssetsEfaAl2Sh, map[string]*bintree{}},
			"efa.managed.boothook": {bindataAssetsEfaManagedBoothook, map[string]*bintree{}},
			"install-ssm.al2.sh": {bindataAssetsInstallSsmAl2Sh, map[string]*bintree{}},
			"kubelet.yaml": {bindataAssetsKubeletYaml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}