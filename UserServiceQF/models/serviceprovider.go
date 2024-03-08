package models

type ServiceProd struct {
	Id            int     `json:"id"`
	Address       string  `json:"address"`
	Contact       string  `json:"contact"`
	Name          string  `json:"name"`
	Experience    string  `json:"experience"`
	Qualification []uint8 `json:"qualification"`
	// Qualification  interface{} `json:"qualification"`
	Location       string `json:"location"`
	Rating         int    `json:"rating"`
	Specialization string `json:"specialization"`
}
