package dto

type CreateUserRequest struct {
	UserName string `json:"user_name" example:"ash"`
	Password string `json:"password"  example:"adsflivqlwbliwef#@#$sdlfkj"`
}

type UpdateUserRequest struct {
	UserName       string `json:"user_name"`
	MembershipType string `json:"membership_type"`
}
