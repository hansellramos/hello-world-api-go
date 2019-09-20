package users

type User struct {
	Id		int		`json:"id"`
	Name	string	`json:"name,omitempty"`
	Email	string  `json:"email,omitempty"`
	Age		byte	`json:"age,omitempty"`
}
