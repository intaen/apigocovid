package domain

import "github.com/go-echarts/go-echarts/v2/opts"

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
