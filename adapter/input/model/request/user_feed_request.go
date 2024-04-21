package request

type UserInfoRequest struct {
	Username string `uri:"username" binding:"required,min=5,max=10,alpha,lowercase"`
}
