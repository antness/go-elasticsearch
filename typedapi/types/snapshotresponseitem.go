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

// SnapshotResponseItem type.
//
// https://github.com/elastic/elasticsearch-specification/blob/e279583a47508af40eb07b84694c5aae7885aa09/specification/snapshot/get/SnapshotGetResponse.ts#L44-L48
type SnapshotResponseItem struct {
	Error      *ErrorCause    `json:"error,omitempty"`
	Repository string         `json:"repository"`
	Snapshots  []SnapshotInfo `json:"snapshots,omitempty"`
}

func (s *SnapshotResponseItem) UnmarshalJSON(data []byte) error {

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

		case "error":
			if err := dec.Decode(&s.Error); err != nil {
				return err
			}

		case "repository":
			if err := dec.Decode(&s.Repository); err != nil {
				return err
			}

		case "snapshots":
			if err := dec.Decode(&s.Snapshots); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewSnapshotResponseItem returns a SnapshotResponseItem.
func NewSnapshotResponseItem() *SnapshotResponseItem {
	r := &SnapshotResponseItem{}

	return r
}
