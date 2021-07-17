package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// index for contacting two md5 string
var contactIndex = []int{2, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 2, 2, 1, 1, 1, 1, 2, 2, 1, 1, 1, 2, 1, 1}

// get password from this index
var passwordLength = []int{5, 24, 19, 23, 2, 11, 20, 9, 3, 29, 17, 27, 6, 13, 31, 16}

func main() {
	var seed string
	for i := 0; i <= 10000; i++ {
		fmt.Printf("The index: [%d]\n", i)

		// get random seed
		seed = RandStringRunes(32)

		// get password and salt
		originalPassword, salt := generateSalt(seed)

		// get md5OfSeed and md5OfSecond from salt
		passwordFromSalt := getPasswordFromSalt(salt)

		// log if password not equals
		if strings.Compare(originalPassword, passwordFromSalt) != 0 {
			fmt.Printf("Password not eqult for seed: [%s]\n", seed)
			fmt.Printf("Salt: [%s]\n", salt)
			fmt.Printf("OriginalPassword: [%s]\n", originalPassword)
			fmt.Printf("PasswordFromSalt: [%s]\n", passwordFromSalt)
			fmt.Println()
		}
	}
}

// generate salt
func generateSalt(seed string) (string, string) {
	// get md5 of seed
	var sha256OfSeedResult [sha256.Size]byte
	for i := 0; i <= 3; i++ {
		sha256OfSeedResult = sha256.Sum256([]byte(seed))
	}
	md5SumOfSeed := md5.Sum([]byte(hex.EncodeToString(sha256OfSeedResult[:])))
	md5HexOfSeed := hex.EncodeToString(md5SumOfSeed[:])

	// get original originalPassword
	originalPassword := getPassword(md5HexOfSeed)

	// get md5 of current nanosecond
	md5OfNanosecond := getMd5OfNanosecond()

	// get split index for md5HexOfSeed
	index := getSplitIndex(md5OfNanosecond)
	// split md5HexOfSeed and recontact
	splitMd5HexOfSeed := md5HexOfSeed[index:] + md5HexOfSeed[0:index]

	var salt string
	currentIndex := 0
	for _, value := range contactIndex {
		// contact salt
		salt = salt +
			splitMd5HexOfSeed[currentIndex : currentIndex+value] +
			md5OfNanosecond[currentIndex : currentIndex+value]

		currentIndex = currentIndex + value
	}

	return originalPassword, salt
}

func getPasswordFromSalt(salt string) string {
	var md5HexOfSeed, md5HexOfNanosecond string

	currentIndex := 0
	for _, value := range contactIndex {
		md5HexOfSeed = md5HexOfSeed + salt[currentIndex:currentIndex+value]
		md5HexOfNanosecond = md5HexOfNanosecond + salt[currentIndex+value:currentIndex+value*2]

		currentIndex = currentIndex + value*2
	}

	subIndex := getSplitIndex(md5HexOfNanosecond)
	originalMd5OfSeed := md5HexOfSeed[int64(len(md5HexOfNanosecond))-subIndex:len(md5HexOfNanosecond)] +
		md5HexOfSeed[0:(int64(len(md5HexOfNanosecond))-subIndex)]

	return getPassword(originalMd5OfSeed)
}

//get password
func getPassword(md5HexOfSeed string) string {
	var password string
	for _, index := range passwordLength {
		password = password + md5HexOfSeed[index:index+1]
	}
	return password
}

// get split index
func getSplitIndex(md5String string) int64 {
	head, _ := strconv.ParseInt(string(md5String[0]), 16, 8)
	tail, _ := strconv.ParseInt(string(md5String[len(md5String)-1]), 16, 8)
	// index to split original md5OfSeed
	return head + tail
}

func getMd5OfNanosecond() string {
	// get current nona second
	secondOfNow := time.Now().UnixNano()

	// calculate md5 of Nanosecond
	md5SumOfSecond := md5.Sum([]byte(strconv.FormatInt(secondOfNow, 10)))
	return hex.EncodeToString(md5SumOfSecond[:])
}
