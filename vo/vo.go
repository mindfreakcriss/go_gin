package vo

type UserReq struct {
	Name     string
	Password string
	NickName string
}

type UserVo struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
	NickName string `json:"nickname,omitempty"`
}
