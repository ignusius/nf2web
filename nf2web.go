package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

type Config struct {
	Address  string `json:"Address"`
	Port     string `json:"Port"`
	Commands []struct {
		Title   string   `json:"Title"`
		Command []string `json:"Command"`
	} `json:"Commands"`
}

func main() {
	// Open config.json
	jsonFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := io.ReadAll(jsonFile)
	defer jsonFile.Close()

	config := Config{}
	json.Unmarshal([]byte(byteValue), &config)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		tmpl, _ := template.ParseFiles("index.html")
		//Make map from data template
		m := make(map[string]string)
		for _, comm := range config.Commands {

			c := comm.Command
			log.Println(c)
			out, _ := exec.Command(c[0], c[1:]...).Output()

			m[comm.Title] = string(out)

		}
		fmt.Println(m)

		tmpl.Execute(w, m)
	})

	fmt.Println("Server is listening " + config.Address + ":" + config.Port)
	http.ListenAndServe(config.Address+":"+config.Port, nil)
}
