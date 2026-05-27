// Copyright 2024 The Sigstore Authors.
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

package ciprovider

import (
	"context"
	"crypto/x509"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/sigstore/fulcio/pkg/config"
	"github.com/sigstore/fulcio/pkg/identity"
)

func mapValuesToString(claims map[string]any) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// A float, but with no fractional part. Treat as an int

func getTokenClaims(token *oidc.IDToken) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// applyTemplateOrReplace performs string interpolation on a given template string.
// It uses Go's text/template syntax (https://pkg.go.dev/text/template).
// The logMetadata parameter is included to provide richer log messages.
func applyTemplateOrReplace(
	extValueTemplate string, tokenClaims map[string]string,
	issuerMetadata map[string]string, logMetadata map[string]string) (string, error) {
	_ = "STUB: not implemented"

	// Merge the data from the ID token claims with the default data
	// from the issuer's metadata configuration.
	// The order of maps.Copy is important here. We copy tokenClaims last
	// to ensure that claim data takes priority over the default metadata.
	return "", nil
}

// The "missingkey=error" option ensures that if a claim referenced
// in the template is not present in the merged data, template
// execution will fail.

// It shouldn't raise error since we already checked all
// templates in validateCIIssuerMetadata functions in config.go

type ciPrincipal struct {
	Token          *oidc.IDToken
	ClaimsMetadata config.IssuerMetadata
}

func WorkflowPrincipalFromIDToken(ctx context.Context, token *oidc.IDToken) (identity.Principal, error) {
	_ = "STUB: not implemented"
	return *new(identity.Principal), nil
}

func (principal ciPrincipal) Name(_ context.Context) string { _ = "STUB: not implemented"; return "" }

func (principal ciPrincipal) Embed(_ context.Context, cert *x509.Certificate) error {
	_ = "STUB: not implemented"
	return nil
}

// We use value.Elem() here because we need an addressable
// reference to the template fields to modify them with SetString().

// The type of the reflect value is needed to get the field name
// for logging and for checking against the "Issuer" field.

// value of each field, e.g the template string
// We check the field name to avoid applying the template for the Issuer.
// The Issuer field should always come from the token issuer.

// Guarantee that the extension issuer and subject are set to the token values,
// regardless of whether these fields have been set before

// Embed additional information into custom extensions
