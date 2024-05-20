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

// Package cluster implements the `create cluster` command
package cluster

import (
	"fmt"
	"io"
	"os"

	"syscall"
	"time"

	term "golang.org/x/term"

	"github.com/spf13/cobra"

	"sigs.k8s.io/kind/pkg/cluster"
	"sigs.k8s.io/kind/pkg/commons"

	"sigs.k8s.io/kind/pkg/cmd"
	"sigs.k8s.io/kind/pkg/errors"
	"sigs.k8s.io/kind/pkg/log"

	"sigs.k8s.io/kind/pkg/internal/cli"
	"sigs.k8s.io/kind/pkg/internal/runtime"
)

type flagpole struct {
	Name           string
	Config         string
	ImageName      string
	Retain         bool
	Wait           time.Duration
	Kubeconfig     string
	VaultPassword  string
	DescriptorPath string
	MoveManagement bool
	AvoidCreation  bool
	ForceDelete    bool
	ValidateOnly   bool
}

const clusterDefaultPath = "./cluster.yaml"
const secretsDefaultPath = "./secrets.yml"

// NewCommand returns a new cobra.Command for cluster creation
func NewCommand(logger log.Logger, streams cmd.IOStreams) *cobra.Command {
	// Print information in different lines: // Added by JANR
	// Relative path: // Added by JANR
	// Brief function goal: // Added by JANR
	// All functions called in order: // Added by JANR
	fmt.Println("File(4) Step(1) Path: Skind/pkg/cmd/kind/create/cluster/createcluster.go - Function: NewCommand()") // Added by JANR
	fmt.Println("File(4) Step(1) Brief function goal: NewCommand returns a new cobra.Command for cluster creation")  // Added by JANR

	flags := &flagpole{}
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "cluster",
		Short: "Creates a local Kubernetes cluster",
		Long:  "Creates a local Kubernetes cluster using Docker container 'nodes'",
		RunE: func(cmd *cobra.Command, args []string) error {
			cli.OverrideDefaultName(cmd.Flags())
			return runE(logger, streams, flags)
		},
	}
	cmd.Flags().StringVarP(
		&flags.Name,
		"name",
		"n",
		"",
		"cluster name, overrides KIND_CLUSTER_NAME, config (default kind)",
	)
	cmd.Flags().StringVar(
		&flags.Config,
		"config",
		"",
		"path to a kind config file",
	)
	cmd.Flags().StringVar(
		&flags.ImageName,
		"image",
		"",
		"node docker image to use for booting the cluster",
	)
	cmd.Flags().BoolVar(
		&flags.Retain,
		"retain",
		false,
		"retain nodes for debugging when cluster creation fails",
	)
	cmd.Flags().DurationVar(
		&flags.Wait,
		"wait",
		time.Duration(0),
		"wait for control plane node to be ready (default 0s)",
	)
	cmd.Flags().StringVar(
		&flags.Kubeconfig,
		"kubeconfig",
		"",
		"sets kubeconfig path instead of $KUBECONFIG or $HOME/.kube/config",
	)
	cmd.Flags().StringVarP(
		&flags.VaultPassword,
		"vault-password",
		"p",
		"",
		"sets vault password to encrypt secrets",
	)
	cmd.Flags().StringVarP(
		&flags.DescriptorPath,
		"descriptor",
		"d",
		"",
		"allows you to indicate the name of the descriptor located in current or other directory. Default: cluster.yaml",
	)
	cmd.Flags().BoolVar(
		&flags.MoveManagement,
		"keep-mgmt",
		false,
		"by setting this flag the cluster management will be kept in the kind",
	)
	cmd.Flags().BoolVar(
		&flags.AvoidCreation,
		"avoid-creation",
		false,
		"by setting this flag the worker cluster won't be created",
	)
	cmd.Flags().BoolVar(
		&flags.ForceDelete,
		"delete-previous",
		false,
		"by setting this flag the local cluster container will be deleted",
	)
	cmd.Flags().BoolVar(
		&flags.ValidateOnly,
		"validate-only",
		false,
		"by setting this flag the descriptor will be validated and the cluster won't be created",
	)

	// Print the command-line arguments // Added by JANR
	//fmt.Println("File(4) Step(1) Path: Skind/pkg/cmd/kind/create/cluster/createcluster.go - Print - Args returned by NewCommand()") // Added by JANR
	//fmt.Println("File(4) Step(1) Command-line arguments:")                                                                          // Added by JANR
	//for i, arg := range os.Args[1:] {                                                                                      // Added by JANR
	//	fmt.Printf("Argument %d: %s\n", i+1, arg) // Added by JANR
	//} // Added by JANR

	return cmd
}

