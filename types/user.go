package types

type UpdateUser struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Birthday  string `json:"birthday"`
	Address   string `json:"address"`
	Nik       string `json:"nik"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
