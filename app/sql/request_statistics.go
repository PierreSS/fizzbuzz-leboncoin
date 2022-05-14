package sql

import (
	"fizzbuzz-leboncoin/app/config"
	"fizzbuzz-leboncoin/app/models"
)

// Upset on request statistics to increase hit_number
func InsertOrIncrementRequestStatistics(c *config.Client, rs *models.RequestStatistics) error {
	_, err := c.DB.NamedQuery(`INSERT INTO request_statistics (name, int1, int2, "limit", str1, str2) 
	VALUES (:name, :int1, :int2, :limit, :str1, :str2)
	ON CONFLICT (name) 
	DO UPDATE SET hit_number = request_statistics.hit_number + 1`, rs)
	if err != nil {
		return err
	}

	return nil
}

// Select max hit_number of request statistics
func GetMaxHitNumberInRequestStatistics(c *config.Client) (models.RequestStatistics, error) {
	var rs models.RequestStatistics

	if err := c.DB.Get(&rs, "SELECT * FROM request_statistics ORDER BY hit_number DESC LIMIT 1"); err != nil {
		return models.RequestStatistics{}, err
	}

	return rs, nil
}
