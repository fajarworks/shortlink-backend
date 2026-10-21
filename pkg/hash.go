package pkg

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type HashConf struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

var (
	errInvalidHash    = errors.New("incorrect hash format")
	errInvalidType    = errors.New("incorrect hash type")
	errInvalidVersion = errors.New("inccorrect hash version")
	errMismatchHash   = errors.New("mismatch hash")
)

func NewRecomHashConfig() *HashConf {
	return &HashConf{
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}

}

func (h *HashConf) GenSalt() ([]byte, error) {
	salt := make([]byte, h.saltLength)
	_, err := rand.Read(salt)

	if err != nil {
		return nil, err
	}

	return salt, nil
}

func (h *HashConf) GenHash(password string) (string, error) {
	salt, err := h.GenSalt()

	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, h.iterations, h.memory, h.parallelism, h.keyLength)

	base64Hash := base64.RawStdEncoding.EncodeToString(hash)
	base64Salt := base64.RawStdEncoding.EncodeToString(salt)

	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, h.memory, h.iterations, h.parallelism, base64Salt, base64Hash)

	return encodedHash, nil

}

func ComparePassAndHash(password, hashedPass string) error {
	result := strings.Split(hashedPass, "$")

	if len(result) != 6 {
		return errInvalidHash
	}
	if result[1] != "argon2id" {
		return errInvalidType
	}
	var version int
	if _, err := fmt.Sscanf(result[2], "v=%d", &version); err != nil {
		return err
	}
	if version != argon2.Version {
		return errInvalidVersion
	}

	var memory, iterations uint32
	var parallelism uint8
	_, err := fmt.Sscanf(result[3], "m=%d,t=%d,p=%d", &memory, &iterations, &parallelism)
	if err != nil {
		return err
	}

	salt, err := base64.RawStdEncoding.DecodeString(result[4])
	if err != nil {
		return err
	}

	hash, err := base64.RawStdEncoding.DecodeString(result[5])
	if err != nil {
		return err
	}

	newHash := argon2.IDKey([]byte(password), salt, iterations, memory, parallelism, uint32(len(hash)))

	if subtle.ConstantTimeCompare(hash, newHash) == 0 {
		return errMismatchHash
	}

	return nil

}
