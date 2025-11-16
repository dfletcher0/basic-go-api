package tools

import "time"

// create struct object containing fake database data for retrieval
type mockDB struct{}

// create a map with string keys, LoginDetails values
var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123ABC",
		Username:  "alex",
	},
	"jason": {
		AuthToken: "456DEF",
		Username:  "jason",
	},
	"marie": {
		AuthToken: "789GHI",
		Username:  "marie",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    100,
		Username: "alex",
	},
	"jason": {
		Coins:    200,
		Username: "jason",
	},
	"marie": {
		Coins:    300,
		Username: "marie",
	},
}

// create methods for the struct which conform to database requirements in database.go
func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	// if key can't be found, return nil
	if !ok {
		return nil
	}

	// return pointer to client LoginDetails
	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	// if key can't be found, return nil
	if !ok {
		return nil
	}

	// return pointer to client LoginDetails
	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
