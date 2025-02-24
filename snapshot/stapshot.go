package snapshot

import (
	"main/state_providers/battery_state"
	"main/state_providers/brightness_state"
	"main/state_providers/cpu_temp"
	"main/state_providers/keyboard_layout"
	"main/state_providers/network_connection_state"
	"main/state_providers/network_stat"
	"main/state_providers/notifications_state"
	"main/state_providers/volume_state"
	"main/state_providers/weather_state"
	"time"
)

type DwmBarStatsSnapshot struct {
	NowDateTime        time.Time
	KeyboardLayout     keyboard_layout.Stats
	BatteryState       battery_state.Stats
	NetworkState       network_connection_state.Stats
	VolumeState        volume_state.Stats
	BrightnessState    brightness_state.Stats
	Mem                uint64
	CpuTemp            cpu_temp.Stats
	Cpu                int
	NetworkStat        network_stat.Stats
	NotificationsState notifications_state.Stats
	WeatherState       *weather_state.Stats
}

func NewDwmBarStatsSnapshot() *DwmBarStatsSnapshot {
	return &DwmBarStatsSnapshot{}
}
