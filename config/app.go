package config

var (
	Email      string
	Shop       string
	UnlockCode string
	Shops      = []string{"Bet9ja", "MerryBet", "NairaBet", "BetKing", "1xBet", "SureBet247"}
	Apibase    = "http://localhost:8888"
	Endpoints  = map[string]string{
		"activation":   Apibase + "/activation",
		"hosts-upload": Apibase + "/hosts-upload",
		"update":       Apibase + "/update",
	}
)
