package ent

type ClassRequest struct {
}
type ClassResponce struct {
	Id          int64  `json:"id"`
	Name        string `json:"class-name"`
	TeacherName string `json:"teacher-name"`
}
