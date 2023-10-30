package domain

type AddUserInput struct {
	CreateBy string
	UpdateBy string
	Username string
	Password string
	NickName string
	Phone    string
	RoleId   string
	Salt     string
	Avatar   string
	Sex      string
	Email    string
	Remark   string
}

type AddUserOutput struct {
	UserId string
}

type DeleteUserInput struct {
	UserId string
}

type DeleteUserOutput struct {
}
