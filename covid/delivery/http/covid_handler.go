package http

import (
	"github.com/gin-gonic/gin"
	"github.com/intaen/apigocovid/domain"
)

type CovidHandler struct {
	chartUsecase domain.ChartUsecase
	covidUsecase domain.CovidUsecase
}

// NewCovidHandlers Comments handlers constructor
func NewCovidHandler(r *gin.Engine, chartUsecase domain.ChartUsecase, covidUsecase domain.CovidUsecase) {
	handler := &CovidHandler{chartUsecase: chartUsecase, covidUsecase: covidUsecase}
	g := r.Group("/covid")
	g.GET("/all", handler.GetListData)
	g.GET("/list", handler.GetDataByKey)
	g.GET("/bar", handler.GetListDataBarChart)
	g.GET("/line", handler.GetListDataLineChart)
}

// Get List Data godoc
// @Tags Covid
// @Summary List Data of Covid
// @Description This is API to get list data of Covid-19
// @Produce json
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.BadRequestResponse
// @Router /covid/all [get]
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
// @Produce json
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.BadRequestResponse
// @Router /covid/list?country [get]
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
// @Router /covid/list/v2 [get]
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
	for _, v := range listData.Detail {
		countries = append(countries, v.CombinedKey)
		confirmed = append(confirmed, v.Confirmed)
		deaths = append(deaths, v.Deaths)
	}

	// Create bar chart
	confirmedBarChart := cv.chartUsecase.CreateBarChart(confirmed, countries, "GO COVID", "This is list of data confirmed covid in the world", "Confirmed")
	deathBarChart := cv.chartUsecase.CreateBarChart(deaths, countries, "", "This is list of data deaths covid in the world", "Deaths")

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
// @Router /covid/list/v3 [get]
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
	for _, v := range listData.Detail {
		countries = append(countries, v.CombinedKey)
		confirmed = append(confirmed, v.Confirmed)
		deaths = append(deaths, v.Deaths)
	}

	// Create line chart
	confirmedLineChart := cv.chartUsecase.CreateLineChart(confirmed, countries, "GO COVID", "This is list of data confirmed covid in the world", "Confirmed")
	deathLineChart := cv.chartUsecase.CreateLineChart(deaths, countries, "", "This is list of data deaths covid in the world", "Deaths")

	// Show single chart
	confirmedLineChart.Render(c.Writer)
	deathLineChart.Render(c.Writer)
}
