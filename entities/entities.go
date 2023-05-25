package entities

type (
	Progress struct {
		ID       uint `gorm:"primaryKey;autoIncrement;not null"`
		UserID   uint
		SchoolID uint
		Status   string
	}
	Testlinks struct {
		ID             uint
		QuizLinkResult string
	}
	TestResult struct {
		Email  string `json:"email"`
		Result string `json:"result"`
	}
	HttpResponse struct {
		Code    int          `json:"code"`
		Message string       `json:"message"`
		Data    []TestResult `json:"data"`
	}
)
