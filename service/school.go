package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/ropel12/scheduler/helper"
	"github.com/ropel12/scheduler/pkg"
	"github.com/ropel12/scheduler/repository"

	"gorm.io/gorm"
)

type (
	service struct {
		db        *gorm.DB
		repo      repository.SchoolRepo
		nsq       *pkg.NSQProducer
		emailmaps map[string]int
	}
	Service interface {
		UpdateTestResult()
	}
)

func NewService(db *gorm.DB, repo repository.SchoolRepo, nsq *pkg.NSQProducer) Service {
	return &service{db: db, repo: repo, nsq: nsq, emailmaps: make(map[string]int)}
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
					mapss := make(map[string][]string)
					for _, val2 := range res.Data {
						mapss[val2.Email] = append(mapss[val2.Email], val2.Result)
					}
					for email, val3 := range mapss {
						leng := len(val3) - 1
						err := s.repo.UpdateTestResult(s.db.WithContext(context.Background()), email, val3[leng], int(val.ID))
						if err != nil {
							log.Printf("Err: %v", err)
						}
						if val3[leng] == "Fail" {
							user, err := s.repo.GetUserDetailByEmail(s.db.WithContext(context.Background()), email)
							if err != nil {
								log.Printf("Error: %v", err)
							} else {
								if s.emailmaps[fmt.Sprintf("%s%d", email, val.ID)] == 0 || s.emailmaps[fmt.Sprintf("%s%d", email, val.ID)] != leng {
									encodeddata, _ := json.Marshal(map[string]any{"email": email, "name": user.FirstName + " " + user.SureName, "school": val.Name, "reason": "Anda tidak berhasil dalam tes tersebut."})
									go func() {
										if err := s.nsq.Publish("1", encodeddata); err != nil {
											log.Printf("Error: %v", err)
										}
									}()
									s.emailmaps[fmt.Sprintf("%s%d", email, val.ID)] = leng
								}
							}
						}

					}
				}
			}
		}
	}
	log.Println("[INFO]SUCCESS UPDATING TEST")
}
