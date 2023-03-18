package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/uptrace/bunrouter"
)

var (
	router = bunrouter.New()
)

func routers(g *bunrouter.Group) {
	g.GET("/users/:id", debugHandler)
	g.GET("/users/current", debugHandler)
	g.GET("/users/*path", debugHandler)
	g.GET("/test", testHandler)
}

func init() {
	router.WithGroup("/api", routers)
}

func testHandler(w http.ResponseWriter, r bunrouter.Request) error {
	form := url.Values{}
	form.Add("version", "2")
	code := `
// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}
`
	form.Add("body", code)
	form.Add("withVet", "true")

	url := "https://go.dev/_/compile?backend="
	res, err := http.PostForm(url, form)
	// req, err := http.NewRequest("POST", url, strings.NewReader(form.Encode()))
	if err != nil {
		return bunrouter.JSON(w, bunrouter.H{
			"error": err.Error(),
		})
	}

	// req.Header.Add("accept", "*/*")
	// req.Header.Add("host", "go.dev")

	// res, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	return bunrouter.JSON(w, bunrouter.H{
	// 		"error": err.Error(),
	// 	})
	// }

	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return bunrouter.JSON(w, bunrouter.H{
			"error": err.Error(),
		})
	}

	var m = make(map[string]interface{})
	err = json.Unmarshal(b, &m)
	if err != nil {
		return bunrouter.JSON(w, bunrouter.H{
			"error": err.Error(),
		})
	}

	return bunrouter.JSON(w, m)
}

func debugHandler(w http.ResponseWriter, req bunrouter.Request) error {
	return bunrouter.JSON(w, bunrouter.H{
		"route":  req.Route(),
		"params": req.Params().Map(),
	})
}

func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
