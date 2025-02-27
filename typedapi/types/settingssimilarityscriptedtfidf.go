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
)

// SettingsSimilarityScriptedTfidf type.
//
// https://github.com/elastic/elasticsearch-specification/blob/e279583a47508af40eb07b84694c5aae7885aa09/specification/indices/_types/IndexSettings.ts#L216-L219
type SettingsSimilarityScriptedTfidf struct {
	Script Script `json:"script"`
	Type   string `json:"type,omitempty"`
}

func (s *SettingsSimilarityScriptedTfidf) UnmarshalJSON(data []byte) error {

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

		case "script":
			if err := dec.Decode(&s.Script); err != nil {
				return err
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return err
			}

		}
	}
	return nil
}

// MarshalJSON override marshalling to include literal value
func (s SettingsSimilarityScriptedTfidf) MarshalJSON() ([]byte, error) {
	type innerSettingsSimilarityScriptedTfidf SettingsSimilarityScriptedTfidf
	tmp := innerSettingsSimilarityScriptedTfidf{
		Script: s.Script,
		Type:   s.Type,
	}

	tmp.Type = "scripted"

	return json.Marshal(tmp)
}

// NewSettingsSimilarityScriptedTfidf returns a SettingsSimilarityScriptedTfidf.
func NewSettingsSimilarityScriptedTfidf() *SettingsSimilarityScriptedTfidf {
	r := &SettingsSimilarityScriptedTfidf{}

	return r
}
