/*
Copyright 2015 Palm Stone Games, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main // import "code.palmstonegames.com/gb-gae"

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/constabulary/gb"
	"github.com/constabulary/gb/cmd"
)

const DocUsage = `gb gae is a tool for managing Google Appengine applications with gb.

Usage:

	gb gae command [arguments]
	
The commands are: 

	serve		starts a local development App Engine server
	deploy		deploys your application to App Engine
	build		compile packages and dependencies
	test		test packages
	raw		Directly call the dev_appserver.py
`

var (
	projectRoot string
)

func main() {
	// Setup flags
	fs := flag.NewFlagSet("gb-gae", flag.ExitOnError)
	fs.StringVar(&projectRoot, "R", os.Getenv("GB_PROJECT_DIR"), "set the project root")
	fs.BoolVar(&gb.Verbose, "v", gb.Verbose, "enable log levels below INFO level")

	err := cmd.RunCommand(fs, &cmd.Command{
		Run: run,
	}, os.Getenv("GB_PROJECT_DIR"), "", os.Args[1:])
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

func run(ctx *gb.Context, args []string) error {
	if len(args) == 0 {
		return errors.New(DocUsage)
	}

	// Build up the fake env
	env := cmd.MergeEnv(os.Environ(), map[string]string{
		"GOPATH": fmt.Sprintf("%s%s%s", projectRoot, string(os.PathListSeparator), path.Join(projectRoot, "vendor")),
	})

	// Switch on the subcommand
	switch args[0] {
	case "serve", "deploy", "build", "test":
		return goapp(ctx, args, env)
	case "raw":
		return raw(ctx, args, env)
	default:
		return fmt.Errorf("Unknown subcommand: %s\n\n%v", args[0], DocUsage)
	}
}

func goapp(ctx *gb.Context, args []string, env []string) error {
	app := exec.Command("goapp", args...)
	app.Stdin = os.Stdin
	app.Stdout = os.Stdout
	app.Stderr = os.Stderr
	app.Env = env

	if err := app.Run(); err != nil {
		return fmt.Errorf("Failed to run goapp command: %v", err)
	}

	return nil
}

func raw(ctx *gb.Context, args []string, env []string) error {
	app := exec.Command("dev_appserver.py", args[1:]...)
	app.Stdin = os.Stdin
	app.Stdout = os.Stdout
	app.Stderr = os.Stderr
	app.Env = env

	if err := app.Run(); err != nil {
		return fmt.Errorf("Failed to run dev_appserver.py: %v", err)
	}

	return nil
}