func runE(logger log.Logger, streams cmd.IOStreams, flags *flagpole) error {
	// Print information in different lines: // Added by JANR
	// Relative path: // Added by JANR
	// Brief function goal: // Added by JANR
	// All functions called in order: // Added by JANR
	fmt.Println("File(4) Step(2) Path: Skind/pkg/cmd/kind/create/cluster/createcluster.go - Function: runE()") // Added by JANR
	fmt.Println("File(4) Step(2) Brief function goal: runE validates the flags and creates the cluster")       // Added by JANR

	err := validateFlags(flags)
	if err != nil {
		return err
	}
	fmt.Println("File(4) Step(2) - Print - Flags validated") // Added by JANR

	if flags.DescriptorPath == "" {
		flags.DescriptorPath = clusterDefaultPath
	}
	fmt.Println("File(4) Step(2) - Print - Descriptor path set") // Added by JANR

	if flags.VaultPassword == "" {
		flags.VaultPassword, err = setPassword(secretsDefaultPath)
		if err != nil {
			return err
		}
	}
	fmt.Println("File(4) Step(2) - Print - Vault password set") // Added by JANR

	keosCluster, clusterConfig, err := commons.GetClusterDescriptor(flags.DescriptorPath)
	if err != nil {
		return errors.Wrap(err, "failed to parse cluster descriptor")
	}
	fmt.Println("File(4) Step(2) - Print - Cluster descriptor parsed") // Added by JANR

	provider := cluster.NewProvider(
		cluster.ProviderWithLogger(logger),
		runtime.GetDefault(logger),
	)
	fmt.Println("File(4) Step(2) - Print - Provider created") // Added by JANR

	clusterCredentials, err := provider.Validate( // Here we validate the cluster
		*keosCluster,
		clusterConfig,
		secretsDefaultPath,
		flags.VaultPassword,
	)
	if err != nil {
		return errors.Wrap(err, "failed to validate cluster")
	}

	dockerRegUrl := ""
	if clusterConfig != nil && clusterConfig.Spec.Private {
		configFile, err := getConfigFile(keosCluster, clusterCredentials)
		if err != nil {
			return errors.Wrap(err, "Error getting private kubeadm config")
		}
		flags.Config = configFile
		for _, dockerReg := range keosCluster.Spec.DockerRegistries {
			if dockerReg.KeosRegistry {
				dockerRegUrl = dockerReg.URL
			}
		}
	}

	if flags.ValidateOnly {
		fmt.Println("Cluster descriptor is valid")
		return nil
	}

	// handle config flag, we might need to read from stdin
	withConfig, err := configOption(flags.Config, streams.In)
	if err != nil {
		return err
	}

	// create the cluster
	if err = provider.Create(
		flags.Name,
		flags.VaultPassword,
		flags.DescriptorPath,
		flags.MoveManagement,
		flags.AvoidCreation,
		dockerRegUrl,
		clusterConfig,
		*keosCluster,
		clusterCredentials,
		withConfig,
		cluster.CreateWithNodeImage(flags.ImageName),
		cluster.CreateWithRetain(flags.Retain),
		cluster.CreateWithMove(flags.MoveManagement),
		cluster.CreateWithAvoidCreation(flags.AvoidCreation),
		cluster.CreateWithForceDelete(flags.ForceDelete),
		cluster.CreateWithWaitForReady(flags.Wait),
		cluster.CreateWithKubeconfigPath(flags.Kubeconfig),
		cluster.CreateWithDisplayUsage(true),
		cluster.CreateWithDisplaySalutation(true),
	); err != nil {
		return errors.Wrap(err, "failed to create cluster")
	}

	return nil
}

