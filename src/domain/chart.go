package domain

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

// Usecase represent the chart's usecases
type ChartUsecase interface {
	CreateBarChart([]int, []string, string, string, string) *charts.Bar
	CreateLineChart([]int, []string, string, string, string) *charts.Line
}

// Repository represent the chart's repository contract
type ChartRepository interface{}

type (
	Chart struct {
		Title    string
		Subtitle string
		AxisX    []string
		Series   ChartSeries
	}

	ChartSeries struct {
		Name     string
		DataBar  []opts.BarData
		DataLine []opts.LineData
	}
)
