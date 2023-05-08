package service

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"

	"github.com/kdl-dev/iConText-test-task/internal/entity"
)

type signature struct {
}

func NewSignature() *signature {
	return &signature{}
}

func (s *signature) SHA512Sign(sha512Input *entity.HMACSHA512DTO) *entity.Signature {
	var token entity.Signature

	hash := hmac.New(sha512.New, []byte(sha512Input.Key))
	hash.Write([]byte(sha512Input.Text))

	token.Value = hex.EncodeToString(hash.Sum(nil))
	return &token
}
