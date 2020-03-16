// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package kubernetes

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "kubernetes", asset.ModuleFieldsPri, AssetKubernetes); err != nil {
		panic(err)
	}
}

// AssetKubernetes returns asset data.
// This is the base64 encoded gzipped contents of ../metricbeat/module/kubernetes.
func AssetKubernetes() string {
	return "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"
}