package model

type Tag struct {
	*Model
	Name string `json:"name"`
	Stat uint8 `json:"stat"`
}

func (t *Tag) TableName() string {
	return "blog_tag"
}
