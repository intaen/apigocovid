package usecase

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/intaen/apigocovid/domain"
	"github.com/intaen/apigocovid/pkg/utils"
)

type chartUsecase struct {
	chartRepo domain.ChartRepository
}

// NewChartUsecase will create new an covidUsecase object representation of domain.ChartUsecase interface
func NewChartUsecase(chartRepo domain.ChartRepository) domain.ChartUsecase {
	return &chartUsecase{chartRepo: chartRepo}
}

func (ch *chartUsecase) CreateBarChart(items []int, values []string, title, subtitle, seriesname string) *charts.Bar {
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

func (ch *chartUsecase) CreateLineChart(items []int, values []string, title, subtitle, seriesname string) *charts.Line {
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
