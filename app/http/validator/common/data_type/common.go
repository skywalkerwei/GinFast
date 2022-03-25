package data_type

type HeaderParams struct {
	Authorization string `header:"Authorization" binding:"required,min=20"`
}

type ID struct {
	Id float64 `form:"id" json:"id"  binding:"required,min=1"`
}

type Status struct {
	Status int `form:"status" json:"status"  binding:"oneof=0 1"`
}

type Page struct {
	Page  float64 `form:"page,default=1" json:"page" binding:"min=1"`    // 页面值>=1
	Limit float64 `form:"limit,default=10" json:"limit" binding:"min=1"` // 每页条数值>=10
	//Page  float64 `form:"page,default=1" json:"page" binding:"-"`   // 必填，页面值>=1
}

type Uid struct {
	Uid float64 `form:"uid" json:"uid"  binding:"required,min=1"`
}

type Mobile struct {
	Mobile string `form:"mobile" json:"mobile"  binding:"numeric,len=11"`
}

type CardNo struct {
	CardNo string `form:"card_no" json:"card_no"  binding:"required,min=1"`
}
