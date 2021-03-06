// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package log

import (
	"github.com/spiral/go-sdk/log"
)

// NopLogger is Logger implementation that doesn't produce any logs.
type NopLogger struct {
}

// NewNopLogger creates new instance of NopLogger.
func NewNopLogger() *NopLogger {
	return &NopLogger{}
}

// Debug does nothing.
func (l *NopLogger) Debug(msg string, keyvals ...interface{}) {
}

// Info does nothing.
func (l *NopLogger) Info(msg string, keyvals ...interface{}) {
}

// Warn does nothing.
func (l *NopLogger) Warn(msg string, keyvals ...interface{}) {
}

// Error does nothing.
func (l *NopLogger) Error(msg string, keyvals ...interface{}) {
}

// With returns new NopLogger.
func (l *NopLogger) With(keyvals ...interface{}) log.Logger {
	return NewNopLogger()
}
