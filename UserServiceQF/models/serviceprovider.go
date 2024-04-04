package models

type ServiceProd struct {
	Address        string  `json:"address"`
	Specialization string  `json:"specialization"`
	Eid            int     `json:"eid"`
	Specz          string  `json:"specz"`
	Rating         float32 `json:"rating"`
	Experience     string  `json:"experience"`
	U_id           int     `json:"u_id"`
	Shopname       string  `json:"shopname"`
	Contact        string  `json:"contact"`
}
