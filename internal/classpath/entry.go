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
	"strings"
)

type Entry interface {
	// readClass find and load class file
	//
	/*
		Example:

		java.lang.Object ->
		className == java/lang/Object.class
	*/
	readClass(className string) ([]byte, Entry, error)
	// String toString()
	String() string
	Close()
}

func newEntry(path string) Entry {
	if strings.Contains(path, PathListSeparator) {
		return newCompositeEntry(path)
	}

	if strings.HasSuffix(path, Star) {
		return newWildcardEntry(path)
	}

	if IsJar(path) || IsZip(path) {
		return newZipEntry(path)
	}

	return newDirEntry(path)
}
