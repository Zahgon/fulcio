//go:build !cgo

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

// Just a placeholder for erroring with a meaningful message if the
// binary has been built with GCO_ENABLED=0 tags.
func newCreateCACmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func runCreateCACmdPlaceholder(cmd *cobra.Command, args []string) {
	_ = "STUB: not implemented"
	return
}
