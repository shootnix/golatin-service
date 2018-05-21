package models

import (
	"log"
	"time"
)

type ApiUser struct {
	ID              int64     `db:"id"`
	Description     string    `db:"description"`
	ApiKey          string    `db:"api_key"`
	TotalRequests   int64     `db:"total_requests"`
	RequestsAllowed int64     `db:"requests_allowed"`
	RequestsCount   int64     `db:"requests_count"`
	UpdatedAt       time.Time `db:"updated_at"`
	Disabled        int64     `db:"disabled"`
	TypeID          int64     `db:"type_id"`
	Type            ApiUserType
}

type ApiUserType struct {
	ID   int64  `db:"id"`
	Name string `db:"string"`
}

func FindApiUser(apiKey string) (*ApiUser, error) {
	sql := `
		SELECT *
		  FROM "api_users"
		 WHERE "api_key" = $1
	`
	apiUser := &ApiUser{}
	if err := Pg.QueryRowx(sql, apiKey).StructScan(apiUser); err != nil {
		log.Println("Find api_user failed: ", err.Error())
		return apiUser, err
	}

	return apiUser, nil
}

// FIXME
func (u *ApiUser) CheckDisabled() bool {
	return false
}

// FIXME
func (u *ApiUser) CheckAllowedRequests() int64 {
	return 1
}

// FIXME
func (u *ApiUser) RegisterRequest() error {
	return nil
}
