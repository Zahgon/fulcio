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
//

package baseca

import (
	"context"
	"crypto"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"

	ct "github.com/google/certificate-transparency-go"
	"github.com/sigstore/fulcio/pkg/ca"
	"github.com/sigstore/fulcio/pkg/identity"
)

var (
	// OIDExtensionCTPoison is defined in RFC 6962 s3.1.
	OIDExtensionCTPoison = asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 11129, 2, 4, 3}
	// OIDExtensionCTSCT is defined in RFC 6962 s3.3.
	OIDExtensionCTSCT = asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 11129, 2, 4, 2}
)

type BaseCA struct {
	// contains the chain of certificates and signer
	ca.SignerWithChain
}

func (bca *BaseCA) CreatePrecertificate(ctx context.Context, principal identity.Principal, publicKey crypto.PublicKey) (*ca.CodeSigningPreCertificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Append poison extension

// From https://github.com/letsencrypt/boulder/blob/54b697d51b9f63cfd6055577cd317d4096aeab08/issuance/issuance.go#L497
func generateSCTListExt(scts []ct.SignedCertificateTimestamp) (pkix.Extension, error) {
	_ = "STUB: not implemented"
	return *new(pkix.Extension), nil
}

func (bca *BaseCA) IssueFinalCertificate(_ context.Context, precert *ca.CodeSigningPreCertificate, sct *ct.SignedCertificateTimestamp) (*ca.CodeSigningCertificate, error) {
	_ = "STUB: not implemented"
	// remove poison extension from precertificate.
	return nil, nil
}

// append SCT extension. Supports multiple SCTs, but Fulcio only writes to one log currently.

func (bca *BaseCA) CreateCertificate(ctx context.Context, principal identity.Principal, publicKey crypto.PublicKey) (*ca.CodeSigningCertificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (bca *BaseCA) TrustBundle(_ context.Context) ([][]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (bca *BaseCA) Close() error { _ = "STUB: not implemented"; return nil }
