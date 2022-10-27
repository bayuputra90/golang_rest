package models

type Students struct {
	Id           int    `form:"id" json:"id"`
	NamaDepan    string `form:"nama_depan" json:"nama_depan"`
	NamaBelakang string `form:"nama_belakang" json:"nama_belakang"`
	NoHp         string `form:"no_hp" json:"no_hp"`
	Gender       string `form:"gender" json:"gender"`
	Jenjang      string `form:"jenjang" json:"jenjang"`
	Hobi         string `form:"hobi" json:"hobi"`
	Alamat       string `form:"alamat" json:"alamat"`
}

type AllResponse struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Data []Students

}
type Response struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Data Students
}

type Responses struct {
	Status int `json:"status"`
	Message string `json:"message"`
}
