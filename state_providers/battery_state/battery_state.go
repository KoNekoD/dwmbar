package battery_state

import (
	"github.com/distatus/battery"
)

type Stats struct {
	State   string
	Percent int
}

type BatteryStateError struct {
	message string
}

func (e *BatteryStateError) Error() string {
	return e.message
}

func Get() (*Stats, error) {
	batteries, err := battery.GetAll()

	for i, b := range batteries {
		if err == nil || err.(battery.Errors)[i] == nil {
			return &Stats{
				State:   b.State.String(),
				Percent: int(b.Current / b.Full * 100),
			}, nil
		}
	}

	return nil, &BatteryStateError{message: "Could not get battery info!"}
}
