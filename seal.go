package main

import (
	"errors"
	"io/ioutil"
	"math/rand"
	"os"

	"golang.org/x/crypto/nacl/secretbox"
)

func init() {
	registerHelp("seal",
		`usage: box seal plain.txt > cipher.txt

Seal reads a plaintext file, encrypts it, and writes the encrypted bytes to stdout.

The environment variable BOX_KEY must be set to a valid secretbox key.

`)
}

func cmdSeal(key *[32]byte, inFile string) error {
	if inFile == "" {
		return errors.New("usage: box seal plain.txt > cipher.txt")
	}
	f, err := os.Open(inFile)
	if err != nil {
		return err
	}
	defer f.Close()
	pb, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	cb, err := encryptBytes(key, pb)
	if err != nil {
		return err
	}
	os.Stdout.Write(cb)
	return nil
}

func encryptBytes(k *[32]byte, b []byte) ([]byte, error) {
	var nonce [24]byte
	_, err := rand.Read(nonce[:])
	if err != nil {
		return nil, err
	}
	out := secretbox.Seal(nonce[:], b, &nonce, k)
	return out, nil
}
