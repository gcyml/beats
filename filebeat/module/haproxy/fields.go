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

package haproxy

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "haproxy", asset.ModuleFieldsPri, Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "eJzMWdtu2zwSvs9TDHKzLZD6AQJsgUW2uy3w/20ucm/Q5FgalCJVcpTUb/+DB8k62lbiovWdRc7MxzkP+QG+4+EeSlE7+/NwA8DEGu/hNn+5vQFQ6KWjmsmae/h4AwDtfvjbqkbjDcCeUCt/Hxc/gBEV9pmGHx9qvIfC2abOX2b4Hhnlv0due2cNo1Hb8LdbHXH5KioEuwcusSOAd9aBJs9o0L2Hl5JkCQ4l0jMqEEZB7axE71FFOmmNQRn4baYodkJ+XwMi75/F8CI8eNQoOUi2UAkjChxhCAvhi0f3jG4GUVq4GJAWnjNNYJ2QjEQmYIZnpLFlobcvgphMsWWqcFv5JblPYTOETUAGKtKaPEprlAdfo2HIfMJqgPAsHNnGw48GG+xzTb6jrSmmiI64I6yrYdpbN2MK9Cx2mnzZ2mVPRuis0AsR7w6MfutQqNMgTVPt0AXDRQpgJ4yviLO7RHCaIuYSkwa1LYA8YNq1uRBQVFlU+q+x5CocbYRlvf8OO94BGakbFcgdsqPLj5Cj8QJlHq3r8EeDnn2bFtBhLyXtcG9dyArkwRpsNZxDOAq6FFyr2l+PLkv6l4dC253QK3HSigybsmowlVDKoffjDH82p6Nz1m0r9F4UiyI/hU2QN4U4K8LpD/D5P4+xEpIBKXxEFflNz8r4k2ccxjZO4vzmJQuV2IlN5J0u7FxUoavIiJghPQtePOKDNYpSgEQH874tBmSOKab9jkaFDDMRV1m17jiBALgUPKzawaFqdCIG8bunh0ewDj4/PT2+P1UDFvP+gzUsyPguNUnbGPZBcz1qEJLpuXPj7OdTW/bbGIBh+zMGl1gOls6EnLRGNs6FLNbHZgegphYJlkqOuRlJmws2mOmtrgeyM+MVUeakcgWQpTBKp/jtN2lXxDppCF4L1TNp3bqlHeT+K8LNVe4U3q99pG0lzXSAP2t0hEa2SiV/RBVhukOIY7Yt9bivPY+3C/bY8yzF+Rezt66K2Q7Ezjbc75OUjSBKbKvaKwK7hUH1zRSx0CSGWqwFl22W3+QKNdhQUeFEws5uUJZbSbV1vF7WhGosaKLYnFjGlXdeWpLVJsgRyXlRpFZKGFKcFaDQcy56S57y3+MWoKPXvMElVhqqh/G8tabSVjlgX9aI8KwuC7QDknm1zE7z6dfV3v+j/fLYVzYUgkt0IWuI0E7mBJfCdUN1UvYwO3wz+gC1wzChAqXGJzH+FIZbkh6FkyXUuinIhD5CPAvSYqcxTPsDXo0f5spTdVxaw2TQ8DhATml/FJQF2s1JPstm7wNpDLvDlrzdymG3tRrKSU6XgNFWjsNmJYgFDpcId1iErvZt9lhmcpExiA9vdYgFFis08HZXOM1oCUp31ce8mAseNYbBSCjV/746zc5kQIe+tsYPwQ5k/y8yAYda9G5NwhQxTzwP5UxeEDU3DtVWWvudTjb6H0eLAN/iitBwG5j9+1noBm8BQ1QCGUUyzT/deJQbmVKo1GQlme280p5pM6e1GcQlCoXuZNs3hfwXeQ5tYCbuuI1BgGqw1XhK18dB9TYTHTcnbrcQJ9QqpvbhJDHfIg4Gy/lOnAU3/jXREZx606l0mc9ScMx67LjtXOOwU9o/3l/zqFKKcCZunEEFovXaF+ISiH17tN/sthHDhV6b9v4yp3XiZZuFbDWZk0Z6sFWtkYdeAoGqf48ZUFXIpVV33R5hVCJ6RhdHNc+OTPFa0PH6Nk2qwwv4CeQnqnCgdIeVoOgb3f1Lfj26y9d55CMBL1/7ohZ1upLkF8yT8Z6c53h93tqwlfgSPbL3CNQ9jrS7u9R0fA6Z00scUs/rJb5NBIe3DW+VYHFORWvvt49vOh5NCLJ9o/Ww1N2BsZzax0AZULz9TFc8xhBzstPe2apf+d4NjrCz6vAexJ7Rje3dN/CaU3YnlIsdzdPDY3xpSWPMWwbGNNzOP6XNwlyj4LmAEFJizUOHl9r6uQFz9LL2R6D8JwAA///iRBtm"
}