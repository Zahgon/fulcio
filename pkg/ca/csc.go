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

package ca

import (
	"crypto/x509"
)

type CodeSigningCertificate struct {
	FinalCertificate *x509.Certificate
	FinalChain       []*x509.Certificate
	finalPEM         string
	finalChainPEM    []string
}

func CreateCSCFromPEM(cert string, chain []string) (*CodeSigningCertificate, error) {
	_ = "STUB: not implemented"
	return nil,

		// convert to X509 and store both formats
		nil
}

// convert to X509 and store both formats

func CreateCSCFromDER(cert []byte, chain []*x509.Certificate) (*CodeSigningCertificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// convert to X509 and store both formats

// convert to X509 and store both formats

func (c *CodeSigningCertificate) CertPEM() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *CodeSigningCertificate) ChainPEM() ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
