package microauth

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserCallback func(user User) (int, []byte, error)
