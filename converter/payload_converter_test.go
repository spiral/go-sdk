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

package converter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	commonpb "github.com/spiral/go-api/common/v1"
)

type testStruct struct {
	Name string
	Age  int
}

func TestProtoJsonPayloadConverter(t *testing.T) {
	pc := NewProtoJSONPayloadConverter()

	wt := &commonpb.WorkflowType{Name: "qwe"}
	payload, err := pc.ToPayload(wt)
	require.NoError(t, err)
	wt2 := &commonpb.WorkflowType{}
	err = pc.FromPayload(payload, &wt2)
	require.NoError(t, err)
	assert.Equal(t, "qwe", wt2.Name)

	var wt3 *commonpb.WorkflowType
	err = pc.FromPayload(payload, &wt3)
	require.NoError(t, err)
	assert.Equal(t, "qwe", wt3.Name)

	s := pc.ToString(payload)
	assert.Equal(t, `{"name":"qwe"}`, s)
}

func TestProtoPayloadConverter(t *testing.T) {
	pc := NewProtoPayloadConverter()

	wt := &commonpb.WorkflowType{Name: "qwe"}
	payload, err := pc.ToPayload(wt)
	require.NoError(t, err)
	wt2 := &commonpb.WorkflowType{}
	err = pc.FromPayload(payload, &wt2)
	require.NoError(t, err)
	assert.Equal(t, "qwe", wt2.Name)

	var wt3 *commonpb.WorkflowType
	err = pc.FromPayload(payload, &wt3)
	require.NoError(t, err)
	assert.Equal(t, "qwe", wt3.Name)

	s := pc.ToString(payload)
	assert.Equal(t, "CgNxd2U", s)
}

func TestJsonPayloadConverter(t *testing.T) {
	pc := NewJSONPayloadConverter()

	wt := testStruct{Name: "qwe"}
	payload, err := pc.ToPayload(wt)
	require.NoError(t, err)
	wt2 := testStruct{}
	err = pc.FromPayload(payload, &wt2)
	require.NoError(t, err)
	assert.Equal(t, "qwe", wt2.Name)

	var wt3 *testStruct
	err = pc.FromPayload(payload, &wt3)
	require.NoError(t, err)
	assert.Equal(t, "qwe", wt3.Name)

	s := pc.ToString(payload)
	assert.Equal(t, "{Age:0 Name:qwe}", s)
}
