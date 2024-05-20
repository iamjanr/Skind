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

package validate

import (
	"fmt" // Added by JANR

	"sigs.k8s.io/kind/pkg/commons"
)

type ValidateParams struct {
	KeosCluster   commons.KeosCluster
	ClusterConfig *commons.ClusterConfig
	SecretsPath   string
	VaultPassword string
}

func Cluster(params *ValidateParams) (commons.ClusterCredentials, error) {
	// Print information in different lines: // Added by JANR
	// Relative path: // Added by JANR
	// Brief function goal: // Added by JANR
	// All functions called in order: // Added by JANR
	fmt.Println("File(7) Step(1) Path: Skind/pkg/cluster/internal/validate/validate.go - Function: Cluster()") // Added by JANR
	fmt.Println("File(7) Step(1) Brief function goal: Validates the credentials using the validateCredentials function.")
	fmt.Println("File(7) Step(1) Brief function goal: If a ClusterConfig is provided, it extracts the Spec from it.")
	fmt.Println("File(7) Step(1) Brief function goal: Validates common parameters between the KeosCluster and the ClusterConfigSpec using the validateCommon function.")
	fmt.Println("File(7) Step(1) Brief function goal: Depending on the infrastructure provider specified in the KeosCluster (aws, gcp, or azure), it validates the specific provider's parameters using the corresponding function (validateAWS, validateGCP, or validateAzure).")

	var err error
	var creds commons.ClusterCredentials

	creds, err = validateCredentials(*params)
	if err != nil {
		return commons.ClusterCredentials{}, err
	}
	fmt.Println("File(7) Step(1) - Print - creds: ", creds) // Added by JANR
	clusterConfigSpec := commons.ClusterConfigSpec{}
	if params.ClusterConfig != nil {
		clusterConfigSpec = params.ClusterConfig.Spec
	}
	if err := validateCommon(params.KeosCluster.Spec, clusterConfigSpec); err != nil {
		return commons.ClusterCredentials{}, err
	}

	switch params.KeosCluster.Spec.InfraProvider {
	case "aws":
		err = validateAWS(params.KeosCluster.Spec, creds.ProviderCredentials)
	case "gcp":
		err = validateGCP(params.KeosCluster.Spec, creds.ProviderCredentials)
	case "azure":
		err = validateAzure(params.KeosCluster.Spec, creds.ProviderCredentials, params.KeosCluster.Metadata.Name)
	}
	if err != nil {
		return commons.ClusterCredentials{}, err
	}

	return creds, nil
}
