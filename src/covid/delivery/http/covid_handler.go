package http

import (
	"time"

	"apigocovid/src/domain"

	"github.com/gin-gonic/gin"
)

type CovidHandler struct {
	chartUsecase domain.ChartUsecase
	covidUsecase domain.CovidUsecase
}

// NewCovidHandlers Comments handlers constructor
func NewCovidHandler(r *gin.Engine, chartUsecase domain.ChartUsecase, covidUsecase domain.CovidUsecase) {
	handler := &CovidHandler{chartUsecase: chartUsecase, covidUsecase: covidUsecase}
	g := r.Group("/covid")
	g.GET("/list", handler.GetListData)
	g.GET("/bar", handler.GetListDataBarChart)
	g.GET("/line", handler.GetListDataLineChart)

	g.GET("/search", handler.GetDataByKey)
	g.GET("/search/bar", handler.GetDataBarChartByKey)
}

// Get List Data godoc
// @Tags Covid
// @Summary List Data of Covid
// @Description This is API to get list data of Covid-19
// @Produce json
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.BadRequestResponse
// @Router /covid/list [get]
func (cv *CovidHandler) GetListData(c *gin.Context) {
	listData, err := cv.covidUsecase.GetAllData()
	if err != nil {
		c.JSON(404, gin.H{
			"code":    "001",
			"message": "Data Not Found",
			"result":  nil,
		})
		return
	}

	if len(listData.Detail) == 0 {
		c.JSON(404, gin.H{
			"code":    "001",
			"message": "Data Not Found",
			"result":  nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    "000",
		"message": "Data Found",
		"result":  listData,
	})
}

// Get Data By Key godoc
// @Tags Covid
// @Summary List Data of Covid
// @Description This is API to get list data of Covid-19 by its country
// @Param        country    query     string  true  "QueryParam"
// @Produce json
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.BadRequestResponse
// @Router /covid/search [get]
func (cv *CovidHandler) GetDataByKey(c *gin.Context) {
	key := c.Query("country")

	listData, err := cv.covidUsecase.GetDataByKey(key)
	if err != nil {
		c.JSON(404, gin.H{
			"code":    "001",
			"message": "Data Not Found",
			"result":  nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    "000",
		"message": "Data Found",
		"result":  listData,
	})
}

// Get List Data Bar Chart godoc
// @Tags Covid
// @Summary List Data of Covid
// @Description This is API to get list data of Covid-19 in bar chart
// @Produce json
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.BadRequestResponse
// @Router /covid/bar [get]
func (cv *CovidHandler) GetListDataBarChart(c *gin.Context) {
	listData, err := cv.covidUsecase.GetAllData()
	if err != nil {
		c.JSON(404, gin.H{
			"code":    "001",
			"message": "Data Not Found",
			"result":  nil,
		})
		return
	}

	// Prepare data
	var countries []string
	var confirmed, deaths []int
	var date string
	for i, v := range listData.Detail {
		countries = append(countries, v.CombinedKey)
		confirmed = append(confirmed, v.Confirmed)
		deaths = append(deaths, v.Deaths)

		if i == len(listData.Detail)-1 {
			dt, _ := time.Parse("2006-01-02 15:04:05", v.DateUpdated)
			date = dt.Format("2 January 2006 15:04:05")
		}
	}

	// Create bar chart
	confirmedBarChart := cv.chartUsecase.CreateBarChart(confirmed, countries, "Last Updated: "+date, "New cases in all over the world", "Confirmed")
	deathBarChart := cv.chartUsecase.CreateBarChart(deaths, countries, "", "Deaths in all over the world", "Deaths")

	// Show single chart in page
	confirmedBarChart.Render(c.Writer)
	deathBarChart.Render(c.Writer)
}

// Get Data Bar Chart By Key godoc
// @Tags Covid
// @Summary List Data of Covid
// @Description This is API to get list data of Covid-19 by its country in Bar Chart
// @Param        country    query     string  true  "QueryParam"
// @Produce json
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.BadRequestResponse
// @Router /covid/search/bar [get]
func (cv *CovidHandler) GetDataBarChartByKey(c *gin.Context) {
	key := c.Query("country")

	listData, err := cv.covidUsecase.GetDataByKey(key)
	if err != nil {
		c.JSON(404, gin.H{
			"code":    "001",
			"message": "Data Not Found",
			"result":  nil,
		})
		return
	}

	// Prepare data
	var countries []string
	var confirmed, deaths []int
	var date string
	for i, v := range listData.Detail {
		countries = append(countries, v.CombinedKey)
		confirmed = append(confirmed, v.Confirmed)
		deaths = append(deaths, v.Deaths)

		if i == len(listData.Detail)-1 {
			dt, _ := time.Parse("2006-01-02 15:04:05", v.DateUpdated)
			date = dt.Format("2 January 2006 15:04:05")
		}
	}

	// Create bar chart
	confirmedBarChart := cv.chartUsecase.CreateBarChart(confirmed, countries, "Last Updated: "+date, "New cases in "+key, "Confirmed")
	deathBarChart := cv.chartUsecase.CreateBarChart(deaths, countries, "", "Deaths in "+key, "Deaths")

	// Show single chart in page
	confirmedBarChart.Render(c.Writer)
	deathBarChart.Render(c.Writer)
}

// Get List Data Line Chart godoc
// @Tags Covid
// @Summary List Data of Covid
// @Description This is API to get list data of Covid-19 in line chart
// @Produce json
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.BadRequestResponse
// @Router /covid/line [get]
func (cv *CovidHandler) GetListDataLineChart(c *gin.Context) {
	listData, err := cv.covidUsecase.GetAllData()
	if err != nil {
		c.JSON(404, gin.H{
			"code":    "001",
			"message": "Data Not Found",
			"result":  nil,
		})
		return
	}

	// Prepare data
	var countries []string
	var confirmed, deaths []int
	var date string
	for i, v := range listData.Detail {
		countries = append(countries, v.CombinedKey)
		confirmed = append(confirmed, v.Confirmed)
		deaths = append(deaths, v.Deaths)

		if i == len(listData.Detail)-1 {
			dt, _ := time.Parse("2006-01-02 15:04:05", v.DateUpdated)
			date = dt.Format("2 January 2006 15:04:05")
		}
	}

	// Create line chart
	confirmedLineChart := cv.chartUsecase.CreateLineChart(confirmed, countries, "Last Updated: "+date, "New cases in all over the world", "Confirmed")
	deathLineChart := cv.chartUsecase.CreateLineChart(deaths, countries, "", "Deaths in all over the world", "Deaths")

	// Show single chart
	confirmedLineChart.Render(c.Writer)
	deathLineChart.Render(c.Writer)
}
