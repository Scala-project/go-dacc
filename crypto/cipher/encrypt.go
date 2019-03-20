// Copyright (C) 2017 go-dacc authors
//
// This file is part of the go-dacc library.
//
// the go-dacc library is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// the go-dacc library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with the go-dacc library.  If not, see <http://www.gnu.org/licenses/>.
//

package cipher

// Encrypt interface for encrypt
type Encrypt interface {

	// Encrypt encrypts data with passphrase,
	Encrypt(data []byte, passphrase []byte) ([]byte, error)

	// EncryptKey encrypt key with address
	EncryptKey(address string, data []byte, passphrase []byte) ([]byte, error)

	// Decrypt decrypts data with passphrase,  returning origin data.
	Decrypt(data []byte, passphrase []byte) ([]byte, error)

	// DecryptKey decrypts a key from a json blob, returning the private key itself.
	DecryptKey(keyjson []byte, passphrase []byte) ([]byte, error)
}
