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

package app

import (
	"context"
	"crypto/tls"
	"sync"

	"github.com/sigstore/sigstore/pkg/signature"

	"github.com/fsnotify/fsnotify"
	ctclient "github.com/google/certificate-transparency-go/client"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sigstore/fulcio/pkg/ca"
	"github.com/sigstore/fulcio/pkg/config"
	gw "github.com/sigstore/fulcio/pkg/generated/protobuf"
	"github.com/sigstore/fulcio/pkg/identity"
	"google.golang.org/grpc"
)

const (
	LegacyUnixDomainSocket = "@fulcio-legacy-grpc-socket"
)

type grpcServer struct {
	*grpc.Server
	grpcServerEndpoint string
	caService          gw.CAServer
	tlsCertWatcher     *fsnotify.Watcher
}

func PassFulcioConfigThruContext(cfg *config.FulcioConfig) grpc.UnaryServerInterceptor {
	_ = "STUB: not implemented"
	return *new(grpc.UnaryServerInterceptor)
}

// For each request, infuse context with our snapshot of the FulcioConfig.
// TODO(mattmoor): Consider periodically (every minute?) refreshing the ConfigMap
// from disk, so that we don't need to cycle pods to pick up config updates.
// Alternately we could take advantage of Knative's configmap watcher.

// Calls the inner handler

type cachedTLSCert struct {
	sync.RWMutex
	certPath string
	keyPath  string
	cert     *tls.Certificate
	Watcher  *fsnotify.Watcher
}

func newCachedTLSCert(certPath, keyPath string) (*cachedTLSCert, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Add a path.

func (c *cachedTLSCert) GetCertificate() *tls.Certificate {
	_ = "STUB: not implemented"
	// get reader lock
	return nil
}

func (c *cachedTLSCert) UpdateCertificate() error {
	_ = "STUB: not implemented"
	// get writer lock
	return nil
}

func (c *cachedTLSCert) GRPCCreds() grpc.ServerOption {
	_ = "STUB: not implemented"
	return *new(grpc.ServerOption)
}

func createGRPCServer(cfg *config.FulcioConfig, ctClient *ctclient.LogClient, baseca ca.CertificateAuthority, algorithmRegistry *signature.AlgorithmRegistryConfig, ip identity.IssuerPool) (*grpcServer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// recovers from per-transaction panics elegantly, so put it first

// Register your gRPC service implementations.

func (g *grpcServer) setupPrometheus(reg *prometheus.Registry) { _ = "STUB: not implemented"; return }

func (g *grpcServer) startTCPListener(wg *sync.WaitGroup) {
	_ = "STUB: not implemented"
	// lis is closed by g.Server.Serve() upon exit
	return
}

// received an interrupt signal, shut down

func (g *grpcServer) startUnixListener() { _ = "STUB: not implemented"; return }

// As MacOS doesn't have abstract unix domain sockets the file
// created by a previous run needs to be explicitly removed

func (g *grpcServer) ExposesGRPCTLS() bool { _ = "STUB: not implemented"; return false }

func createLegacyGRPCServer(cfg *config.FulcioConfig, unixDomainSocket string, v2Server gw.CAServer) (*grpcServer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// recovers from per-transaction panics elegantly, so put it first

// Register your gRPC service implementations.

func panicRecoveryHandler(ctx context.Context, p any) error { _ = "STUB: not implemented"; return nil }
