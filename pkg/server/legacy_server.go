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

	fulciogrpc "github.com/sigstore/fulcio/pkg/generated/protobuf"
	"github.com/sigstore/fulcio/pkg/generated/protobuf/legacy"
	"google.golang.org/genproto/googleapis/api/httpbody"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

const (
	PEMCertificateChain         = "application/pem-certificate-chain"
	SCTMetadataKey              = "x-sct"
	HTTPResponseCodeMetadataKey = "x-http-code"
)

type legacyGRPCCAServer struct {
	legacy.UnimplementedCAServer
	v2Server fulciogrpc.CAServer
}

func NewLegacyGRPCCAServer(v2Server fulciogrpc.CAServer) legacy.CAServer {
	_ = "STUB: not implemented"
	return *new(legacy.CAServer)
}

func (l *legacyGRPCCAServer) CreateSigningCertificate(ctx context.Context, request *legacy.CreateSigningCertificateRequest) (*httpbody.HttpBody, error) {
	_ = "STUB: not implemented"
	// OIDC token either is passed in gRPC field or was extracted from HTTP headers
	return nil, nil
}

//lint:ignore SA1019 this is valid because we're converting from v1beta to v1 API

// the CSR and the public key have not been set

// create new CA request mapping fields from legacy to actual
//lint:ignore SA1019 this is valid because we're converting from v1beta to v1 API

//lint:ignore SA1019 this is valid because we're converting from v1beta to v1 API

//lint:ignore SA1019 this is valid because we're converting from v1beta to v1 API,

// we need to return a HTTP 201 Created response code to be backward compliant

// the SCT for the certificate needs to be returned in a HTTP response header

func (l *legacyGRPCCAServer) GetRootCertificate(ctx context.Context, _ *emptypb.Empty) (*httpbody.HttpBody, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
