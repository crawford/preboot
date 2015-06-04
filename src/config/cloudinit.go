// Copyright 2015 CoreOS, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// These functions are copied from github.com/coreos/coreos-cloudinit/config.

package config

import (
	"strings"
)

func isCloudConfig(userdata []byte) bool {
	header := strings.SplitN(string(userdata), "\n", 2)[0]

	// Explicitly trim the header so we can handle user-data from
	// non-unix operating systems. The rest of the file is parsed
	// by yaml, which correctly handles CRLF.
	header = strings.TrimSuffix(header, "\r")

	return (header == "#cloud-config")
}

func isScript(userdata []byte) bool {
	header := strings.SplitN(string(userdata), "\n", 2)[0]
	return strings.HasPrefix(header, "#!")
}