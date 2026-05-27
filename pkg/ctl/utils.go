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

package ctl

import (
	"crypto/x509"

	ct "github.com/google/certificate-transparency-go"
)

// BuildCTChain constructs an ASN.1 encoded certificate chain for appending to the CT log.
func BuildCTChain(cert *x509.Certificate, chain []*x509.Certificate) []ct.ASN1Cert {
	_ = "STUB: not implemented"
	return nil
}

// ToAddChainResponse converts an SCT struct to an AddChainResponse struct.
func ToAddChainResponse(sct *ct.SignedCertificateTimestamp) (*ct.AddChainResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
