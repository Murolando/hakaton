package ent

type User struct {
	Name         *string `json:"name"`
	Login        *string `json:"login"`
	PasswordHash *string `json:"password"`
	RoleId       *int    `json:"role_id"`
}
type UserRequest struct {
	Name         *string `json:"name"`
	Login        *string `json:"login"`
	PasswordHash *string `json:"password"`
	RoleId       *int    `json:"role_id"`
}
type Session struct {
	RefreshToken string `json:"refres_token"`
	ExpiredAt    int64  `json:"expired_at"`
}

type Auth struct {
	Login        *string `json:"login"`
	PasswordHash *string `json:"password"`
}
