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
	"strings"
)

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func Parse(jreOpt, cpOpt string) *Classpath {
	cp := &Classpath{}
	cp.parseBootClasspath(jreOpt)
	cp.parseExtClasspath(jreOpt)
	cp.parseUserClasspath(cpOpt)

	return cp
}

func (cp *Classpath) parseBootClasspath(jreOpt string) {
	jreDir := determineJreDir(jreOpt)

	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, JRELib, Star)

	cp.bootClasspath = newWildcardEntry(jreLibPath)
}

func (cp *Classpath) parseExtClasspath(jreOpt string) {
	jreDir := determineJreDir(jreOpt)
	jreLibExtPath := filepath.Join(jreDir, JRELib, JRELibExt, Star)

	// jre/lib/ext/*
	cp.extClasspath = newWildcardEntry(jreLibExtPath)
}

func (cp *Classpath) parseUserClasspath(cpOpt string) {
	if cpOpt == EmptyString {
		cpOpt = Dot
	}

	cp.userClasspath = newEntry(cpOpt)
}

func (cp *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	if !strings.HasSuffix(className, SuffixClass) {
		className = className + SuffixClass
	}

	// boot
	if data, entry, err := cp.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	// ext
	if data, entry, err := cp.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}

	// app
	return cp.userClasspath.readClass(className)
}

func (cp *Classpath) String() string {
	return cp.userClasspath.String()
}

func determineJreDir(jreOpt string) string {
	if jreOpt != EmptyString && IsExists(jreOpt) {
		return jreOpt
	}

	if IsExists(JREDir) {
		return JREDir
	}

	if javaHome := os.Getenv(JavaHome); javaHome != EmptyString {
		return filepath.Join(javaHome, JRE)
	}

	panic("classpath: Can't find jre folder!")
}
