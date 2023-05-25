package log

import (
	"encoding/json"
	"go.uber.org/zap/zapcore"
)

const (
	flagLevel             = "log.level"
	flagDisableCaller     = "log.disable-caller"
	flagDisableStacktrace = "log.disable-stacktrace"
	flagFormat            = "log.format"
	flagEnableColor       = "log.enable-color"
	flagOutputPaths       = "log.output-paths"
	flagDevelopment       = "log.development"
	flagName              = "log.name"

	consoleFormat = "console"
	jsonFormat    = "json"
)

type Options struct {
	OutputPaths       []string `json:"output_paths"`
	ErrorOutputPaths  []string `json:"error_output_paths"`
	Level             string   `json:"level"`
	Format            string   `json:"format"`
	DisableCaller     bool     `json:"disable_caller"`
	DisableStacktrace bool     `json:"disable_stacktrace"`
	EnableColor       bool     `json:"enable_color"`
	Development       bool     `json:"development"`
	Name              string   `json:"name"`
}

// NewOptions create with default params
func NewOptions() *Options {
	return &Options{
		Level:             zapcore.InfoLevel.String(),
		DisableCaller:     false,
		DisableStacktrace: false,
		Format:            consoleFormat,
		EnableColor:       false,
		Development:       false,
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
	}
}

func (o *Options) String() string {
	data, _ := json.Marshal(o)
	return string(data)
}
