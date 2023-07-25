package main

import (
	"fmt"
	"os"

	"github.com/WilfredAlmeida/unkey-go/features"
	"github.com/joho/godotenv"
)

func main() {
	// var isKeyValid, _ = features.VerifyRequest("key_3ZZ7faUrkfv1YAhffAcnKW74")
	// fmt.Println("Is key valid:", isKeyValid)

	// key := features.CreateKeyRequest{
	// 	APIId:      "api_ThrZA83W5BeUE3xm8onQ45",
	// 	Prefix:     "xyz",
	// 	ByteLength: 16,
	// 	OwnerId:    "wilfred",
	// 	Meta: features.Meta{
	// 		"hello": "world",
	// 	},
	// 	Expires: 1696951966471,
	// 	RateLimit: features.CreateKeyRateLimit{
	// 		Type:           "fast",
	// 		Limit:          10,
	// 		RefillRate:     1,
	// 		RefillInterval: 1000,
	// 	},
	// }

	// authToken := "unkey_3ZeRMoXP9NFgsSxjnNtjSGLF"

	// response, err := features.CreateKey(key, authToken)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// fmt.Println("Key: ", response.Key)
	// fmt.Println("Key ID: ", response.KeyId)

	// response, err := features.DeleteKey("key_VNZez78oQTwyB6ggSaGtTF", authToken)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// fmt.Println("Response: ", response)

	// response, err := features.KeyVerify("key_3ZZSawUTYL1DdsgCycdp7Xdu")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// fmt.Printf("Response.ratelimit: %+v\n", response.Ratelimit)

	err := godotenv.Load("./.env")

	if err != nil {
		fmt.Println("Error loading .env file")
		fmt.Println(err)
	}

	// response, err := features.APIGet(os.Getenv("API_ID"),os.Getenv("AUTH_TOKEN"))
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// fmt.Printf("Response: %+v\n", response)
	// response, err := features.APIListKeys(os.Getenv("API_ID"), os.Getenv("AUTH_TOKEN"))
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// fmt.Printf("Response: %+v\n", response)

	keyUpdateReq := features.KeyUpdateRequest{
		KeyId: "key_5KCLmfb2HY5czAfmEFNP3h",
		Name: features.NullableField[string]{Defined: false, Value: new(string)},
		OwnerId: features.NullableField[string]{Defined: false, Value: nil},
		Meta: features.NullableField[map[string]interface{}]{Defined: false, Value: new(map[string]interface{})},
		Expires: features.NullableField[int64]{Defined: false, Value: new(int64)},
		Ratelimit: features.NullableField[features.RateLimitSchema]{Defined: false, Value: nil},
		Remaining: features.NullableField[int64]{Defined: false, Value: new(int64)},
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

	// Set the value for Expires
	// *keyUpdateReq.Expires.Value = 123456789

	// Set the value for Remaining
	// *keyUpdateReq.Remaining.Value = 100

	response, err := features.KeyUpdate("xyz_3ZMvC1CMJ9AwVgob6aG1NaPd",keyUpdateReq, os.Getenv("AUTH_TOKEN"))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Response: %+v\n", response)

}
