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
		UpdateTestResult(db *gorm.DB, email, status string, schoolid int) (int, error)
		GetUserDetailByEmail(db *gorm.DB, email string) (*entities.User, error)
		GetAllSchedules(db *gorm.DB) ([]entities.BillingSchedule, error)
		DeleteSchedule(db *gorm.DB, id int) error
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
		return nil, errors.New("Test Url Doesn't exist")
	}
	return res, nil
}
func (t *school) UpdateTestResult(db *gorm.DB, email, status string, schoolid int) (int, error) {
	progress := entities.Progress{}
	err := db.Transaction(func(db *gorm.DB) error {
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
		if err := db.Model(&progress).Where("school_id=? AND user_id=? AND status='Send Test Link'", schoolid, user.ID).Scan(&progress).Update("status", status).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return int(progress.ID), nil
}

func (t *school) GetUserDetailByEmail(db *gorm.DB, email string) (*entities.User, error) {
	res := entities.User{}
	if err := db.Table("users").Where("email=? AND is_verified=1 AND deleted_at IS NULL", email).Select("username,first_name,sure_name").Scan(&res).Error; err != nil {
		return nil, err
	}
	if res.FirstName == "" {
		return nil, errors.New("Data Not Found")
	}
	return &res, nil
}

func (t *school) GetAllSchedules(db *gorm.DB) ([]entities.BillingSchedule, error) {
	res := []entities.BillingSchedule{}
	if err := db.Where("date <= NOW() AND deleted_at IS NULL").Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
func (t *school) DeleteSchedule(db *gorm.DB, id int) error {
	if err := db.Where("id=?", id).Delete(&entities.BillingSchedule{}).Error; err != nil {
		return err
	}
	return nil
}
