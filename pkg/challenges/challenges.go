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

package challenges

import (
	"context"
	"crypto"

	"github.com/sigstore/fulcio/pkg/identity"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/sigstore/sigstore/pkg/signature"
)

// CheckSignature verifies a challenge, a signature over the subject or email
// of an OIDC token
func CheckSignature(pub crypto.PublicKey, proof []byte, subject string) error {
	_ = "STUB: not implemented"
	return nil
}

// CheckSignatureWithVerifier verifies a challenge, a signature over the subject
// or email of an OIDC token
func CheckSignatureWithVerifier(verifier signature.Verifier, proof []byte, subject string) error {
	_ = "STUB: not implemented"
	return nil
}

func PrincipalFromIDToken(ctx context.Context, tok *oidc.IDToken) (identity.Principal, error) {
	_ = "STUB: not implemented"
	return *new(identity.Principal), nil
}

// nolint

// nolint

// nolint

// ParsePublicKey parses a PEM or DER encoded public key. Returns an error if
// decoding fails or if no public key is found.
func ParsePublicKey(encodedPubKey string) (crypto.PublicKey, error) {
	_ = "STUB: not implemented"
	return *new(crypto.PublicKey), nil
}

// try to unmarshal as PEM

// try to unmarshal as DER
