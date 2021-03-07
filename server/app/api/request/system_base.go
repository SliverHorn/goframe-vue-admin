package request

type AdminLogin struct {
	Username  string `p:"username" v:"required|length:1,30#请输入用户名称|您输入用户名称长度非法"`
	Password  string `p:"password" v:"required|length:6,30#请输入密码|密码长度为:min到:max位"`
	Captcha   string `json:"captcha" valid:"required#请输入正确的验证码"`
	CaptchaId string `json:"captchaId" valid:"required|length:20,20#请输入captchaId|您输入captchaId长度非法"`
}

type InitDB struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	UserName string `json:"userName" binding:"required"`
	Password string `json:"password"`
	DBName   string `json:"dbName" binding:"required"`
}
