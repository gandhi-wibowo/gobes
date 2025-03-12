package logger

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/rotisserie/eris"
	"github.com/sirupsen/logrus"
)

var (
	once     sync.Once
	instance *Writer
)

type Writer struct {
	code         string
	domain       string
	hint         string
	instance     *logrus.Entry
	message      string
	owner        any
	stackEnabled bool
	stacktrace   map[string]any
	tags         []string
	user         any
	with         map[string]any
}

func NewWriter() WriterProvider {
	once.Do(func() {
		logrusInstance := logrus.New()
		logrusInstance.SetLevel(logrus.DebugLevel)
		logrusInstance.SetOutput(os.Stdout)
		ctx := logrusInstance.WithContext(context.Background())
		instance = &Writer{
			code:         "",
			domain:       "",
			hint:         "",
			instance:     ctx,
			message:      "",
			owner:        nil,
			stackEnabled: false,
			stacktrace:   nil,
			tags:         []string{},
			user:         nil,
			with:         map[string]any{},
		}
	})
	return instance
}

func (r *Writer) Debug(args ...any) {
	r.instance.WithField("root", r.toMap()).Debug(args...)
}

func (r *Writer) Debugf(format string, args ...any) {
	r.instance.WithField("root", r.toMap()).Debugf(format, args...)
}

func (r *Writer) Info(args ...any) {
	r.instance.WithField("root", r.toMap()).Info(args...)
}

func (r *Writer) Infof(format string, args ...any) {
	r.instance.WithField("root", r.toMap()).Infof(format, args...)
}

func (r *Writer) Warning(args ...any) {
	r.instance.WithField("root", r.toMap()).Warning(args...)
}

func (r *Writer) Warningf(format string, args ...any) {
	r.instance.WithField("root", r.toMap()).Warningf(format, args...)
}

func (r *Writer) Error(args ...any) {
	r.withStackTrace(fmt.Sprint(args...))
	r.instance.WithField("root", r.toMap()).Error(args...)
}

func (r *Writer) Errorf(format string, args ...any) {
	r.withStackTrace(fmt.Sprintf(format, args...))
	r.instance.WithField("root", r.toMap()).Errorf(format, args...)
}

func (r *Writer) Fatal(args ...any) {
	r.withStackTrace(fmt.Sprint(args...))
	r.instance.WithField("root", r.toMap()).Fatal(args...)
}

func (r *Writer) Fatalf(format string, args ...any) {
	r.withStackTrace(fmt.Sprintf(format, args...))
	r.instance.WithField("root", r.toMap()).Fatalf(format, args...)
}

func (r *Writer) Panic(args ...any) {
	r.withStackTrace(fmt.Sprint(args...))
	r.instance.WithField("root", r.toMap()).Panic(args...)
}

func (r *Writer) Panicf(format string, args ...any) {
	r.withStackTrace(fmt.Sprintf(format, args...))
	r.instance.WithField("root", r.toMap()).Panicf(format, args...)
}

// Code set a code or slug that describes the error.
// Error messages are intended to be read by humans, but such code is expected to
// be read by machines and even transported over different services.
func (r *Writer) Code(code string) WriterProvider {
	r.code = code
	return r
}

// Hint set a hint for faster debugging.
func (r *Writer) Hint(hint string) WriterProvider {
	r.hint = hint

	return r
}

// In sets the feature category or domain in which the log entry is relevant.
func (r *Writer) In(domain string) WriterProvider {
	r.domain = domain

	return r
}

// Owner set the name/email of the colleague/team responsible for handling this error.
// Useful for alerting purpose.
func (r *Writer) Owner(owner any) WriterProvider {
	r.owner = owner

	return r
}

// Tags add multiple tags, describing the feature returning an error.
func (r *Writer) Tags(tags ...string) WriterProvider {
	r.tags = append(r.tags, tags...)

	return r
}

// User sets the user associated with the log entry.
func (r *Writer) User(user any) WriterProvider {
	r.user = user
	return r
}

// With adds key-value pairs to the context of the log entry
func (r *Writer) With(data map[string]any) WriterProvider {
	for k, v := range data {
		r.with[k] = v
	}

	return r
}

// WithTrace adds a stack trace to the log entry.
func (r *Writer) WithTrace() WriterProvider {
	r.withStackTrace("")
	return r
}

func (r *Writer) withStackTrace(message string) {
	erisNew := eris.New(message)
	if erisNew == nil {
		return
	}

	r.message = erisNew.Error()
	format := eris.NewDefaultJSONFormat(eris.FormatOptions{
		InvertOutput: true,
		WithTrace:    true,
		InvertTrace:  true,
	})
	r.stacktrace = eris.ToCustomJSON(erisNew, format)
	r.stackEnabled = true
}

// resetAll resets all properties of the WriterProvider for a new log entry.
func (r *Writer) resetAll() {
	r.code = ""
	r.domain = ""
	r.hint = ""
	r.message = ""
	r.owner = nil
	r.stacktrace = nil
	r.stackEnabled = false
	r.tags = []string{}
	r.user = nil
	r.with = map[string]any{}
}

// toMap returns a map representation of the error.
func (r *Writer) toMap() map[string]any {
	payload := map[string]any{}

	if code := r.code; code != "" {
		payload["code"] = code
	}
	if ctx := r.instance.Context; ctx != nil {
		values := make(map[any]any)
		getContextValues(ctx, values)
		if len(values) > 0 {
			payload["context"] = values
		}
	}
	if domain := r.domain; domain != "" {
		payload["domain"] = domain
	}
	if hint := r.hint; hint != "" {
		payload["hint"] = hint
	}
	if message := r.message; message != "" {
		payload["message"] = message
	}
	if owner := r.owner; owner != nil {
		payload["owner"] = owner
	}
	if stacktrace := r.stacktrace; stacktrace != nil || r.stackEnabled {
		payload["stacktrace"] = stacktrace
	}
	if tags := r.tags; len(tags) > 0 {
		payload["tags"] = tags
	}
	if r.user != nil {
		payload["user"] = r.user
	}
	if with := r.with; len(with) > 0 {
		payload["with"] = with
	}

	// reset all properties for a new log entry
	r.resetAll()

	return payload
}