// configOption converts the raw --config flag value to a cluster creation
// option matching it. it will read from stdin if the flag value is `-`
func configOption(rawConfigFlag string, stdin io.Reader) (cluster.CreateOption, error) {
	// if not - then we are using a real file
	if rawConfigFlag != "-" {
		return cluster.CreateWithConfigFile(rawConfigFlag), nil
	}
	// otherwise read from stdin
	raw, err := io.ReadAll(stdin)
	if err != nil {
		return nil, errors.Wrap(err, "error reading config from stdin")
	}
	return cluster.CreateWithRawConfig(raw), nil
}

func setPassword(secretsDefaultPath string) (string, error) {
	// Print information in different lines: // Added by JANR
	// Relative path: // Added by JANR
	// Brief function goal: // Added by JANR
	// All functions called in order: // Added by JANR
	fmt.Println("File(4) Step(4) Path: Skind/pkg/cmd/kind/create/cluster/createcluster.go - Function: setPassword()") // Added by JANR
	fmt.Println("File(4) Step(4) Brief function goal: setPassword sets the vault password")                           // Added by JANR
	fmt.Println("File(4) Step(4) All functions called in order: requestPassword")                                     // Added by JANR
	firstPassword, err := requestPassword("Vault Password: ")
	if err != nil {
		return "", err
	}

	if _, err := os.Stat(secretsDefaultPath); os.IsNotExist(err) {
		secondPassword, err := requestPassword("Rewrite Vault Password:")
		if err != nil {
			return "", err
		}
		if firstPassword != secondPassword {
			return "", errors.New("The passwords do not match.")
		}
	}

	return firstPassword, nil
}

func requestPassword(request string) (string, error) {
	// Print information in different lines: // Added by JANR
	// Relative path: // Added by JANR
	// Brief function goal: // Added by JANR
	// All functions called in order: // Added by JANR
	fmt.Println("File(4) Step(5) Path: Skind/pkg/cmd/kind/create/cluster/createcluster.go - Function: requestPassword()") // Added by JANR
	fmt.Println("File(4) Step(5) Brief function goal: requestPassword requests the password")                             // Added by JANR
	fmt.Println("File(4) Step(5) All functions called in order: term.ReadPassword")                                       // Added by JANR
	fmt.Print(request)
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}
	fmt.Print("\n")
	return string(bytePassword), nil
}

func validateFlags(flags *flagpole) error {
	// Print information in different lines: // Added by JANR
	// Relative path: // Added by JANR
	// Brief function goal: // Added by JANR
	// All functions called in order: // Added by JANR
	fmt.Println("File(4) Step(3) Path: Skind/pkg/cmd/kind/create/cluster/createcluster.go - Function: validateFlags()") // Added by JANR
	fmt.Println("File(4) Step(3) Brief function goal: validateFlags validates the flags for the cluster creation")      // Added by JANR
	count := 0
	if flags.AvoidCreation {
		count++
	}
	if flags.Retain {
		count++
	}
	if flags.MoveManagement {
		count++
	}
	if count > 1 {
		return errors.New("Flags --retain, --avoid-creation, and --keep-mgmt are mutually exclusive")
	}
	return nil
}

// JANR: Improvement PDTE

//func validateFlags(flags *flagpole) error {
//	// Print information in different lines: // Added by JANR
//	// Relative path: // Added by JANR
//	// Brief function goal: // Added by JANR
//	// All functions called in order: // Added by JANR
//	fmt.Println("(4) Path: skin/pkg/cmd/kind/create/cluster/createcluster.go - validateFlags()")       // Added by JANR
//	fmt.Println("(4) Brief function goal: validateFlags validates the flags for the cluster creation") // Added by JANR
//	var conflictingFlags []string
//	if flags.AvoidCreation {
//		conflictingFlags = append(conflictingFlags, "--avoid-creation")
//	}
//	if flags.Retain {
//		conflictingFlags = append(conflictingFlags, "--retain")
//	}
//	if flags.MoveManagement {
//		conflictingFlags = append(conflictingFlags, "--keep-mgmt")
//	}
//	if len(conflictingFlags) > 1 {
//		return fmt.Errorf("flags %v are mutually exclusive", conflictingFlags)
//	}
//	return nil
//}
