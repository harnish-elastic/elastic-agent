// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package decode_cef

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "decode_cef", asset.ModuleFieldsPri, AssetDecodeCef); err != nil {
		panic(err)
	}
}

// AssetDecodeCef returns asset data.
// This is the base64 encoded gzipped contents of processors/decode_cef.
func AssetDecodeCef() string {
	return "eJy0lE9v00AQxe/5FHMEKY6KQPyJFA6ERFQCCam013S7+xwPsXfN7jipvz1aO3ZcEqSogT1Eynj85jdvPZPQBvWUNNIRkbDkmNJnaGdA88WSSu80QnCeUkZuwojIIGjPpbCzU/o4IiKau6JwlhZbWKGl84USejFfLF+SUaImI9q/PW2yE7KqQFczHqlLTGntXVXuIyeKxPOpJoNUVbmQZKB705CuNNL7AerOsyCQyvOmPqXeFU3+fLHspQqEoNYgcSQZB7pvRNzDT2iZ0LWQdlYU29C9SRmUQWcEKWvik14PjwIb2Nm+53iGfQ9738LH3D7eebBBvXPeDOJ/cSKeu1aEXNozhhKaU9Yq5lMVYOihbp7u+52MjlgMtqwx2cIa5y8lihodUCtMkimJt2MqDXMeS5stl8F8b0Uup/kfl7Xv8Dk4iGO20rkKYcXmMqpby78qEBtY4ZTR311TpFE8ARKwhWepz6iNR1WUcancwdfJF15n54FdF6XzoqzGE6IJ/chAW5WzoSCe7Tr+qeK0e9Ct3Vi3s2P66nbjJ3LfYLgqxhQBxs3s9jxDSbaCNfxQ8yp5PTuSe5O8nXWS7xJ6PzvofkheXc0O4sfmxd/LLu0mc16GCU8tOi7Z76ZwVLhdeINwG1g9A2vu8hy649mgThobqVTsA2nlPSNa3C+rw8ps1uQevP/IXOU1JgG++eoHtp1iG3L1VDftu+2QcbvJW9lTjh3GLAjbZoP+8+qi/BryR/XfAQAA//+1OzaZ"
}