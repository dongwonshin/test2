package model

type (
	user2 struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

var (
	users = map[int]*user2{}
	seq   = 1
)
