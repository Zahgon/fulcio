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

package certmaker

import (
	"crypto/x509"
	"errors"

	"github.com/sigstore/sigstore/pkg/signature"
)

var (
	// returned when certificate file is not found
	ErrCertificateNotFound = errors.New("certificate file not found")

	// returned when PEM decoding fails or wrong block type
	ErrInvalidPEM = errors.New("invalid PEM format")

	// returned when PEM block contains no certificate
	ErrNoCertificateData = errors.New("no certificate data in PEM block")

	// returned when certificate public key doesn't match KMS key
	ErrKeyMismatch = errors.New("certificate public key does not match KMS key")
)

// reads a PEM-encoded X.509 certificate from the
// specified file path and returns the parsed certificate.
func LoadCertificateFromFile(path string) (*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Minimal validation to preserve existing error semantics

// Use cryptoutils to parse one or more certs, return the first

// verifies that the public key in the given certificate
// matches the public key from the KMS SignerVerifier.
func ValidateCertificateKeyMatch(cert *x509.Certificate, sv signature.SignerVerifier) error {
	_ = "STUB: not implemented"
	return nil
}

// get public key from KMS

// get public key from cert

// compare public keys using sigstore cryptoutils
