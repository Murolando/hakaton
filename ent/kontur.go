package ent

type KonturResponse struct {
	Id       int64     `json:"id"`
	Name     [4]string `json:"name"`
	ImageSrc string    `json:"image_src"`
	Correct  string    `json:"correct_name"`
}

type ProcessRequest struct {
	Answers  []bool `json:"answer"`
	LessonId int64  `json:"lesson_id"`
}
type ProcessResponse struct {
	Grade int64 `json:"grade"`
}
