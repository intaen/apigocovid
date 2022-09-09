package domain

import (
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
)

// Usecase represent the covid's usecases
type CovidUsecase interface {
	// Scheduler
	AddMasterStatistic()

	// API
	GetAllData() (*Result, error)

	// Data
	CreateBarChart([]int, []string, string, string, string) *charts.Bar
	CreateLineChart([]int, []string, string, string, string) *charts.Line
}

// Repository represent the covid's repository contract
type CovidRepository interface {
	// API
	GetAllData() (*CoronaVirus, error)

	// Database
	AddData([]CoronaVirusStatistic) error
	FindAllData() ([]CoronaVirusStatistic, error)
}

// Hero represent model
type CoronaVirusStatistic struct {
	CoronaVirusStatisticID uint      `gorm:"primarykey;autoIncrement:true"`
	Country                string    `gorm:"column:country"`
	ProvinceState          string    `gorm:"column:province_state"`
	CombinedKey            string    `gorm:"column:combined_key;unique"`
	Active                 int       `gorm:"column:active"`
	Confirmed              int       `gorm:"column:confirmed"`
	Deaths                 int       `gorm:"column:death"`
	Recovered              int       `gorm:"column:recovered"`
	DataSource             string    `gorm:"column:data_source"`
	PublishedBy            string    `gorm:"column:published_by"`
	DateCreated            time.Time `gorm:"column:date_created;autoUpdateTime:nano"`
	DateUpdated            string    `gorm:"column:date_updated"`
}

func (s CoronaVirusStatistic) TableName() string {
	return "coronavirus_statistic"
}

type (
	//
	Result struct {
		Count  int            `json:"count"`
		Detail []DetailResult `json:"detail"`
	}

	DetailResult struct {
		Country     string `json:"-"`
		CombinedKey string `json:"country_region"`
		Active      int    `json:"active"`
		Confirmed   int    `json:"confirmed"`
		Deaths      int    `json:"deaths"`
		Recovered   int    `json:"recovered"`
		DataSource  string `json:"-"`
		PublishedBy string `json:"-"`
		DateUpdated string `json:"-"`
	}

	// API
	CoronaVirus struct {
		SummaryStats  Summary    `json:"summaryStats"`
		Cache         Cache      `json:"cache"`
		DataSource    DataSource `json:"dataSource"`
		APISourceCode string     `json:"apiSourceCode"`
		RawData       []RawData  `json:"rawData"`
	}

	Summary struct {
		Global   Status `json:"global"`
		China    Status `json:"china"`
		NonChina Status `json:"nonChina"`
	}

	Status struct {
		Confirmed int `json:"confirmed"`
		Recovered int `json:"recovered"`
		Deaths    int `json:"deaths"`
	}

	Cache struct {
		LastUpdated          string `json:"lastUpdated"`
		Expired              string `json:"expires"`
		LastUpdatedTimestamp int64  `json:"lastUpdatedTimestamp"`
		ExpiresTimestamp     int64  `json:"expiresTimestamp"`
	}

	DataSource struct {
		URL              string `json:"url"`
		LastGithubCommit string `json:"lastGithubCommit"`
		PublishedBy      string `json:"publishedBy"`
		Ref              string `json:"ref"`
	}

	RawData struct {
		FIPS                string `json:"FIPS"`
		Admin               string `json:"Admin2"`
		ProvinceState       string `json:"Province_State"`
		CountryRegion       string `json:"Country_Region"`
		LastUpdate          string `json:"Last_Update"`
		Lat                 string `json:"Lat"`
		Long                string `json:"Long"`
		Confirmed           string `json:"Confirmed"`
		Deaths              string `json:"Deaths"`
		Recovered           string `json:"Recovered"`
		Active              string `json:"Active"`
		Combined_Key        string `json:"Combined_Key"`
		Incident_Rate       string `json:"Incident_Rate"`
		Case_Fatality_Ratio string `json:"Case_Fatality_Ratio"`
	}
)
