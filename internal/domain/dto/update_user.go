package dto

type UpdateUserDTO struct {
	ID                    uint   `json:"ID"`
	FirstName             string `json:"FirstName"`
	Email                 string `json:"Email"`
	Phone                 string `json:"Phone"`
	Username              string `json:"Username"`
	Password              string `json:"Password"`
	Role                  string `json:"Role"`
	CreatedQuestionsCount int64  `json:"Created"`
	SolvedQuestionsCount  int64  `json:"Solved"`
	SubmissionsCount      int64  `json:"Submissions"`
}
