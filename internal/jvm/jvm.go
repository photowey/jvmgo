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

package jvm

import (
	"fmt"
	"strings"

	"github.com/photowey/jvmgo/internal/classpath"
)

func Start(cmd *Cmd) {
	cp := classpath.Parse(cmd.xjreOpt, cmd.cpOpt)
	fmt.Printf("jvm: classpath: %s class: %s args: %v\n", cp, cmd.class, cmd.args)

	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("jvm: Could not find or load main class %s\n", cmd.class)
		return
	}

	fmt.Printf("jvm: class data:%v\n", classData)
}
