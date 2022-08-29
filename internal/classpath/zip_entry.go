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
	"archive/zip"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
	zipRc   *zip.ReadCloser
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return &ZipEntry{
		absPath: absPath,
	}
}

func (entry *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	if entry.zipRc == nil {
		err := entry.openJar()
		if err != nil {
			return nil, nil, err
		}
	}

	zf := entry.findClass(className)
	if zf == nil {
		return nil, nil, fmt.Errorf("classpath: Class not found: %s", className)
	}

	data, err := loadClass(zf)

	return data, entry, err
}

func (entry *ZipEntry) openJar() error {
	reader, err := zip.OpenReader(entry.absPath)
	if err == nil {
		entry.zipRc = reader
	}

	return err
}

func (entry *ZipEntry) findClass(className string) *zip.File {
	for _, zf := range entry.zipRc.File {
		if zf.Name == className {
			return zf
		}
	}

	return nil
}

func loadClass(zf *zip.File) ([]byte, error) {
	cf, err := zf.Open()
	if err != nil {
		return nil, err
	}

	defer cf.Close()
	data, err := ioutil.ReadAll(cf)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (entry *ZipEntry) String() string {
	return entry.absPath
}

func (entry *ZipEntry) Close() {
	_ = entry.zipRc.Close()
}
