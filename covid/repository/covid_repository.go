package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"apigocovid/domain"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type covidRepo struct {
	rd domain.ClientRedis
	db *gorm.DB
}

// NewCovidRepository will create an implementation of covid.Repository
func NewCovidRepository(rd domain.ClientRedis, db *gorm.DB) domain.CovidRepository {
	return &covidRepo{rd: rd, db: db}
}

// ---- API

func (cv *covidRepo) GetAllData() (*domain.CoronaVirus, error) {
	// Consume third API
	client := resty.New()       // Create client
	var resp domain.CoronaVirus // Initialize new variable to catch response from 3rd party
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&resp).
		Get(viper.GetString("url.covid"))
	log.Println(fmt.Sprintf("Request Method: %s, URI: %s, Response: %s", res.Request.Method, res.Request.URL, "panjang")) //string(res.Body())))
	if err != nil {
		fmt.Println("ERR", err)
		return nil, err
	}

	if res.IsError() {
		return nil, errors.New("data not found")
	} else if res.IsSuccess() {
		return &resp, nil
	}

	return nil, nil
}

// ---- Database

func (cv *covidRepo) AddData(data []domain.CoronaVirusStatistic) error {
	err := cv.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "combined_key"}},
		DoUpdates: clause.AssignmentColumns([]string{"confirmed", "death", "recovered", "date_updated"}),
	}).Create(&data).Error
	if err != nil {
		return err
	}

	return nil
}

func (cv *covidRepo) FindAllData() ([]domain.CoronaVirusStatistic, error) {
	var data []domain.CoronaVirusStatistic

	result, err := cv.rd.GetRedisValue("find_covids")
	if err != nil {
		fmt.Println("Database")
		err := cv.db.Find(&data).Error
		if err != nil {
			return nil, err
		}
		cv.rd.SetRedisValue("find_covids", data, 300)
		return data, nil
	}

	fmt.Println("Redis")
	err = json.Unmarshal(result, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (cv *covidRepo) FindDataByKey(key string) ([]domain.CoronaVirusStatistic, error) {
	var data []domain.CoronaVirusStatistic

	result, err := cv.rd.GetRedisValue("find_covids_by_" + key)
	if err != nil {
		fmt.Println("Database")
		err := cv.db.Where("combined_key ILIKE $1", "%"+key+"%").Find(&data).Error
		if err != nil {
			return nil, err
		}
		cv.rd.SetRedisValue("find_covids_by_"+key, data, 300)
		return data, nil
	}

	fmt.Println("Redis")
	err = json.Unmarshal(result, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
