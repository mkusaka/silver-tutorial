package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		n := c.Query("n")
		if n == "" {
			// render view
			c.HTML(http.StatusOK, "index.tmpl.html", nil)
		} else {
			i, err := strconv.Atoi(n)
			if err != nil || i < 1 || i > 10000 {
				// render error
				c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
					"error": "Please submit a valid number between 1 and 10000.",
				})
			} else {
				p := calculatePrime(i)
				// rende prime
				c.HTML(http.StatusOK, "index.tmpl.html", gin.H{"n": i, "prime": p})
			}
		}
	})

	router.Run(":" + port)
}

func calculatePrime(n int) int {
	prime := 1
	for i := n; i > 1; i-- {
		isPrime := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			prime = i
			break
		}
	}
	return prime
}
