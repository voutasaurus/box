package main

import (
	"errors"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/nacl/secretbox"
)

var ErrInvalidCipher = errors.New("Invalid Cipher: BOX_KEY could not decrypt file provided")

func init() {
	registerHelp("open",
		`usage: box open cipher.txt > plain.txt

Open reads a ciphertext file, decrypts it, and writes the decrypted bytes to stdout.

The environment variable BOX_KEY must be set to a valid secretbox key.

`)
}

func cmdOpen(key *[32]byte, inFile string) error {
	if inFile == "" {
		return errors.New("usage: box open cipher.txt > plain.txt")
	}
	f, err := os.Open(inFile)
	if err != nil {
		return err
	}
	defer f.Close()
	cb, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	pb, err := decryptBytes(key, cb)
	if err != nil {
		return err
	}
	os.Stdout.Write(pb)
	return nil
}

func decryptBytes(key *[32]byte, b []byte) ([]byte, error) {
	if len(b) < 24 {
		return nil, ErrInvalidCipher
	}
	var nonce [24]byte
	copy(nonce[:], b)
	out, ok := secretbox.Open(nil, b[len(nonce):], &nonce, key)
	if !ok {
		return nil, ErrInvalidCipher
	}
	return out, nil
}
