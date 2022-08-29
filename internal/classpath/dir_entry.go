/*
 * Copyright 2022 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package classpath

import (
	"io/ioutil"
	"path/filepath"
)

var _ Entry = (*DirEntry)(nil)

type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return &DirEntry{
		absDir: absDir,
	}
}

func (entry *DirEntry) readClass(className string) ([]byte, Entry, error) {
	cf := filepath.Join(entry.absDir, className)
	data, err := ioutil.ReadFile(cf)

	return data, entry, err
}

func (entry *DirEntry) String() string {
	return entry.absDir
}

func (entry *DirEntry) Close() {
}
