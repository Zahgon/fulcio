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

package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"

	fulciopb "github.com/sigstore/fulcio/pkg/generated/protobuf"
	"github.com/sigstore/sigstore/pkg/signature"
)

var (
	fulcioURL    = "https://fulcio.sigstore.dev"
	oidcIssuer   = "https://oauth2.sigstore.dev/auth"
	oidcClientID = "sigstore"
)

// Some of this is just ripped from cosign
func GetCert(signer *signature.ECDSASignerVerifier, fc fulciopb.CAClient, oidcIssuer string, oidcClientID string) (*fulciopb.SigningCertificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Sign the email address as part of the request

func NewClient(fulcioURL string) (fulciopb.CAClient, error) {
	_ = "STUB: not implemented"
	return *new(fulciopb.CAClient), nil
}

func main() {
	signer, _, err := signature.NewDefaultECDSASignerVerifier()
	if err != nil {
		log.Fatal(err)
	}

	fClient, err := NewClient(fulcioURL)
	if err != nil {
		log.Fatal(err)
	}

	certResp, err := GetCert(signer, fClient, oidcIssuer, oidcClientID)
	if err != nil {
		log.Fatal(err)
	}

	var chain *fulciopb.CertificateChain
	switch cert := certResp.Certificate.(type) {
	case *fulciopb.SigningCertificate_SignedCertificateDetachedSct:
		chain = cert.SignedCertificateDetachedSct.GetChain()
	case *fulciopb.SigningCertificate_SignedCertificateEmbeddedSct:
		chain = cert.SignedCertificateEmbeddedSct.GetChain()
	}
	clientPEM, _ := pem.Decode([]byte(chain.Certificates[0]))
	cert, err := x509.ParseCertificate(clientPEM.Bytes)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Received signing cerificate with serial number: ", cert.SerialNumber)
}
