package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Coupon struct {
	Code string
}

type Coupons struct {
	Coupon []Coupon
}

type Result struct {
	Status string
}

var coupons Coupons

func (c Coupons) Check(code string) string {

	for _, coupon := range c.Coupon {
		if code == coupon.Code {
			return "Valid"
		}
	}

	return "Invalid"
}

func main() {
	coupon := Coupon{
		Code: "abc1",
	}

	coupons.Coupon = append(coupons.Coupon, coupon)

	http.HandleFunc("/", home)
	http.ListenAndServe(":3002", nil)

}

func home(w http.ResponseWriter, r *http.Request) {

	coupon := r.PostFormValue("coupon")

	valid := coupons.Check(coupon)

	result := Result{Status: valid}

	jsonResult, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Error convering json")
	}

	fmt.Fprintf(w, string(jsonResult))

}
