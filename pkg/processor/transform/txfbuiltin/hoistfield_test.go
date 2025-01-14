// Copyright © 2022 Meroxa, Inc.
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

package txfbuiltin

import (
	"reflect"
	"testing"

	"github.com/conduitio/conduit/pkg/foundation/assert"
	"github.com/conduitio/conduit/pkg/processor/transform"
	"github.com/conduitio/conduit/pkg/record"
	"github.com/conduitio/conduit/pkg/record/schema/mock"
)

func TestHoistFieldKey_Build(t *testing.T) {
	type args struct {
		config transform.Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{{
		name:    "nil config returns error",
		args:    args{config: nil},
		wantErr: true,
	}, {
		name:    "empty config returns error",
		args:    args{config: map[string]string{}},
		wantErr: true,
	}, {
		name:    "empty field returns error",
		args:    args{config: map[string]string{hoistFieldConfigField: ""}},
		wantErr: true,
	}, {
		name:    "non-empty field returns transform",
		args:    args{config: map[string]string{hoistFieldConfigField: "foo"}},
		wantErr: false,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := HoistFieldKey(tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("HoistFieldKey() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestHoistFieldKey_Transform(t *testing.T) {
	type args struct {
		r record.Record
	}
	tests := []struct {
		name    string
		config  transform.Config
		args    args
		want    record.Record
		wantErr bool
	}{{
		name:   "structured data",
		config: map[string]string{hoistFieldConfigField: "foo"},
		args: args{r: record.Record{
			Key: record.StructuredData{
				"bar": 123,
				"baz": nil,
			},
		}},
		want: record.Record{
			Key: record.StructuredData{
				"foo": map[string]interface{}{
					"bar": 123,
					"baz": nil,
				},
			},
		},
		wantErr: false,
	}, {
		name:   "raw data without schema",
		config: map[string]string{hoistFieldConfigField: "foo"},
		args: args{r: record.Record{
			Key: record.RawData{
				Raw:    []byte("raw data"),
				Schema: nil,
			},
		}},
		want: record.Record{
			Key: record.StructuredData{
				"foo": []byte("raw data"),
			},
		},
		wantErr: false,
	}, {
		name:   "raw data with schema",
		config: map[string]string{hoistFieldConfigField: "foo"},
		args: args{r: record.Record{
			Key: record.RawData{
				Raw:    []byte("raw data"),
				Schema: mock.NewSchema(nil),
			},
		}},
		want:    record.Record{},
		wantErr: true, // TODO not implemented
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			txfFunc, err := HoistFieldKey(tt.config)
			assert.Ok(t, err)
			got, err := txfFunc(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Transform() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transform() got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestHoistFieldPayload_Build(t *testing.T) {
	type args struct {
		config transform.Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{{
		name:    "nil config returns error",
		args:    args{config: nil},
		wantErr: true,
	}, {
		name:    "empty config returns error",
		args:    args{config: map[string]string{}},
		wantErr: true,
	}, {
		name:    "empty field returns error",
		args:    args{config: map[string]string{hoistFieldConfigField: ""}},
		wantErr: true,
	}, {
		name:    "non-empty field returns transform",
		args:    args{config: map[string]string{hoistFieldConfigField: "foo"}},
		wantErr: false,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := HoistFieldPayload(tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("HoistFieldPayload() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestHoistFieldPayload_Transform(t *testing.T) {
	type args struct {
		r record.Record
	}
	tests := []struct {
		name    string
		config  transform.Config
		args    args
		want    record.Record
		wantErr bool
	}{{
		name:   "structured data",
		config: map[string]string{hoistFieldConfigField: "foo"},
		args: args{r: record.Record{
			Payload: record.StructuredData{
				"bar": 123,
				"baz": nil,
			},
		}},
		want: record.Record{
			Payload: record.StructuredData{
				"foo": map[string]interface{}{
					"bar": 123,
					"baz": nil,
				},
			},
		},
		wantErr: false,
	}, {
		name:   "raw data without schema",
		config: map[string]string{hoistFieldConfigField: "foo"},
		args: args{r: record.Record{
			Payload: record.RawData{
				Raw:    []byte("raw data"),
				Schema: nil,
			},
		}},
		want: record.Record{
			Payload: record.StructuredData{
				"foo": []byte("raw data"),
			},
		},
		wantErr: false,
	}, {
		name:   "raw data with schema",
		config: map[string]string{hoistFieldConfigField: "foo"},
		args: args{r: record.Record{
			Payload: record.RawData{
				Raw:    []byte("raw data"),
				Schema: mock.NewSchema(nil),
			},
		}},
		want:    record.Record{},
		wantErr: true, // TODO not implemented
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			txfFunc, err := HoistFieldPayload(tt.config)
			assert.Ok(t, err)
			got, err := txfFunc(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Transform() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transform() got = %v, want = %v", got, tt.want)
			}
		})
	}
}
