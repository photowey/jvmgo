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
	"fmt"
	"strings"
)

type CompositeEntry struct {
	entries []Entry
}

func newEmptyCompositeEntry() *CompositeEntry {
	return &CompositeEntry{
		entries: make([]Entry, 0),
	}
}

func newCompositeEntry(pathList string) Entry {
	compositeEntry := newEmptyCompositeEntry()
	for _, path := range strings.Split(pathList, PathListSeparator) {
		entry := newEntry(path)
		compositeEntry.entries = append(compositeEntry.entries, entry)
	}

	return compositeEntry
}

func (entry CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, sub := range entry.entries {
		data, from, err := sub.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}

	return nil, nil, fmt.Errorf("classpath: Class not found: %s", className)
}

func (entry CompositeEntry) String() string {
	dirs := make([]string, 0, len(entry.entries))
	for _, sub := range entry.entries {
		dirs = append(dirs, sub.String())
	}

	return strings.Join(dirs, PathListSeparator)
}

func (entry CompositeEntry) Close() {
}
