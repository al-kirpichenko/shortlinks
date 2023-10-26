package keygen

import "fmt"

func Example() {
	key := GenerateKey()
	fmt.Println(key)
}
