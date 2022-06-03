package model

type Post struct {
	ID      int64
	Slug    string
	Title   string
	Content string
	Author  *Author
}

type Author struct {
	ID           int64
	Name         string
	ProfileImage string
}
