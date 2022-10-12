package http

import (
	"apigocovid/domain"

	"github.com/gin-gonic/gin"
)

type ChartHandler struct {
	chartUsecase domain.ChartUsecase
}

// NewChartHandlers Comments handlers constructor
func NewChartHandler(r *gin.Engine, chartUsecase domain.ChartUsecase) {
	handler := &ChartHandler{chartUsecase: chartUsecase}
	g := r.Group("/chart")
	g.GET("/example/bar", handler.GetListExampleBarChart)
}

// Get List Example Bar Chart godoc
// @Tags Example
// @Summary Example of Data in Bar Chart
// @Description This is API to get example of bar chart
// @Produce json
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.BadRequestResponse
// @Router /chart/example [get]
func (ch *ChartHandler) GetListExampleBarChart(c *gin.Context) {
	// Prepare data
	var months = []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	var days = []string{"Sun", "Mon", "Tue", "Wed", "Thur", "Fri", "Sat"}
	var values1 = []int{1997, 1998, 1999, 2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008}
	var values2 = []int{300, 200, 100, 20, 400, 530, 109}

	// Create bar chart
	confirmedBarChart := ch.chartUsecase.CreateBarChart(values1, months, "GO Example", "This is example of bar chart", "Month")
	deathBarChart := ch.chartUsecase.CreateBarChart(values2, days, "", "This is example of bar chart", "Day")

	// Show single chart in page
	confirmedBarChart.Render(c.Writer)
	deathBarChart.Render(c.Writer)
}
