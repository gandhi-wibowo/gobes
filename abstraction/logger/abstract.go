package logger

import "context"

type Log interface {
	// WithContext adds a context to the logger.
	WithContext(ctx context.Context) WriterProvider
	// Channel return a writer for a specific channel.
	Channel(channel string) WriterProvider
	// Stack return a writer for multiple channels.
	Stack(channels []string) WriterProvider
	WriterProvider
}

type WriterProvider interface {
	// Debug logs a message at DebugLevel.
	Debug(args ...any)
	// Debugf is equivalent to Debug, but with support for fmt.Printf-style arguments.
	Debugf(format string, args ...any)
	// Info logs a message at InfoLevel.
	Info(args ...any)
	// Infof is equivalent to Info, but with support for fmt.Printf-style arguments.
	Infof(format string, args ...any)
	// Warning logs a message at WarningLevel.
	Warning(args ...any)
	// Warningf is equivalent to Warning, but with support for fmt.Printf-style arguments.
	Warningf(format string, args ...any)
	// Error logs a message at ErrorLevel.
	Error(args ...any)
	// Errorf is equivalent to Error, but with support for fmt.Printf-style arguments.
	Errorf(format string, args ...any)
	// Fatal logs a message at FatalLevel.
	Fatal(args ...any)
	// Fatalf is equivalent to Fatal, but with support for fmt.Printf-style arguments.
	Fatalf(format string, args ...any)
	// Panic logs a message at PanicLevel.
	Panic(args ...any)
	// Panicf is equivalent to Panic, but with support for fmt.Printf-style arguments.
	Panicf(format string, args ...any)
	// Code set a code or slug that describes the error.
	// Error messages are intended to be read by humans, but such code is expected to
	// be read by machines and even transported over different services.
	Code(code string) WriterProvider
	// Hint set a hint for faster debugging.
	Hint(hint string) WriterProvider
	// In sets the feature category or domain in which the log entry is relevant.
	In(domain string) WriterProvider
	// Owner set the name/email of the colleague/team responsible for handling this error.
	// Useful for alerting purpose.
	Owner(owner any) WriterProvider
	// Tags add multiple tags, describing the feature returning an error.
	Tags(tags ...string) WriterProvider
	// User sets the user associated with the log entry.
	User(user any) WriterProvider
	// With adds key-value pairs to the context of the log entry
	With(data map[string]any) WriterProvider
	// WithTrace adds a stack trace to the log entry.
	WithTrace() WriterProvider
}
