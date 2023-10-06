package ent

type LessonRating struct {
}
type LessonInfoResponce struct {
	Name           *string `json:"name"`
	Video          *string `json:"video"`
	LessonTypeName string  `json:"lesson_type_name"`
	CreatedAt      string  `json:"created_at"`
	ExpiredAt      string  `json:"expired_at"`
	LessonAccess   bool    `json:"access"`
}
