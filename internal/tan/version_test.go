// Copyright 2012 The LevelDB-Go and Pebble Authors. All rights reserved. Use
// of this source code is governed by a BSD-style license that can be found in
// the LICENSE file.
//
// Copyright 2017-2019 Lei Ni (nilei81@gmail.com) and other contributors.
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

package tan

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVersionUnref(t *testing.T) {
	list := &versionList{}
	list.init(&sync.Mutex{})
	v := &version{deleted: func([]*fileMetadata) {}}
	v.ref()
	list.pushBack(v)
	v.unref()
	require.True(t, list.empty(), "expected version list to be empty")
}
