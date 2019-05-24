package entity

type UserAddDto struct {
	Loginname string `form:"loginname" valid:"Required"`
	Username  string `form:"username" valid:"Required"`
	Password  string `form:"password" valid:"Required"`
	Email     string `form:"email"`
	Phone     string `form:"phone"`
	Avatar    string `form:"avatar"`
	DeptId    int    `form:"deptid"`
	Status    int    `form:"status"`
	Sex       int    `form:"sex"`
}
