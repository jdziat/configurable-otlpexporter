// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configurableotlpexporter // import "github.com/jdziat/configurableotlpexporter"

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/config/configretry"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

// Config defines configuration for OTLP exporter.
type Config struct {
	exporterhelper.TimeoutSettings `mapstructure:",squash"`     // squash ensures fields are correctly decoded in embedded struct.
	QueueConfig                    exporterhelper.QueueSettings `mapstructure:"sending_queue"`
	RetryConfig                    configretry.BackOffConfig    `mapstructure:"retry_on_failure"`
	MaxSendSize                    int                          `mapstructure:"max_send_size"`
	MaxReceiveSize                 int                          `mapstructure:"max_receive_size"`
	configgrpc.GRPCClientSettings  `mapstructure:",squash"`     // squash ensures fields are correctly decoded in embedded struct.
}

var _ component.Config = (*Config)(nil)
