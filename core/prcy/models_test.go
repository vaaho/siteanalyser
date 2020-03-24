package prcy

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrCyModelsParsing_1(t *testing.T) {
	data := []byte(`{
		"error":"Ошибка анализа"
	}`)

	var model Analysis
	err := json.Unmarshal(data, &model)

	assert.Nil(t, err)
	assert.Equal(t, model.Error, "Ошибка анализа")
	assert.Nil(t, model.PublicStatistics, nil)
}

func TestPrCyModelsParsing_2(t *testing.T) {
	data := []byte(`{
		"publicStatistics": {
			"publicStatisticsSourceType": "prcy",
			"publicStatisticsSourceLink": "#",
			"publicStatisticsVisitsDaily": 950000,
			"publicStatisticsVisitsWeekly": 6650000,
			"publicStatisticsVisitsMonthly": 27550000,
			"publicStatisticsPageViewsDaily": 12110000,
			"publicStatisticsPageViewsWeekly": 84720000,
			"publicStatisticsPageViewsMonthly": 350950000,
			"publicStatisticsAlexaVisits": 30700,
			"publicStatisticsPrcyVisits": 950000,
			"name": "publicStatistics",
			"updated": "2020-03-23T17:16:48+03:00"
		}
	}`)

	var model Analysis
	err := json.Unmarshal(data, &model)

	assert.Nil(t, err)
	assert.Empty(t, model.Error)
	assert.NotNil(t, model.PublicStatistics)
	assert.Equal(t, model.PublicStatistics.Name, "publicStatistics")
	assert.Equal(t, model.PublicStatistics.Updated, "2020-03-23T17:16:48+03:00")
	assert.Equal(t, model.PublicStatistics.SourceType, "prcy")
	assert.Equal(t, model.PublicStatistics.SourceLink, "#")
	assert.Equal(t, model.PublicStatistics.VisitsDaily, 950000)
	assert.Equal(t, model.PublicStatistics.VisitsWeekly, 6650000)
	assert.Equal(t, model.PublicStatistics.VisitsMonthly, 27550000)
	assert.Equal(t, model.PublicStatistics.PageViewsDaily, 12110000)
	assert.Equal(t, model.PublicStatistics.PageViewsWeekly, 84720000)
	assert.Equal(t, model.PublicStatistics.PageViewsMonthly, 350950000)
}
