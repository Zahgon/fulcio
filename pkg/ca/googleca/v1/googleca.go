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

package v1

import (
	"context"
	"crypto"
	"crypto/x509"
	"encoding/asn1"
	"sync"

	privateca "cloud.google.com/go/security/privateca/apiv1"
	"cloud.google.com/go/security/privateca/apiv1/privatecapb"
	"github.com/sigstore/fulcio/pkg/ca"
	"github.com/sigstore/fulcio/pkg/identity"
	"google.golang.org/api/option"
)

type CertAuthorityService struct {
	certAuthorityID       string
	certAuthorityResource string
	caPoolResource        string
	client                *privateca.CertificateAuthorityClient

	// protected by once
	cachedRoots     [][]*x509.Certificate
	cachedRootsOnce sync.Once
}

func NewCertAuthorityService(ctx context.Context, parent string, opts ...option.ClientOption) (ca.CertificateAuthority, error) {
	_ = "STUB: not implemented"
	return *new(ca.CertificateAuthority), nil
}

// parent should be in the form projects/*/locations/*/caPools/*/certificateAuthorities/*
// to create a cert, we only want projects/*/locations/*/caPools/*

func (c *CertAuthorityService) Close() error { _ = "STUB: not implemented"; return nil }

// getPubKeyFormat Returns the PublicKey KeyFormat required by gcp privateca.
// https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/security/privateca/v1#PublicKey_KeyType
func getPubKeyFormat(pemBytes []byte) (privatecapb.PublicKey_KeyFormat, error) {
	_ = "STUB: not implemented"
	return *new(privatecapb.PublicKey_KeyFormat), nil
}

func convertID(id asn1.ObjectIdentifier) []int32 { _ = "STUB: not implemented"; return nil }

func Req(parent, certAuthority string, pemBytes []byte, cert *x509.Certificate) (*privatecapb.CreateCertificateRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Translate the x509 certificate's subject to Google proto.

func (c *CertAuthorityService) TrustBundle(ctx context.Context) ([][]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	// if we've already successfully fetched the CA info, just use the cached value
	return nil, nil
}

// if a specific certificate authority was specified, use that one

// otherwise, get certs from all of the CAs in the pool

func (c *CertAuthorityService) getCertificateAuthorityTrustBundle(ctx context.Context) ([][]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if we fail to parse the PEM content, return an error

func (c *CertAuthorityService) listCertificateAuthorityTrustBundle(ctx context.Context) ([][]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	// fetch the latest values for the specified CA pool
	return nil, nil
}

// if the iterator returns an issue for some reason, exit

// if we fail to parse the PEM content, return an error

func (c *CertAuthorityService) CreateCertificate(ctx context.Context, principal identity.Principal, publicKey crypto.PublicKey) (*ca.CodeSigningCertificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
