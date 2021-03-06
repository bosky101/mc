/*
 * Minio Client (C) 2015 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package client

import (
	"io"
	"os"
	"time"
)

// Client - client interface
type Client interface {
	// Common operations
	Stat() (content *Content, err error)
	List(recursive bool) <-chan ContentOnChannel

	// Bucket operations
	MakeBucket() error
	SetBucketACL(acl string) error

	// Object operations
	GetObject(offset, length int64) (body io.ReadCloser, size int64, err error)
	PutObject(size int64, data io.Reader) error

	// URL returns back internal url
	URL() *URL
}

// ContentOnChannel - List contents on channel
type ContentOnChannel struct {
	Content *Content
	Err     error
}

// Content container for content metadata
type Content struct {
	Name string
	Time time.Time
	Size int64
	Type os.FileMode
}
