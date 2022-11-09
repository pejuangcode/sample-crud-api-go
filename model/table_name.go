package model

func (t *User) TableName() string {
	return "user"
}

func (t *Post) TableName() string {
	return "post"
}
