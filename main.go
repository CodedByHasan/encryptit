package main

import (
    "fmt"
    "github.com/CodedByHasan/encryptit/decrypt"
    "github.com/CodedByHasan/encryptit/encrypt"
)

func main() {
    encryptStr := encrypt.Nimbus("Hello World")
    fmt.Println("Encrypted String: ", encryptStr)
    fmt.Println("Decrypting the string...")
    fmt.Println(decrypt.Nimbus(encryptStr))
}
