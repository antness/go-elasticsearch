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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/noderole"
)

// UpdateByQueryRethrottleNode type.
//
// https://github.com/elastic/elasticsearch-specification/blob/e279583a47508af40eb07b84694c5aae7885aa09/specification/_global/update_by_query_rethrottle/UpdateByQueryRethrottleNode.ts#L25-L27
type UpdateByQueryRethrottleNode struct {
	Attributes       map[string]string   `json:"attributes"`
	Host             string              `json:"host"`
	Ip               string              `json:"ip"`
	Name             string              `json:"name"`
	Roles            []noderole.NodeRole `json:"roles,omitempty"`
	Tasks            map[string]TaskInfo `json:"tasks"`
	TransportAddress string              `json:"transport_address"`
}

func (s *UpdateByQueryRethrottleNode) UnmarshalJSON(data []byte) error {

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

		case "attributes":
			if s.Attributes == nil {
				s.Attributes = make(map[string]string, 0)
			}
			if err := dec.Decode(&s.Attributes); err != nil {
				return err
			}

		case "host":
			if err := dec.Decode(&s.Host); err != nil {
				return err
			}

		case "ip":
			if err := dec.Decode(&s.Ip); err != nil {
				return err
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return err
			}

		case "roles":
			if err := dec.Decode(&s.Roles); err != nil {
				return err
			}

		case "tasks":
			if s.Tasks == nil {
				s.Tasks = make(map[string]TaskInfo, 0)
			}
			if err := dec.Decode(&s.Tasks); err != nil {
				return err
			}

		case "transport_address":
			if err := dec.Decode(&s.TransportAddress); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewUpdateByQueryRethrottleNode returns a UpdateByQueryRethrottleNode.
func NewUpdateByQueryRethrottleNode() *UpdateByQueryRethrottleNode {
	r := &UpdateByQueryRethrottleNode{
		Attributes: make(map[string]string, 0),
		Tasks:      make(map[string]TaskInfo, 0),
	}

	return r
}
