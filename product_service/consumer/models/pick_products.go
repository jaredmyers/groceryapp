package models

type Products struct {
	ID        int    `json:"id"`
	Name      string `json:"item"`
	Price     string `json:"price"`
	Nutrition []Nutrition
	Store     string `json:"store"`
	Location  string `json:"location"`
}

type Nutrition struct {
	Calories        int     `json:"calories"`
	Servings        float64 `json:"servings"`
	Servingsize     float64 `json:"servingsize"`
	ServingSizeUnit string  `json:"servingsizeunit"`
	Fat             float64 `json:"fat"`
	Cholestrol      int     `json:"cholestrol"`
	Sodium          int     `json:"sodium"`
	TotalCarb       int     `json:"totalcarb"`
	Fiber           int     `json:"fiber"`
	Sugars          int     `json:"sugars"`
	Protein         int     `json:"protein"`
}

type GroceryListPage struct {
	ID   int    `json:"id"`
	Item string `json:"item"`
}
