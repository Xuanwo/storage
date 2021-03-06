// Code generated by go generate internal/cmd; DO NOT EDIT.
package types

import (
	"fmt"
	"sync"
	"time"
)

// Field index in object bit
const (
	objectIndexContentLength          uint64 = 1 << 0
	objectIndexContentMd5             uint64 = 1 << 1
	objectIndexContentType            uint64 = 1 << 2
	objectIndexEtag                   uint64 = 1 << 3
	objectIndexID                     uint64 = 1 << 4
	objectIndexLastModified           uint64 = 1 << 5
	objectIndexLinkTarget             uint64 = 1 << 6
	objectIndexMode                   uint64 = 1 << 7
	objectIndexMultipartID            uint64 = 1 << 8
	objectIndexMultipartNumberMaximum uint64 = 1 << 9
	objectIndexMultipartSizeMaximum   uint64 = 1 << 10
	objectIndexMultipartSizeMinimum   uint64 = 1 << 11
	objectIndexPath                   uint64 = 1 << 12
	objectIndexServiceMetadata        uint64 = 1 << 13
	objectIndexUserMetadata           uint64 = 1 << 14
)

// Object is the smallest unit in go-storage.
//
// NOTES:
//   - Object's fields SHOULD not be changed outside services.
//   - Object CANNOT be copied
//   - Object is concurrent safe.
type Object struct {
	contentLength int64
	contentMd5    string
	contentType   string
	etag          string
	// ID is the unique key in storage.
	ID           string
	lastModified time.Time
	// LinkTarget is the symlink target for link object.
	linkTarget string
	Mode       ObjectMode
	// MultipartID is the part id of part object.
	multipartID string
	// Maximum part number in multipart operation
	multipartNumberMaximum int
	// Maximum part size defined by storager
	multipartSizeMaximum int64
	// Minimum part size defined by storager
	multipartSizeMinimum int64
	// Path is either the absolute path or the relative path towards storage's WorkDir depends on user's input.
	Path string
	// ServiceMetadata stores service defined metadata
	serviceMetadata map[string]string
	// UserMetadata stores user defined metadata
	userMetadata map[string]string

	// client is the client in which Object is alive.
	client Storager

	// bit used as a bitmap for object value, 0 means not set, 1 means set
	bit  uint64
	done uint32
	m    sync.Mutex
}

func (o *Object) GetContentLength() (int64, bool) {
	o.stat()

	if o.bit&objectIndexContentLength != 0 {
		return o.contentLength, true
	}

	return 0, false
}

func (o *Object) MustGetContentLength() int64 {
	o.stat()

	if o.bit&objectIndexContentLength == 0 {
		panic(fmt.Sprintf("object content-length is not set"))
	}
	return o.contentLength
}

func (o *Object) SetContentLength(v int64) *Object {
	o.contentLength = v
	o.bit |= objectIndexContentLength
	return o
}

func (o *Object) GetContentMd5() (string, bool) {
	o.stat()

	if o.bit&objectIndexContentMd5 != 0 {
		return o.contentMd5, true
	}

	return "", false
}

func (o *Object) MustGetContentMd5() string {
	o.stat()

	if o.bit&objectIndexContentMd5 == 0 {
		panic(fmt.Sprintf("object content-md5 is not set"))
	}
	return o.contentMd5
}

func (o *Object) SetContentMd5(v string) *Object {
	o.contentMd5 = v
	o.bit |= objectIndexContentMd5
	return o
}

func (o *Object) GetContentType() (string, bool) {
	o.stat()

	if o.bit&objectIndexContentType != 0 {
		return o.contentType, true
	}

	return "", false
}

func (o *Object) MustGetContentType() string {
	o.stat()

	if o.bit&objectIndexContentType == 0 {
		panic(fmt.Sprintf("object content-type is not set"))
	}
	return o.contentType
}

func (o *Object) SetContentType(v string) *Object {
	o.contentType = v
	o.bit |= objectIndexContentType
	return o
}

func (o *Object) GetEtag() (string, bool) {
	o.stat()

	if o.bit&objectIndexEtag != 0 {
		return o.etag, true
	}

	return "", false
}

func (o *Object) MustGetEtag() string {
	o.stat()

	if o.bit&objectIndexEtag == 0 {
		panic(fmt.Sprintf("object etag is not set"))
	}
	return o.etag
}

func (o *Object) SetEtag(v string) *Object {
	o.etag = v
	o.bit |= objectIndexEtag
	return o
}
func (o *Object) GetID() string {
	return o.ID
}

func (o *Object) SetID(v string) *Object {
	o.ID = v
	return o
}

func (o *Object) GetLastModified() (time.Time, bool) {
	o.stat()

	if o.bit&objectIndexLastModified != 0 {
		return o.lastModified, true
	}

	return time.Time{}, false
}

func (o *Object) MustGetLastModified() time.Time {
	o.stat()

	if o.bit&objectIndexLastModified == 0 {
		panic(fmt.Sprintf("object last-modified is not set"))
	}
	return o.lastModified
}

func (o *Object) SetLastModified(v time.Time) *Object {
	o.lastModified = v
	o.bit |= objectIndexLastModified
	return o
}

func (o *Object) GetLinkTarget() (string, bool) {
	o.stat()

	if o.bit&objectIndexLinkTarget != 0 {
		return o.linkTarget, true
	}

	return "", false
}

