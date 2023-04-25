package models

type Blog struct {
	ID      int    `json:"id" form:"id"`
	UserId  int    `json:"userid" form:"userid"`
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
}
