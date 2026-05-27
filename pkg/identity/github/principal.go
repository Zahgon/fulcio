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

package github

import (
	"context"
	"crypto/x509"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/sigstore/fulcio/pkg/identity"
)

// Deprecated: Use ciprovider.ciPrincipal instead
type workflowPrincipal struct {
	// Subject matches the 'sub' claim from the OIDC ID token this is what is
	// signed as proof of possession for Github workflow identities
	subject string

	// OIDC Issuer URL. Matches 'iss' claim from ID token. The real issuer URL is
	// https://token.actions.githubusercontent.com/.well-known/openid-configution
	issuer string

	// URL of issuer
	url string

	// Commit SHA being built
	sha string

	// Event that triggered this workflow run. E.g "push", "tag"
	eventName string

	// Name of repository being built
	repository string

	// Deprecated
	// Name of workflow that is running (mutable)
	workflow string

	// Git ref being built
	ref string

	// Specific build instructions (i.e. reusable workflow)
	jobWorkflowRef string

	// Commit SHA to specific build instructions
	jobWorkflowSha string

	// Whether the build took place in cloud or self-hosted infrastructure
	runnerEnvironment string

	// ID to the source repo
	repositoryID string

	// Owner of the source repo (mutable)
	repositoryOwner string

	// ID of the source repo
	repositoryOwnerID string

	// Visibility of the source repo
	repositoryVisibility string

	// Ref of top-level workflow that is running
	workflowRef string

	// Commit SHA of top-level workflow that is running
	workflowSha string

	// ID of workflow run
	runID string

	// Attempt number of workflow run
	runAttempt string
}

// Deprecated: Use ciprovider.WorkflowPrincipalFromIDToken instead
func WorkflowPrincipalFromIDToken(_ context.Context, token *oidc.IDToken) (identity.Principal, error) {
	_ = "STUB: not implemented"
	return *new(identity.Principal), nil
}

func (w workflowPrincipal) Name(_ context.Context) string { _ = "STUB: not implemented"; return "" }

func (w workflowPrincipal) Embed(_ context.Context, cert *x509.Certificate) error {
	_ = "STUB: not implemented"
	return nil
}

// Set workflow ref URL to SubjectAlternativeName on certificate

// Embed additional information into custom extensions

// BEGIN: Deprecated

// END: Deprecated
