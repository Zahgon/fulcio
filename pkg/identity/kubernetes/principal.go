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

package kubernetes

import (
	"context"
	"crypto/x509"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/sigstore/fulcio/pkg/identity"
)

type principal struct {
	// Subject ('sub') from ID token
	subject string

	// Issuer ('iss') from ID token
	issuer string

	// URI to be set in certificate. URI is of the form
	// https://kubernetes.io/namespaces/<namespace>/serviceaccounts/<serviceaccount>.
	uri string
}

func PrincipalFromIDToken(_ context.Context, token *oidc.IDToken) (identity.Principal, error) {
	_ = "STUB: not implemented"
	return *new(identity.Principal), nil
}

func (p principal) Name(context.Context) string { _ = "STUB: not implemented"; return "" }

func (p principal) Embed(_ context.Context, cert *x509.Certificate) error {
	_ = "STUB: not implemented"
	return nil
}

func kubernetesToken(token *oidc.IDToken) (string, error) {
	_ = "STUB: not implemented"
	// Extract custom claims
	return "", nil
}

// "kubernetes.io": {
//   "namespace": "default",
//   "pod": {
// 	    "name": "oidc-test",
// 	    "uid": "49ad3572-b3dd-43a6-8d77-5858d3660275"
//   },
//   "serviceaccount": {
// 	    "name": "default",
//      "uid": "f5720c1d-e152-4356-a897-11b07aff165d"
//   }
// }

// We use this in URIs, so it has to be a URI.
