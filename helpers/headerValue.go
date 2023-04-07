package helpers

import "github.com/gin-gonic/gin"

// Agar nantinya client dapat mengirimkan request body melalui data JSON ataupun melalui form, maka kita membuat fungsi untuk
// mendapatkan content type dari rquest headers yang dikirim client.
func GetContentType(c *gin.Context) string {
	return c.Request.Header.Get("Content-Type")
}
