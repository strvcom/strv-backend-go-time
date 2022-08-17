# STRV time

![Latest release][release]
[![codecov][codecov-img]][codecov]
![GitHub][license]

Custom Go time & duration implementation.

## Duration
Duration is a wrapper around [time](https://pkg.go.dev/time) that simplifies `serialization`/`deserialization` of time.

## Examples
```go
package main

import (
	timex "go.strv.io/time"
)

type config struct {
	AccessTokenExpiration  timex.Duration `json:"access_token_expiration"`
	RefreshTokenExpiration timex.Duration `json:"refresh_token_expiration"`
}

func main() {
	cfg := config{}
	data, _ := os.ReadFile("config.json")
	_ = json.Unmarshal(data, &cfg)
}
```
```json
// Content of config.json
{
  "access_token_expiration": "1h",
  "refresh_token_expiration": "30d"
}
```
As can be seen, there is an option to use days directly, so there is no need to use `720h` in case of `refresh_token_expiration` in the example. 

[release]: https://img.shields.io/github/v/release/strvcom/strv-backend-go-time
[codecov]: https://codecov.io/gh/strvcom/strv-backend-go-time
[codecov-img]: https://codecov.io/gh/strvcom/strv-backend-go-time/branch/ci/github-actions/graph/badge.svg?token=FM3Y66NLKX
[license]: https://img.shields.io/github/license/strvcom/strv-backend-go-time
