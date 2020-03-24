package analyser

import (
	"siteanalyser/core"
	"time"
)

func sleep(config *core.Config, multi int) {
	time.Sleep(time.Duration(multi*config.PrCyDelay) * time.Millisecond)
}
