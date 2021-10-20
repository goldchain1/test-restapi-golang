package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	routeRequest()
}
func routeRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/student", student)
	http.HandleFunc("/subject", subject)
	http.ListenAndServe(":1111", nil)
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome")
}
func student(w http.ResponseWriter, r *http.Request) {
	m := []struct {
		Code       int
		Name       string
		Major      string
		Course     string
		University string
	}{
		{
			Code:       111,
			Name:       "เชน",
			Major:      "it",
			Course:     "",
			University: "Rajabhat",
		},
		{
			Code:       222,
			Name:       "แดง",
			Major:      "it",
			Course:     "",
			University: "Rajabhat",
		},
		{
			Code:       333,
			Name:       "เขียว",
			Major:      "it",
			Course:     "",
			University: "Rajabhat",
		},
		{
			Code:       444,
			Name:       "ขาว",
			Major:      "it",
			Course:     "",
			University: "Rajabhat",
		},
	}

	js, _ := json.Marshal(m)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
func subject(w http.ResponseWriter, r *http.Request) {
	m := []struct {
		Subject string
	}{
		{
			Subject: "คอมพิวเตอร์",
		},
		{
			Subject: "อังกฤษ",
		},
		{
			Subject: "ไทย",
		},
		{
			Subject: "คณิต",
		},
	}

	js, _ := json.Marshal(m)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
