package controller

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"bytes"
	"QChangTest/model"
)

func TestCashier(t *testing.T) {
	// test cashier
	t.Run("test success", testCashier(2000, 1250, http.StatusOK))
	t.Run("test cash in cashier desk not enough", testCashier(1000000, 10, http.StatusBadRequest))
	t.Run("test customer pay less than product price", testCashier(100, 1000, http.StatusBadRequest))

	// test add cash
	var cashierDesk model.CashierDesk
	cashierDesk.BankNote1000 = 10
	cashierDesk.BankNote500 = 20
	cashierDesk.BankNote100 = 15
	cashierDesk.BankNote50 = 20
	cashierDesk.BankNote20 = 30
	cashierDesk.Coin10 = 20
	cashierDesk.Coin5 = 20
	cashierDesk.Coin1 = 20
	cashierDesk.Coin025 = 50

	t.Run("test add cash", testAddCash(cashierDesk, http.StatusOK))
}

func testCashier(customerPay, productPrice float64, expectCode int) func(*testing.T) {
	return func(t *testing.T) {
		reqBody := map[string]float64{"customer_pay": customerPay, "product_price": productPrice}
		body, _ := json.Marshal(reqBody)
		req := httptest.NewRequest("POST", "/cashier", bytes.NewReader(body))
		res := httptest.NewRecorder()

		Cashier(res, req)
	
		if res.Code != expectCode {
			t.Error(res.Body.String())
		} else {
			t.Log(res.Body.String())
		}
	}
}

func testAddCash(cashierDesk model.CashierDesk, expectCode int) func(*testing.T) {
	return func(t *testing.T) {
		body, _ := json.Marshal(cashierDesk)
		req := httptest.NewRequest("POST", "/addCash", bytes.NewReader(body))
		res := httptest.NewRecorder()

		AddCash(res, req)
	
		if res.Code != expectCode {
			t.Error(res.Body.String())
		} else {
			t.Log(res.Body.String())
		}
	}
}