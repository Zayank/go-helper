package controllers

import (
	"encoding/json"
	"golang-crud-rest-api/helpers"
	"math"
	"net/http"
)

type response_struct struct {
	Emi             float64   `json:"emi"`
	Sum             float64   `json:"sum"`
	Total           float64   `json:"total"`
	Percent         float64   `json:"percent"`
	Principal_total float64   `json:"principal_total"`
	Interest_total  float64   `json:"interest_total"`
	Closing_bal     []float64 `json:"closing_bal"`
	Interest        []float64 `json:"interest"`
	Principal       []float64 `json:"principal"`
}

func powInt(x, y float64) float64 {
	return math.Pow(float64(x), float64(y))
}

func calculateemi(amount float64, tenure float64, interest_rate float64, month_year string) (map[string]float64, map[string][]float64) {
	tenure = tenure * 12

	var intrest float64 = (float64(interest_rate) / (100 * 12))
	var one float64 = 1

	var precal = powInt(one+intrest, tenure)

	if precal == 0 {
		precal = 1
	}
	var emi float64 = (float64(amount) * intrest * (precal / (precal - one)))
	var sum float64 = 0
	var interest_total float64 = 0
	var interest_total_old float64 = 0
	var principal_total float64 = 0
	var principal_total_old float64 = 0
	var calculatedInterest, calculatedPrincipal, calculatedBalance float64
	var interest []float64
	var principal []float64
	var closing_bal []float64
	var yearly_interest = make(map[int]float64)
	var yearly_principal = make(map[int]float64)
	var yearly_closing_bal = make(map[int]float64)

	for i := 0; i < int(tenure); i++ {

		calculatedInterest = roundFloat(amount*intrest, 2)

		calculatedPrincipal = roundFloat(emi-calculatedInterest, 2)

		amount = amount - calculatedPrincipal

		calculatedBalance = roundFloat(amount, 2)

		sum = (sum + calculatedInterest)

		interest_total += calculatedInterest
		principal_total += calculatedPrincipal

		interest = append(interest, calculatedInterest)
		principal = append(principal, calculatedPrincipal)
		closing_bal = append(closing_bal, calculatedBalance)

		if i%11 == 0 {
			yearly_closing_bal[i/11] = amount
			yearly_interest[i/11] = interest_total - interest_total_old
			yearly_principal[i/11] = principal_total - principal_total_old
			interest_total_old = interest_total
			principal_total_old = principal_total
		}

	}

	var total = (interest_total + principal_total)
	var percent = 100 - (sum/total)*100

	data := map[string]float64{"emi": roundFloat(emi, 2), "sum": roundFloat(sum, 2), "total": roundFloat(total, 2), "percent": roundFloat(percent, 2), "principal_total": roundFloat(principal_total, 2), "interest_total": roundFloat(interest_total, 2)}
	data2 := make(map[string][]float64)
	data2["closing_bal"] = closing_bal
	data2["interest"] = interest
	data2["principal"] = principal
	return data, data2
}

func EmiCalculator(w http.ResponseWriter, r *http.Request) {

	errors := make(map[string]string)

	amount := r.URL.Query().Get("amount")
	tenure := r.URL.Query().Get("tenure")
	interest := r.URL.Query().Get("interest")

	loan_amount := helpers.VarToFloat64(amount)
	tenure_number := helpers.VarToFloat64(tenure)
	interest_number := helpers.VarToFloat64(interest)

	errors["amount"] = helpers.SetRules(amount, "required|greater_than[10000]|less_than[100000000]")
	errors["tenure"] = helpers.SetRules(tenure, "required|is_alphanumeric")
	errors["interest"] = helpers.SetRules(interest, "Required|Is_alpha")

	errors = helpers.DeleteEmptyMapValues(errors)

	w.Header().Set("Content-Type", "application/json")

	if len(errors) != 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors)
		return
	}
	data1, data2 := calculateemi(loan_amount, tenure_number, interest_number, "month")

	response := new(response_struct)
	response.Emi = data1["emi"]
	response.Sum = data1["sum"]
	response.Total = data1["total"]
	response.Percent = data1["percent"]
	response.Principal_total = data1["principal_total"]
	response.Interest_total = data1["interest_total"]
	response.Closing_bal = data2["closing_bal"]
	response.Interest = data2["interest"]
	response.Principal = data2["principal"]

	json.NewEncoder(w).Encode(&response)

}

func roundFloat(n float64, decimals uint32) float64 {
	return math.Round(n*math.Pow(10, float64(decimals))) / math.Pow(10, float64(decimals))
}
