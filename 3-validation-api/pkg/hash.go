package pkg

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io"
	mr "math/rand"
	"net/url"
	"os"
	"strconv"
	"time"
)

type VerifyFile struct {
	Email string `json:"email"`
	Hash  string `json:"hash"`
}

func MakeHash(mail string) (hash string, err error) {
	t := time.Now()
	rn := mr.Intn(999999999999999999)
	block, err := aes.NewCipher([]byte("q3IoYlWkgoPWUqDJtxk0WQ=="))
	if err != nil {
		return "", err

	}
	aesGsm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, aesGsm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return "", err
	}
	text := t.String() + mail + strconv.Itoa(rn)
	ciphertext := aesGsm.Seal(nonce, nonce, []byte(text), nil)
	encoded := base64.StdEncoding.EncodeToString(ciphertext)
	return encoded, nil
}

func MakeFile(email string, hash string) (urlString string, err error) {
	file, err := os.OpenFile("verify.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {

		return "", err
	}
	param := url.PathEscape(hash)
	urlString = "http://127.0.0.1:8082/verify/" + param
	var vf = []VerifyFile{}
	fileOpen, _ := os.ReadFile("verify.json")
	json.Unmarshal(fileOpen, &vf)
	vf = append(vf, VerifyFile{
		Email: email,
		Hash:  hash,
	})

	content, _ := json.Marshal(vf)
	err = os.WriteFile("verify.json", content, 06444)
	if err != nil {
		return "", err
	}
	defer file.Close()
	return urlString, nil
}
