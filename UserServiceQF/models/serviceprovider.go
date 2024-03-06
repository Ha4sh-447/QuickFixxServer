package models

type ServiceProd struct {
	Id            int      `json:"id"`
	Address       string   `json:"address"`
	Contact       string   `json:"contact"`
	Name          string   `json:"name"`
	Experience    string   `json:"experience"`
	Qualification []string `json:"qualification"`
	Location      string   `json:"location"`
	Rating        int      `json:"rating"`
}
