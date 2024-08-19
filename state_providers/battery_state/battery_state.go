package battery_state

import (
	"errors"
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
	if err != nil {
		return nil, &BatteryStateError{
			message: "Could not get battery info!",
		}
	}

	for _, bat := range batteries {
		return &Stats{
			State:   bat.State.String(),
			Percent: int(bat.Current / bat.Full * 100),
		}, nil
	}

	return nil, errors.New("can't find battery")
}
