package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"myvrf/crypto-fork"

	"github.com/algorand/go-algorand/protocol"
)

type word string

func (m word) ToBeHashed() (protocol.HashID, []byte) {
	return "", []byte(m)
}

func main() {
	var hidden word
	var guess word

	fmt.Printf("Enter the word to be guessed: ")
	fmt.Scanln(&hidden)

	pk, sk := crypto.VrfKeygen()
	pkStr := hex.EncodeToString(pk[:])
	skStr := hex.EncodeToString(sk[:])

	proof, ok := sk.Prove(hidden)
	if !ok {
		fmt.Printf("Failed to generate proof\n")
		os.Exit(1)
	}
	proofStr := hex.EncodeToString(proof[:])

	fmt.Printf("\n== PUBLIC ==========\n")
	fmt.Printf("pk: %s\n", pkStr)
	fmt.Printf("proof: %s\n", proofStr)
	fmt.Printf("======================\n")

	fmt.Printf("\nEnter the word to be guessed: ")
	fmt.Scanln(&guess)

	fmt.Printf("\n== SECRET ==========\n")
	fmt.Printf("sk: %s\n", skStr)
	fmt.Printf("hidden word: %s\n", hidden)
	fmt.Printf("======================\n")

	fmt.Printf("\n== VERIFY ==========\n")
	fmt.Printf("guess word: %s\n", guess)
	fmt.Printf("pk: %s\n", pkStr)
	fmt.Printf("proof: %s\n", proofStr)

	pkB, _ := hex.DecodeString(pkStr)
	copy(pk[:], pkB)

	proofB, _ := hex.DecodeString(proofStr)
	copy(proof[:], proofB)

	fmt.Printf("\n")
	ok, out := pk.Verify(proof, guess)
	if !ok {
		fmt.Printf("Guess failed\n")
	} else {
		outStr := hex.EncodeToString(out[:])
		fmt.Printf("Guess ok: %s\n", outStr)
	}
	fmt.Printf("======================\n")

}
