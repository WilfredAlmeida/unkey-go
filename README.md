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
			Type:           "fast",
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


** Update Key**
```go
package main

import (
	"fmt"

	unkey "github.com/WilfredAlmeida/unkey-go/features"
)

func main() {

	// The API for update has optional params. To determine which params to update, a custome type is made which mimics the behaviour of `undefined`/`null`.
	// To update, define a object in the following manner & set the value for the fields you want to update following on.

	keyUpdateReq := unkey.KeyUpdateRequest{
		KeyId: "someKey",
		Name: unkey.NullableField[string]{Defined: false, Value: new(string)},
		OwnerId: unkey.NullableField[string]{Defined: false, Value: nil},
		Meta: unkey.NullableField[map[string]interface{}]{Defined: false, Value: new(map[string]interface{})},
		Expires: unkey.NullableField[int64]{Defined: false, Value: new(int64)},
		Ratelimit: unkey.NullableField[unkey.RateLimitSchema]{Defined: false, Value: nil},
		Remaining: unkey.NullableField[int64]{Defined: false, Value: new(int64)},
	}

	// Set the value for Name
	keyUpdateReq.Name.Defined = true
	*keyUpdateReq.Name.Value = "John Doe"

	// Set the value for Meta
	keyUpdateReq.Meta.Defined = true
	metaData := map[string]interface{}{
		"field1": "value1",
		"field2": 42,
	}
	*keyUpdateReq.Meta.Value = metaData

	response, err := features.KeyUpdate("someKey",keyUpdateReq, os.Getenv("AUTH_TOKEN"))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Response: %+v\n", response)

}
```

**Get API Information**
```go
package main

import (
	"fmt"

	unkey "github.com/WilfredAlmeida/unkey-go/features"
)

func main() {
	response, err := features.APIGet("API_ID","AUTH_TOKEN")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Response: %+v\n", response)
}
```

**Get List Keys**
```go
package main

import (
	"fmt"

	unkey "github.com/WilfredAlmeida/unkey-go/features"
)

func main() {
	response, err := features.APIListKeys("API_ID","AUTH_TOKEN")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Response: %+v\n", response)
}
```

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
    Valid       bool                   `json:"valid"`
    OwnerId     string                 `json:"ownerId,omitempty"`
    Meta        map[string]interface{} `json:"meta,omitempty"`
    Expires     int64                  `json:"expires,omitempty"`
    Remaining   int64                  `json:"remaining,omitempty"`
    Ratelimit   *ratelimitResponse     `json:"ratelimit,omitempty"`
    Code        string                 `json:"code,omitempty"`
    Enabled     bool                   `json:"enabled"`
    Environment string                 `json:"environment"`
    KeyId       string                 `json:"keyId"`
    Name        string                 `json:"name"`
    Permissions []string               `json:"permissions"`
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

**KeyUpdate**
```go
type NullableField[T any] struct {
	Defined bool
	Value   *T
}

func (m *NullableField[T]) UnmarshalJSON(data []byte) error {
	m.Defined = true
	return json.Unmarshal(data, &m.Value)
}

type RateLimitSchema struct {
	Type           string `json:"type" validate:"required"`
	Limit          int64  `json:"limit" validate:"required"`
	RefillRate     int64  `json:"refillRate" validate:"required"`
	RefillInterval int64  `json:"refillInterval" validate:"required"`
}

type KeyUpdateRequest struct {
	KeyId     string                         `json:"keyId" validate:"required"`
	Name      NullableField[string]          `json:"name"`
	OwnerId   NullableField[string]          `json:"ownerId"`
	Meta      NullableField[map[string]any]  `json:"meta"`
	Expires   NullableField[int64]           `json:"expires"`
	Ratelimit NullableField[RateLimitSchema] `json:"ratelimit"`
	Remaining NullableField[int64]           `json:"remaining"`
}
```

**APIGet**
```go
type APIGetResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	WorkspaceID string `json:"workspaceId"`
}
```

**APIListKeys**
```go
type Key struct {
	ID          string            `json:"id,omitempty"`
	APIID       string            `json:"apiId,omitempty"`
	WorkspaceID string            `json:"workspaceId,omitempty"`
	Start       string            `json:"start,omitempty"`
	CreatedAt   int64             `json:"createdAt,omitempty"`
	Expires     *int64            `json:"expires,omitempty"`
	Ratelimit   ratelimitResponse `json:"ratelimit,omitempty"`
}

type APIListKeysResponse struct {
	Keys  []Key `json:"keys,omitempty"`
	Total int   `json:"total,omitempty"`
}
```
> Note: `ratelimitResponse` is defined in `KeyVerify` struct reference

---
## Contributing
Refer the [CONTRIBUTING.md](https://github.com/WilfredAlmeida/unkey-go/blob/main/CONTRIBUTING.md) file for more information.
