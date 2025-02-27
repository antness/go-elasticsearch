// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/e279583a47508af40eb07b84694c5aae7885aa09

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// TrainedModelInferenceFeatureImportance type.
//
// https://github.com/elastic/elasticsearch-specification/blob/e279583a47508af40eb07b84694c5aae7885aa09/specification/ml/_types/inference.ts#L451-L455
type TrainedModelInferenceFeatureImportance struct {
	Classes     []TrainedModelInferenceClassImportance `json:"classes,omitempty"`
	FeatureName string                                 `json:"feature_name"`
	Importance  *Float64                               `json:"importance,omitempty"`
}

func (s *TrainedModelInferenceFeatureImportance) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "classes":
			if err := dec.Decode(&s.Classes); err != nil {
				return err
			}

		case "feature_name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.FeatureName = o

		case "importance":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Importance = &f
			case float64:
				f := Float64(v)
				s.Importance = &f
			}

		}
	}
	return nil
}

// NewTrainedModelInferenceFeatureImportance returns a TrainedModelInferenceFeatureImportance.
func NewTrainedModelInferenceFeatureImportance() *TrainedModelInferenceFeatureImportance {
	r := &TrainedModelInferenceFeatureImportance{}

	return r
}
