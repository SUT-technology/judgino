package service

type Service struct{
	AuthSrvc AuthService
	QuestionsSrvc QuestionsService
	UserSrvc UserService
	SubmissionSrvc SubmissionService
}