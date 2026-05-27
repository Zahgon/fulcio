//
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

package api

import (
	"net/http"
	"net/url"
	"time"
)

type CertificateResponse struct {
	CertPEM  []byte
	ChainPEM []byte
	SCT      []byte
}

type RootResponse struct {
	ChainPEM []byte
}

type Key struct {
	// +required
	Content   []byte `json:"content"`
	Algorithm string `json:"algorithm,omitempty"`
}

type CertificateRequest struct {
	// +optional
	PublicKey Key `json:"publicKey"`

	// +optional
	SignedEmailAddress []byte `json:"signedEmailAddress"`

	// +optional
	CertificateSigningRequest []byte `json:"certificateSigningRequest"`
}

const (
	signingCertPath = "/api/v1/signingCert"
	rootCertPath    = "/api/v1/rootCert"
)

// SigstorePublicServerURL is the URL of Sigstore's public Fulcio service.
const SigstorePublicServerURL = "https://fulcio.sigstore.dev"

// LegacyClient is the interface for accessing the Fulcio API.
type LegacyClient interface {
	// SigningCert sends the provided CertificateRequest to the /api/v1/signingCert
	// endpoint of a Fulcio API, authenticated with the provided bearer token.
	SigningCert(cr CertificateRequest, token string) (*CertificateResponse, error)
	// RootCert sends a request to get the current CA used by Fulcio.
	RootCert() (*RootResponse, error)
}

// ClientOption is a functional option for customizing static signatures.
type ClientOption func(*clientOptions)

// NewClient creates a new Fulcio API client talking to the provided URL.
func NewClient(url *url.URL, opts ...ClientOption) LegacyClient {
	_ = "STUB: not implemented"
	return *new(LegacyClient)
}

type client struct {
	baseURL *url.URL
	client  *http.Client
}

var _ LegacyClient = (*client)(nil)

// SigningCert implements Client
func (c *client) SigningCert(cr CertificateRequest, token string) (*CertificateResponse, error) {
	_ = "STUB: not implemented"
	// Construct the API endpoint for this handler
	return nil, nil
}

// Set the authorization header to our OIDC bearer token.

// Set the content-type to reflect we're sending JSON.

// The API should return a 201 Created on success.  If we see anything else,
// then turn the response body into an error.

// Extract the SCT from the response header.

// Split the cert and the chain

func (c *client) RootCert() (*RootResponse, error) {
	_ = "STUB: not implemented"
	// Construct the API endpoint for this handler
	return nil, nil
}

type clientOptions struct {
	UserAgent string
	Timeout   time.Duration
}

func makeOptions(opts ...ClientOption) *clientOptions { _ = "STUB: not implemented"; return nil }

// WithTimeout sets the request timeout for the client
func WithTimeout(timeout time.Duration) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

// WithUserAgent sets the media type of the signature.
func WithUserAgent(userAgent string) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

type roundTripper struct {
	http.RoundTripper
	UserAgent string
}

// RoundTrip implements `http.RoundTripper`
func (rt *roundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createRoundTripper(inner http.RoundTripper, o *clientOptions) http.RoundTripper {
	_ = "STUB: not implemented"
	return *new(http.RoundTripper)
}

// There's nothing to do...
