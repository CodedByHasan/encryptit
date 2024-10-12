package main

import (
    "fmt"
    "github.com/codedByHasan/encryptit/decrypt"
    "github.com/codedByHasan/encryptit/encrypt"
)

func main() {
    encryptStr := encrypt.Nimbus("Hello World")
    fmt.Println("Encrypted String: ", encryptStr)
    fmt.Println("Decrypting the string...")
    fmt.Println(decrypt.Nimbus(encryptStr))
}
