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
	"context"
	"crypto/x509"
	"flag"
	"log"
	"os"

	privateca "cloud.google.com/go/security/privateca/apiv1"
	"github.com/sigstore/sigstore/pkg/cryptoutils"

	// Register the provider-specific plugins

	_ "github.com/sigstore/sigstore/pkg/signature/kms/aws"
	_ "github.com/sigstore/sigstore/pkg/signature/kms/azure"
	_ "github.com/sigstore/sigstore/pkg/signature/kms/gcp"
	_ "github.com/sigstore/sigstore/pkg/signature/kms/hashivault"
)

/*
To run:
go run cmd/fetch_ca_cert/fetch_ca_cert.go \
  --kms-resource="gcpkms://projects/<project>/locations/<region>/keyRings/<key-ring>/cryptoKeys/<key>/versions/1" \
  --gcp-ca-parent="projects/<project>/locations/<region>/caPools/<ca-pool>" \
  --output="chain.crt.pem"

go run cmd/fetch_ca_cert/fetch_ca_cert.go \
  --tink-kms-resource="gcp-kms://projects/<project>/locations/<region>/keyRings/<key-ring>/cryptoKeys/<key>" \
  --tink-keyset-path="enc-keyset.cfg" \
  --gcp-ca-parent="projects/<project>/locations/<region>/caPools/<ca-pool>" \
  --output="chain.crt.pem"

You must have the permissions to read the KMS key, and create a certificate in the CA pool.
*/

var (
	gcpCaParent    = flag.String("gcp-ca-parent", "", "Resource path to GCP CA Service CA")
	kmsKey         = flag.String("kms-resource", "", "Resource path to KMS key, starting with gcpkms://, awskms://, azurekms:// or hashivault://")
	tinkKeysetPath = flag.String("tink-keyset-path", "", "Path to Tink keyset")
	tinkKmsKey     = flag.String("tink-kms-resource", "", "Resource path to KMS key to decrypt Tink keyset, starting with gcp-kms:// or aws-kms://")
	outputPath     = flag.String("output", "", "Path to the output file")
)

func fetchCACertificate(ctx context.Context, parent, kmsKey, tinkKeysetPath, tinkKmsKey string,
	client *privateca.CertificateAuthorityClient) ([]*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// default value of 0 for int32

// Default to a very large lifetime - CA Service will truncate the
// lifetime to be no longer than the root's lifetime.
// 20 years (24 hours * 365 days * 20)

func main() {
	flag.Parse()

	if *gcpCaParent == "" {
		log.Fatal("gcp-ca-parent must be set")
	}
	if *kmsKey == "" && *tinkKeysetPath == "" {
		log.Fatal("either kms-resource or tink-keyset-path must be set")
	}
	if *tinkKeysetPath != "" && *tinkKmsKey == "" {
		log.Fatal("tink-keyset-path must be set with tink-kms-resource must be set")
	}
	if *outputPath == "" {
		log.Fatal("output must be set")
	}

	client, err := privateca.NewCertificateAuthorityClient(context.Background())
	if err != nil {
		client.Close()
		log.Fatal(err)
	}
	parsedCerts, err := fetchCACertificate(context.Background(), *gcpCaParent, *kmsKey, *tinkKeysetPath, *tinkKmsKey, client)
	if err != nil {
		client.Close()
		log.Fatal(err)
	}
	pemCerts, err := cryptoutils.MarshalCertificatesToPEM(parsedCerts)
	if err != nil {
		client.Close()
		log.Fatal(err)
	}

	err = os.WriteFile(*outputPath, pemCerts, 0600)
	if err != nil {
		client.Close()
		log.Fatal(err)
	}
	defer client.Close()
}
