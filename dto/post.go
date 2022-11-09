package dto

type NewPost struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	UserId      int    `json:"user_id"`
}

type Updatepost struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
