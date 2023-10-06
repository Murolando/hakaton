package ent

type ClassRequest struct {
}

type ChildDashClassResponce struct {
	ClassProgress        []*ClassProgressResponce `json:"class_progress"`
	ExProgressBar        int                      `json:"ex_progress_bar"`
	MaxExProgressBar     int                      `json:"max_ex_progress_bar"`
	TheoryProgressBar    int                      `json:"theory_progress_bar"`
	MaxTheoryProgressBar int                      `json:"max_theory_progress_bar"`
}

type ChildMyClassResponce struct {
	Id            int64                  `json:"id"`
	Name          string                 `json:"class-name"`
	TeacherName   string                 `json:"teacher-name"`
	ChildCount    int64                  `json:"child-count"`
	ClassProgress *ClassProgressResponce `json:"class_progress"`
}

type ClassProgressResponce struct {
	Id             int64  `json:"id"`
	Name           string `json:"class-name"`
	ProgressBar    int    `json:"progress_now"`
	MaxProgressBar int    `json:"max_count"`
}
