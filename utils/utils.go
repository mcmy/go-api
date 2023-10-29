package utils

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const CharsetDefault = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandomStr(length int, charset ...string) string {
	ct := ""
	if len(charset) == 0 || charset[0] == "" {
		ct = CharsetDefault
	} else {
		ct = charset[0]
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	str := make([]byte, length)
	for i := range str {
		str[i] = ct[r.Intn(len(ct))]
	}

	return string(str)
}

func CalculateMD5(input string) string {
	hasher := md5.New()
	hasher.Write([]byte(input))
	return hex.EncodeToString(hasher.Sum(nil))
}

func CalculateSHA1(input string) string {
	hasher := sha1.New()
	hasher.Write([]byte(input))
	return hex.EncodeToString(hasher.Sum(nil))
}

func CalculateSHA256(input string) string {
	hasher := sha256.New()
	hasher.Write([]byte(input))
	return hex.EncodeToString(hasher.Sum(nil))
}

func ReadFileMd5(filePath string) (string, error) {
	f, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	hasher := md5.New()
	hasher.Write(f)
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

// ByteSerialize 序列化数据
func ByteSerialize(data interface{}) ([]byte, error) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	if err := encoder.Encode(data); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

// ByteDeserialize 反序列化数据
func ByteDeserialize(data []byte, target interface{}) error {
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	return decoder.Decode(target)
}

func FileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func GetExistFile(def string, files ...string) string {
	f := ""
	for _, val := range files {
		if FileExist(val) {
			f = val
			break
		}
	}
	if f == "" {
		return def
	}
	return f
}

func WaitQuit() {
	var str string
	for {
		_, err := fmt.Scan(&str)
		if err != nil {
			continue
		}
		switch str {
		case "quit":
			os.Exit(0)
		default:
			fmt.Println(`input "quit" to quit`)
		}
	}
}
