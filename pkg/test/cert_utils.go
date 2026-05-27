// Copyright 2022 The Sigstore Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package test

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/x509"
)

/*
To use:

rootCert, rootKey, _ := GenerateRootCA()
subCert, subKey, _ := GenerateSubordinateCa(rootCert, rootKey)
leafCert, _, _ := GenerateLeafCert("subject@example.com", "oidc-issuer", subCert, subKey)

roots := x509.NewCertPool()
subs := x509.NewCertPool()
roots.AddCert(rootCert)
subs.AddCert(subCert)
opts := x509.VerifyOptions{
	Roots:         roots,
	Intermediates: subs,
	KeyUsages: []x509.ExtKeyUsage{
		x509.ExtKeyUsageCodeSigning,
	},
}
_, err := leafCert.Verify(opts)
*/

func createCertificate(template *x509.Certificate, parent *x509.Certificate, pub any, priv crypto.Signer) (*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GenerateRootCA() (*x509.Certificate, *ecdsa.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func GenerateRootCAFromSigner(signer crypto.Signer) (*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GenerateSubordinateCA(rootTemplate *x509.Certificate, rootPriv crypto.Signer) (*x509.Certificate, *ecdsa.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func GenerateWeakSubordinateCA(rootTemplate *x509.Certificate, rootPriv crypto.Signer) (*x509.Certificate, *ecdsa.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func GenerateSubordinateCAWithoutEKU(rootTemplate *x509.Certificate, rootPriv crypto.Signer) (*x509.Certificate, *ecdsa.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func GenerateLeafCert(subject string, oidcIssuer string, parentTemplate *x509.Certificate, parentPriv crypto.Signer) (*x509.Certificate, *ecdsa.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// OID for OIDC Issuer extension
