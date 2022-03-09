// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../../../contracts/LeofyCoin.cdc (6.821kB)
// ../../../contracts/LeofyNFT.cdc (17.637kB)

package assets

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

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	clErr := gz.Close()
	if clErr != nil {
		return nil, clErr
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

var _leofycoinCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x59\xdd\x6e\xdc\xbc\x11\xbd\xf7\x53\x4c\x7d\xd1\xae\x11\x7b\xed\x02\x45\x2f\x0c\x3b\x89\x93\x26\x40\x80\x36\x08\x92\xfe\x5c\x73\xa5\xd9\x15\x6b\x89\x14\x48\x6a\xd7\x1b\xc3\xef\xfe\x81\xc3\x1f\x91\x5c\x69\x6d\x23\x09\x02\x64\x25\x71\x86\xc3\xe1\x99\x33\x87\x0c\xef\x7a\xa9\x0c\x7c\x1e\xc4\x86\xaf\x5a\xfc\xb7\xbc\x47\x01\x6b\x25\x3b\x38\x5d\x5e\x6a\xc3\x44\xcd\x54\x7d\x99\x7d\x5e\x56\x75\x75\x7a\x72\xd2\x0f\x2b\xa8\xa4\x30\x8a\x55\x06\xfe\x89\x72\xbd\xff\x28\xb9\xb8\x2e\x5c\x3d\x9e\x00\x00\x5c\x5e\x02\x3d\xea\x2f\x82\x1b\xce\x5a\xfe\x13\x6b\xff\x21\x7e\x6f\x10\x70\x8b\xc2\x80\x69\x98\x01\xae\x01\x3b\x6e\x0c\xd6\xb0\x6b\x50\x80\x69\x70\x9c\x8d\x6b\xa8\x14\x32\xe3\x9d\xd8\x48\x9c\xe9\xc1\x24\x0b\xee\x7e\xff\x18\xfa\xbe\xdd\x5f\xc3\x7f\x3e\xf3\x87\xbf\xff\xed\xec\x24\x8f\xea\x7f\xdc\x34\xb5\x62\x3b\xf1\xca\x98\xc8\x18\x98\x42\xd8\x05\x0f\x2e\x77\x0c\xfe\xcb\x86\xd6\x4c\x46\x17\x27\x5b\xb0\x4e\x0e\xc2\x84\xa0\xce\xc9\xf4\x1a\xee\xea\x5a\xa1\xd6\xef\xca\x20\xff\x81\xbd\xd4\xdc\xbc\x3a\x71\x63\x90\x75\xf0\x00\x46\x1e\x0d\x31\x4e\x75\x10\xa2\x91\xf3\x01\xfe\x8b\x8b\x57\x47\x27\x70\x97\x46\xd8\x8d\x2e\xca\x98\x9c\xf7\x22\xa0\x32\x84\x0f\x83\x12\xbf\x94\x20\x6d\x94\xdc\xcf\x44\xe0\x9c\xcf\x46\x40\x01\xaa\x8f\x09\x2e\x5f\x1c\x02\xa3\x3c\xd0\xe2\x15\x28\xd4\x72\x50\x15\xce\xa3\x3c\x9b\x69\xc1\xda\x56\xee\xb0\xbe\x9b\x09\xeb\x2b\xeb\xb0\x86\x9e\x99\x46\xa7\x41\x59\x6f\x2d\x1a\x07\x82\x1f\x46\x2a\xb6\xc1\x6f\xcc\x34\xd7\x90\x3c\x64\x23\xbf\x63\x85\x7c\x8b\xea\xdb\xb0\x6a\x79\xe5\xc6\x8e\xbf\xb3\xa1\x1f\x58\xcb\x44\x85\x2f\x18\x79\x57\x77\x5c\xcc\x4e\x3f\x6e\xae\x61\x2d\x68\xaa\x61\x90\xeb\x91\x6d\x80\x0b\xc0\x07\xae\x0d\x8a\x0a\xa3\xe3\x2d\x53\x60\xac\x49\x5e\xf5\xd1\xdd\x08\xfc\x71\x87\x3e\xb1\xaa\x81\x41\xa3\x02\x6d\xa4\x42\x0d\xcc\x3a\xb7\x04\x58\xa1\x9d\x52\x8a\x76\x4f\x24\x44\xc6\x76\x5e\xd3\x20\x77\xa3\xd9\x06\xd3\x8d\x5e\x0f\xa2\x32\x5c\x0a\xed\x47\x79\x13\x26\x6a\xd8\xc8\x2d\x5a\x10\xc1\xca\x39\xeb\x15\xd2\xfb\x5e\x6a\x63\xe9\xad\xe6\x64\x18\xbc\x71\x51\xb0\x69\xa0\xc2\x3d\xc1\xb5\x62\x6d\x8b\xf5\x32\x9d\xbb\x6a\xb0\xba\xd7\xd0\xb0\xbe\xb7\xb8\x32\xa0\x06\x61\x78\x87\x64\x89\x5b\x54\xc0\x62\x7c\x04\xb0\xcc\x45\xf0\xf4\xdd\x43\xd0\x7e\x17\x6e\xe9\x2b\x0c\x60\x0c\xab\xb2\x6c\x8c\x0f\xc6\x26\x27\x23\x67\x42\xb8\x8d\x31\x78\x73\x95\xb5\xe6\x82\x6c\xcf\x41\x4b\xfb\x59\x11\xc0\x85\x84\x1d\xdb\xc3\x5a\xda\xc0\x3a\xd6\xf2\x8a\xcb\x41\xbb\x8d\x30\xd2\x4f\xe9\x12\x18\xb3\x22\x07\x3f\x29\x17\xc0\xb8\x5a\xc2\x1d\xe8\x1e\x2b\xce\x5a\x5f\x19\x63\x09\x09\xc4\x5a\x5b\x47\xab\x31\x04\x23\xa9\xd2\x82\xb7\x91\x81\x96\x65\x7d\x44\x37\x34\x7f\xd1\xd9\x96\xdf\x94\xdc\xf2\x1a\xd5\x79\xf1\x3e\xd4\x49\xf9\xde\x17\x05\x3c\x3a\x1c\x26\x9b\x46\x58\x85\x95\xff\xee\x96\xa6\x61\x1b\x51\x9a\xe2\xda\x8f\xca\x31\x1d\xe1\x12\xfa\x1e\xed\x48\x70\x68\x51\x10\x56\x42\x09\xb5\x7b\x6f\x41\x11\x6d\xad\xe1\xa2\xf0\x7c\xe6\x5b\x77\xf8\xa3\xb1\x5d\x2f\x83\xcb\xdb\xe0\x3c\x0e\x79\xca\x22\x09\xed\x30\x79\x97\x7e\xfe\x1c\x10\xe8\xb0\xc2\xee\x5d\xbd\x39\x66\x05\xe6\x1e\xd4\x66\xe8\x50\x98\xd4\xce\x96\x4a\x70\xad\x9d\xb1\xb7\xa1\xae\x1b\x6b\x6d\x39\x33\xef\x17\xe3\x11\xa5\x3d\xe5\x1a\xb4\xf2\x87\xa9\xbd\xaf\xd1\xc0\xce\x83\x76\x40\x69\x64\x5b\xa7\x0e\xec\x0c\x9d\x14\xb8\x8f\x23\x57\xc8\xc5\x06\x8c\x62\x42\xaf\x51\x29\xac\x97\x76\x16\x85\x66\x50\x42\xd3\x78\x81\xbb\x76\x9f\x3a\x09\x75\xe4\xa7\x94\x59\x35\x91\x5f\x57\x95\xb6\x50\xb8\xa1\x12\x5c\x25\xcd\x3b\x75\x85\xad\xc6\x9d\xad\xa5\xc9\x15\x5b\xcc\xac\x07\x11\x53\x56\xb6\xae\x6b\x78\x9f\x43\xd4\x45\x74\x74\xdf\xb3\xc7\x0b\x9f\xfe\xcc\xc0\x76\xb6\x59\xa9\xe3\xfe\x0d\x52\x87\x9c\xc9\x9d\x40\xf5\x6e\xc9\x9c\xaa\x38\xcb\x7c\xb9\x3c\xc2\xcd\x45\x4a\x04\x23\x52\x9d\xb7\xb3\x19\x10\xfa\x8c\xbd\x06\x83\x7e\x4f\xe4\xea\xff\x58\x95\x40\x24\xf4\xb1\xba\xd6\x59\xc9\x19\x1d\xeb\xcc\xef\x64\x56\xc7\x08\xb4\x3c\xfd\x2c\x2e\xb9\x06\xdf\xc8\xad\x1f\x2f\x44\xc8\x81\xb6\x73\xbb\xb8\x56\x58\xb1\x41\xe3\x08\xf4\xac\xe6\x6c\xb8\x09\xa2\x2d\x76\x51\x85\x30\x3c\xc5\x11\xdf\x90\xe9\x5f\xc6\xc0\x1b\x96\xad\x69\x85\x28\x2c\x1e\xf5\x60\x95\x83\x5d\x35\xd1\xf5\x5a\x52\xcb\xf1\x60\xf4\x42\xe9\x28\xee\xfc\x06\x2c\xdc\x66\x4f\x61\xad\x24\x19\x2b\x09\x88\xf7\xe0\xe6\xc2\x2b\x69\xfd\x27\x78\x1f\xfb\xfd\x32\x5f\xf5\x73\xf8\x7c\xe3\x9c\x2d\x4b\xb2\x2a\x60\x7a\x28\x77\x33\x33\xa7\x7a\x9f\xc5\x6a\x66\x03\xb7\x70\xb5\xbc\xca\xbe\x87\x3d\xcd\x79\x3d\x81\xac\x1f\xb0\x28\x93\x32\xae\x3e\x11\x35\x70\x3b\xf3\xfe\x22\x4b\x41\xe6\x88\xaf\x17\x59\x7a\xde\xda\x18\xcb\xd9\x8a\xdc\x14\xa2\x37\xb5\xcf\x97\xff\x74\x92\xff\x7a\x8a\x72\xcb\xd5\xee\xa7\xae\x37\xfb\x29\xe5\x95\x17\x62\xce\xcf\x0e\xf6\x96\xbf\x80\xa5\x85\xf5\x13\x95\x8c\xf2\x42\xd4\x91\x6f\xf9\xc8\xa7\xac\x6d\x2d\x33\x7b\x5e\xb5\x22\x81\x44\x45\x37\x68\xc7\xaf\xae\xc9\x06\x2d\x94\x3a\x23\x01\x48\x4e\x9c\xdb\x48\xd5\xa5\xe8\xb3\x2f\xa4\xaa\x9d\x54\xa1\xea\x75\xdf\xa3\xb3\xaa\xa2\xe6\xe4\xf4\x07\x5b\xb5\xc4\x11\xca\xe9\x83\x50\x1d\x3a\x36\x7c\x2a\x51\x30\xfb\x1e\x0f\x94\x88\xad\xa6\x32\x8d\x0b\xcb\xdf\x25\x63\x3f\x43\x98\x76\xbf\xd3\xdd\xc9\x54\x0e\x69\x71\xae\x8d\x62\x46\xaa\x42\xa6\x38\x77\x5f\x71\xe7\x14\xd6\x8b\x38\x35\x6e\x65\xb2\x41\x93\x27\x9d\x63\x1c\x52\x4c\x3c\x73\xda\xb9\x86\xf7\x5e\xfa\x3d\x1e\x16\xf9\xd1\xe3\x52\xf6\x78\xbc\xf9\x4c\x47\x30\xe3\xa0\x2c\x82\x24\x6f\x87\x3a\x3b\x34\x1c\xd7\x8a\x08\x06\xcc\x6e\x46\x40\x90\xd3\xe1\x96\xd0\x83\x7a\x7d\x91\x6c\x8d\x29\x29\x05\xa7\x17\x4d\x16\x78\xee\xf0\x1b\x34\x7b\xd8\x99\xbc\x1b\x45\xb9\x0c\x89\x10\x9d\xdc\x88\x6c\x26\x6b\xe6\x48\xe4\x45\x70\xb1\xc3\x75\xb2\xae\x73\x6a\xb6\x36\xaa\x2e\x14\xb5\x49\xce\x81\xe7\xa5\x30\x4c\x34\x57\x37\xc7\x02\xc7\x90\x36\x86\x3b\xa1\x93\x8a\x0e\x54\xc0\xcc\x1e\xe3\x0e\x59\xd4\x67\x99\x58\xf6\x1a\x4e\x5d\xa6\xfc\x15\x87\xe3\xa1\x15\xc2\x86\xb0\xa5\x6c\x0a\x04\xd1\xda\xe9\x9c\x9f\x1b\xdf\xdc\x8a\xc4\xcf\xf8\x6d\x51\x6b\xe7\xd4\xe6\x21\x6c\xa6\x73\x75\x3a\x43\xdc\xf0\xea\x66\xf3\x66\x4a\x00\x1e\x46\x09\x53\xa1\x3f\xab\x1e\x8b\x1b\x9f\x52\xec\xc1\x2f\xe9\x43\x3a\xf2\x4c\x93\xc9\x94\x00\x2e\x97\x93\x3d\x4f\x17\x3d\xcd\x90\x3a\x23\x47\xe5\x7d\x0b\xdc\xc2\xa5\x6f\x27\x97\x31\xc7\xb9\xca\x21\xbb\xc3\xdb\x17\x6b\xd9\xd3\xd3\x68\x18\x46\xe5\xb6\x07\xd7\x31\x53\xa6\x1f\x0a\xc1\x40\x96\xe5\xf5\xcc\x64\xb4\x34\x88\xec\x32\x4d\x3b\x1e\x46\xe3\xd5\x80\x36\xcc\x24\x47\x15\x9a\x22\x87\xd9\x5f\xaf\xae\xac\x6e\x4a\x1d\x39\xe2\x06\x06\x2e\x62\xa8\x58\xcf\x56\xbc\xe5\x66\x1f\xaa\x3c\x39\xbb\xd1\x5d\x05\x3e\xf4\x52\xa3\x2e\x0f\x6e\xbe\xe1\x8e\xf7\x1f\xa6\x51\x72\xd8\x34\xf4\x31\x64\x0e\x88\x00\xd7\xac\xcc\x84\x67\xe2\xa5\x66\x5b\x5c\xcc\xe1\xad\x5c\xd0\x99\x93\x8e\x05\x79\x24\xf9\x3c\x9b\x9e\xa4\xe5\xe2\xfe\xe6\xcf\x85\xd9\xe3\xf4\x35\xc3\xd3\xdb\xc5\x4c\x01\x1f\x42\xe6\x3c\x1b\x69\x98\xda\xa0\x39\x16\x5e\x1c\x7e\xf6\x9b\x77\x24\xd9\x95\xa0\xea\xd6\x1c\x6d\x7b\x4b\xb6\x24\x5c\x98\x3c\xb3\x23\x2f\x49\x96\x77\x35\x9f\xab\x83\x12\xf9\xbd\xa9\xa2\x83\xa0\x40\x30\x6a\xf0\x3a\x2b\x9e\x31\x45\x1d\x81\x69\xff\x8a\xe4\x3e\x80\x6a\x26\x2c\x33\x7a\xb5\x67\x24\xa7\x0e\x6e\x2e\xc2\x05\x59\x26\xdd\x16\x33\xa8\xf2\xd0\x25\xdb\xe4\x4c\x53\xd6\x78\xbe\x80\x4f\x96\x8f\x99\x48\x6f\xaf\x75\x23\x77\x89\x64\x88\x91\xda\x03\xe8\x78\x05\x95\x1c\x0d\x13\x4e\x3f\xf2\x5f\x33\x07\xc5\xe3\xc9\xf4\xe9\x04\xfe\x08\x00\x00\xff\xff\xa0\xb3\x88\x73\xa5\x1a\x00\x00"

func leofycoinCdcBytes() ([]byte, error) {
	return bindataRead(
		_leofycoinCdc,
		"LeofyCoin.cdc",
	)
}

func leofycoinCdc() (*asset, error) {
	bytes, err := leofycoinCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "LeofyCoin.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x24, 0x49, 0x6d, 0xb1, 0x4, 0xb3, 0x2d, 0xfe, 0x87, 0x3, 0x14, 0x50, 0x1e, 0x80, 0x2c, 0x5b, 0x49, 0x6c, 0xf0, 0x83, 0x1a, 0x47, 0x79, 0x1e, 0x14, 0x17, 0x46, 0xdf, 0xd9, 0x61, 0x8d, 0xfc}}
	return a, nil
}

var _leofynftCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\x5f\x73\x1b\x39\x72\x7f\x5e\x7f\x8a\x36\x1f\x6c\x72\x23\x51\xce\x9f\xca\x03\x4f\x5a\xdb\x91\xac\x44\x29\xaf\xce\xb5\xe2\xed\x55\xca\xe5\xca\x81\x33\x20\x89\xf3\x10\xe0\x02\x20\xb9\x3c\x9d\xbe\x7b\xaa\x1b\x03\x0c\x30\x83\xa1\xe8\x5d\x29\x4f\xc7\x07\x9b\xe2\x00\x8d\x46\xa3\xff\xfc\xba\xd1\x73\xf6\xfd\x0b\x00\x80\x2b\x6e\x0a\x2d\xd6\x56\x28\x39\x81\x4b\x2e\xad\x66\x15\xdc\xad\x98\xb6\x70\xa9\xf0\xaf\xc2\xc2\x5c\x69\xf8\xc8\xd5\x7c\xff\x82\xa6\x4c\x97\xc2\x80\xa1\x21\x85\x1f\x82\x5f\x98\x90\x06\xec\x92\x43\xa1\x34\x87\xf9\x46\x16\x48\x95\x55\xc2\xee\x89\x04\x4d\x26\x3a\x27\x50\x68\xce\x2c\x2f\x61\xb6\x87\x8f\x1f\xfe\x78\xfd\x3f\x70\x75\xf3\x9f\x37\xd3\xf7\x1f\xe1\x6e\xfc\x71\xec\x97\xe1\x0d\xfd\x15\x93\x6c\xc1\x1d\xf9\x92\x59\x06\xcc\x18\x55\x08\xa2\xb1\x13\x76\x09\xac\xaa\xe8\xa1\xb0\x7c\x65\x68\xbe\x5d\x32\x0b\x4c\x73\xd8\x18\x5e\x02\x33\x60\xf9\x6a\x5d\x31\xcb\x0d\x71\x83\x83\x6f\xaf\xa7\x26\x2c\x26\x81\x49\x78\x5f\xae\x84\x84\x82\xc9\x9a\x43\x90\x7c\x07\x37\x48\x73\xec\xfe\x43\x96\x8c\x30\x16\xd4\x1c\x18\xac\x37\xb3\x4a\x14\x60\xac\xde\x14\xd6\xad\x48\xe4\x82\x38\xea\x01\x42\xce\x95\x5e\x31\x94\x07\xb0\x99\xda\x58\x60\xc4\xe9\x09\x30\x59\x22\x1d\x2d\xb6\xb8\x9a\xe6\x46\x6d\x74\xe1\x78\x76\xbb\x50\xb0\x12\xd2\x12\x1f\xb7\xd7\xd3\xd7\x06\x2a\x21\xbf\xf2\x12\x1f\xe0\x16\x90\xa9\x48\x60\x8c\xf8\x0f\x64\x96\xcc\x89\x6c\xad\x76\x5c\xe3\x94\x52\x91\xa4\xd4\xdc\x09\x6b\xb5\x56\xda\x32\x69\x81\xd1\x59\x39\xc1\x09\x49\x0f\xd3\x13\x1e\xc3\x9f\x49\x44\x48\xdf\xc0\x0e\xe7\x58\x05\x05\x12\xf3\x27\x6d\x70\x26\x23\x8e\x4e\xea\x13\xe0\x7b\x37\xc4\x2e\xb9\xd0\x30\x53\x5a\xab\x1d\x3e\x0f\x73\x90\xc8\x82\xa3\x38\x34\x9f\x73\xcd\x65\xc1\xc1\xef\xdb\x89\xc8\xb3\xd3\x30\x82\x67\x75\xe2\x89\xcb\x36\x0f\x4a\x06\x3d\x80\x8d\x11\x72\xe1\x8e\x25\x90\x1f\x13\x79\xfa\x87\x76\x84\x3a\x40\x6a\x82\x62\xe6\x65\x4d\x18\x7f\x10\x52\x58\xc1\x2a\xf1\xb7\xa0\x63\xb4\xb7\x9b\x2b\x3c\x34\x22\x80\xa3\x34\xb7\x1b\x2d\x9d\x2a\xe3\xc2\x44\x46\xe7\x74\x98\x55\x46\x41\xc9\xe7\x42\x72\x03\x0c\x2e\x55\x55\x71\x27\x04\x7f\x5e\x63\x67\x5b\xc2\xa0\x2a\xaa\xd9\x5f\x79\xa2\x54\x7c\xcb\xf5\xde\x19\x10\x72\x0d\x6a\x27\xb9\x86\x9d\xa8\x2a\x30\x56\x11\xc3\xb5\xa0\x59\x51\xa8\x8d\xb4\x41\x81\xc8\x78\xea\x67\x38\xb3\x08\x4b\x47\x7c\xae\x98\x90\x35\xf5\x7a\xbe\xa3\x4d\x6c\x2f\xd9\x16\x65\x6a\x70\x51\xf8\x51\xad\xb8\xb4\x11\x95\x5a\x6f\x2c\x9d\xc7\xc6\x70\x5c\x74\xa9\xaa\x32\xcc\x58\xd1\x0c\xe3\x36\x43\xb4\xa4\xb2\xb0\xe7\x16\x66\x9c\x4b\x30\xdc\xa9\x13\xc3\xc9\x5e\x76\xb7\xca\xf2\x09\xbc\xa7\xdd\xa1\x6d\x14\x4b\x26\x17\x78\x9e\xcd\x59\x13\x7f\x6b\x26\xd1\xc2\xe6\x28\x33\x21\xb7\xac\x12\x25\x30\xbd\xd8\x10\x8f\xc2\xb1\xb6\xd6\x6a\x2b\x4a\x5e\x82\xd2\xa0\x24\x47\xfd\x47\xd6\xd6\x9a\x9f\x16\x4a\x96\xa2\x56\x1d\x0d\x6b\x65\x48\xe7\xfd\x4f\x4c\x73\xf9\xda\xc2\x8a\x5b\xa7\x37\xd7\x61\x6d\xda\x4a\xa9\xe8\xa9\x2a\xc5\x7c\x5f\xb3\xe9\xce\x43\xac\xd6\xd5\xbe\xd6\x0d\x78\x83\x94\xa5\xa8\xdc\x31\xa2\xc5\xdb\xa5\x32\x1c\x0a\x66\xb8\x01\xc9\x9d\x35\xcf\xd0\x5c\x65\x59\x35\xaa\x84\x9a\x4d\xe2\xf8\xfe\xec\xc5\x0b\x67\xab\x70\xab\xe4\xf5\x46\x2e\xc4\xac\xe2\x53\xf5\x95\x4b\x98\x6b\xb5\x82\xc1\xf8\xcc\x58\x26\x4b\xa6\xcb\xb3\xf6\x88\x71\x51\x16\x03\x3f\xfd\x47\x6e\x19\xfa\xcf\x9f\x05\xdf\x99\xee\xdc\xe4\x71\x32\xf1\x91\x45\xfb\x57\x24\x8d\xba\x54\xa2\x99\x14\x7e\x71\xe3\x5e\xac\x37\xb3\xc6\x48\xe8\xe1\xed\xf5\x74\xd2\xdd\xe9\xbd\xd3\x8b\xb3\x33\x38\x7d\x9a\x8f\x27\xe7\x94\x3e\xb0\xf0\x61\x8b\xba\xfa\xd4\x6b\x79\x7a\x1f\x56\xc2\x52\xd4\x42\xe7\x83\xa7\xec\x77\xdc\x30\x20\x8c\x8f\x8f\x4e\x79\x37\x33\xb4\x7d\xd9\x44\xe4\x9b\xc6\x33\x0d\x47\x79\xca\x2c\x44\x2e\x1f\x9d\x7a\xa9\xe2\xa0\x4b\xf7\x64\x28\xca\x09\xfc\xe9\x46\xda\x7f\xff\xb7\x13\x54\x7a\x52\x86\x09\xdc\xdf\x59\x2d\xe4\x62\xe2\xfe\x7b\x18\xb5\x08\xdc\x71\x9b\x9b\x2f\xd9\x8a\x4f\xc0\xcd\xa9\xb9\x6c\xe6\xfc\x59\xd8\x65\xa9\xd9\x2e\x99\x81\x1a\x32\x81\xf7\x65\xa9\xb9\x31\x6f\xdb\xcb\x5c\xf1\xb5\x32\xc2\x26\x33\xac\xea\x1f\xff\x23\x39\xf4\x64\xb8\x20\xff\xdd\xfc\x6d\xb8\x16\xac\xba\xdd\xac\x66\x5c\xbb\x5f\xff\xf5\x5f\x1a\x81\xde\xb2\x15\x2f\xe1\x13\xb3\x4b\xaf\x0d\x61\x85\x8a\x3b\xb9\xdd\x59\xa5\xd9\x82\xe3\x18\xdc\x6a\xf8\xa3\x33\xf0\x13\x01\x01\x37\xae\xf9\xfe\x22\x19\xd7\x44\x84\xa3\xc8\x36\xc3\x7b\x88\xc7\x83\x09\xd8\xf4\x92\x7d\x26\xc3\x9a\xaa\xf5\xdd\x52\x35\x20\xe2\xb4\xe2\x5b\x5e\xc1\x5c\xf0\xaa\x34\xe3\x30\x6a\xc9\x0d\xf7\x90\x09\x91\xc8\x86\x55\xb0\x65\xd5\x86\x9b\x06\xc4\x51\x84\x2b\x7b\xb0\xc9\x73\xd9\xea\xcf\x4c\x0b\x36\xab\x38\x18\xf1\x37\x0e\xa5\x70\x80\x56\xef\x31\x7e\x44\xa6\xe5\xb5\x83\x15\x05\x37\x66\x68\x78\x35\x1f\xc1\x96\x69\x07\x46\x27\xf0\xee\xde\xe9\xdb\x84\x26\x3d\x34\x87\x8e\x63\xac\xb2\xac\xba\xdb\xac\xd7\xd5\xde\xeb\x65\xf7\x39\xa9\x5a\x3a\xe6\x39\x7d\x61\xec\x8d\xea\x33\xbb\x54\x2b\xb2\x3e\x0e\xd3\xfd\x9a\x3b\x18\x23\x9a\xd8\xff\x0c\xba\x43\x5a\x81\x47\xff\xd7\x8d\xb1\xf0\x7d\xb4\xe2\xf7\x04\xe0\x91\x8f\x5a\x41\x2c\xa2\x26\xcf\xb0\x27\x80\x71\x56\xd9\x25\x0f\x70\xc8\x78\x78\x32\xae\x89\x47\x24\x11\x18\x4b\x15\xf2\x98\x40\xc3\x29\x63\xad\x7c\x4e\x27\x4f\x60\x86\xf8\x1d\xc1\x06\x86\xbf\x82\xc3\x10\x51\x05\xa1\xb5\x11\x6a\x46\x0d\x30\x2c\xad\x41\x5c\x7a\x72\xb8\xfe\x8c\xc7\xc9\x4f\xc2\xb9\xdb\x4c\xc8\x1d\x92\x55\x9f\x47\xc7\x3d\x51\x52\x66\x84\x9d\xf0\x93\x4f\x1d\x1c\x54\x53\x55\x69\x42\x20\x88\x93\x2e\x3f\xb3\x06\xc6\x66\xcd\x0b\x31\x17\x05\xbc\xd7\x16\xb3\x23\x4a\x02\xa0\x12\x5f\xb9\xcb\x3f\x44\x61\x37\x98\x14\x22\x04\xa8\x87\xfc\xb7\x5a\x4a\xb8\x52\x3c\x76\xad\x21\x18\x13\x2a\xaf\xe1\x67\x15\xe5\x06\x0c\x10\xd2\x57\x35\xbe\xaf\xb3\x1b\x07\x84\xd5\xdc\x93\x40\x68\xe7\x79\x1e\x7b\x4e\x1b\xab\x0b\xd9\x11\x21\xf5\x39\x2b\x5c\x16\xd5\x76\xa7\x70\x4f\x33\xfc\xac\xf9\x46\x62\xa6\x42\x49\xe0\x47\x2e\x17\x76\x39\x1c\x4d\xe0\xa6\x46\xd9\x99\x51\x88\xa5\xae\x98\x65\xff\xb1\xbf\x46\x7f\x37\x6c\xc7\x1e\xf2\x82\x21\x3e\xfa\x2f\x6f\x3b\xe4\x9a\xac\xa9\x45\x62\x34\x81\x57\xf8\xf3\x7d\x13\x5f\x1e\x68\xf6\x43\x66\xaf\xe9\x0e\x27\xc7\xec\x38\xeb\xc1\xbc\x7f\x18\x47\xae\xcc\x25\x8d\xc2\xc2\x70\x14\x91\xc0\x0f\xba\xc2\x31\xd1\x80\xf3\x53\xb8\x7f\x08\x0f\xa3\x99\x7e\x9f\xce\x2a\x68\x9f\x5d\xe4\x51\x4b\xe7\xe1\x04\x53\xe5\x82\x4f\xe0\x4f\xd7\xe2\x57\x12\x81\x63\xcc\x83\x43\xff\x39\x3b\x03\x07\x49\x48\x41\x10\x0a\x7d\xaa\xd8\x3e\x19\x82\xdb\x93\xdc\xa5\xa3\xe7\xa7\x3e\xdf\xa7\xf5\x93\x71\x00\x11\x12\xf2\xdf\x4e\xda\x43\x6a\xb6\xe8\xbf\xe4\xd9\x28\xf9\x2b\xf9\x03\xe3\x32\x72\x70\x05\x17\x9e\x93\xb1\x3b\xe2\xce\x66\xee\x5c\x82\x67\xdb\xd9\x30\x39\x09\xb6\xe0\x3d\x52\xff\x4c\xe4\xbf\xc0\xf9\xe9\x4b\xbf\x42\x32\x92\xaf\x44\x17\xff\x85\x23\x6e\x85\x9e\x08\x11\xfa\x2f\xa3\x0e\xa3\x37\xb2\xd0\x9c\x52\x2f\xaa\x4e\x5c\x81\x51\xce\x95\x20\xef\x06\xb3\x25\x57\x8e\x59\x78\x2f\xeb\x3f\x7d\xab\xc2\x45\xff\xa3\x7f\x82\x7f\x4e\x19\xa8\x13\x2e\x27\xd4\x9c\xd8\x33\x7a\xf7\x4d\xf6\xd5\xd2\xef\xb5\xe6\xad\x5f\x5a\xf2\x77\x14\xbf\xc0\xcb\x0b\x4c\x00\x27\x30\xb8\x64\x12\xe3\x8c\x5b\x94\x44\x3f\xa1\xec\x9b\xd4\xb0\x54\x9c\x44\xc4\x7f\x15\xc6\x0e\x12\xba\x0f\xd9\x8d\xbe\xca\x2c\xc5\x4c\x97\xed\x3f\xbc\xc8\x49\xe0\xec\xac\xe5\xd1\x20\x7e\xf4\x13\x2d\x61\xc8\xc7\x41\xe5\x9e\x53\xde\x4c\x55\xb0\x28\x93\x80\xc7\x3c\x64\x4b\x44\x35\xef\x0d\xeb\x63\x47\xfd\x30\x8b\x2d\x77\x5a\x53\x71\x01\x20\x17\xa0\xea\xc0\x14\x53\xea\xf9\x84\xd8\x45\x1e\xd9\xd7\xc6\x3c\xcd\x23\x08\x7c\xf8\x75\xe2\xbd\xf9\xc0\x05\xb7\x81\x8b\x5d\x7e\xa7\x6a\xc5\xed\x52\xc8\xc5\x11\xb4\x28\x64\x0e\x7c\x6c\x1c\xc4\x33\xe2\xef\x9f\x98\x66\x2b\x6e\xb9\x36\x93\x90\xd4\xa0\x1a\x89\xc0\x3f\x29\x94\xb3\x3c\x03\x33\x2e\xe4\x02\x0c\x67\xba\x58\x46\xa7\xd6\x62\xa3\xde\x03\xd2\x71\xa2\xb0\xaa\x9e\x83\x80\x2b\x9a\x95\x55\x93\x69\x7c\x10\x6e\x3e\x43\x44\xe1\x3c\x37\xfc\x71\xed\xaa\xc1\x4f\x1c\x2e\x5b\xba\x75\x76\x06\x57\x54\x93\x99\x2b\x4d\x70\x41\xf3\x2d\xd7\x16\xc4\x3c\x54\x05\x6f\xae\x40\xe9\x9a\x41\x61\x7c\xc9\x28\xa1\x22\xe6\xc3\x7e\x23\x1e\xdd\x77\x9c\x38\x81\x91\x8b\x83\xd6\x98\xb3\x00\x1c\x33\xf6\x22\xfb\x4c\x1c\x7d\x69\x19\x7d\xfc\x17\xaf\x0c\xef\x3a\x1b\xef\xf2\x44\xf5\x87\x9e\xb9\xcd\xb7\xf0\xa5\xe4\xc6\x6a\xb5\x1f\xb6\xf6\x52\xff\x1c\xed\xbd\x45\x24\x87\x2c\x52\x14\xe5\x1c\x4e\x0a\x24\xbc\x88\x9a\xe3\xec\x3c\xee\x8f\xf8\x1d\x48\x22\x29\x5b\x77\xc9\xbd\xcf\xd9\x3b\x83\x12\x94\xd0\xe7\xf3\x1b\xfc\x83\x8e\xea\x55\x08\x33\xcd\xef\xf7\x75\xad\x2a\xc5\x49\x0f\x1d\x7a\xeb\x8d\x2e\x96\xcc\xf0\xe1\x9a\xed\x31\xf8\x4d\xe0\x5d\x5a\x11\xfb\x99\x6d\x2a\x3b\x9a\xc0\xbb\x4e\x75\xee\xf6\x7a\x7a\x10\xb4\x4d\x22\xb1\xc6\x28\xc7\xa5\x48\xb0\x91\xe2\x97\x0d\x05\x5a\x7f\xa9\x91\x68\x5b\x8f\xf8\x63\x2a\x04\x2d\x4c\xb8\x40\x69\x7c\x29\xdd\x53\x84\x5a\x3a\xd9\xb2\x71\xb6\xbc\x62\xeb\x75\xcb\x9d\xf9\xc2\x35\x06\x37\x9c\x53\x29\xb9\x00\xcb\xf5\x0a\x76\xcc\x95\xab\x03\x61\x72\x8d\x33\x9f\xd2\x8f\xe1\xc6\xbe\x46\xda\x96\xaf\xd6\x4a\x33\xbd\x8f\xc9\x16\x4a\xd6\xe5\xab\xdd\x52\x54\x1c\x76\xe8\x9a\x16\x98\x4a\xb8\x4b\x94\x19\xb7\x96\x6b\x5a\xc3\xdd\x6e\x04\xdc\x9f\x73\x57\x47\x68\xdb\x71\xea\xd6\x19\x85\xf9\x4a\x0c\xaf\xdf\x65\x94\xa9\x3b\xab\x47\x4b\x11\x4b\x7f\x03\x06\x3e\x0a\x93\x04\xb9\xd4\x91\xfc\xe5\x05\xbc\x99\xc0\xe0\xd6\x17\x09\xc3\xe9\x14\x35\x40\xe1\xc0\x57\x6b\xbb\x6f\xa3\x90\xf8\xaf\xe0\x26\x08\xc1\xf6\xa1\xb4\xee\x94\xb0\xd6\x45\x37\xca\x86\x41\xb4\x47\xb8\xc8\x20\x6a\x7a\x1c\x1f\x0b\x5c\xc0\x9b\xee\x88\xf4\x48\x22\x88\x1f\x9b\xfd\x21\x60\xb8\x12\xd2\xde\x5e\x4f\x3b\x49\x0d\xaa\xa5\xa3\xc4\xfc\xb5\x5c\x2e\xb1\x40\x9d\x6f\x16\x45\x42\x9d\x43\xe9\xa2\xed\x1a\x69\x77\x47\xd6\xf6\x1b\x89\xbc\x3b\x28\xad\x66\x76\xc5\x84\x70\x39\x1e\xdf\x45\xef\xa5\x2b\xb1\x46\x89\x86\xe6\x85\x58\x0b\x2e\xc9\x46\xeb\x6b\x21\x7f\xad\xc6\x85\x6e\x52\xf3\xc7\x0e\x60\x5c\xd3\x1e\x5a\x74\x7b\x13\x38\x3f\x75\x32\x6a\x31\x41\x19\x49\x54\xb9\xcd\x8a\x27\x2b\x8e\x63\xb6\xdf\x5a\x2c\xa7\x49\x47\xc8\x2d\xc7\x53\xc7\x00\xa2\x14\xe5\x80\x8e\xcd\x98\x2d\x96\x3f\xd6\x8a\xf6\xcb\x86\x49\x2b\x6c\xa8\xf1\xb5\xa2\x33\xe5\xe3\x21\xe1\x6d\xeb\xbc\xf3\x8f\x02\xce\xc1\x93\xe9\xcb\x4c\x82\x5e\x77\xb5\x0c\x2e\x40\x20\xcb\xfd\x50\xe2\xa9\x43\x69\x3e\x37\x78\x95\xb3\x60\x84\x52\xc7\x87\xe6\x0c\xab\xbf\x2f\x4a\x1f\x9f\xfa\xb5\x14\x1f\x01\xee\x95\x19\x8e\xbc\xeb\xfd\x81\x3c\x6f\x25\x8c\x45\x33\x5a\x62\x60\xad\x34\x67\xe5\xde\x5d\x86\x7a\x2e\xcb\x41\x87\x7a\xcd\xf7\x58\x98\x9b\xba\xf6\x38\x9c\xee\xd7\xfc\xfc\x5d\x73\xa5\x46\x5b\xf8\x61\x38\x1a\x4d\x60\x50\x0f\x87\x2d\xfe\xe6\x83\xb3\xe6\xbf\x6c\xb8\x41\xb5\x9e\xd7\x3b\x04\xb2\xc8\xfe\xd5\x66\xac\xa2\x32\xe7\xc5\x45\xe4\x99\x3b\xe4\x31\x77\x8d\x4b\xa8\xd1\x42\x34\xe1\x60\x3a\x4b\x95\x90\xb9\x45\x7f\x99\x93\xe0\xce\x5f\x16\xf9\x2f\xc1\xf6\x7b\x44\xfd\xf9\xcd\x97\x51\x67\x01\xc7\x67\x64\xa7\x0b\x6e\x83\xe0\x48\x6e\x2d\x93\xa0\x09\xc1\x6f\xb9\xcb\xa9\xf3\x53\x2f\x97\x51\x4e\x71\xcf\x4f\x71\x1f\x39\x25\x0c\x90\x1b\x0e\x60\xee\x74\x3b\x79\xf0\x1d\x81\x2d\x26\x41\xac\xd6\x15\xd5\x5b\x5c\x73\x09\x75\xa5\x14\x1b\x63\x55\x14\xd2\xb7\x82\xef\x9a\x26\x9e\x71\x42\x85\x9e\x25\xdd\x3b\x09\x1c\xf0\x83\x03\x36\xad\xd1\x98\x97\x61\x7c\x61\x0c\x5d\xd4\xcf\x36\x76\xa9\xb4\x47\x2e\x9d\xc7\xf1\x0d\x61\xe7\x61\x19\x77\x26\xf5\x8c\xb1\xcb\xcd\x6a\x26\x99\xa8\x7a\x47\x3c\x92\x77\xe4\x2e\x00\x5b\x40\x2c\x39\xad\x74\x43\x69\x08\x8e\x77\x73\xd2\x3e\xe3\xf6\x56\xd2\x01\xd1\x3e\xde\xcb\xfd\x1d\xc9\xf8\x3e\xbd\x8c\xbf\x16\x15\x7f\x48\x67\xb5\x52\xe4\x56\x68\xcb\x6c\xcc\x3f\x1c\x65\xaa\xb4\x6e\x67\x70\x51\x6f\x31\x13\x27\xd9\x0a\xf1\x18\xfe\xd7\x7d\x18\xed\x10\x2e\xe2\xfd\x76\x87\x86\xbd\xc2\x45\xb3\xef\xf1\x46\x8b\x96\xf5\xa5\xe8\xb2\xae\x8b\x76\x06\xc4\xfb\x84\x8b\x64\xdb\x47\xe4\xae\xd9\x46\x83\xf1\xcd\xed\xf5\xf4\x24\x6d\x95\x18\xff\xc4\x8d\xaa\xb6\x5c\x43\x46\xcb\x9b\x1b\xe6\xe7\x56\xbf\xe8\x2e\xfb\x89\x34\x21\x5b\xb0\x47\x1c\xd4\x2a\x89\x3c\xe5\x69\x74\x02\x73\xa4\x2f\xc3\x50\xda\x81\x7c\x9d\x25\x02\x04\xa9\x2b\xcf\xdd\x68\x0c\xb3\x4e\x7a\x40\xc7\xfe\xba\x1b\xf1\xc6\x85\x92\x05\xb3\xc3\x74\x9d\xf1\x23\x15\xaa\x04\x84\xfa\x3a\x20\x9a\xc9\x60\x04\x6f\xdf\xc2\xe0\xf5\xeb\x41\x17\x6a\xf9\xa5\x06\xaf\xeb\x4b\xb1\xda\xb1\x1c\x1a\xfb\x04\x6c\x39\xe3\x3e\x92\x31\x2a\xa1\xba\x33\xac\xd3\x60\x38\x30\xa1\xa3\x00\x63\xab\xdc\x49\x0e\x47\x07\xf3\xac\x05\xb7\x64\x63\x78\xf4\x9f\x11\xcf\x7c\xc9\x83\xc2\xcf\x9d\xa5\x09\xfc\xa4\x76\x7a\x25\xcc\xba\x62\xfb\x1f\x86\xa3\x6e\x7e\x44\xc3\x73\x91\xeb\x87\x96\x9e\x7c\x39\xc4\xae\x76\x9e\x00\xe7\x0d\xff\x97\xe2\xe7\x84\x28\x8f\x22\xdf\xdd\x2e\x4b\x3e\x85\xf2\x9a\x9d\xb0\xc5\xd2\x05\xec\x2e\xf0\x2c\x58\x7d\xd9\xdc\x2b\x8f\x49\x67\x4e\x24\xdb\xec\xa4\x6e\xc2\xea\x3f\x2e\xd4\x3d\xad\x99\x0c\xba\x07\xe6\x3f\x49\xfc\x6c\xc7\x9b\xdc\x49\xfb\x4f\x14\x57\xd3\x1d\xfe\xd7\x74\xfa\x09\x23\xea\x53\xd8\x54\x58\xc4\x6f\xa4\x6b\x23\xd0\xb9\x08\x4c\xcf\xac\x47\x29\x0f\x1e\x59\x6e\x4e\xff\x89\x79\xe0\xf2\xd4\x3e\x64\x70\x02\xff\xd0\x92\xdf\xa6\x25\xfd\x0c\x1d\x55\xed\xf1\x9f\x4c\xd9\x23\x1b\x77\xe3\xcf\xe8\x50\x32\xd6\xdc\x28\x1c\x81\xa0\x9a\xea\x7f\x36\x1f\xcf\x34\x51\xb4\xea\x40\xd9\x5c\x7b\x94\x0b\x12\x94\xd9\x4d\xe0\xb3\x03\x39\x5f\x7a\x2a\x11\xb7\xd7\xd3\xa8\xe5\x6f\x34\x81\x57\xbd\x25\xf7\xee\x64\x6f\x53\x2d\x0a\xc1\x53\xdf\x5e\x4f\xdb\xae\x9d\x3a\x86\xbb\xfe\x78\xa8\xb9\xa1\x74\xd3\xdd\x19\xc1\xdf\xff\xee\x7f\x7a\x4b\xf8\x0a\x01\x56\x8f\x7d\xb7\xee\x88\x43\x47\x56\x28\xb5\x4d\xfc\xdd\x7a\x7d\xd5\x17\x5a\xd0\x9b\x3e\x19\xba\xd8\x2a\x94\xd6\xbc\xe8\x5c\x24\x3f\x7e\xa8\x71\x25\x3b\x7b\xac\x27\x5d\xd0\xfc\xc9\xb5\x58\xeb\xcc\xa3\x9f\x78\xc1\xc5\x36\xfb\xa8\x4b\x38\x0f\xbb\xa3\xc0\xd9\xc8\xfa\xec\xac\xd5\x93\x57\xf7\xad\xcd\x95\x5e\x51\x85\x12\x97\x30\xf1\x70\x1c\x40\x6d\x4e\x61\xab\x76\xbf\xe6\xf5\x5d\xb1\x84\xbf\xb8\x23\xff\x0b\x5d\xa7\xa0\xdd\x26\x6a\xb2\x65\x9a\x5a\x8d\x4a\xcc\xd2\xe3\x9e\x98\x9c\x82\x1d\xd7\x1a\x13\xc8\xf5\xb7\xc7\xd4\x3d\x56\xa5\x66\x3b\xd0\x7c\xa5\xb6\x9c\xf2\x7e\xdc\x09\x01\x47\xd7\x0d\xd2\x14\xc9\x64\x09\x6e\x90\xb0\xfe\x1d\x11\xd7\x47\xde\x51\xf9\x6c\x55\x25\x28\xfd\x31\x35\x30\x4a\xc1\xa9\x37\xdb\x17\x6e\xc2\x86\xc6\x8e\xd7\xe1\x57\xbe\x9f\x40\xb3\x00\xf9\x3e\xea\xd8\x1f\x0e\x56\xc2\x50\x21\x99\xd0\xf8\xc0\xc3\xc8\x66\x6c\x8c\x1f\x73\x55\xe2\xa4\x85\x98\xd8\x18\x8b\xd2\x37\x11\x07\x6e\xf4\xdb\x31\x73\x0d\xc2\xa3\xac\x9b\x3b\x3f\xa5\xa9\x3d\xa2\xf7\x05\x71\xcb\xbe\xd2\x4b\x1b\x28\x03\x7a\x65\xa7\x2c\x13\x09\x37\xaf\x42\x44\x2a\x19\x13\x0a\x93\x6a\xd3\xad\x27\xd2\xbb\x0a\x3a\xea\x4c\xfa\x26\x2f\x79\xe8\x38\xdc\x17\x66\x5e\x46\x77\x51\xe8\xf8\xba\xb8\xb4\x8c\x6a\xca\x5e\x8c\x9d\x9b\x01\x56\x96\xa1\x95\xca\x91\xae\x77\x10\x59\xe0\x6e\x29\x8a\x65\x50\x52\xea\xcc\xab\x4a\x50\x92\x77\xd6\x54\x55\x39\xcd\xeb\xcd\x67\x51\x7e\x09\xec\x67\x0e\x3d\x6e\x02\xc7\xd3\xb6\xea\x98\xb3\xf6\x45\x36\xbf\x6e\xcf\x69\xbb\x20\x13\x5a\x4b\x98\x74\x87\x13\x5a\x2a\xae\xa2\x86\xe4\xd0\x88\xd5\xa9\xd6\xf5\x87\xac\x03\x8d\x30\x8d\xe1\x7c\xe5\x7b\xd3\xc3\x5f\x88\x70\x48\xdb\x24\x2f\x4d\x59\xe5\xbd\x42\x3f\x63\x67\x67\xa1\x1b\xab\x71\x0b\xd4\x94\xaa\x39\x2b\x93\xb6\x49\xd2\x57\x7a\xb3\xaa\xfe\x75\xa9\x4a\xf3\xbb\x22\xee\xa1\x42\x7f\x7a\xfa\xcc\x3c\x12\xb3\x7b\x6f\x21\x7e\x43\xf8\x16\xf3\x9c\x02\xba\x4e\x8f\x4c\x58\x6f\xfa\x0a\x51\x39\x08\x08\xd3\xfb\x61\xe9\x41\x54\x95\xda\x41\xa9\x76\xb2\x60\x54\xe4\xef\x90\x41\x1b\xd0\x7c\x1e\x7a\x46\x3a\x02\x40\xd2\x8f\x48\xa1\x25\x49\x24\x87\xd6\x9e\xec\xf6\x9b\x61\x5e\x8f\x60\x31\x18\xfb\x50\xdc\x12\xee\x7b\xb9\xf7\x3d\xc3\xf7\xf9\xd8\xdd\xbe\xe3\xf1\xc5\xfe\xdf\xbb\x79\xbf\x81\xb9\x3d\xb0\xf1\x6f\xad\xc0\x07\x66\xf2\x30\xe9\xff\xaf\xe1\x3e\xbc\x21\xf9\x8c\xad\xf6\x9e\x5e\xfd\x96\x6a\xf3\x52\x26\x39\x39\xb9\x57\x92\x37\xaf\x58\x5a\x95\x5e\x7b\x53\x7b\x40\xdb\xcd\xa4\xcd\xbc\x1f\x70\x48\x7a\x35\x98\x6b\x8c\x88\xce\xa3\xb9\x46\xe9\xbb\xac\x8f\xf0\x6a\xd4\x3d\xd6\xad\x9d\xb4\x54\x33\xe9\x5b\x3e\x74\x0b\x19\x3b\xe5\xfa\xc2\x1b\xf3\xbd\x4b\xb6\x66\x33\x51\x09\xbb\x1f\x26\xa4\x9a\xd7\x6d\x9a\x9c\x65\xec\x8c\xe6\xfc\xd5\x23\xab\xc6\x65\xa7\x06\x16\x5d\xaa\x4d\x55\x42\x94\x00\x14\x61\x69\x87\xf8\xea\xd3\xa2\x7b\x9a\x86\xe6\xa0\x4f\x3c\xed\x2b\x2e\xef\x10\x9b\xeb\xc2\x47\x76\xef\x77\xd3\xb9\x64\x74\x48\xab\xf5\x73\xf4\x5a\xd1\xe8\xc5\x77\xdf\x7d\x77\x60\x63\x89\xd3\x0c\x8d\xfc\xaf\x0d\x10\xa1\x97\x7e\x47\xf0\xec\xa6\x17\xde\xf4\x75\x97\x68\xde\x0e\x9e\xc5\xdc\xa8\xa8\x1f\xbb\x20\x92\x75\xeb\x5d\x32\xb8\x80\xb3\xba\xb3\xdb\xbd\x2e\x99\xea\x4f\x77\x6e\xa3\x86\x38\xd5\x69\x48\x76\x66\x3a\x35\xfb\xc6\x59\x67\xf1\xdb\xeb\x69\xdf\xda\xb9\x77\xd0\xda\x1c\x1c\x98\xde\x7e\x2b\x2d\xb7\x34\x35\x61\xe8\x04\x07\x35\xef\x3f\x92\xda\x50\xbf\x05\x98\xb4\xdb\xc8\xdd\x35\x25\x0d\x1a\x6f\x32\x0f\x93\x46\xf3\x37\x9d\x60\x91\x5a\x42\xa5\x58\x79\xfe\x2e\x95\xa8\x37\x83\xdc\x31\x8e\x62\xa6\x83\xf3\xac\xed\x37\x36\xeb\xa8\x75\x2f\x07\xdc\x12\x26\x0c\xdb\xf2\xe1\xf9\x69\xf4\xbe\x42\xec\x23\x23\x44\xdc\xcb\x4c\xba\x27\x21\xbf\x9e\xbf\xea\x71\x54\x8f\x3a\xb0\x6e\x4a\x9b\x6a\x63\xeb\x9a\x93\xe9\x05\xb7\x79\xf6\xc2\xc0\x51\x72\xd4\x1e\x6d\xe5\x5e\x97\x27\x90\x6a\xdc\x9b\xe9\xd4\xab\xdc\x7a\x17\xe2\xc0\x19\xf6\x9c\x5f\xd6\x1e\x22\x86\x10\xba\x14\x47\xf7\x94\x65\x8f\x2d\x0c\x8c\x4e\xea\xb1\x65\x8f\xd3\x9e\xc7\x14\xa7\x75\xd2\xd1\x29\x1f\x51\x96\x09\xd3\xf2\x6d\x37\x39\x4d\xc8\xf9\x86\x03\xfa\x90\x15\x42\xa4\x15\xfe\x5b\xf8\x42\xf9\x60\xfe\x95\x68\xa0\x48\xf8\xf0\xe2\xff\x02\x00\x00\xff\xff\x24\x1a\xd9\x92\xe5\x44\x00\x00"

func leofynftCdcBytes() ([]byte, error) {
	return bindataRead(
		_leofynftCdc,
		"LeofyNFT.cdc",
	)
}

func leofynftCdc() (*asset, error) {
	bytes, err := leofynftCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "LeofyNFT.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x69, 0xe6, 0x55, 0x77, 0xb6, 0xb7, 0x7b, 0xe3, 0xb3, 0x18, 0xe8, 0x15, 0x96, 0xab, 0x2, 0xa5, 0x6a, 0x61, 0x83, 0x57, 0xa9, 0xc3, 0x24, 0x81, 0x90, 0x8, 0x62, 0x15, 0x2a, 0x12, 0x16, 0xab}}
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
	"LeofyCoin.cdc": leofycoinCdc,
	"LeofyNFT.cdc":  leofynftCdc,
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
	"LeofyCoin.cdc": {leofycoinCdc, map[string]*bintree{}},
	"LeofyNFT.cdc": {leofynftCdc, map[string]*bintree{}},
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
