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
//

// Package certmaker provides template parsing and certificate generation functionality
// for creating X.509 certificates from JSON templates per RFC3161 standards.
package certmaker

import (
	"crypto"
	"crypto/x509"
	_ "embed"
	"time"
)

//go:embed templates/root-template.json
var rootTemplate string

//go:embed templates/intermediate-template.json
var intermediateTemplate string

//go:embed templates/leaf-template.json
var leafTemplate string

func ParseTemplate(input any, parent *x509.Certificate, notAfter time.Time, publicKey crypto.PublicKey, commonName string) (*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get cert life and subject from template

// Get CN from template

// Create base cert with public key

// Set parent cert info

// Ensure cert life is set

func determinePublicKeyAlgorithm(publicKey crypto.PublicKey) x509.PublicKeyAlgorithm {
	_ = "STUB: not implemented"
	return *new(x509.PublicKeyAlgorithm)
}

// Default to ECDSA if key type is unknown

// Performs validation checks on the cert template
func ValidateTemplate(filename string, _ *x509.Certificate, _ string) error {
	_ = "STUB: not implemented"
	return nil
}

// Parse template via x509util to avoid issues with templating

// Returns a default JSON template string for the specified cert type
func GetDefaultTemplate(certType string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Both intermediate and leaf are optional - return empty if not found

// Ensures that required templates are present
func ValidateTemplateRequirements() error { _ = "STUB: not implemented"; return nil }
