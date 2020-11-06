package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Result struct {
	Status string
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/process", process)
	fmt.Println("CHECKOUT SERVICE LISTEN ON PORT 3000")
	http.ListenAndServe(":3000", nil)

}

func home(w http.ResponseWriter, r *http.Request) {

	// Working Directory
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("error working directory")
	}

	fmt.Println("\n\n\n\n wd", wd)

	tmpl := template.Must(template.ParseFiles("microservice-a/ui/html/home.html"))
	log.Println("/")
	tmpl.Execute(w, Result{})
}

func process(w http.ResponseWriter, r *http.Request) {

	serviceUrlPayment := "http://localhost:3001"
	coupon := r.FormValue("coupon")
	ccNumber := r.FormValue("cc-number")

	log.Println("form_data", coupon, ccNumber)

	result := makeHttpCall(serviceUrlPayment, coupon, ccNumber)

	tmpl := template.Must(template.ParseFiles("microservice-a/ui/html/home.html"))
	tmpl.Execute(w, result)

}

func makeHttpCall(urlMicroservice, coupon, ccNumber string) Result {

	values := url.Values{}
	values.Add("coupon", coupon)
	values.Add("ccNumber", ccNumber)

	res, err := http.PostForm(urlMicroservice, values)
	if err != nil {
		msg := "PAYMENT SERVICE OUT"
		log.Println(msg)
		return Result{Status: msg}
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error reading response body ")
	}

	result := Result{}

	json.Unmarshal(data, &result)

	log.Println("json result: ", result)

	return result

}
