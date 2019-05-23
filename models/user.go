package models

import "time"

type User struct {
	Id         int       `json:"id" orm:"column(id);pk" description:"主键"`
	Loginname  string    `json:"login_name" orm:"column(login_name)" description:"登录名"`
	Username   string    `json:"user_name" orm:"column(user_name);null" description:"用户名"`
	Password   string    `json:"password" description:"密码"`
	Phone      string    `json:"phone" orm:"size(11);null" description:"手机号"`
	Avatar     string    `json:"avatar" orm:"null" description:"头像"`
	DeptId     int       `json:"dept_id" orm:"column(dept_id);null" description:"手机号"`
	Status     int       `json:"status" orm:"column(status);size(1);default(0)" description:"用户状态：0-正常，1-锁定"`
	DelFlag    int       `json:"del_flag" orm:"column(del_flag);size(1);default(0)" description:"删除标记：0-正常，1-删除"`
	CreateTime time.Time `json:"create_time" orm:"column(create_time);auto_now_add;type(datetime)"`
	updateTime time.Time `json:"update_time" orm:"column(update_time);auto_now;type(datetime)"`
}
