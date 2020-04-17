package controllers

type LoginController struct {
}

type Wxlogin struct {
	Code      string `json:"code"`
	NickName  string `json:"nickName"`
	AvatarUrl string `json:"avatarUrl"`
	Gender    int    `json:"gender"`
}
