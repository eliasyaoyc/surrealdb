// Copyright © 2016 Abcum Ltd
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

package cryp

import (
	"crypto/aes"
	"crypto/cipher"
	"github.com/abcum/surreal/util/rand"
)

func Encrypt(key []byte, src []byte) (dst []byte, err error) {

	if len(key) == 0 {
		return src, nil
	}

	// Initiate AES256
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	// Initiate cipher
	cipher, err := cipher.NewGCM(block)
	if err != nil {
		return
	}

	nonce := rand.New(12)

	dst = cipher.Seal(nil, nonce, src, nil)

	dst = append(nonce[:], dst[:]...)

	return

}

func Decrypt(key []byte, src []byte) (dst []byte, err error) {

	if len(key) == 0 {
		return src, nil
	}

	// Initiate AES256
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	// Initiate cipher
	cipher, err := cipher.NewGCM(block)
	if err != nil {
		return
	}

	return cipher.Open(nil, src[:12], src[12:], nil)

}