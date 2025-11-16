package tools

import (
	log "github.com/sirupsen/logrus"
)

// Database objects
type LoginDetails struct {
	AuthToken string
	Username  string
}
type CoinDetails struct {
	Coins    int64
	Username string
}

// create DatabaseInterface to make it easy to switch databases in future
type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetUserCoins(username string) *CoinDetails
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {

	// set up mock database
	var database DatabaseInterface = &mockDB{}

	var err error = database.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// return memory pointer to db
	return &database, nil
}
