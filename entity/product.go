package entity

type Product struct {
	Id          int    `sql:",pk" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageLink   string `json:"imageLink"`
	Price       string `json:"price"`
	Rating      string `json:"rating"`
	Merchant    string `json:"merchant"`

	// nolint:structcheck,unused
	tableName struct{} `sql:"products,alias:p,discard_unknown_columns"`
}
