//go:build cgo

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

package app

import (
	"github.com/spf13/cobra"
)

const (
	LABEL = "PKCS11CA"
)

func newCreateCACmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func runCreateCACmd(cmd *cobra.Command, args []string) {
	_ = "STUB: not implemented" //nolint: revive
	return
}

// Check if CA already exists (or a cert within the provided ID)

// Find the existing Key Pair
// TODO: We could make the TAG customizable

// Import the root CA into the HSM
// TODO: We could make the TAG customizable

// Save out the file in pem format for easy import to CTL chain

//nolint
