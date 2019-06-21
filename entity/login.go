package entity

type LoginRegisterDto struct {
	Loginname string `form:"loginname" valid:"Required"`
	Password  string `form:"password" valid:"Required"`
}

type LoginDto struct {
	Loginname string `form:"loginname" valid:"Required"`
	Password  string `form:"password" valid:"Required"`
}
