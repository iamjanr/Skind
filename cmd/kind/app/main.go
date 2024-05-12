/*
Copyright 2019 The Kubernetes Authors.

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

package app

import (
	"fmt" // Import the fmt package for printing  // Added by JANR
	"io/ioutil"
	"os"

	"github.com/spf13/pflag"

	"sigs.k8s.io/kind/pkg/cmd"
	"sigs.k8s.io/kind/pkg/cmd/kind"
	"sigs.k8s.io/kind/pkg/errors"
	"sigs.k8s.io/kind/pkg/exec"
	"sigs.k8s.io/kind/pkg/log"
)

// Main is the kind main(), it will invoke Run(), if an error is returned
// it will then call os.Exit
func Main() {
	// Print information in different lines: // Added by JANR
	// Relative path: // Added by JANR
	// Brief function goal: // Added by JANR
	// All functions called in order: // Added by JANR
	fmt.Println("(1)(1) Path: skin/cmd/kind/app/main.go - Main() ")                                                                             // Added by JANR
	fmt.Println("(1)(1) Brief function goal: Main is the kind main(), it will invoke Run(), if an error is returned it will then call os.Exit") // Added by JANR
	fmt.Println("(1)(1) All functions called in order: Run()")                                                                                  // Added by JANR

	if err := Run(cmd.NewLogger(), cmd.StandardIOStreams(), os.Args[1:]); err != nil {
		os.Exit(1)
	}
}

// Run invokes the kind root command, returning the error.
// See: sigs.k8s.io/kind/pkg/cmd/kind
func Run(logger log.Logger, streams cmd.IOStreams, args []string) error {
	// Print information in different lines: // Added by JANR
	// Relative path: // Added by JANR
	// Brief function goal: // Added by JANR
	// All functions called in order: // Added by JANR
	fmt.Println("(1)(2) Path: skin/cmd/kind/app/main.go - Run()")                                                      // Added by JANR
	fmt.Println("(1)(2) Brief function goal: Run invokes the kind root command, returning the error")                  // Added by JANR
	fmt.Println("(1)(2) All functions called in order: checkQuiet(), kind.NewCommand(), c.SetArgs(args), c.Execute()") // Added by JANR
	// NOTE: we handle the quiet flag here so we can fully silence cobra
	if checkQuiet(args) {
		//Print if we a re on quiet mode // Added by JANR
		fmt.Println("Quiet Mode") // Added by JANR
		// if we are in quiet mode, we want to suppress all status output
		// only streams.Out should be written to (program output)
		logger = log.NoopLogger{}
		streams.ErrOut = ioutil.Discard
	}
	// actually run the command
	c := kind.NewCommand(logger, streams)
	c.SetArgs(args)                     // Set the arguments for the command // Added by JANR
	if err := c.Execute(); err != nil { // Execute the command
		logError(logger, err)
		return err
	}
	return nil
}

// checkQuiet returns true if -q / --quiet was set in args
func checkQuiet(args []string) bool {
	// Print information in different lines: // Added by JANR
	// Relative path: // Added by JANR
	// Brief function goal: // Added by JANR
	// All functions called in order: // Added by JANR
	fmt.Println("(1)(3) Path: skin/cmd/kind/app/main.go - checkQuiet()")                                                                                           // Added by JANR
	fmt.Println("(1)(3) Brief function goal: checkQuiet returns true if -q / --quiet was set in args")                                                             // Added by JANR
	fmt.Println("(1)(3) All functions called in order: pflag.NewFlagSet(), flags.ParseErrorsWhitelist.UnknownFlags, flags.BoolVarP(), flags.Usage, flags.Parse()") // Added by JANR
	// create a new flag set	// Added by JANR
	flags := pflag.NewFlagSet("persistent-quiet", pflag.ContinueOnError)
	flags.ParseErrorsWhitelist.UnknownFlags = true
	quiet := false
	flags.BoolVarP(
		&quiet,
		"quiet",
		"q",
		false,
		"silence all stderr output",
	)
	// NOTE: pflag will error if -h / --help is specified
	// We don't care here. That will be handled downstream
	// It will also call flags.Usage so we're making that no-op
	flags.Usage = func() {}
	_ = flags.Parse(args)
	return quiet
}

// logError logs the error and the root stacktrace if there is one
func logError(logger log.Logger, err error) {
	// Print information in different lines: // Added by JANR
	// Relative path: // Added by JANR
	// Brief function goal: // Added by JANR
	// All functions called in order: // Added by JANR
	fmt.Println("(1)(4) Path: skin/cmd/kind/app/main.go - logError()") // Added by JANR
	colorEnabled := cmd.ColorEnabled(logger)
	if colorEnabled {
		logger.Errorf("\x1b[31mERROR\x1b[0m: %v", err)
	} else {
		logger.Errorf("ERROR: %v", err)
	}
	// Display Output if the error was from running a command ...
	if err := exec.RunErrorForError(err); err != nil {
		if colorEnabled {
			logger.Errorf("\x1b[31mCommand Output\x1b[0m: %s", err.Output)
		} else {
			logger.Errorf("\nCommand Output: %s", err.Output)
		}
	}
	// TODO: stacktrace should probably be guarded by a higher level ...?
	if logger.V(1).Enabled() {
		// Then display stack trace if any (there should be one...)
		if trace := errors.StackTrace(err); trace != nil {
			if colorEnabled {
				logger.Errorf("\x1b[31mStack Trace\x1b[0m: %+v", trace)
			} else {
				logger.Errorf("\nStack Trace: %+v", trace)
			}
		}
	}
}
