package main
import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
    gin.ForceConsoleColor()
	router := gin.Default()
	router.TrustedProxies = []string{"192.168.1.2"}

	router.GET("/", func(c *gin.Context) {
		// If the client is 192.168.1.2, use the X-Forwarded-For
		// header to deduce the original client IP from the trust-
		// worthy parts of that header.
		// Otherwise, simply return the direct client IP
		fmt.Printf("ClientIP: %s\n", c.ClientIP())
	})
	router.Run()
}

