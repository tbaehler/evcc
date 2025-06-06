package charger

// Code generated by github.com/evcc-io/evcc/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/evcc-io/evcc/api"
)

func decorateHomeAssistantSwitch(base *HomeAssistantSwitch, meter func() (float64, error)) api.Charger {
	switch {
	case meter == nil:
		return base

	case meter != nil:
		return &struct {
			*HomeAssistantSwitch
			api.Meter
		}{
			HomeAssistantSwitch: base,
			Meter: &decorateHomeAssistantSwitchMeterImpl{
				meter: meter,
			},
		}
	}

	return nil
}

type decorateHomeAssistantSwitchMeterImpl struct {
	meter func() (float64, error)
}

func (impl *decorateHomeAssistantSwitchMeterImpl) CurrentPower() (float64, error) {
	return impl.meter()
}
