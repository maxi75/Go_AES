package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"time"
)

func main() {
	//Festlegung eines einheitlichen Keys
	key := "vzxgk9PH/Zj31vjRpAzyolTNJjIQ+/FZ"

	//Festlegung eines zu ver- und entschlüsselnden Textes
	data := "Lorem ipsum dolor sit amet, consectetur adipisici elit, sed eiusmod tempor incidunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquid ex ea commodi consequat. Quis aute iure reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint obcaecat cupiditat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

	//Zeitmessung und iteratives Ver- und Entschlüsseln
	start := time.Now()
	for i := 10000000; i > 0; i-- {
		data = AESEncrypt(data, []byte(key))
		data = AESDecrypt(data, []byte(key))
	}
	ende := time.Since(start) / time.Millisecond
	fmt.Printf("Dauer in ms: %dms", ende)
}

//Verschlüsselungsfunktion
//Berechnet einen verschlüsselten Text aus einem Eingabetext und einem Schlüssel
func AESEncrypt(data string, key []byte) string {
	bytes := []byte(data)
	block, _ := aes.NewCipher(key)
	cipherText := make([]byte, aes.BlockSize+len(data))
	iv := cipherText[:aes.BlockSize]
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], bytes)
	return base64.StdEncoding.EncodeToString(cipherText)
}

//Entschlüsselungsfunktion
//Berechnet einen Text aus einem verschlüsselten Eingabetext und einem Schlüssel
func AESDecrypt(data string, key []byte) string {
	cipherText, _ := base64.StdEncoding.DecodeString(data)
	block, _ := aes.NewCipher(key)
	if len(cipherText) < aes.BlockSize {
		err := "Ciphertext to short"
		return err
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)
	return string(cipherText)
}
