package prcy

var TestNames = []string{
	"publicStatistics",
}

type Analysis struct {
	HttpStatus       int                    `json:"httpStatus"`
	Error            string                 `json:"error,omitempty"`
	PublicStatistics *PublicStatisticsModel `json:"publicStatistics,omitempty"`
}

func NewAnalysis() *Analysis {
	return &Analysis{}
}

type ErrorModel struct {
	Error string `json:"error"`
}

type TestModel struct {
	Name    string `json:"name"`    // ex: publicStatistics
	Updated string `json:"updated"` // ex: 2019-04-03T18:20:28+03:00
}

type PublicStatisticsModel struct {
	*TestModel
	SourceLink       string `json:"publicStatisticsSourceLink"` // ex: http://www.alexa.com/siteinfo/plandex.ru
	SourceType       string `json:"publicStatisticsSourceType"` // ex: alexa
	VisitsDaily      int    `json:"publicStatisticsVisitsDaily"`
	VisitsWeekly     int    `json:"publicStatisticsVisitsWeekly"`
	VisitsMonthly    int    `json:"publicStatisticsVisitsMonthly"`
	PageViewsDaily   int    `json:"publicStatisticsPageViewsDaily"`
	PageViewsWeekly  int    `json:"publicStatisticsPageViewsWeekly"`
	PageViewsMonthly int    `json:"publicStatisticsPageViewsMonthly"`
}
