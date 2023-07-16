# unkey-go

`unkey-go` is a Go client for [unkey](https://github.com/unkeyed/unkey/) - a simple, secure, and open-source API key management service.

Get started with [unkey](https://unkey.dev) & check the official [docs](https://docs.unkey.dev) for more information.

## Installation

To use the `unkey-go` SDK in your Go project, you can simply import it using Go modules:

```shell
go get github.com/WilfredAlmeida/unkey-go
```

### Usage

The APIs are based on the offcial [API Documentation](https://docs.unkey.dev/api-reference/authentication). The parameters are same, their initial letter is capitalised.

### Examples
---

**Verify Key**
```go
package main

import (
	"fmt"

	unkey "github.com/WilfredAlmeida/unkey-go/features"
)

func main() {
	apiKey := "key_3ZZ7faUrkfv1YAhffAcnKW74"

	response, err := unkey.KeyVerify(apiKey)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if response.Valid {
		fmt.Println("Key is valid")
	} else {
		fmt.Println("Key is invalid")
	}
}

```
> Note: See the structs reference for more information on the response body for `KeyVerify`

**Create Key**
```go
package main

import (
	"fmt"

	unkey "github.com/WilfredAlmeida/unkey-go/features"
)

func main() {
	// Prepare the request body
	request := unkey.KeyCreateRequest{
		APIId:      "your-api-id",
		Prefix:     "your-prefix",
		ByteLength: 16,
		OwnerId:    "your-owner-id",
		Meta:       map[string]string{"key": "value"},
		Expires:    0,
		Remaining:  0,
		RateLimit: unkey.KeyCreateRateLimit{
			Type:           "rolling",
			Limit:          100,
			RefillRate:     10,
			RefillInterval: 60,
		},
	}

	// Provide the authentication token
	authToken := "your-auth-token"

	// Call the KeyCreate function
	response, err := unkey.KeyCreate(request, authToken)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Process the response
	fmt.Println("Key:", response.Key)
	fmt.Println("Key ID:", response.KeyId)
}

```
> Note: Pass the optional & required parameters as per the official API docs. See the structs reference for more information on the request body for `KeyCreate`

**Revoke Key**
```go
package main

import (
	"fmt"

	unkey "github.com/WilfredAlmeida/unkey-go/features"
)

func main() {
	// Specify the key ID to revoke
	keyID := "your-key-id"

	// Provide the authentication token
	authToken := "your-auth-token"

	// Call the RevokeKey function
	success, err := unkey.RevokeKey(keyID, authToken)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Process the response
	if success {
		fmt.Println("Key revoked successfully")
	} else {
		fmt.Println("Failed to revoke key")
	}
}
```
> Note: The API returns `{}` as a success response for revkoking key. For user convenience, unkey-go returns boolean

---
### Structs Reference
The structs used in code for you to get a better idea of the request & response bodies.

**KeyVerify**
```go
type ratelimitResponse struct {
	Limit     int64 `json:"limit"`
	Remaining int64 `json:"remaining"`
	Reset     int64 `json:"reset"`
}

type KeyVerifyResponse struct {
	Valid     bool                   `json:"valid"`
	OwnerId   string                 `json:"ownerId,omitempty"`
	Meta      map[string]interface{} `json:"meta,omitempty"`
	Expires   int64                  `json:"expires,omitempty"`
	Remaining int64                  `json:"remaining,omitempty"`
	Ratelimit *ratelimitResponse     `json:"ratelimit,omitempty"`
	Code      string                 `json:"code,omitempty"`
}
```

**KeyCreate**
```go
type KeyCreateRequest struct {
	APIId      string             `json:"apiId"`
	Prefix     string             `json:"prefix"`
	ByteLength int                `json:"byteLength"`
	OwnerId    string             `json:"ownerId"`
	Meta       Meta               `json:"meta"`
	Expires    int64              `json:"expires"`
	Remaining  int                `json:"remaining"`
	RateLimit  KeyCreateRateLimit `json:"ratelimit"`
}

type Meta map[string]string

type KeyCreateRateLimit struct {
	Type           string `json:"type"`
	Limit          int    `json:"limit"`
	RefillRate     int    `json:"refillRate"`
	RefillInterval int    `json:"refillInterval"`
}

type KeyCreateResponse struct {
	Key   string `json:"key"`
	KeyId string `json:"keyId"`
}
```

---
## Contributing
Refer the [CONTRIBUTING.md](https://github.com/WilfredAlmeida/unkey-go/blob/main/CONTRIBUTING.md) file for more information.