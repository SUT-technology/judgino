package dto

type ProfileResponeDTO struct {
	UserId    uint
	CurrentUserId uint
	Username string
	Phone string
	Email string
	Role string
	NotAccepted int64
	Accepted int64
	Total int64
	SolvedPercentage int
	IsCurrentUserAdmin bool
	err error
}