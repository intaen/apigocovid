package repository

import (
	"apigocovid/domain"

	"gorm.io/gorm"
)

type chartRepo struct {
	rd domain.ClientRedis
	db *gorm.DB
}

// NewChartRepository will create an implementation of chart.Repository
func NewChartRepository(rd domain.ClientRedis, db *gorm.DB) domain.ChartRepository {
	return &chartRepo{rd: rd, db: db}
}
