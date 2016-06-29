package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"os"
)

var envBoxKey = os.Getenv("BOX_KEY")

var keyEncoding = base64.URLEncoding

func init() {
	registerHelp("generate",
		`usage: box generate

Generate securely generates a secretbox key to use for encrypting and decrypting files.

`)
}

func cmdGenerate() error {
	var k [32]byte
	if _, err := rand.Read(k[:]); err != nil {
		return err
	}
	fmt.Println(keyEncoding.EncodeToString(k[:]))
	return nil
}

func mustGetKey() *[32]byte {
	if envBoxKey == "" {
		log.Fatal("BOX_KEY environment variable must be set")
	}
	var key [32]byte
	raw, err := keyEncoding.DecodeString(envBoxKey)
	if err != nil {
		log.Fatal("invalid key in BOX_KEY: DecodeString:", err)
	}
	if n := copy(key[:], raw); n < len(key) {
		log.Fatal("invalid key in BOX_KEY: decoded less than 32 bytes")
	}
	return &key
}
