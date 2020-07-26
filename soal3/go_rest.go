package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func index(w http.ResponseWriter, r *http.Request) {
	msgData := map[string]interface{}{
		"message": "welcome to rest",
	}

	// JSON encoding is done the same way as before
	data, _ := json.Marshal(msgData)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(data))
}

func CountPalindrom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	i1, err := strconv.Atoi(vars["start"])
	if err == nil {
		fmt.Println(i1)
	}
	i2, err := strconv.Atoi(vars["finish"])
	if err == nil {
		fmt.Println(i2)
	}
	var total int = CountTotalPalindrom(i1, i2)
	msgData := map[string]interface{}{
		//"birdSounds": map[string]string{
		//	"pigeon": "coo",
		//	"eagle":  "squak",
		//},
		"message": "Total Palindrom",
		"value":   total,
	}

	// JSON encoding is done the same way as before
	data, _ := json.Marshal(msgData)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(data))
}

func CountTotalPalindrom(start int, end int) int {
	var totalpalindrom int = 0
	for i := start; i < end; i++ {
		var originalNumber int = i
		var reverseNumber int = 0
		var temp = originalNumber

		for temp > 0 {
			var remainder = temp % 10
			reverseNumber = (reverseNumber * 10) + remainder
			temp = temp / 10
		}

		if originalNumber == reverseNumber {
			totalpalindrom++;
		}
	}
	return totalpalindrom
}

func handleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", index);
	r.HandleFunc("/countpalindrom/{start}/{finish}", CountPalindrom);
	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	handleRequest();
}
