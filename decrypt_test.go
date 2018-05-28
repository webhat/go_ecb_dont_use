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
	"log"
	"testing"
)

func TestDecrypt_base64(t *testing.T) {
	CIPHER_KEY := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	msg := []byte{0xC4, 0x26, 0xA6, 0xD3, 0x40, 0x6F, 0x56, 0x54, 0xD8, 0x20, 0x26, 0xFB, 0x02, 0x60, 0x85, 0xC2, 0xE0, 0xFC, 0xB6, 0xEA, 0x8A, 0xDF, 0xBB, 0xE7, 0xE6, 0x99, 0xBC, 0xA8, 0x80, 0x67, 0x26, 0x2F, 0x7C, 0x24, 0x28, 0xCB, 0x17, 0x5F, 0xAB, 0x9C, 0x4B, 0x48, 0x83, 0xE1, 0xEB, 0xFD, 0x00, 0x4C}

	if decrypted, err := decrypt(CIPHER_KEY, msg); err != nil {
		t.Errorf("%v", err)
	} else {
		log.Printf("DECRYPTED: %s\n", decrypted)
	}

}

func TestDecrypt(t *testing.T) {
	CIPHER_KEY := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	base64_msg := "/6q7AMQmptNAb1ZU2CAm+wJghcLg/Lbqit+75+aZvKiAZyYvfCQoyxdfq5xLSIPh6/0ATA=="
	if decrypted, err := decrypt_base64(CIPHER_KEY, base64_msg); err != nil {
		t.Errorf("%v", err)
	} else {
		log.Printf("DECRYPTED: %s\n", decrypted)
	}
}
