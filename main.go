package main

import (
	"fmt"
	"github.com/WilfredAlmeida/unkey-go/features"
)



func main() {
	var isKeyValid, _ = features.VerifyRequest("key_3ZZ7faUrkfv1YAhffAcnKW74")
	fmt.Println("Is key valid:", isKeyValid)

}
