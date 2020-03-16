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

package prometheus

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "prometheus", asset.ModuleFieldsPri, AssetPrometheus); err != nil {
		panic(err)
	}
}

// AssetPrometheus returns asset data.
// This is the base64 encoded gzipped contents of ../metricbeat/module/prometheus.
func AssetPrometheus() string {
	return "eJyckU1Ow0AMhfc5xdOwq9oeIAuuABJLhKppxklM50+2q6q3RyFtCYUNePmen/2NZ4MDnVtUKYlspKM2gLFFauGeb6JrgEDaCVfjkls8NgDwYt4U2omvFNBLSfD4SoFyqIWzbRtAxyK260rueWjR+6jUAEKRvFKLwU89ZMZ50BavTjW6NdxoVt1bA/RMMWj7uXeD7BPdUU9l5zrNknKsF2UZm+oBTxJIwApOtYj5bBhJaI3o9xQVJ44RyVs3omdRW8NGgpAavBBCOe4j3eZdUebwdnUzrjBl/06dLeRZ2M3ugc6nImFh/3Lmay0um8iEu8vWHzCz+3eau7d9c3fJ18p5uLS6lfsn9IJ2+fsfAQAA//+kwcDt"
}