package structs

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Token string `json:"-"`
}

type Fpasswordstruct struct{
	Username string `json:"email"` 
}

type Newpassword struct {
	Password string `json:"password" binding:"required"`
	Repeatepassword string `json:"repeatepassword"`
}
