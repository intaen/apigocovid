package usecase

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/intaen/apigocovid/domain"
	"github.com/intaen/apigocovid/pkg/utils"
)

type covidUsecase struct {
	covidRepo domain.CovidRepository
}

// NewCovidUsecase will create new an covidUsecase object representation of domain.CovidUsecase interface
func NewCovidUsecase(covidRepo domain.CovidRepository) domain.CovidUsecase {
	return &covidUsecase{covidRepo: covidRepo}
}

// ---- Scheduler

func (cv *covidUsecase) AddMasterStatistic() {
	listData, err := cv.covidRepo.GetAllData()
	if err != nil {
		return
	}

	var listResult []domain.CoronaVirusStatistic
	for _, v := range listData.RawData {
		act, _ := strconv.Atoi(v.Active)
		conf, _ := strconv.Atoi(v.Confirmed)
		deaths, _ := strconv.Atoi(v.Deaths)
		rec, _ := strconv.Atoi(v.Recovered)
		// Assign if province state is empty
		if v.ProvinceState == "" {
			v.ProvinceState = v.CountryRegion
		}

		result := domain.CoronaVirusStatistic{
			Country:       v.CountryRegion,
			ProvinceState: v.ProvinceState,
			CombinedKey:   v.Combined_Key,
			Active:        act,
			Confirmed:     conf,
			Deaths:        deaths,
			Recovered:     rec,
			DataSource:    listData.DataSource.URL,
			PublishedBy:   listData.DataSource.PublishedBy,
			DateUpdated:   v.LastUpdate,
		}
		listResult = append(listResult, result)
	}
	if err != nil {
		log.Println(fmt.Sprintf("GetAllData is Error: %v", err.Error()))
		return
	}

	if len(listResult) == 0 {
		log.Println(fmt.Sprintf("GetAllData is empty, TotalData: %v", len(listResult)))
		return
	}

	// Insert
	err = cv.covidRepo.AddData(listResult)
	if err != nil {
		log.Println(fmt.Sprintf("AddData is Error: %v", err.Error()))
		return
	}
}

// ---- API

func (cv *covidUsecase) GetAllData() (*domain.Result, error) {
	listData, err := cv.covidRepo.FindAllData()
	if err != nil {
		return nil, err
	}

	var result domain.Result
	result.Count = len(listData)
	for _, v := range listData {

		// For json
		detailResult := domain.DetailResult{
			Country:     v.Country,
			CombinedKey: v.CombinedKey,
			Active:      v.Active,
			Confirmed:   v.Confirmed,
			Deaths:      v.Deaths,
			Recovered:   v.Recovered,
		}
		result.Detail = append(result.Detail, detailResult)
	}

	// Sort by country name
	sort.SliceStable(result.Detail, func(i, j int) bool {
		return result.Detail[i].Country < result.Detail[j].Country
	})

	return &result, nil
}

// ---- Data

func (cv *covidUsecase) CreateBarChart(items []int, values []string, title, subtitle, seriesname string) *charts.Bar {
	// Get total data
	totalItem := utils.GetTotalBarItems(items)

	// Assign it to struct of chart
	var chart domain.Chart
	chart.Title = title
	chart.Subtitle = subtitle
	chart.AxisX = values
	chart.Series.Name = seriesname
	chart.Series.DataBar = totalItem

	return utils.GenerateBarChart(chart)
}

func (cv *covidUsecase) CreateLineChart(items []int, values []string, title, subtitle, seriesname string) *charts.Line {
	// Get total data
	totalItem := utils.GetTotalLineItems(items)

	// Assign it to struct of chart
	var chart domain.Chart
	chart.Title = title
	chart.Subtitle = subtitle
	chart.AxisX = values
	chart.Series.Name = seriesname
	chart.Series.DataLine = totalItem

	return utils.GenerateLineChart(chart)
}
