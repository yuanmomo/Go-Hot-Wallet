package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

func main() {
	stringInDb := "b-6S,c%Xxt!QmbZj56RjrM@3p"
	indexArray := []uint8{1, 2, 4, 8, 16}

	// calculate sha256
	sha256Sum := sha256.Sum256([]byte(stringInDb))
	hexOfSha256 := hex.EncodeToString(sha256Sum[:])
	fmt.Printf("SHA256 of [%s] is %s\n", stringInDb, hexOfSha256)

	// calculate md5
	md5Sum := md5.Sum([]byte(hexOfSha256))
	hexOfMd5 := hex.EncodeToString(md5Sum[:])
	fmt.Printf("MD5 of [%s] is %s\n", hexOfSha256, hexOfMd5)

	// get current nona second
	secondOfNow := time.Now().UnixNano()
	// calculate md5 of Nanosecond
	md5SumOfSecond := md5.Sum([]byte(strconv.FormatInt(secondOfNow, 10)))
	hexOfSecond := hex.EncodeToString(md5SumOfSecond[:])
	fmt.Printf("MD5 of [%d] is %s\n", secondOfNow, hexOfSecond)

	// get substrate index
	index, _ := strconv.ParseInt(string(hexOfSecond[0]), 16, 8)
	subIndex := index + 1

	//
	newMd5HexOfSecond := hexOfSecond[subIndex:] + hexOfSecond[1:subIndex]
	fmt.Printf("Substrate index: [%d] \n", index)
	fmt.Printf("Firtst: [%s] \n", hexOfSecond[1:subIndex])
	fmt.Printf("Second: [%s] \n", hexOfSecond[subIndex:])
	fmt.Printf("New   : [%s] \n", newMd5HexOfSecond)
	fmt.Printf("Len of new   : [%d] \n", len(newMd5HexOfSecond))


}

func getPasswordFromRandom(random string) {

}
