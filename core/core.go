package core

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"github.com/pborman/uuid"
	"golang.org/x/crypto/pbkdf2"
)

const (
	// TimeTemplate is used for time template
	TimeTemplate = "2006-01-02T15:04:05Z"
)

// NewUUID return a random string ID
func NewUUID() string {
	return uuid.NewRandom().String()
}

// Vector is used for generate password
type Vector struct {
	password string
	salt     string
	iter     int
	length   int
}

// Bytes2Password generate password
func Bytes2Password(bytes []byte) string {
	h := sha256.New()
	h.Write(bytes)
	md := h.Sum(nil)
	password := hex.EncodeToString(md)

	return password
}

// GetSalt return the password's salt
func GetSalt(password string) string {
	sp := strings.Split(password, ";")
	salt := ""
	if len(sp) == 3 {
		salt = sp[1]
	}

	return salt
}

// GenerateSalt return a salt
func GenerateSalt(length int) string {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return ""
	}

	salt := bytes2String(bytes)
	return salt
}

func bytes2String(bytes []byte) string {
	for i := 0; i < len(bytes); i++ {
		bytes[i] = bytes[i]%93 + '!'
		if bytes[i] == ';' {
			bytes[i] = byte(i) + byte(1) + ';'
		}
	}
	return string(bytes)
}

// EncryptPassword encrypt password
func EncryptPassword(password string, salt string) string {
	var v Vector
	v.password = password
	v.salt = salt
	v.iter = 4096
	v.length = 25

	bytes := pbkdf2.Key([]byte(v.password), []byte(v.salt), v.iter, v.length, sha256.New)
	encryptPassword := "PBKDF2" + ";" + salt + ";" + Bytes2Password(bytes)

	return encryptPassword
}
