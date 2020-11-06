package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Result struct {
	Status string
}

func main() {

	http.HandleFunc("/", home)
	fmt.Println("PAYMENT SERVICE LISTEN ON PORT 3001")
	http.ListenAndServe(":3001", nil)

}

func home(w http.ResponseWriter, r *http.Request) {

	serviceUrlCoupon := "http://localhost:3002"
	coupon := r.PostFormValue("coupon")
	ccNumber := r.PostFormValue("ccNumber")
	result := Result{Status: "Declined"}

	fmt.Printf("PAYMENT SERVICE RECEIVED. [Coupon: %s] [ccNumber: %s]", coupon, ccNumber)

	resultCoupon := makeHttpCall(serviceUrlCoupon, coupon, ccNumber)

	if coupon != "" && resultCoupon.Status == "Invalid" {
		result.Status = "Invalid coupon"
	}

	if ccNumber == "1" {
		result.Status = "Approved"
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Error processing json")
	}

	fmt.Fprintf(w, string(jsonData))
}

func makeHttpCall(urlMicroservice, coupon, ccNumber string) Result {

	values := url.Values{}
	values.Add("coupon", coupon)
	values.Add("ccNumber", ccNumber)

	res, err := http.PostForm(urlMicroservice, values)
	if err != nil {
		msg := "COUPON SERVICE OUT"
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
