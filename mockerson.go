package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type jsonResponse struct {
	Path         string `json:"path"`
	Method       string `json:"method"`
	ResponseCode int    `json:"code"`
	ResponseBody string `json:"body"`
	File         string `json:"file"`
}

var allResponses []jsonResponse

func mockersonHandler(w http.ResponseWriter, r *http.Request) {
	success := false
	for _, resp := range allResponses {
		if r.Method == resp.Method && r.RequestURI == resp.Path {
			w.WriteHeader(resp.ResponseCode)
			if resp.File != "" {
				file, err := ioutil.ReadFile(resp.File)
				if err == nil {
					fmt.Fprintf(w, string(file))
					log.Printf(
						"Response %d %s for %s (%s) from file %s",
						resp.ResponseCode,
						http.StatusText(resp.ResponseCode),
						r.RequestURI,
						r.Method,
						resp.File,
					)
				} else {
					log.Printf(
						"Error loading file %s (%s)",
						resp.File,
						err.Error(),
					)
				}
			} else {
				fmt.Fprintf(w, resp.ResponseBody)
				log.Printf(
					"Response %d %s for %s (%s)",
					resp.ResponseCode,
					http.StatusText(resp.ResponseCode),
					r.RequestURI,
					r.Method,
				)
			}
			return
		}
	}

	if !success {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Not found in json")
	}
}

func main() {
	var port int
	var jsonFile string
	flag.IntVar(&port, "port", 8080, "Add port to run")
	flag.StringVar(&jsonFile, "json", "default.json", "Change json to read responses")
	flag.Parse()
	http.HandleFunc("/", mockersonHandler)

	runningOn := fmt.Sprintf("localhost:%d", port)
	log.Printf("Running on %s", runningOn)
	log.Printf("Responses are in %s", jsonFile)
	file, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		panic(fmt.Sprintf("File: %s", err.Error()))
	}
	if e := json.Unmarshal([]byte(file), &allResponses); e != nil {
		panic(fmt.Sprintf("Unmarshal: %s", e.Error()))
	}

	http.ListenAndServe(runningOn, nil)
}
