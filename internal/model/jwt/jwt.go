package jwt

type LoginReq struct {
	UserName string `json:"user_name,omitempty"`
	Password string `json:"password,omitempty"`
}

type User struct {
	Id        int64  `json:"id"`
	Name      string `json:"name,omitempty"`
	TimeStamp int64  `json:"time_stamp"`
	Password  string `json:"password,omitempty"`
}

type LoginResult struct {
	Id    int64  `json:"id"`
	Name  string `json:"name,omitempty"`
	Token string `json:"token,omitempty"`
}

type RegisterReq struct {
	Id       int64  `json:"id"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password"`
	Phone    string `json:"phone,omitempty"`
}

type UserCreate struct {
	Id       int64  `json:"id,omitempty"`
	UserName string `json:"user_name,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
