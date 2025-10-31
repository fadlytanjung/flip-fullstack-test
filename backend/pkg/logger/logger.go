package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger wraps zap.Logger and provides convenience methods
type Logger struct {
	*zap.Logger
}

// Info logs an info message
func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.Logger.Info(msg, fields...)
}

// Debug logs a debug message
func (l *Logger) Debug(msg string, fields ...zap.Field) {
	l.Logger.Debug(msg, fields...)
}

// Warn logs a warning message
func (l *Logger) Warn(msg string, fields ...zap.Field) {
	l.Logger.Warn(msg, fields...)
}

// Error logs an error message
func (l *Logger) Error(msg string, fields ...zap.Field) {
	l.Logger.Error(msg, fields...)
}

// Fatal logs a fatal message and exits
func (l *Logger) Fatal(msg string, fields ...zap.Field) {
	l.Logger.Fatal(msg, fields...)
}

// With creates a child logger with additional fields
func (l *Logger) With(fields ...zap.Field) *Logger {
	return &Logger{Logger: l.Logger.With(fields...)}
}

// Sync flushes any buffered log entries
func (l *Logger) Sync() error {
	return l.Logger.Sync()
}

// String creates a string field
func String(key string, val string) zap.Field {
	return zap.String(key, val)
}

// Int creates an int field
func Int(key string, val int) zap.Field {
	return zap.Int(key, val)
}

// Int64 creates an int64 field
func Int64(key string, val int64) zap.Field {
	return zap.Int64(key, val)
}

// Error creates an error field
func Error(err error) zap.Field {
	return zap.Error(err)
}

// Any creates a field from any value
func Any(key string, val interface{}) zap.Field {
	return zap.Any(key, val)
}

// Duration creates a duration field
func Duration(key string, val interface{}) zap.Field {
	if d, ok := val.(interface{ Sub(interface{}) interface{} }); ok {
		_ = d // type assertion placeholder
	}
	return zap.Any(key, val)
}

type LoggerBuilderOption struct {
	UDPIP       string
	UDPPort     int
	PrettyPrint bool
}

// WithUdpSyncer adds UDP syncer to the logger sink so that logs can be sent to UDP server
//
// Parameters:
//   - ip: IP address of the UDP server
//   - port: port of the UDP server
//
// Returns:
//   - func(*LoggerBuilderOption): option function
func WithUdpSyncer(ip string, port int) func(*LoggerBuilderOption) {
	return func(config *LoggerBuilderOption) {
		config.UDPIP = ip
		config.UDPPort = port
	}
}

// WithPrettyPrint enables pretty print for the logger
// Use this for local development only
//
// Returns:
//   - func(*LoggerBuilderOption): option function
func WithPrettyPrint() func(*LoggerBuilderOption) {
	return func(config *LoggerBuilderOption) {
		config.PrettyPrint = true
	}
}

// NewLogger creates a new logger instance with the given service name, log level, and options
//
// Parameters:
//   - serviceName: name of the service (e.g., "flip-fullstack-test-backend")
//   - level: log level (e.g., "info", "debug", "warn", "error")
//   - options: optional configuration functions
//
// Returns:
//   - *Logger: configured logger instance
func NewLogger(serviceName string, level string, options ...func(*LoggerBuilderOption)) *Logger {
	// build config
	cfg := &LoggerBuilderOption{}
	for _, option := range options {
		option(cfg)
	}

	// create multiple sync target if UDP logging is enabled
	syncer := zapcore.AddSync(os.Stdout)
	if cfg.UDPIP != "" && cfg.UDPPort > 0 {
		syncer = zapcore.NewMultiWriteSyncer(os.Stdout, NewUDPSyncer(cfg.UDPIP, cfg.UDPPort))
	}

	// create encoder config
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "severity",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// create new core with appropriate encoder
	var core zapcore.Core
	if cfg.PrettyPrint {
		// Development: Colorful console output
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		core = zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			syncer,
			getLogLevel(level),
		)
	} else {
		// Production: JSON format for Cloud Run
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			syncer,
			getLogLevel(level),
		)
	}

	// create new log instance
	zapLogger := zap.New(core, zap.AddCaller())
	zapLogger = zapLogger.With(zap.String("service_name", serviceName))

	return &Logger{Logger: zapLogger}
}

// getLogLevel converts string log level to zapcore.Level
func getLogLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

// WithId creates a child logger with context identifiers
// This is useful for adding domain/handler context to logs
func WithId(l *Logger, domain string, handler string) *Logger {
	return l.With(
		String("domain", domain),
		String("handler", handler),
	)
}


