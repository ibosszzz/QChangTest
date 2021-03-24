package controller

import (
	"net/http"
	"QChangTest/model"
	"encoding/json"
	"strconv"
)

func FindValue(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var res model.Response

	question := []string{"1", "X", "8", "17", "Y", "Z", "78", "113"}
	ans := []int{}
	diff := 0
	for index, value := range question {
		i, err := strconv.Atoi(value)
		if err != nil {
			i = ans[len(ans)-1] + diff + (index+1)
			ans = append(ans, i)
		} else {
			ans = append(ans, i)
		}
		if index != 0 {
			diff = i - ans[len(ans)-2]
		}
	}
	dataReturn := make(map[string]interface{})
	dataReturn["question"] = question
	dataReturn["answers"] = ans

	res.Error = nil
	res.Data = dataReturn
	res.Status = http.StatusOK
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
	return
}