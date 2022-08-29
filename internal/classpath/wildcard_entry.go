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
	"os"
	"path/filepath"
)

type WildcardEntry struct {
	composite *CompositeEntry
}

func newEmptyWildcardEntry() *WildcardEntry {
	return &WildcardEntry{
		composite: newEmptyCompositeEntry(),
	}
}

func newWildcardEntry(path string) *WildcardEntry {
	baseDir := path[:len(path)-1] // remove *
	wildcardEntry := newEmptyWildcardEntry()

	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if IsJar(path) {
			jarEntry := newZipEntry(path)
			wildcardEntry.composite.entries = append(wildcardEntry.composite.entries, jarEntry)
		}

		return nil
	}

	_ = filepath.Walk(baseDir, walkFn)

	return wildcardEntry
}

func (entry WildcardEntry) readClass(className string) ([]byte, Entry, error) {
	return entry.composite.readClass(className)
}

func (entry WildcardEntry) String() string {
	return entry.composite.String()
}

func (entry WildcardEntry) Close() {
}
