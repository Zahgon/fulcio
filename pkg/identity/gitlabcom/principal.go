// Copyright 2023 The Sigstore Authors.
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

package gitlabcom

import (
	"context"
	"crypto/x509"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/sigstore/fulcio/pkg/identity"
)

// Deprecated: Use ciprovider.ciPrincipal instead
type jobPrincipal struct {
	// Subject matches the 'sub' claim from the OIDC ID token this is what is
	// signed as proof of possession for Buildkite job identities
	subject string

	// OIDC Issuer URL. Matches 'iss' claim from ID token. The real issuer URL is
	// https://agent.buildkite.com/.well-known/openid-configuration
	issuer string

	// The URL of the GitLab instance. https://gitlab.com
	url string

	// Event that triggered this workflow run. E.g "push", "tag" etc
	eventName string

	// Pipeline ID
	pipelineID string

	// Ref of top-level pipeline definition. E.g. gitlab.com/my-group/my-project//.gitlab-ci.yml@refs/heads/main
	ciConfigRefURI string

	// Commit sha of top-level pipeline definition, and is
	// only populated when `ciConfigRefURI` is local to the GitLab instance
	ciConfigSha string

	// Repository building built
	repository string

	// ID to the source repo
	repositoryID string

	// Owner of the source repo (mutable)
	repositoryOwner string

	// ID of the source repo
	repositoryOwnerID string

	// job ID
	jobID string

	// Git ref being built
	ref string

	// Commit SHA being built
	sha string

	// ID of the runner
	runnerID int64

	// The type of runner used by the job. May be one of gitlab-hosted or self-hosted.
	runnerEnvironment string

	// Visibility of the source project
	projectVisibility string
}

// Deprecated: Use ciprovider.WorkflowPrincipalFromIDToken instead
func JobPrincipalFromIDToken(_ context.Context, token *oidc.IDToken) (identity.Principal, error) {
	_ = "STUB: not implemented"
	return *new(identity.Principal), nil
}

func (p jobPrincipal) Name(_ context.Context) string { _ = "STUB: not implemented"; return "" }

func (p jobPrincipal) Embed(_ context.Context, cert *x509.Certificate) error {
	_ = "STUB: not implemented"
	return nil
}

// ci_config_ref_uri claim is a URI that does not include protocol scheme so we need to normalize it

// default to https

// or use scheme from issuer if from the same host

// Set workflow ref URL to SubjectAlternativeName on certificate

// Embed additional information into custom extensions
