package model

type Building struct {
	ID        int    `json:"id" pg:"id,pk"`
	Name      string `json:"name" binding:"required"`
	City      string `json:"city" binding:"required"`
	YearBuilt int    `json:"year_built" binding:"required"`
	Floors    int    `json:"floors" binding:"required"`
}
