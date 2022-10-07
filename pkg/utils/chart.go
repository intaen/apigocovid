package utils

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"github.com/intaen/apigocovid/domain"
)

func GetTotalBarItems(list []int) []opts.BarData {
	items := make([]opts.BarData, 0)
	for _, v := range list {
		items = append(items, opts.BarData{Value: v})
	}
	return items
}

func GenerateBarChart(chartz domain.Chart) *charts.Bar {
	// create a new bar instance
	bar := charts.NewBar()
	// set some global options like Title/Legend/ToolTip or anything else
	bar.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeInfographic}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true}),
		charts.WithTitleOpts(opts.Title{
			Title:    chartz.Title,
			Subtitle: chartz.Subtitle}),
		charts.WithLegendOpts(opts.Legend{Right: "80%"}))

	// Put data into instance
	bar.SetXAxis(chartz.AxisX).
		AddSeries(chartz.Series.Name, chartz.Series.DataBar)

	return bar
}

func GetTotalLineItems(list []int) []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < len(list); i++ {
		items = append(items, opts.LineData{Value: list[i]})
	}
	return items
}

func GenerateLineChart(chartz domain.Chart) *charts.Line {
	// create a new line instance
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeChalk}),
		charts.WithTitleOpts(opts.Title{
			Title:    chartz.Title,
			Subtitle: chartz.Subtitle,
		}))

	// Put data into instance
	line.SetXAxis(chartz.AxisX).
		AddSeries(chartz.Series.Name, chartz.Series.DataLine)

	return line
}
