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

package tinkca

import (
	"context"

	"github.com/sigstore/fulcio/pkg/ca"
	"github.com/sigstore/fulcio/pkg/ca/baseca"
	"github.com/tink-crypto/tink-go/v2/tink"
)

type tinkCA struct {
	baseca.BaseCA
}

// NewTinkCA creates a signer from an encrypted Tink keyset, encrypted with a GCP KMS key.
func NewTinkCA(ctx context.Context, kmsKey, tinkKeysetPath, certPath string) (ca.CertificateAuthority, error) {
	_ = "STUB: not implemented"
	return *new(ca.CertificateAuthority), nil
}

// NewTinkCAFromHandle creates a signer from an encrypted Tink keyset, encrypted with an AEAD key.
func NewTinkCAFromHandle(_ context.Context, tinkKeysetPath, certPath string, primaryKey tink.AEAD) (ca.CertificateAuthority, error) {
	_ = "STUB: not implemented"
	return *new(ca.CertificateAuthority), nil
}

// GetPrimaryKey returns a Tink AEAD encryption key from KMS
// Supports GCP and AWS
func GetPrimaryKey(ctx context.Context, kmsKey string) (tink.AEAD, error) {
	_ = "STUB: not implemented"
	return *new(tink.AEAD), nil
}
