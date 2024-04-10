package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os/exec"
	"strconv"
)

var port int

func init() {
	flag.IntVar(&port, "port", 9427, "Port number to listen on")
	flag.Parse()
}

func main() {
	http.HandleFunc("/squeue", func(w http.ResponseWriter, r *http.Request) {
		// Run squeue command and capture output
		cmdOut, err := exec.Command("squeue", "--json").Output()
		if err != nil {
			fmt.Fprintf(w, "Error getting job information: %v", err)
			return
		}

		// Convert output to JSON object (optional)
		var jobs map[string]interface{}
		err = json.Unmarshal(cmdOut, &jobs)
		if err != nil {
			fmt.Fprintf(w, "Error parsing job data: %v", err)
			return
		}

		// Write JSON data (or raw output) to response
		w.Header().Set("Content-Type", "application/json")
		if jobs != nil {
			json.NewEncoder(w).Encode(jobs)
		} else {
			w.Write(cmdOut)
		}
	})

	fmt.Println("Starting server on port", port)
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
