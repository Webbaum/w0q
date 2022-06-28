package main

import (
	"github.com/Webbaum/w0q/config"
	"github.com/Webbaum/w0q/models"
	"github.com/fsuhrau/anonymizer"
	"github.com/gin-gonic/gin"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func RandStringBytesMask(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func printAscii() {
	var ascii string = "" +
		"             ___     ___________\n" +
		"            / _ \\   //_________/    \n" +
		" __________| | | |_// _ \n" +
		" \\\\ \\ /\\ / / | | |//_` |\n" +
		"  \\\\ V  V /| |_| |/(_| |\n" +
		"   \\\\_/\\_/__\\___//\\__, |\n" +
		"                     | |\n" +
		"                     |_|\n"
	print(ascii)
}

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("public/legal.html")

	config.ReadConfig()
	models.ConnectDatabase()

	// handle static files
	r.StaticFile("/", "./public/index.html")
	r.GET("/legal", func(c *gin.Context) {
		c.HTML(http.StatusOK, "public/legal.html", gin.H{
			"address": template.HTML(config.MyConfig.LegalAddress),
			"mail":    template.HTML(config.MyConfig.LegalMail),
			"phone":   template.HTML(config.MyConfig.LegalPhone),
		})
	})
	r.StaticFile("/favicon.png", "./public/favicon.png")
	r.StaticFile("/robots.txt", "./public/robots.txt")
	r.StaticFile("/styles.css", "./public/styles.css")

	r.POST("/url", func(c *gin.Context) {
		var input models.UrlInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var id string
		for {
			id = RandStringBytesMask(5)
			var result = models.DB.First(&models.Url{}, "id = ?", id)
			if result.RowsAffected == 0 {
				break
			}
		}

		url := models.Url{
			ID:   id,
			Url:  input.Url,
			Date: time.Now(),
			IP:   anonymizer.AnonymizeIP(c.ClientIP()),
		}
		models.DB.Create(&url)

		c.JSON(http.StatusOK, gin.H{"data": url})
	})

	r.GET("/:id", func(c *gin.Context) {
		var url models.Url

		if err := models.DB.First(&url, "id = ?", c.Param("id")).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
			return
		}

		c.Redirect(http.StatusMovedPermanently, url.Url)
	})
	printAscii()
	r.Run()
}
