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

package cloudinit

import (
	"fmt"

	"gitlab.com/eupraxialabs/cluster-api-k0s/pkg/secret"
)

// rewrite to pass in version of k0s
const (
	controlPlaneCloudInit = `{{.Header}}
{{template "files" .WriteFiles}}
runcmd:
{{- template "commands" .PreK0sCommands }}
  - 'curl -sSLf https://get.k0s.sh | sudo K0S_VERSION=v1.22.3+k0s.0 sh'
{{- template "commands" .PostK0sCommands }}
`
)

// ControlPlaneInput defines the context to generate a controlplane instance user data.
type ControlPlaneInput struct {
	BaseUserData
	secret.Certificates
}

// NewInitControlPlane returns the user data string to be used on a controlplane instance.
func NewInitControlPlane(input *ControlPlaneInput) ([]byte, error) {
	input.Header = cloudConfigHeader
	input.WriteFiles = input.Certificates.AsFiles()
	input.WriteFiles = append(input.WriteFiles, input.AdditionalFiles...)
	input.WriteFiles = append(input.WriteFiles, input.ConfigFile)

	//	controlPlaneCloudJoinWithVersion := fmt.Sprintf(controlPlaneCloudInit, input.K0sVersion)
	controlPlaneCloudJoinWithVersion := fmt.Sprintf("K0s Version: %s", input.K0sVersion)
	userData, err := generate("InitControlplane", controlPlaneCloudJoinWithVersion, input)
	if err != nil {
		return nil, err
	}

	return userData, nil
}
