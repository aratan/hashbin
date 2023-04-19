package main

import (
    "crypto/sha512"
    "fmt"
    "io"
    "os"
)

// Función para calcular el hash de un archivo
func calculateHash(filename string) ([]byte, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    hash := sha512.New()
    if _, err := io.Copy(hash, file); err != nil {
        return nil, err
    }

    return hash.Sum(nil), nil
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Debes proporcionar el nombre del archivo como argumento.")
        return
    }
    filename := os.Args[1]

    hash, err := calculateHash(filename)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Printf("%x\n", hash)
}
