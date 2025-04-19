package model

type UserMessage string

const (
	OKMessage           UserMessage = "با موفقیت انجام شد"
	BadRequestMessage   UserMessage = "درخواست ورودی نامعتبر است"
	UnauthorizedMessage UserMessage = "وارد حساب خود شوید"
	NotFoundMessage     UserMessage = "پیدا نشد"
	InternalMessage     UserMessage = "خطا هنگام ارسال اطلاعات"
	TimeoutMessage      UserMessage = "مشکلی در ارتباط رخ داد"
	TooManyMessage      UserMessage = "تعداد درخواست های شما بیش از حد مجاز است. لطفا بعدا دوباره تلاش کنید"
	ForbiddenMessage    UserMessage = "شما مجوز دسترسی به این بخش را ندارید"
)

type Response struct {
	Msg          UserMessage `json:"userMsg"`
	Data         any         `json:"data"`
	TrackingCode string      `json:"trackingCode,omitempty"`
}
