package checks_helpers

import "time"

type Check interface {
	Measure()
	MeasureFrequency() time.Duration
}
