package ent

type User struct {
	Name         *string `json:"name"`
	Login        *string `json:"login"`
	PasswordHash *string `json:"password"`
	RoleId       *int    `json:"role-id"`
}
type UserRequest struct {
	Name         *string `json:"name"`
	Login        *string `json:"login"`
	PasswordHash *string `json:"password"`
	RoleId       *int    `json:"role-id"`
}
type Session struct {
	RefreshToken string `json:"refres-token"`
	ExpiredAt    int64  `json:"expired-at"`
}

type Auth struct {
	Login        *string `json:"login"`
	PasswordHash *string `json:"password"`
}
