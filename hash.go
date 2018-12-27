package gopp

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"hash"
	"io"
	"io/ioutil"
	"os"
)

func Md5(b []byte) []byte {
	sum := md5.Sum(b)
	return sum[:]
}

func Md5AsStr(b []byte) string {
	sum := md5.Sum(b)
	return hex.EncodeToString(sum[:])
}

func Md5Str(s string) string { return Md5AsStr([]byte(s)) }

func Md5File(filePath string) string {
	var returnMd5String string
	file, err := os.Open(filePath)
	if err != nil {
		return returnMd5String
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return returnMd5String
	}
	hashInBytes := hash.Sum(nil)[:16]
	returnMd5String = hex.EncodeToString(hashInBytes)
	return returnMd5String

}
func Md5File_dep(p string) string {
	bcc, err := ioutil.ReadFile(p)
	if err != nil {
		return ""
	}
	return Md5AsStr(bcc)
}

func Sha1(b []byte) []byte {
	sum := sha1.Sum(b)
	return sum[:]
}

func Sha1AsStr(b []byte) string {
	sum := sha1.Sum(b)
	return hex.EncodeToString(sum[:])
}

func Sha1Str(s string) string { return Md5AsStr([]byte(s)) }

func Sha1File(filePath string) string {
	return ShaxFile(filePath, sha1.New())
}

func Sha256AsStr(b []byte) string {
	sum := sha256.Sum256(b)
	return hex.EncodeToString(sum[:])
}

func Sha256File(filePath string) string {
	return ShaxFile(filePath, sha256.New())
}

func ShaxFile(filePath string, hash hash.Hash) string {
	//Initialize variable returnMD5String now in case an error has to be returned
	var returnSHA1String string

	//Open the filepath passed by the argument and check for any error
	file, err := os.Open(filePath)
	if err != nil {
		return returnSHA1String
	}

	//Tell the program to call the following function when the current function returns
	defer file.Close()

	//Open a new SHA1 hash interface to write to
	// hash := sha1.New()

	//Copy the file in the hash interface and check for any error
	if _, err := io.Copy(hash, file); err != nil {
		return returnSHA1String
	}

	//Get the 20 bytes hash
	hashInBytes := hash.Sum(nil)

	//Convert the bytes to a string
	returnSHA1String = hex.EncodeToString(hashInBytes)

	return returnSHA1String

}
func Sha1File_dep(p string) string {
	bcc, err := ioutil.ReadFile(p)
	if err != nil {
		return ""
	}
	return Sha1AsStr(bcc)
}
