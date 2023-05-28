package entities

import "gorm.io/gorm"

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
		Username  string
		FirstName string
		SureName  string
	}
	BillingSchedule struct {
		ID           uint `gorm:"primaryKey;not null;autoIncrement"`
		StudentName  string
		StudentEmail string
		SchoolName   string
		DeletedAt    gorm.DeletedAt `gorm:"index"`
		Date         string         `gorm:"type:timestamp;not null"`
		Total        int
	}
	HttpResponse struct {
		Code    int          `json:"code"`
		Message string       `json:"message"`
		Data    []TestResult `json:"data"`
	}
)
