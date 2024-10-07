package checks

import (
	"log"
	"paropd/checks/check_packages"
	"time"
)

type CheckManager struct {
	tickers []*time.Ticker
}

func (cm *CheckManager) Start() {
	cm.Close()

	log.Println("Starting checks")
	check_packages := &check_packages.CheckPackages{}
	cm.addTicker(schedule(check_packages.Measure, check_packages.MeasureFrequency()))
}

func (cm *CheckManager) Close() {
	cm.ensureInit()
	cm.stopTickers()
}

// PRIVATE

func (cm *CheckManager) ensureInit() {
	if cm.tickers == nil {
		cm.tickers = []*time.Ticker{}
	}
}

func (cm *CheckManager) addTicker(ticker *time.Ticker) {
	cm.tickers = append(cm.tickers, ticker)
}

func (cm *CheckManager) stopTickers() {
	for _, ticker := range cm.tickers {
		ticker.Stop()
	}
	// Clear the slice
	cm.tickers = cm.tickers[:0]
}

func schedule(f func(), interval time.Duration) *time.Ticker {
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			f()
		}
	}()
	return ticker
}
