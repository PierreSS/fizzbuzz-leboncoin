package handlers

import (
	"encoding/json"
	"fizzbuzz-leboncoin/app/config"
	"fizzbuzz-leboncoin/app/models"
	"fizzbuzz-leboncoin/app/sql"
	"net/http"
)

// RequestStatisticsOutput to make it conform to the given instructions
type RequestStatisticsOutput struct {
	Int1      int    `db:"int1" json:"int1"`
	Int2      int    `db:"int2" json:"int2"`
	Limit     int    `db:"limit" json:"limit"`
	Str1      string `db:"str1" json:"str1"`
	Str2      string `db:"str2" json:"str2"`
	HitNumber int64  `db:"hit_number" json:"hit_number"`
}

// @Summary Get the name of the max hit number of requests for a fizzbuzz
// @Success 200 {object} RequestStatisticsOutput
// @Failure 500 {object} models.APIError
// @Router /statistics [get]
func requestNameForMaxHitNumber(c *config.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		rs, err := sql.GetMaxHitNumberInRequestStatistics(c)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.APIError{Code: 500, Message: "Error getting max hit number in request statistics : " + err.Error()})
			return
		}

		rso := RequestStatisticsOutput{
			Int1:      rs.Int1,
			Int2:      rs.Int2,
			Limit:     rs.Limit,
			Str1:      rs.Str1,
			Str2:      rs.Str2,
			HitNumber: rs.HitNumber,
		}

		json.NewEncoder(w).Encode(rso)
	}
}
