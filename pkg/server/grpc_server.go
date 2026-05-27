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

package server

import (
	"context"
	"crypto"
	"crypto/x509"

	ctclient "github.com/google/certificate-transparency-go/client"
	health "google.golang.org/grpc/health/grpc_health_v1"

	certauth "github.com/sigstore/fulcio/pkg/ca"
	fulciogrpc "github.com/sigstore/fulcio/pkg/generated/protobuf"
	"github.com/sigstore/fulcio/pkg/identity"
	"github.com/sigstore/sigstore/pkg/signature"
)

type GRPCCAServer interface {
	fulciogrpc.CAServer
	health.HealthServer
}

func NewGRPCCAServer(ct *ctclient.LogClient, ca certauth.CertificateAuthority, algorithmRegistry *signature.AlgorithmRegistryConfig, ip identity.IssuerPool) GRPCCAServer {
	_ = "STUB: not implemented"
	return *new(GRPCCAServer)
}

const (
	MetadataOIDCTokenKey = "oidcidentitytoken"
)

type grpcaCAServer struct {
	fulciogrpc.UnimplementedCAServer
	ct                *ctclient.LogClient
	ca                certauth.CertificateAuthority
	algorithmRegistry *signature.AlgorithmRegistryConfig
	identity.IssuerPool
}

func (g *grpcaCAServer) CreateSigningCertificate(ctx context.Context, request *fulciogrpc.CreateSigningCertificateRequest) (*fulciogrpc.SigningCertificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// OIDC token either is passed in gRPC field or was extracted from HTTP headers

// Authenticate OIDC ID token by checking signature

// Verify caller is in possession of their private key and extract
// public key from request.

// Option 1: Verify CSR

// Parse public key and check for weak key parameters

// Option 2: Check the signature for proof of possession of a private key

// Parse public key and check for weak parameters

// TODO: Ideally this comes from the verifier

// Check proof of possession signature

// Check whether the public-key/hash algorithm combination is allowed

// For CAs that do not support embedded SCTs or if the CT log is not configured

// currently configured CA doesn't support pre-certificate flow required to embed SCT in final certificate

// if the error was due to invalid input in the request, return HTTP 400

// otherwise return a 500 error to reflect that it is a transient server issue that the client can't resolve

// Submit to CTL

// convert to AddChainResponse because Cosign expects this struct.

// if the error was due to invalid input in the request, return HTTP 400

// otherwise return a 500 error to reflect that it is a transient server issue that the client can't resolve

// submit precertificate and chain to CT log

func (g *grpcaCAServer) GetTrustBundle(ctx context.Context, _ *fulciogrpc.GetTrustBundleRequest) (*fulciogrpc.TrustBundle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *grpcaCAServer) GetConfiguration(ctx context.Context, _ *fulciogrpc.GetConfigurationRequest) (*fulciogrpc.Configuration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *grpcaCAServer) Check(_ context.Context, _ *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *grpcaCAServer) List(_ context.Context, _ *health.HealthListRequest) (*health.HealthListResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *grpcaCAServer) Watch(_ *health.HealthCheckRequest, _ health.Health_WatchServer) error {
	_ = "STUB: not implemented"
	return nil
}

func getHashFuncForSignatureAlgorithm(signatureAlgorithm x509.SignatureAlgorithm) (crypto.Hash, error) {
	_ = "STUB: not implemented"
	return *new(crypto.Hash), nil
}
