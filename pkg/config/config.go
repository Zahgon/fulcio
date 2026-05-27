// Copyright 2021 The Sigstore Authors.
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
//

package config

import (
	"context"
	"net/http"
	"net/url"
	"regexp"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/sigstore/fulcio/pkg/certificate"
	fulciogrpc "github.com/sigstore/fulcio/pkg/generated/protobuf"
)

const defaultOIDCDiscoveryTimeout = 10 * time.Second

// All hostnames for subject and issuer OIDC claims must have at least a
// top-level and second-level domain
const minimumHostnameLength = 2

type verifierWithConfig struct {
	*oidc.IDTokenVerifier
	*oidc.Config
}

type bearerTokenTransport struct {
	Transport http.RoundTripper
	Token     string
	IssuerURL string
}

func (t *bearerTokenTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type FulcioConfig struct {
	OIDCIssuers map[string]OIDCIssuer `json:"OIDCIssuers,omitempty" yaml:"oidc-issuers,omitempty"`

	// A meta issuer has a templated URL of the form:
	//   https://oidc.eks.*.amazonaws.com/id/*
	// Where * can match a single hostname or URI path parts
	// (in particular, no '.' or '/' are permitted, among
	// other special characters)  Some examples we want to match:
	// * https://oidc.eks.us-west-2.amazonaws.com/id/B02C93B6A2D30341AD01E1B6D48164CB
	// * https://container.googleapis.com/v1/projects/mattmoor-credit/locations/us-west1-b/clusters/tenant-cluster
	MetaIssuers map[string]OIDCIssuer `json:"MetaIssuers,omitempty" yaml:"meta-issuers,omitempty"`

	// It defines metadata to be used for the CIProvider identity provider principal.
	// The CI provider has a generic logic for ci providers, this metadata is used
	// to define the right behavior for each ci provider that is defined
	// on the configuration file
	CIIssuerMetadata map[string]IssuerMetadata `json:"CIIssuerMetadata,omitempty" yaml:"ci-issuer-metadata,omitempty"`

	// verifiers is a fixed mapping from our OIDCIssuers to their OIDC verifiers.
	verifiers map[string][]*verifierWithConfig
	// lru is an LRU cache of recently used verifiers for our meta issuers.
	lru *lru.TwoQueueCache[string, []*verifierWithConfig]
}

type IssuerMetadata struct {
	// Defaults contains key-value pairs that can be used for filling the templates from ExtensionTemplates
	// If a key cannot be found on the token claims, the template will use the defaults
	DefaultTemplateValues map[string]string `json:"DefaultTemplateValues,omitempty" yaml:"default-template-values,omitempty"`
	// ExtensionTemplates contains a mapping between certificate extension and token claim
	// Provide either strings following https://pkg.go.dev/text/template syntax,
	// e.g "{{ .url }}/{{ .repository }}"
	// or non-templated strings with token claim keys to be replaced,
	// e.g "job_workflow_sha"
	ExtensionTemplates certificate.Extensions `json:"ExtensionTemplates,omitempty" yaml:"extension-templates,omitempty"`
	// Template for the Subject Alternative Name extension
	// It's typically the same value as Build Signer URI
	SubjectAlternativeNameTemplate string `json:"SubjectAlternativeNameTemplate,omitempty" yaml:"subject-alternative-name-template,omitempty"`
}

type OIDCIssuer struct {
	// The expected issuer of an OIDC token
	IssuerURL string `json:"IssuerURL,omitempty" yaml:"issuer-url,omitempty"`
	// The expected client ID of the OIDC token
	ClientID string `json:"ClientID" yaml:"client-id,omitempty"`
	// Used to determine the subject of the certificate and if additional
	// certificate values are needed
	Type IssuerType `json:"Type" yaml:"type,omitempty"`
	// CIProvider is an optional configuration to map token claims to extensions for CI workflows
	CIProvider string `json:"CIProvider,omitempty" yaml:"ci-provider,omitempty"`
	// Optional, if the issuer is in a different claim in the OIDC token
	IssuerClaim string `json:"IssuerClaim,omitempty" yaml:"issuer-claim,omitempty"`
	// The domain that must be present in the subject for 'uri' issuer types
	// Also used to create an email for 'username' issuer types
	SubjectDomain string `json:"SubjectDomain,omitempty" yaml:"subject-domain,omitempty"`
	// SPIFFETrustDomain specifies the trust domain that 'spiffe' issuer types
	// issue ID tokens for. Tokens with a different trust domain will be
	// rejected.
	SPIFFETrustDomain string `json:"SPIFFETrustDomain,omitempty" yaml:"spiffe-trust-domain,omitempty"`
	// Optional, the challenge claim expected for the issuer
	// Set if using a custom issuer
	ChallengeClaim string `json:"ChallengeClaim,omitempty" yaml:"challenge-claim,omitempty"`
	// Optional, the description for the issuer
	Description string `json:"Description,omitempty" yaml:"description,omitempty"`
	// Optional, the contact for the issuer team
	// Usually it is a email
	Contact string `json:"Contact,omitempty" yaml:"contact,omitempty"`

	// CACert is an optional parameter that holds the CA certificate in PEM format.
	// This is used to trust the TLS certificate signed by an internal CA when interacting
	// with some OIDC providers, preventing x509 certificate verification failures.
	CACert string `json:"CACert,omitempty" yaml:"ca-cert,omitempty"`

	// SkipEmailVerification skips the email_verified claim check for email-type issuers.
	// This should only be set to true for trusted internal identity providers (e.g., Microsoft Entra, ADFS)
	// that perform email verification through their own processes but don't include the email_verified claim.
	SkipEmailVerification bool `json:"SkipEmailVerification,omitempty" yaml:"skip-email-verification,omitempty"`
}

func MetaRegex(issuer string) (*regexp.Regexp, error) {
	_ = "STUB: not implemented"
	// Quote all of the "meta" characters like `.` to avoid
	// those literal characters in the URL matching any character.
	// This will ALSO quote `*`, so we replace the quoted version.
	return nil, nil
}

// Replace the quoted `*` with a regular expression that
// will match alpha-numeric parts with common additional
// "special" characters.

// Add anchors to the beginning and end of the regular expression
// to prevent matching URLs where the issuer is not the host of the URL,
// e.g. http://localhost:3000?https://meta-url-issuer.com/*
// Resolves GHSA-59jp-pj84-45mr

// Compile into a regular expression.

// GetIssuer looks up the issuer configuration for an `issuerURL`
// coming from an incoming OIDC token.  If no matching configuration
// is found, then it returns `false`.
func (fc *FulcioConfig) GetIssuer(issuerURL string) (OIDCIssuer, bool) {
	_ = "STUB: not implemented"
	return *new(OIDCIssuer), false
}

// Shouldn't happen, we check parsing the config

// If it matches, then return a concrete OIDCIssuer
// configuration for this issuer URL.

// GetVerifier fetches a token verifier for the given `issuerURL`
// coming from an incoming OIDC token.  If no matching configuration
// is found, then it returns `false`.
func (fc *FulcioConfig) GetVerifier(issuerURL string, opts ...InsecureOIDCConfigOption) (*oidc.IDTokenVerifier, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Look up our fixed issuer verifiers

// Look in the LRU cache for a verifier

// If this issuer hasn't been recently used, or we have special config options, then create a new verifier
// and add it to the LRU cache.

type InsecureOIDCConfigOption func(opt *oidc.Config)

func WithSkipExpiryCheck() InsecureOIDCConfigOption {
	_ = "STUB: not implemented"
	return *new(InsecureOIDCConfigOption)
}

// ToIssuers returns a proto representation of the OIDC issuer configuration.
func (fc *FulcioConfig) ToIssuers() []*fulciogrpc.OIDCIssuer { _ = "STUB: not implemented"; return nil }

func noRedirectClient(client *http.Client) *http.Client { _ = "STUB: not implemented"; return nil }

func httpClientForIssuer(fc *FulcioConfig, iss OIDCIssuer) (*http.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Add the Kubernetes cluster's CA to the client's CA pool

// add the authentication header

func (fc *FulcioConfig) prepare() error { _ = "STUB: not implemented"; return nil }

/* size */

var (
	k8sCA = "/var/run/fulcio/ca.crt"
	// k8sTokenFile specifies the standard path where Kubernetes automatically
	// mounts the projected service account token for a pod.
	// This path is publicly known and not a sensitive credential itself,
	// hence the Gosec G101 warning is a false positive and is suppressed.
	// #nosec G101
	k8sTokenFile = "/var/run/secrets/kubernetes.io/serviceaccount/token"
	k8sIssuerURL = "https://kubernetes.default.svc"
)

func (fc *FulcioConfig) insertVerifier(iss OIDCIssuer) error { _ = "STUB: not implemented"; return nil }

type IssuerType string

func (it IssuerType) String() string { _ = "STUB: not implemented"; return "" }

const (
	IssuerTypeBuildkiteJob      = "buildkite-job"
	IssuerTypeEmail             = "email"
	IssuerTypeGithubWorkflow    = "github-workflow"
	IssuerTypeCodefreshWorkflow = "codefresh-workflow"
	IssuerTypeGitLabPipeline    = "gitlab-pipeline"
	IssuerTypeChainguard        = "chainguard-identity"
	IssuerTypeKubernetes        = "kubernetes"
	IssuerTypeSpiffe            = "spiffe"
	IssuerTypeURI               = "uri"
	IssuerTypeUsername          = "username"
	IssuerTypeCIProvider        = "ci-provider"
)

func parseConfig(b []byte) (cfg *FulcioConfig, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateConfig(conf *FulcioConfig) error { _ = "STUB: not implemented"; return nil }

// verify that trust domain is valid

// The domain in the configuration must match the domain (excluding the subdomain) of the issuer
// In order to declare this configuration, a test must have been done to prove ownership
// over both the issuer and domain configuration values.
// Valid examples:
// * SubjectDomain = https://example.com, IssuerURL = https://accounts.example.com
// * SubjectDomain = https://accounts.example.com, IssuerURL = https://accounts.example.com
// * SubjectDomain = https://users.example.com, IssuerURL = https://accounts.example.com

// The domain in the configuration must match the domain (excluding the subdomain) of the issuer
// In order to declare this configuration, a test must have been done to prove ownership
// over both the issuer and domain configuration values.
// Valid examples:
// * SubjectDomain = example.com, IssuerURL = https://accounts.example.com
// * SubjectDomain = accounts.example.com, IssuerURL = https://accounts.example.com
// * SubjectDomain = users.example.com, IssuerURL = https://accounts.example.com

// This would establish a many to one relationship for OIDC issuers
// to trust domains so we fail early and reject this configuration.

var DefaultConfig = &FulcioConfig{
	OIDCIssuers: map[string]OIDCIssuer{
		"https://oauth2.sigstore.dev/auth": {
			IssuerURL:   "https://oauth2.sigstore.dev/auth",
			ClientID:    "sigstore",
			IssuerClaim: "$.federated_claims.connector_id",
			Type:        IssuerTypeEmail,
		},
		"https://accounts.google.com": {
			IssuerURL: "https://accounts.google.com",
			ClientID:  "sigstore",
			Type:      IssuerTypeEmail,
		},
		"https://token.actions.githubusercontent.com": {
			IssuerURL: "https://token.actions.githubusercontent.com",
			ClientID:  "sigstore",
			Type:      IssuerTypeGithubWorkflow,
		},
	},
}

type configKey struct{}

func With(ctx context.Context, cfg *FulcioConfig) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func FromContext(ctx context.Context) *FulcioConfig { _ = "STUB: not implemented"; return nil }

// It checks that the templates defined are parseable
// We should check it during the service bootstrap to avoid errors further
func validateCIIssuerMetadata(fulcioConfig *FulcioConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// Load a config from disk, or use defaults
func Load(configPath string) (*FulcioConfig, error) { _ = "STUB: not implemented"; return nil, nil }

// Read parses the bytes of a config
func Read(b []byte) (*FulcioConfig, error) { _ = "STUB: not implemented"; return nil, nil }

// isURISubjectAllowed compares the subject and issuer URIs,
// returning an error if the scheme or the hostnames do not match
func isURISubjectAllowed(subject, issuer *url.URL) error { _ = "STUB: not implemented"; return nil }

// validateAllowedDomain compares two hostnames, returning an error if the
// top-level and second-level domains do not match
// TODO: This does not work for domains that end in co.jp or co.uk. We should consider
// using eTLDs, or removing this validation when we can challenge domain ownership.
func validateAllowedDomain(subjectHostname, issuerHostname string) error {
	_ = "STUB: not implemented"
	// If the hostnames exactly match, return early
	return nil
}

// Compare the top level and second level domains

func issuerToChallengeClaim(issType IssuerType, challengeClaim string) string {
	_ = "STUB: not implemented"
	return ""
}
