// Copyright 2018, Special Brands BV
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"crypto/aes"
	"encoding/base64"
	"errors"
)

func decrypt_base64(key []byte, secureMessage string) (decodedMessage string, err error) {
	cipherText, err := base64.StdEncoding.DecodeString(secureMessage)
	if err != nil {
		return
	}
	decodedMessage, err = decrypt(key, cipherText[4:])

	return
}

func decrypt(key []byte, cipherText []byte) (decodedmess string, err error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	if len(cipherText) < aes.BlockSize {
		err = errors.New("Ciphertext block size is too short!")
		return
	}

	decrypter := NewECBDecrypter(block)
	p := make([]byte, len(cipherText))
	decrypter.CryptBlocks(p, []byte(cipherText))

	decodedmess = string(p)
	return
}
