package dto

type LoginDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterDTO struct {
	Username   string `json:"username" binding:"required,min=3"`
	Age        int    `json:"age" binding:"required,min=1,max=100"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"` // 跨字段验证
	//CheckIn  time.Time `form:"check_in" json:"check_in" binding:"required,bookabledate" time_format:"2006-01-02" label:"输入时间"`
	//CheckOut time.Time `form:"check_out" json:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02" label:"输出时间"`
}
