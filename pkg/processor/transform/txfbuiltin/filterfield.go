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
	"bytes"

	"github.com/conduitio/conduit/pkg/foundation/cerrors"
	"github.com/conduitio/conduit/pkg/processor"
	"github.com/conduitio/conduit/pkg/processor/transform"
	"github.com/conduitio/conduit/pkg/record"

	"github.com/antchfx/jsonquery"
)

const (
	// Names for the Filter in global builder registry
	filterFieldKeyName     = "filterfieldkey"
	filterFieldPayloadName = "filterfieldpayload"

	// Config Fields for each Transform
	filterFieldConfigType          = "type"
	filterFieldConfigCondition     = "condition"
	filterFieldConfigMissingOrNull = "missingornull"
	filterFieldConfigExists        = "exists"
)

// `type` sets the behavior to "include" or "exclude" the record based on the
//  result of the condition.
// `condition` is an XPath query expression that the user defines to forward
// or drop a record on its results.
// `missingornull` defines how to handle the record in the event the fields
// the query would use don't exist.
// `exists` field in the config gives the user a chance to define an existence
// query for a given filter.
// * If `condition` passes, then it will immediately handle the record as
// `type` dictates.
// * If `condition` doesn't match, and `exists` matches nothing too, then it
// will handle the record as `missingornull` specifies.
//
// example transform config with noted possible values
// {
// 	"type": "include", // [include, exclude]
// 	"condition":"<xpath expression>",
// 	"exists": "<xpath expression>",
// 	"missingornull": "fail" // [fail, include, exclude]
// }

var (
	// ErrDropRecord is returned when a Record is not being forwarded.
	ErrDropRecord = cerrors.New("ErrDropRecord")
)

func init() {
	processor.GlobalBuilderRegistry.MustRegister(filterFieldKeyName, transform.NewBuilder(FilterFieldKey))
	processor.GlobalBuilderRegistry.MustRegister(filterFieldPayloadName, transform.NewBuilder(FilterFieldPayload))
}

func FilterFieldKey(config transform.Config) (transform.Transform, error) {
	return filterField(filterFieldKeyName, recordKeyGetSetter{}, config)
}

func FilterFieldPayload(config transform.Config) (transform.Transform, error) {
	return filterField(filterFieldPayloadName, recordPayloadGetSetter{}, config)
}

func filterField(
	transformName string,
	getSetter recordDataGetSetter,
	config transform.Config,
) (transform.Transform, error) {
	if config == nil {
		return nil, cerrors.New("must provide a transform config")
	}
	if len(config) == 0 {
		return nil, cerrors.New("must provide non-empty transform config")
	}
	var (
		filtertype      string
		filtercondition string
		filternull      string
		filterexists    string
	)

	// assign the values from our config
	filtertype = config[filterFieldConfigType]
	filtercondition = config[filterFieldConfigCondition]
	filternull = config[filterFieldConfigMissingOrNull]
	filterexists = config[filterFieldConfigExists]

	if filtertype == "" {
		return nil, cerrors.New("must specify include or exclude filter type")
	}
	if filtercondition == "" {
		return nil, cerrors.New("must specify filter condition")
	}
	// if filternull is not provided, filternull should fail loudly
	if filternull == "" {
		filternull = "fail"
	}

	return func(r record.Record) (record.Record, error) {
		data := getSetter.Get(r)
		switch d := data.(type) {
		case record.RawData:
			if d.Schema == nil {
				return record.Record{}, cerrors.Errorf("%s: schemaless raw data not supported", transformName)
			}
			return record.Record{}, cerrors.Errorf("%s: data with schema not supported yet", transformName) // TODO
		case record.StructuredData:
			doc, err := jsonquery.Parse(bytes.NewReader(d.Bytes()))
			if err != nil {
				return record.Record{}, cerrors.Errorf("filterfield failed to parse path: %w", err)
			}
			match := jsonquery.FindOne(doc, filtercondition)
			if match == nil {
				// check the filterexists query if one is set.
				if filterexists != "" {
					exists := jsonquery.Find(doc, filterexists)
					// if it matches, handle normal drop record behavior.
					if exists == nil {
						// if it doesn't match, defer to filternull behavior
						switch filternull {
						case "include":
							return r, nil
						case "exclude":
							return record.Record{}, ErrDropRecord
						case "fail":
							// fail should fail loudly with an existence error
							return record.Record{}, cerrors.Errorf("field does not exist: %s", filterexists)
						}
					}
				}
				return record.Record{}, ErrDropRecord
			}

			// handle matches based on filtertype as normal
			switch filtertype {
			case "include":
				return r, nil
			case "exclude":
				return record.Record{}, ErrDropRecord
			default:
				return record.Record{}, cerrors.Errorf("invalid filtertype: %s", filtertype)
			}

		default:
			return record.Record{}, cerrors.Errorf("%s: unexpected data type %T", transformName, data)
		}
	}, nil
}
