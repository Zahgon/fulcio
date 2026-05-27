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

package log

import (
	"context"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"goa.design/goa/v3/grpc/middleware"
)

// Logger set the default logger to development mode
var Logger *zap.SugaredLogger

func init() {
	ConfigureLogger("dev")
}

func ConfigureLogger(logType string) { _ = "STUB: not implemented"; return }

var CliLogger = createCliLogger()

func createCliLogger() *zap.SugaredLogger { _ = "STUB: not implemented"; return nil }

type requestIDMetadataKeyType string

const (
	requestIDMetadataKey requestIDMetadataKeyType = middleware.RequestIDMetadataKey
)

func ContextLogger(ctx context.Context) *zap.SugaredLogger { _ = "STUB: not implemented"; return nil }

func SetupGRPCLogging() (*zap.Logger, []grpc_zap.Option) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: implement filters to eliminate health check log statements
