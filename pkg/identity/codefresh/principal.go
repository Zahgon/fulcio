// Copyright 2022 The Sigstore Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package codefresh

import (
	"context"
	"crypto/x509"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/sigstore/fulcio/pkg/identity"
)

const (
	DefaultPlatformURL string = "https://g.codefresh.io"
)

// Deprecated: Use ciprovider.ciPrincipal instead
type workflowPrincipal struct {
	// Subject matches the 'sub' claim from the OIDC ID token this is what is signed as proof of possession for Codefresh workflow identities
	subject string

	// OIDC Issuer URL. Matches 'iss' claim from ID token. The real issuer URL is
	// https://oidc.codefresh.io/.well-known/openid-configuration
	issuer string

	// Codefresh account id
	accountID string

	// Codefresh account name
	accountName string

	// Codefresh Pipeline id
	pipelineID string

	// Codefresh pipline name (project/pipeline)
	pipelineName string

	// The ID of the specific workflow authorized in the claim. For example, 64f447c02199f903000gh20.
	workflowID string

	// 	Applies to manual trigger types, and is the username of the user manually triggered the pipeline
	initiator string

	// Applies to Git push, PR, and manual Git trigger types. The SCM name of the user who initiated the Git action.
	scmUsername string

	// Applies to Git push, PR, and manual Git trigger types. The SCM URL specifying the Git repository’s location. For example, https://github.com/codefresh-user/oidc-test
	scmRepoURL string

	// Applies to Git push, PR, and manual Git trigger types. The SCM name of the branch or tag within the Git repository for which the workflow should execute. For example, main or v1.0.0.
	scmRef string

	// 	Applies to Git PR trigger types. The SCM target branch the pull request should merge into. For example, production
	scmPullRequestTargetBranch string

	// Whether the build took place in cloud or self-hosted infrastructure
	runnerEnvironment string

	// Codefresh platform url
	platformURL string
}

func (w workflowPrincipal) Name(_ context.Context) string {
	_ = "STUB: not implemented"

	// Deprecated: Use ciprovider.WorkflowPrincipalFromIDToken instead
	return ""
}

func WorkflowPrincipalFromIDToken(_ context.Context, token *oidc.IDToken) (identity.Principal, error) {
	_ = "STUB: not implemented"
	return *new(identity.Principal), nil
}

// Set default platform url in case it is missing in the token

func (w workflowPrincipal) Embed(_ context.Context, cert *x509.Certificate) error {
	_ = "STUB: not implemented"
	return nil
}

// Set SAN to the <platform url>/<account name>/<pipeline name>:<account id>/<pipeline id> - for example https://g.codefresh.io/codefresh-account/oidc-test/get-token:628a80b693a15c0f9c13ab75/65e5a53e52853dc51a5b0cc1
// In Codefresh account names and pipeline names may be changed where as IDs do not.
// This pattern will give users the possibility to verify the signature using various forms of `cosign verify --certificate-identity-regexp` i.e https://g.codefresh.io/codefresh-account/oidc-test/get-token:* or https://g.codefresh.io/*:628a80b693a15c0f9c13ab75/65e5a53e52853dc51a5b0cc1

// URL of the build in Codefresh.
// The workflow url is used for build signer in Codefresh because for public builds unauthenticated users only have access to the workflow, not the pipeline definition.
// Also, the workflow contains the definition of the pipeline that was used at the time of the build, making it ideal to be used as the signer url.
