/*
Copyright 2018 The Kubernetes Authors.

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

// Package create implements the `create` command
package create

import (
	"errors"
	"fmt" // Import the fmt package for printing  // Added by JANR

	"github.com/spf13/cobra"

	"sigs.k8s.io/kind/pkg/cmd"
	createcluster "sigs.k8s.io/kind/pkg/cmd/kind/create/cluster"
	"sigs.k8s.io/kind/pkg/log"
)

// NewCommand returns a new cobra.Command for cluster creation
func NewCommand(logger log.Logger, streams cmd.IOStreams) *cobra.Command {
	// Print information in different lines: // Added by JANR
	// Relative path: // Added by JANR
	// Brief function goal: // Added by JANR
	// All functions called in order: // Added by JANR
	fmt.Println("File(3) Step(1) Path: Skind/pkg/cmd/kind/create/create.go - Function: NewCommand()")               // Added by JANR
	fmt.Println("File(3) Step(1) Brief function goal: NewCommand returns a new cobra.Command for cluster creation") // Added by JANR

	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "create",
		Short: "Creates one of [cluster]",
		Long:  "Creates one of local Kubernetes cluster (cluster)",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cmd.Help()
			if err != nil {
				return err
			}
			return errors.New("Subcommand is required")
		},
	}
	cmd.AddCommand(createcluster.NewCommand(logger, streams)) // createcluster.NewCommand refers to import "sigs.k8s.io/kind/pkg/cmd/kind/create/cluster"
	return cmd
}
