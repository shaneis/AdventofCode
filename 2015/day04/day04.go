package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func getMD5Hash(key string, zeroLen int) int {
	i := 1
	numZeros := strings.Repeat("0", zeroLen)
	for {
		joinedKey := fmt.Sprintf("%s%d", key, i)
		hash := md5.Sum([]byte(joinedKey))
		hashToString := hex.EncodeToString(hash[:])
		if hashToString[:zeroLen] == numZeros {
			return i
		}
		i++
	}
}

func main() {
	fileName := "puzzle_input.txt"
	f, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Printf(
		"The lowest positive number that produces a hash with %d leading zeros for %s is %d\n",
		5,
		string(f),
		getMD5Hash(string(f), 5),
	)
	fmt.Printf(
		"The lowest positive number that produces a hash with %d leading zeros for %s is %d\n",
		6,
		string(f),
		getMD5Hash(string(f), 6),
	)
}
