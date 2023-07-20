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
	response, err := features.APIListKeys(os.Getenv("API_ID"), os.Getenv("AUTH_TOKEN"))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Response: %+v\n", response)

}
