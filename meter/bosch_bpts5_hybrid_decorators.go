package meter

// Code generated by github.com/evcc-io/evcc/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/evcc-io/evcc/api"
)

func decorateBoschBpts5Hybrid(base api.Meter, battery func() (float64, error), batteryCapacity func() float64) api.Meter {
	switch {
	case battery == nil:
		return base

	case battery != nil && batteryCapacity == nil:
		return &struct {
			api.Meter
			api.Battery
		}{
			Meter: base,
			Battery: &decorateBoschBpts5HybridBatteryImpl{
				battery: battery,
			},
		}

	case battery != nil && batteryCapacity != nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
		}{
			Meter: base,
			Battery: &decorateBoschBpts5HybridBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateBoschBpts5HybridBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
		}
	}

	return nil
}

type decorateBoschBpts5HybridBatteryImpl struct {
	battery func() (float64, error)
}

func (impl *decorateBoschBpts5HybridBatteryImpl) Soc() (float64, error) {
	return impl.battery()
}

type decorateBoschBpts5HybridBatteryCapacityImpl struct {
	batteryCapacity func() float64
}

func (impl *decorateBoschBpts5HybridBatteryCapacityImpl) Capacity() float64 {
	return impl.batteryCapacity()
}
