package service

import (
	"context"
	"log"

	"github.com/ropel12/scheduler/helper"
	"github.com/ropel12/scheduler/repository"

	"gorm.io/gorm"
)

type (
	service struct {
		db   *gorm.DB
		repo repository.SchoolRepo
	}
	Service interface {
		UpdateTestResult()
	}
)

func NewService(db *gorm.DB, repo repository.SchoolRepo) Service {
	return &service{db: db, repo: repo}
}

func (s *service) UpdateTestResult() {
	datas, err := s.repo.GetAllTestUrl(s.db)
	if err != nil {
		log.Printf("Err: %v", err)
	}
	if err == nil {
		for _, val := range datas {
			res := helper.ApiCall(val.QuizLinkResult)
			if res.Code == 200 {
				if len(res.Data) != 0 {
					for _, val2 := range res.Data {
						err := s.repo.UpdateTestResult(s.db.WithContext(context.Background()), val2.Email, val2.Result, int(val.ID))
						if err != nil {
							log.Printf("Err: %v", err)
						}
					}
				}
			}
		}
	}
	log.Println("[INFO]SUCCESS UPDATING TEST")
}
