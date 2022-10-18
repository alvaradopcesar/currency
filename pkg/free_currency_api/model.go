package free_currency_api

type Config struct {
	FreeCurrencyTime   int
	FreeCurrencyApiKey string
}

type Data struct {
	Data map[string]float64 `json:"data"`
}

type QueryInfo struct {
	Method  string
	Address string
	Code    int
	Time    float64
}
