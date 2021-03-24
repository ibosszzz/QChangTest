package controller

import (
	"net/http"
	"QChangTest/model"
	"encoding/json"
	"io/ioutil"
	"errors"
	"math"

)

var CashierDesk model.CashierDesk

func init() {
	CashierDesk.BankNote1000 = 10
	CashierDesk.BankNote500 = 20
	CashierDesk.BankNote100 = 15
	CashierDesk.BankNote50 = 20
	CashierDesk.BankNote20 = 30
	CashierDesk.Coin10 = 20
	CashierDesk.Coin5 = 20
	CashierDesk.Coin1 = 20
	CashierDesk.Coin025 = 50
}

func Cashier(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var res model.Response

	var customerPay model.CustomerPay
	// get data from body
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &customerPay)
	if err != nil {
		res.Error = err.Error()
		res.Data = nil
		res.Status = http.StatusBadRequest
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	cashierDesk, err := Calulate(customerPay.CustomerPay, customerPay.ProductPrice)
	if err != nil {
		res.Error = err.Error()
		res.Data = nil
		res.Status = http.StatusBadRequest
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	res.Error = nil
	res.Data = cashierDesk
	res.Status = http.StatusOK
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
	return
}

func CheckCashierDesk(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var res model.Response

	res.Error = nil
	res.Data = CashierDesk
	res.Status = http.StatusOK
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
	return
}

func AddCash(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var res model.Response

	var dataInput model.CashierDesk
	// get data from body
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &dataInput)
	if err != nil {
		res.Error = err.Error()
		res.Data = nil
		res.Status = http.StatusBadRequest
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	CashierDesk.BankNote1000 += dataInput.BankNote1000
	CashierDesk.BankNote500 += dataInput.BankNote500
	CashierDesk.BankNote100 += dataInput.BankNote100
	CashierDesk.BankNote50 += dataInput.BankNote50
	CashierDesk.BankNote20 += dataInput.BankNote20
	CashierDesk.Coin10 += dataInput.Coin10
	CashierDesk.Coin5 += dataInput.Coin5
	CashierDesk.Coin1 += dataInput.Coin1
	CashierDesk.Coin025 += dataInput.Coin025

	res.Error = nil
	res.Data = CashierDesk
	res.Status = http.StatusOK
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
	return
}

func Calulate(customerPay, productPrice float64) (model.CashierDesk, error) {
	var cashierDesk model.CashierDesk
	change := customerPay - productPrice
	if change < 0 {
		// return error customer pay less than product price
		return cashierDesk, errors.New("Please pay another cash!")
	}
	if change > 1000 && CashierDesk.BankNote1000 > 0 {
		change, cashierDesk.BankNote1000 = CheckBackNoteOrCoinEnough(change, CashierDesk.BankNote1000, 1000)
	}
	if change > 500 && CashierDesk.BankNote500 > 0 {
		change, cashierDesk.BankNote500 = CheckBackNoteOrCoinEnough(change, CashierDesk.BankNote500, 500)
	}
	if change > 100 && CashierDesk.BankNote100 > 0 {
		change, cashierDesk.BankNote100 = CheckBackNoteOrCoinEnough(change, CashierDesk.BankNote100, 100)
	}
	if change > 50 && CashierDesk.BankNote50 > 0 {
		change, cashierDesk.BankNote50 = CheckBackNoteOrCoinEnough(change, CashierDesk.BankNote50, 50)
	}
	if change > 20 && CashierDesk.BankNote20 > 0 {
		change, cashierDesk.BankNote20 = CheckBackNoteOrCoinEnough(change, CashierDesk.BankNote20, 20)
	}
	if change > 10 && CashierDesk.Coin10 > 0 {
		change, cashierDesk.Coin10 = CheckBackNoteOrCoinEnough(change, CashierDesk.Coin10, 10)
	}
	if change > 5 && CashierDesk.Coin5 > 0 {
		change, cashierDesk.Coin5 = CheckBackNoteOrCoinEnough(change, CashierDesk.Coin5, 5)
	}
	if change > 1 && CashierDesk.Coin1 > 0 {
		change, cashierDesk.Coin1 = CheckBackNoteOrCoinEnough(change, CashierDesk.Coin1, 1)
	}
	if change > 0.25 && CashierDesk.Coin025 > 0 {
		change, cashierDesk.Coin025 = CheckBackNoteOrCoinEnough(change, CashierDesk.Coin025, 0.25)
	}
	if change != 0 {
		// cash in cashier desk not enough for return to customer
		return cashierDesk, errors.New("Cash in cashier desk not enough!")
	}
	// remove cash in cashier desk
	CashierDesk.BankNote1000 -= cashierDesk.BankNote1000
	CashierDesk.BankNote500 -= cashierDesk.BankNote500
	CashierDesk.BankNote100 -= cashierDesk.BankNote100
	CashierDesk.BankNote50 -= cashierDesk.BankNote50
	CashierDesk.BankNote20 -= cashierDesk.BankNote20
	CashierDesk.Coin10 -= cashierDesk.Coin10
	CashierDesk.Coin5 -= cashierDesk.Coin5
	CashierDesk.Coin1 -= cashierDesk.Coin1
	CashierDesk.Coin025 -= cashierDesk.Coin025

	return cashierDesk, nil
}

func CheckBackNoteOrCoinEnough(change float64, amountBanknoteOrCoin int, banknoteOrCoinValue float64) (float64, int) {
	amount := int(change/banknoteOrCoinValue)
	change = math.Mod(change, banknoteOrCoinValue)
	if amount > amountBanknoteOrCoin {
		// banknote banknoteOrCoinValue not enough
		diff := amount - amountBanknoteOrCoin
		amount = amountBanknoteOrCoin
		change += float64(diff) * banknoteOrCoinValue
	}
	return change, amount
}