package main

import (
	"embed"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/webview/webview"
)

//go:embed web/public
var web embed.FS

func server(debug bool) {
	if debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	mdRegexp := ".md$"
	matcher, err := regexp.Compile(mdRegexp)

	if err != nil {
		log.Fatal("Error compiling regex (" + mdRegexp + ")")
	}

	app, err := fs.Sub(web, "web/public/app")
	if err != nil {
		log.Fatal("Error looking for subdirectory (app)")
	}
	r.StaticFS("/app", http.FS(app))

	r.GET("/api/list", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")

		files, _ := ioutil.ReadDir("./")
		markdown := []string{}

		for _, v := range files {
			match := matcher.MatchString(v.Name())

			if !v.IsDir() && match {
				markdown = append(markdown, v.Name())
			}
		}

		c.JSON(200, markdown)
	})

	r.GET("/api/read/:name", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")

		name, _ := c.Params.Get("name")
		content, _ := ioutil.ReadFile("./" + name)

		var json struct {
			Name    string
			Content string
		}

		json.Name = name
		json.Content = string(content)

		c.JSON(200, json)
	})

	r.POST("/api/write", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		name := c.PostForm("name")
		content := c.PostForm("content")

		os.WriteFile(name, []byte(content), 0644)
	})

	r.Run(":2001")
}

func app(debug bool) {
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Marked")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate("http://localhost:2001/app")
	w.Run()
}

func main() {
	debug := false

	if os.Getenv("DEV") == "true" {
		debug = true
	}

	if debug {
		server(debug)
	} else {
		go server(debug)
		app(debug)
	}
}
