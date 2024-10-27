package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Codec struct {
	urlToCode map[string]string
	codeToUrl map[string]string
	baseURL   string
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Constructor() Codec {
	return Codec{
		urlToCode: make(map[string]string),
		codeToUrl: make(map[string]string),
		baseURL:   "http://tinyurl.com/",
	}
}

func generateCode() string {
	code := make([]byte, 6)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}
	return string(code)
}

func (this *Codec) encode(longUrl string) string {
	if code, exists := this.urlToCode[longUrl]; exists {
		return this.baseURL + code
	}

	code := generateCode()
	for _, exists := this.codeToUrl[code]; exists; code = generateCode() {
		// Regenerate code in case of collision
	}

	this.urlToCode[longUrl] = code
	this.codeToUrl[code] = longUrl
	return this.baseURL + code
}

func (this *Codec) decode(shortUrl string) string {
	code := shortUrl[len(this.baseURL):]
	return this.codeToUrl[code]
}

func main() {
	obj := Constructor()

	fmt.Print("Enter the long URL to encode: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	longUrl := strings.TrimSpace(scanner.Text())

	encodedUrl := obj.encode(longUrl)
	decodedUrl := obj.decode(encodedUrl)

	fmt.Println("Original URL:", longUrl)
	fmt.Println("Encoded URL:", encodedUrl)
	fmt.Println("Decoded URL:", decodedUrl)
}
