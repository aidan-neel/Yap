package schemas

import "server/models"

type CreatePostReq struct {
	Content string `json:"content" validate:"required"`
	UserId string `json:"user_id" validate:"required"`
}

type PostResponse struct {
	models.Post
	Upvoted   bool `json:"upvoted"`
	Downvoted bool `json:"downvoted"`
}

type CreateVoteReq struct {
	PostId int    `json:"post_id" validate:"required"`
	UserId string `json:"user_id" validate:"required"`
	IsUp   *bool  `json:"is_up" validate:"required"`
}

type DeleteVoteReq struct {
	PostId int    `json:"post_id" validate:"required"`
	UserId string `json:"user_id" validate:"required"`
}