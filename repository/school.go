package repository

import (
	"errors"

	"github.com/ropel12/scheduler/entities"

	"gorm.io/gorm"
)

type (
	udata struct {
		ID uint
	}
	school     struct{}
	SchoolRepo interface {
		GetAllTestUrl(db *gorm.DB) ([]entities.Testlinks, error)
		UpdateTestResult(db *gorm.DB, email, status string, schoolid int) error
	}
)

func NewRepo() SchoolRepo {
	return &school{}
}
func (t *school) GetAllTestUrl(db *gorm.DB) ([]entities.Testlinks, error) {
	res := []entities.Testlinks{}
	if err := db.Table("schools").Where("quiz_link_result !=''").Scan(&res).Error; err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, errors.New("Url Test Doesn't exist")
	}
	return res, nil
}
func (t *school) UpdateTestResult(db *gorm.DB, email, status string, schoolid int) error {
	return db.Transaction(func(db *gorm.DB) error {
		var user = udata{}
		if err := db.Table("users").Where("email=? AND is_verified=1 AND deleted_at IS NULL", email).Select("id").Scan(&user).Error; err != nil {
			return err
		}
		if user.ID == 0 {
			return errors.New("Data Not Found")
		}
		if status == "Fail" {
			status = "Failed Test Result"
		} else {
			status = "Test Result"
		}
		if err := db.Model(&entities.Progress{}).Where("school_id=? AND user_id=? AND status='Send Test Link'", schoolid, user.ID).Update("status", status).Error; err != nil {
			return err
		}
		return nil
	})
}
