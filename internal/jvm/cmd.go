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
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// Cmd java [-options] class [args...]
type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOpt       string
	xjreOpt     string
	class       string
	args        []string
}

func (cmd *Cmd) Version() bool {
	return cmd.versionFlag
}

func (cmd *Cmd) Help() bool {
	return cmd.helpFlag || cmd.class == ""
}

func ParseArgs() *Cmd {
	cmd := &Cmd{}

	flag.Usage = Usage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print helpFlag message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print helpFlag message")

	flag.BoolVar(&cmd.versionFlag, "version", false, "print versionFlag and exit")

	flag.StringVar(&cmd.cpOpt, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOpt, "cp", "", "classpath")

	flag.StringVar(&cmd.xjreOpt, "Xjre", "", "path to jre")

	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd
}

func Usage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", parseExec())
}

func parseExec() string {
	execPath, _ := os.Executable()
	_, exec := filepath.Split(execPath)
	suffix := path.Ext(exec)

	return strings.TrimSuffix(exec, suffix)
}
