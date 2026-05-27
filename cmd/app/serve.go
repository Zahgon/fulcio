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
	"context"

	v1 "github.com/sigstore/protobuf-specs/gen/pb-go/common/v1"
	"github.com/sigstore/sigstore/pkg/signature"

	ctclient "github.com/google/certificate-transparency-go/client"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	certauth "github.com/sigstore/fulcio/pkg/ca"
	"github.com/sigstore/fulcio/pkg/config"
	"github.com/sigstore/fulcio/pkg/identity"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const (
	serveCmdEnvPrefix        = "FULCIO_SERVE"
	defaultConfigPath string = "/etc/fulcio-config/config.yaml"
)

var serveCmdConfigFilePath string

func newServeCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// convert "http-host" flag to "host" and "http-port" flag to be "port"

const (
	maxMsgSize int64 = 1 << 22 // 4MiB
)

// Adaptor for logging with the CT log
type logAdaptor struct {
	logger *zap.SugaredLogger
}

func (la logAdaptor) Printf(s string, args ...any) { _ = "STUB: not implemented"; return }

func runServeCmd(cmd *cobra.Command, args []string) {
	_ = "STUB: not implemented" //nolint: revive
	return
}

// If a config file is provided, modify the viper config to locate and read it

// Allow recognition of environment variables such as FULCIO_SERVE_CA etc.

// There's a MarkDeprecated function in cobra/pflags, but it doesn't use log.Logger

// this is a no-op since this is a self-signed in-memory CA for testing

// Setup the logger to dev/prod

// from https://github.com/golang/glog/commit/fca8c8854093a154ff1eb580aae10276ad6b1b5f

// optionally add CT log public key to verify SCTs

// StartDuplexServer will always return an error, log fatally if it's non-nil

// waiting for http and grpc servers to shutdown gracefully

// received an interrupt signal, shut down

// error from closing listeners, or context timeout

// wait for http and grpc servers to shutdown

func checkServeCmdConfigFile() error { _ = "STUB: not implemented"; return nil }

func duplexHealthz(_ context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	_ = "STUB: not implemented"
	return nil
}

func StartDuplexServer(ctx context.Context, cfg *config.FulcioConfig, ctClient *ctclient.LogClient, baseca certauth.CertificateAuthority, algorithmRegistry *signature.AlgorithmRegistryConfig, host string, port, metricsPort int, ip identity.IssuerPool) error {
	_ = "STUB: not implemented"
	return nil
}

// recovers from per-transaction panics elegantly, so put it first

// GRPC server

// Legacy server

// Prometheus

// Healthz

// Register prometheus handle.

func buildDefaultClientSigningAlgorithms(allowedAlgorithms []v1.PublicKeyDetails) []string {
	_ = "STUB: not implemented"
	return nil
}
