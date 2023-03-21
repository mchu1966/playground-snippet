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
	g.POST("/run", runHandler)
	g.POST("/format", formatHandler)
}

func init() {
	router.WithGroup("/api", routers)
}

func runHandler(w http.ResponseWriter, r bunrouter.Request) error {
	defer r.Body.Close()

	rByte, err := io.ReadAll(r.Body)
	if err != nil {
		return bunrouter.JSON(w, bunrouter.H{
			"error": err.Error(),
		})
	}

	var body = make(map[string]interface{})
	err = json.Unmarshal(rByte, &body)
	if err != nil {
		return bunrouter.JSON(w, bunrouter.H{
			"error": err.Error(),
		})
	}

	form := url.Values{}
	form.Add("version", "2")
	form.Add("body", body["code"].(string))
	form.Add("withVet", "true")
	url := "https://go.dev/_/compile?backend="

	res, err := http.PostForm(url, form)
	if err != nil {
		return bunrouter.JSON(w, bunrouter.H{
			"error": err.Error(),
		})
	}

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

func formatHandler(w http.ResponseWriter, r bunrouter.Request) error {
	defer r.Body.Close()

	rByte, err := io.ReadAll(r.Body)
	if err != nil {
		return bunrouter.JSON(w, bunrouter.H{
			"error": err.Error(),
		})
	}

	var body = make(map[string]interface{})
	err = json.Unmarshal(rByte, &body)
	if err != nil {
		return bunrouter.JSON(w, bunrouter.H{
			"error": err.Error(),
		})
	}

	form := url.Values{}
	form.Add("body", body["code"].(string))
	form.Add("imports", "true")
	url := "https://go.dev/_/fmt?backend="

	res, err := http.PostForm(url, form)
	if err != nil {
		return bunrouter.JSON(w, bunrouter.H{
			"error": err.Error(),
		})
	}

	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return bunrouter.JSON(w, bunrouter.H{
			"error": err.Error(),
		})
	}

	var m = make(map[string]string)
	err = json.Unmarshal(b, &m)
	if err != nil {
		return bunrouter.JSON(w, bunrouter.H{
			"error": err.Error(),
		})
	}

	return bunrouter.JSON(w, m)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
