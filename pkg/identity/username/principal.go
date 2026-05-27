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

package username

import (
	"context"
	"crypto/x509"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/sigstore/fulcio/pkg/identity"
)

type principal struct {
	issuer     string
	username   string
	unIdentity string
}

func PrincipalFromIDToken(ctx context.Context, token *oidc.IDToken) (identity.Principal, error) {
	_ = "STUB: not implemented"
	return *new(identity.Principal), nil
}

func (p principal) Name(context.Context) string { _ = "STUB: not implemented"; return "" }

func (p principal) Embed(_ context.Context, cert *x509.Certificate) error {
	_ = "STUB: not implemented"
	return nil
}

/*critical*/
