package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"runtime"
	"strings"

	"golang.org/x/crypto/argon2"
)

const structure string = "$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s"
const saltLength byte = 16
const time uint32 = 1
const memory uint32 = 64 * 1024 // 64MB
const keyLen uint32 = 32

var threads uint8 = uint8(runtime.NumCPU()) / 2 // balance load

func genSalt() []byte {
	salt := make([]byte, saltLength)

	rand.Read(salt)

	return salt
}

func VerifyPassword(pass string, hashedPass string) error {
	parts := strings.Split(hashedPass, "$")

	var m, t uint32
	var p uint8
	hash, err := base64.RawURLEncoding.DecodeString(parts[4])

	if err != nil {
		return err
	}

	salt, err := base64.RawURLEncoding.DecodeString(parts[5])

	if err != nil {
		return err
	}

	keyLen := uint32(len(hash))

	fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &m, &t, &p)

	computeHash := argon2.IDKey([]byte(pass), salt, t, m, p, keyLen)

	if subtle.ConstantTimeCompare([]byte(hashedPass), computeHash) == 0 {
		return errors.New("password doesn't match")
	}

	return nil
}

func HashPass(password string) string {
	salt := genSalt()
	hash := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLen)

	return fmt.Sprintf(structure,
		argon2.Version, memory, time, threads,
		base64.RawURLEncoding.EncodeToString(salt),
		base64.RawURLEncoding.EncodeToString(hash))
}
