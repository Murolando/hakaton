package ent

type ClassRequest struct {
}
type ChildDashClassResponce struct {
	Id   int64  `json:"id"`
	Name string `json:"class-name"`
	// TeacherName string `json:"teacher-name"`
}
type ChildMyClassResponce struct {
	Id          int64  `json:"id"`
	Name        string `json:"class-name"`
	TeacherName string `json:"teacher-name"`
	ChildCount  int64  `json:"child-count"`
}