func (o *Object) MustGetLinkTarget() string {
	o.stat()

	if o.bit&objectIndexLinkTarget == 0 {
		panic(fmt.Sprintf("object link-target is not set"))
	}
	return o.linkTarget
}

func (o *Object) SetLinkTarget(v string) *Object {
	o.linkTarget = v
	o.bit |= objectIndexLinkTarget
	return o
}
func (o *Object) GetMode() ObjectMode {
	return o.Mode
}

func (o *Object) SetMode(v ObjectMode) *Object {
	o.Mode = v
	return o
}

func (o *Object) GetMultipartID() (string, bool) {
	o.stat()

	if o.bit&objectIndexMultipartID != 0 {
		return o.multipartID, true
	}

	return "", false
}

func (o *Object) MustGetMultipartID() string {
	o.stat()

	if o.bit&objectIndexMultipartID == 0 {
		panic(fmt.Sprintf("object multipart-id is not set"))
	}
	return o.multipartID
}

func (o *Object) SetMultipartID(v string) *Object {
	o.multipartID = v
	o.bit |= objectIndexMultipartID
	return o
}

func (o *Object) GetMultipartNumberMaximum() (int, bool) {
	o.stat()

	if o.bit&objectIndexMultipartNumberMaximum != 0 {
		return o.multipartNumberMaximum, true
	}

	return 0, false
}

func (o *Object) MustGetMultipartNumberMaximum() int {
	o.stat()

	if o.bit&objectIndexMultipartNumberMaximum == 0 {
		panic(fmt.Sprintf("object multipart-number-maximum is not set"))
	}
	return o.multipartNumberMaximum
}

func (o *Object) SetMultipartNumberMaximum(v int) *Object {
	o.multipartNumberMaximum = v
	o.bit |= objectIndexMultipartNumberMaximum
	return o
}

func (o *Object) GetMultipartSizeMaximum() (int64, bool) {
	o.stat()

	if o.bit&objectIndexMultipartSizeMaximum != 0 {
		return o.multipartSizeMaximum, true
	}

	return 0, false
}

func (o *Object) MustGetMultipartSizeMaximum() int64 {
	o.stat()

	if o.bit&objectIndexMultipartSizeMaximum == 0 {
		panic(fmt.Sprintf("object multipart-size-maximum is not set"))
	}
	return o.multipartSizeMaximum
}

func (o *Object) SetMultipartSizeMaximum(v int64) *Object {
	o.multipartSizeMaximum = v
	o.bit |= objectIndexMultipartSizeMaximum
	return o
}

func (o *Object) GetMultipartSizeMinimum() (int64, bool) {
	o.stat()

	if o.bit&objectIndexMultipartSizeMinimum != 0 {
		return o.multipartSizeMinimum, true
	}

	return 0, false
}

func (o *Object) MustGetMultipartSizeMinimum() int64 {
	o.stat()

	if o.bit&objectIndexMultipartSizeMinimum == 0 {
		panic(fmt.Sprintf("object multipart-size-minimum is not set"))
	}
	return o.multipartSizeMinimum
}

func (o *Object) SetMultipartSizeMinimum(v int64) *Object {
	o.multipartSizeMinimum = v
	o.bit |= objectIndexMultipartSizeMinimum
	return o
}
func (o *Object) GetPath() string {
	return o.Path
}

func (o *Object) SetPath(v string) *Object {
	o.Path = v
	return o
}

func (o *Object) GetServiceMetadata() (map[string]string, bool) {
	o.stat()

	if o.bit&objectIndexServiceMetadata != 0 {
		return o.serviceMetadata, true
	}

	return map[string]string{}, false
}

func (o *Object) MustGetServiceMetadata() map[string]string {
	o.stat()

	if o.bit&objectIndexServiceMetadata == 0 {
		panic(fmt.Sprintf("object service-metadata is not set"))
	}
	return o.serviceMetadata
}

func (o *Object) SetServiceMetadata(v map[string]string) *Object {
	o.serviceMetadata = v
	o.bit |= objectIndexServiceMetadata
	return o
}

func (o *Object) GetUserMetadata() (map[string]string, bool) {
	o.stat()

	if o.bit&objectIndexUserMetadata != 0 {
		return o.userMetadata, true
	}

	return map[string]string{}, false
}

func (o *Object) MustGetUserMetadata() map[string]string {
	o.stat()

	if o.bit&objectIndexUserMetadata == 0 {
		panic(fmt.Sprintf("object user-metadata is not set"))
	}
	return o.userMetadata
}

func (o *Object) SetUserMetadata(v map[string]string) *Object {
	o.userMetadata = v
	o.bit |= objectIndexUserMetadata
	return o
}

func (o *Object) clone(xo *Object) {
	o.contentLength = xo.contentLength
	o.contentMd5 = xo.contentMd5
	o.contentType = xo.contentType
	o.etag = xo.etag
	o.ID = xo.ID
	o.lastModified = xo.lastModified
	o.linkTarget = xo.linkTarget
	o.Mode = xo.Mode
	o.multipartID = xo.multipartID
	o.multipartNumberMaximum = xo.multipartNumberMaximum
	o.multipartSizeMaximum = xo.multipartSizeMaximum
	o.multipartSizeMinimum = xo.multipartSizeMinimum
	o.Path = xo.Path
	o.serviceMetadata = xo.serviceMetadata
	o.userMetadata = xo.userMetadata

	o.bit = xo.bit
}
