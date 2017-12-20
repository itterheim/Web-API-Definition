package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()

	files, err := ioutil.ReadDir(".")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fileName := file.Name()
		if strings.Contains(fileName, ".mad") {
			loadDefinition(fileName)
		}
	}

	endTime := time.Since(startTime)
	fmt.Println("Done in:", endTime.Seconds(), "s")
}

func loadDefinition(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	parseDefinition(string(data))
}

func parseDefinition(definition string) {
	parser := NewParser(strings.NewReader(definition))

	app := parser.GetApp()

	createServer(&app)
}

func createServer(app *App) {

	for _, endpoint := range app.Endpoints {
		handler := createHandler(endpoint)
		http.HandleFunc(app.Options.RoutePrefix+endpoint.Route, handler)
	}

	port := ":"
	if len(app.Options.Port) > 0 {
		port += app.Options.Port
	} else {
		port += "80"
	}

	http.ListenAndServe(port, nil)
}

func createHandler(e Endpoint) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if len(e.Options.ContentType) > 0 {
			w.Header().Set("Content-Type", e.Options.ContentType)
		}
		w.Write([]byte(e.Options.Response))
	}
}
