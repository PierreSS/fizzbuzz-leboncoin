package models

// RequestStatistics define the request_statistics table
type RequestStatistics struct {
	ID        int    `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	HitNumber int64  `db:"hit_number" json:"hit_number"`
	Int1      int    `db:"int1" json:"int1"`
	Int2      int    `db:"int2" json:"int2"`
	Limit     int    `db:"limit" json:"limit"`
	Str1      string `db:"str1" json:"str1"`
	Str2      string `db:"str2" json:"str2"`
}
