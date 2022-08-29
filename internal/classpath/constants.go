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
)

const (
	EmptyString = ""
	Dot         = "."
)

const (
	// PathListSeparator :(linux/unix) or ;(windows)
	PathListSeparator = string(os.PathListSeparator)
	Star              = "*"
)

const (
	Jar       = "jar"
	JAR       = "JAR"
	Zip       = "zip"
	ZIP       = "ZIP"
	SuffixJar = ".jar"
	SuffixJAR = ".JAR"
	SuffixZip = ".zip"
	SuffixZIP = ".ZIP"
)

const (
	JavaHome  = "JAVA_HOME"
	JRE       = "jre"
	JRELib    = "lib"
	JRELibExt = "ext"
	JREDir    = "./jre"
)

const (
	SuffixClass = ".class"
)
