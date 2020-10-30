/*
Copyright 2020 The Kubernetes Authors.

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

package main

import (
	"fmt"
	"os"
	"k8s.io/cloud-provider-gcp/cmd/auth-provider-gcp/app"
)

// Command-line argument names that can be used to configure base URLs. The
// keys correspond to the constants in credentialprovider/gcp/metadata.go.
const (
	metadataUrlArg              = "metadataURL"
	storageScopePrefixArg       = "storageScopePrefix"
	cloudPlatformScopePrefixArg = "cloudPlatformScope"
)

func main() {
	cmd, err := app.NewAuthProviderCommand()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if err := cmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
