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
	"net/http"
	"sync"

	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

type httpServer struct {
	*http.Server
	httpServerEndpoint string
}

func extractOIDCTokenFromAuthHeader(_ context.Context, req *http.Request) metadata.MD {
	_ = "STUB: not implemented"
	return *new(metadata.MD)
}

func createHTTPServer(ctx context.Context, serverEndpoint string, grpcServer, legacyGRPCServer *grpcServer) httpServer {
	_ = "STUB: not implemented"
	return *new(httpServer)
}

/* #nosec G402 */ // InsecureSkipVerify is only used for the HTTP server to call the TLS-enabled grpc endpoint.

// we are connecting over a unix domain socket, therefore we won't ever need TLS

// Limit request size

// enable CORS
// cors.Default() configures to accept requests for all domains

// Timeouts

func (h httpServer) startListener(wg *sync.WaitGroup) { _ = "STUB: not implemented"; return }

// received an interrupt signal, shut down

// error from closing listeners, or context timeout

func setResponseCodeModifier(ctx context.Context, w http.ResponseWriter, _ proto.Message) error {
	_ = "STUB: not implemented"
	return nil
}

// set SCT if present ahead of modifying response code

// strip all GRPC response headers

// set http status code

// delete the headers to not expose any grpc-metadata in http response
