package service

type Service struct{
	AuthSrvc AuthService
	PrflSrvc ProfileService
	QuestionsSrvc QuestionsService
	UserSrvc UserService
	SubmissionSrvc SubmissionService
}