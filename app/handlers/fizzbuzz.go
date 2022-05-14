package handlers

import (
	"fizzbuzz-leboncoin/app/config"
	"fizzbuzz-leboncoin/app/helpers"
	"fizzbuzz-leboncoin/app/models"
	"fizzbuzz-leboncoin/app/sql"
	"strconv"
	"strings"

	"encoding/json"
	"net/http"
)

// Execute fizzbuzz exercise
func ExecFizzBuzz(int1, int2, limit int, str1, str2 string) string {
	var output []string

	for x := 1; x <= limit; x++ {
		switch {
		case x%(int1*int2) == 0:
			output = append(output, str1+str2)
		case x%int1 == 0:
			output = append(output, str1)
		case x%int2 == 0:
			output = append(output, str2)
		default:
			output = append(output, strconv.Itoa(x))
		}
	}

	return strings.Join(output, ",")
}

// @Summary Replace int1 and int2 by str1 and str2 for x in limit % int1 and int2
// @Param int1 query integer true "3"
// @Param int2 query integer true "5"
// @Param limit query integer true "100"
// @Param str1 query string true "fizz"
// @Param str2 query string true "buzz"
// @Produce  json
// @Success 200 {string} string
// @Failure 500 {object} models.APIError
// @Failure 400 {object} models.APIError
// @Router /fizzbuzz [get]
func fizzbuzz(c *config.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// Check if one input is empty
		input := helpers.IsInvalidInput(r.URL.Query(), "int1", "int2", "limit", "str1", "str2")
		if input != "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.APIError{Code: 400, Message: "invalid " + input})
			return
		}

		// Get all input is strings
		int1, int2, limit, str1, str2 := r.URL.Query().Get("int1"), r.URL.Query().Get("int2"), r.URL.Query().Get("limit"), r.URL.Query().Get("str1"), r.URL.Query().Get("str2")

		// Put all inputs in request statistics model
		var rs models.RequestStatistics
		var err1, err2, err3 error

		rs.Name = int1 + int2 + limit + str1 + str2

		rs.Int1, err1 = strconv.Atoi(int1)
		rs.Int2, err2 = strconv.Atoi(int2)
		rs.Limit, err3 = strconv.Atoi(limit)
		if err1 != nil || err2 != nil || err3 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.APIError{Code: 500, Message: "Converting from string to int failed in fizzbuzz."})
			return

		}

		rs.Str1 = str1
		rs.Str2 = str2

		// Call Insert or update sql on request statistics
		if err := sql.InsertOrIncrementRequestStatistics(c, &rs); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.APIError{Code: 500, Message: "Update request statitics failed : " + err.Error()})
			return
		}

		// Execute the fizzbuzz and return the formatted string
		json.NewEncoder(w).Encode(ExecFizzBuzz(rs.Int1, rs.Int2, rs.Limit, rs.Str1, rs.Str2))
	}
}
