package medium

import (
	"encoding/base64"
	"fmt"
	"strings"
)

type Codec struct {
	prefix       string
	equalsAmount int
}

func Constructor() Codec {
	prefix := "http://tinyurl.com/"

	return Codec{prefix, 0}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	encodedUrl := base64.StdEncoding.EncodeToString([]byte(longUrl))

	this.equalsAmount = strings.Count(encodedUrl, "=")

	return fmt.Sprintf("%s%s", this.prefix, encodedUrl[:len(encodedUrl)-this.equalsAmount])
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	prefixLen := len(this.prefix)

	equals := ""
	for i := 0; i < this.equalsAmount; i++ {
		equals += "="
	}

	decodedUrl, _ := base64.StdEncoding.DecodeString(fmt.Sprintf("%s%s", shortUrl[prefixLen:], equals))

	return fmt.Sprintf("%s", decodedUrl)
}
