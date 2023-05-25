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
		Name           string
	}

	TestResult struct {
		Email  string `json:"email"`
		Result string `json:"result"`
	}
	School struct {
		Name string
	}
	User struct {
		FirstName string
		SureName  string
	}

	HttpResponse struct {
		Code    int          `json:"code"`
		Message string       `json:"message"`
		Data    []TestResult `json:"data"`
	}
)
