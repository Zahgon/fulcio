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

// Package certmaker implements a certificate creation utility for Fulcio.
// It supports creating root, intermediate, and leaf certs using (AWS, GCP, Azure, HashiVault).
package certmaker

import (
	"context"
	"crypto"
	"crypto/x509"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/sigstore/sigstore/pkg/signature"
	"github.com/sigstore/sigstore/pkg/signature/kms"

	// Initialize AWS KMS provider
	_ "github.com/sigstore/sigstore/pkg/signature/kms/aws"
	// Initialize Azure KMS provider
	_ "github.com/sigstore/sigstore/pkg/signature/kms/azure"
	// Initialize GCP KMS provider
	_ "github.com/sigstore/sigstore/pkg/signature/kms/gcp"
	// Initialize HashiVault KMS provider
	_ "github.com/sigstore/sigstore/pkg/signature/kms/hashivault"
)

// variable to allow stubbing in tests
var kmsGet = kms.Get

// CryptoSignerVerifier extends SignerVerifier with CryptoSigner capability
type CryptoSignerVerifier interface {
	signature.SignerVerifier
	CryptoSigner(context.Context, func(error)) (crypto.Signer, crypto.SignerOpts, error)
}

// KMSConfig holds config for KMS providers.
type KMSConfig struct {
	CommonName string
	Type       string
	KeyID      string
	Options    map[string]string
}

// InitKMS initializes KMS provider based on the given config, KMSConfig.
var InitKMS = func(ctx context.Context, config KMSConfig) (signature.SignerVerifier, error) {
	if err := ValidateKMSConfig(config); err != nil {
		return nil, fmt.Errorf("invalid KMS configuration: %w", err)
	}

	var sv signature.SignerVerifier
	var err error

	switch config.Type {
	case "awskms":
		ref := fmt.Sprintf("awskms:///%s", config.KeyID)
		if awsRegion := config.Options["aws-region"]; awsRegion != "" {
			os.Setenv("AWS_REGION", awsRegion)
		}
		sv, err = kmsGet(ctx, ref, crypto.SHA256)
		if err != nil {
			return nil, fmt.Errorf("failed to initialize AWS KMS: %w", err)
		}

	case "gcpkms":
		ref := fmt.Sprintf("gcpkms://%s", config.KeyID)
		if gcpCredsFile := config.Options["gcp-credentials-file"]; gcpCredsFile != "" {
			os.Setenv("GCP_CREDENTIALS_FILE", gcpCredsFile)
		}
		sv, err = kmsGet(ctx, ref, crypto.SHA256)
		if err != nil {
			return nil, fmt.Errorf("failed to initialize GCP KMS: %w", err)
		}

	case "azurekms":
		keyURI := config.KeyID
		if strings.HasPrefix(config.KeyID, "azurekms:name=") {
			nameStart := strings.Index(config.KeyID, "name=") + 5
			vaultIndex := strings.Index(config.KeyID, ";vault=")
			if vaultIndex != -1 {
				keyName := strings.TrimSpace(config.KeyID[nameStart:vaultIndex])
				vaultName := strings.TrimSpace(config.KeyID[vaultIndex+7:])
				keyURI = fmt.Sprintf("azurekms://%s.vault.azure.net/%s", vaultName, keyName)
			}
		}
		if config.Options != nil && config.Options["azure-tenant-id"] != "" {
			azureTenantID := config.Options["azure-tenant-id"]
			os.Setenv("AZURE_TENANT_ID", azureTenantID)
			os.Setenv("AZURE_ADDITIONALLY_ALLOWED_TENANTS", "*")
		}
		os.Setenv("AZURE_AUTHORITY_HOST", "https://login.microsoftonline.com/")

		sv, err = kmsGet(ctx, keyURI, crypto.SHA256)
		if err != nil {
			return nil, fmt.Errorf("failed to initialize Azure KMS: %w", err)
		}

	case "hashivault":
		keyURI := fmt.Sprintf("hashivault://%s", config.KeyID)
		if config.Options != nil {
			if vaultToken := config.Options["vault-token"]; vaultToken != "" {
				os.Setenv("VAULT_TOKEN", vaultToken)
			}
			if vaultAddr := config.Options["vault-address"]; vaultAddr != "" {
				os.Setenv("VAULT_ADDR", vaultAddr)
			}
			if vaultNamespace := config.Options["vault-namespace"]; vaultNamespace != "" {
				os.Setenv("VAULT_NAMESPACE", vaultNamespace)
			}
		}

		sv, err = kmsGet(ctx, keyURI, crypto.SHA256)
		if err != nil {
			return nil, fmt.Errorf("failed to initialize HashiVault KMS: %w", err)
		}

	default:
		return nil, fmt.Errorf("unsupported KMS type: %s", config.Type)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to get KMS signer: %w", err)
	}
	if sv == nil {
		return nil, fmt.Errorf("KMS returned nil signer")
	}

	return sv, nil
}

// CreateCertificates creates certificates using the provided KMS and templates.
// Root certificate is always required (either generated or loaded from existingRootCertPath).
// Intermediate and leaf certificates are optional based on provided key IDs and templates.
// if existingRootCertPath or existingIntermediateCertPath are provided, those certificates
// will be loaded and used instead of generating new ones.
func CreateCertificates(config KMSConfig,
	rootTemplatePath, leafTemplatePath string,
	rootCertPath, leafCertPath string,
	intermediateKeyID, intermediateTemplatePath, intermediateCertPath string,
	leafKeyID string,
	rootLifetime, intermediateLifetime, leafLifetime time.Duration,
	existingRootCertPath, existingIntermediateCertPath string) error {
	_ = "STUB: not implemented"

	// Initialize root KMS signer
	return nil
}

// Get crypto.Signer for root

// Load existing root certificate or generate new one

// Load existing root certificate

// Validate that loaded certificate matches root KMS key

// Generate new root certificate

// Use default root template if none provided

// Read from FS if path is provided

// Create or load intermediate cert (optional)

// Load existing intermediate certificate

// If intermediate key ID is provided, validate cert matches the key

// Get signer for intermediate

// No intermediate key provided - can't sign with this cert
// This would be an error case - you need the key to sign

// Generate new intermediate certificate

// Read from FS if path is provided

// Create leaf cert (optional)

// Read from FS if path is provided

// Writes cert to a PEM-encoded file
func WriteCertificateToFile(cert *x509.Certificate, filename string) error {
	_ = "STUB: not implemented"
	return nil
}

// Get certificate type

// Ensures all required KMS config params are present
func ValidateKMSConfig(config KMSConfig) error { _ = "STUB: not implemented"; return nil }

// Root key is always required

// AWS KMS validation

// GCP KMS validation

// Azure KMS validation

// HashiVault KMS validation
