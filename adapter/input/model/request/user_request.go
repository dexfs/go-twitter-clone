package request

type UserInfoRequest struct {
	Username string `uri:"username" binding:"required,min=5,max=10,alphanum,lowercase"`
}

type UserFeedRequest struct {
	Username string `uri:"username" binding:"required,min=5,max=10,alphanum,lowercase"`
}
