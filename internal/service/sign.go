package service

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"

	"github.com/kdl-dev/iConText-test-task/internal/entity"
)

type signatureService struct {
}

func NewSignatureService() *signatureService {
	return &signatureService{}
}

func (s *signatureService) SHA512Sign(sha512Input *entity.HMACSHA512DTO) *entity.Signature {
	var token entity.Signature

	hash := hmac.New(sha512.New, []byte(sha512Input.Key))
	hash.Write([]byte(sha512Input.Text))

	token.Value = hex.EncodeToString(hash.Sum(nil))
	return &token
}
