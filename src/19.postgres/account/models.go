package account

// User schema of the user table
type Account struct {
	ID       string `json:"id"`
	Owner    string `json:"owner"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
}
