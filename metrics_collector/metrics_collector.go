package metrics_collector

import (
	"github.com/mackerelio/go-osstat/memory"
	"go.uber.org/zap"
	"main/snapshot"
	"main/state_providers/battery_state"
	"main/state_providers/brightness_state"
	"main/state_providers/cpu_stat"
	"main/state_providers/cpu_temp"
	"main/state_providers/keyboard_layout"
	"main/state_providers/network_connection_state"
	"main/state_providers/network_stat"
	"main/state_providers/volume_state"
	"main/util"
	"sync"
	"time"
)

type DwmBarMetricsCollector struct {
	snapshot *snapshot.DwmBarStatsSnapshot
	checker  *util.ErrorChecker
	logger   *zap.Logger
	config   *util.Config
}

func NewDwmBarMetricsCollector(
	snapshot *snapshot.DwmBarStatsSnapshot,
	checker *util.ErrorChecker,
	logger *zap.Logger,
	config *util.Config,
) *DwmBarMetricsCollector {
	return &DwmBarMetricsCollector{
		snapshot: snapshot,
		checker:  checker,
		logger:   logger,
		config:   config,
	}
}

func (c *DwmBarMetricsCollector) collectAllMetrics() {
	c.updateCpu()
	c.updateMemory()
	c.updateNetworkStateAndStats()
	c.updatePowerState()
	c.updateNowDateTime()
	c.updateTemp()
	c.updateKeyboardLayout()
	c.updateVolumeState()
	c.updateBrightnessState()
}

func (c *DwmBarMetricsCollector) callMethods(methods []func()) {
	var wg sync.WaitGroup
	for i := range methods {
		method := methods[i]

		wg.Add(1)
		go func() {
			defer wg.Done()
			method()
		}()
	}
	wg.Wait()
}

func (c *DwmBarMetricsCollector) collectEverySecondsMetrics() {
	methods := []func(){
		c.updateCpu,
		c.updateMemory,
		c.updateNetworkStateAndStats,
		c.updatePowerState,
		c.updateTemp,
		c.updateKeyboardLayout,
		c.updateVolumeState,
		c.updateBrightnessState,
	}

	c.callMethods(methods)
}

func (c *DwmBarMetricsCollector) collectEveryMinutesMetrics() {
	methods := []func(){
		c.updateNowDateTime,
	}

	c.callMethods(methods)
}

func (c *DwmBarMetricsCollector) Run() {
	tickerSecond := time.NewTicker(time.Second)
	defer tickerSecond.Stop()

	tickerMinute := time.NewTicker(time.Minute)
	defer tickerMinute.Stop()

	for {
		select {
		case <-tickerSecond.C:
			c.collectEverySecondsMetrics()
		case <-tickerMinute.C:
			c.collectEveryMinutesMetrics()
		}
	}
}

// FirstCollect Proxy method for extra case
func (c *DwmBarMetricsCollector) FirstCollect() {
	c.collectAllMetrics()
}

func (c *DwmBarMetricsCollector) updateCpu() {
	if c.config.NoCpu {
		return
	}

	cpu, err := cpu_stat.Get()
	if c.checker.ErrorFound(err) {
		return
	}

	c.snapshot.Cpu = cpu
}

func (c *DwmBarMetricsCollector) updateMemory() {
	if c.config.NoMemory {
		return
	}

	stats, err := memory.Get()
	if c.checker.ErrorFound(err) {
		return
	}

	c.snapshot.Mem = stats.Used
}

func (c *DwmBarMetricsCollector) updateNetworkStateAndStats() {
	if c.config.NoNetworkState && c.config.NoNetworkStats {
		return
	}

	connectionState, err := network_connection_state.Get()
	if c.checker.ErrorFound(err) {
		return
	}

	c.snapshot.NetworkState = *connectionState

	if connectionState.IsNotConnected() {
		return
	}

	interfaceName, err1 := connectionState.GetActiveInterfaceName()
	if c.checker.ErrorFound(err1) {
		return
	}

	networkStat, err2 := network_stat.Get(interfaceName)
	if c.checker.ErrorFound(err2) {
		return
	}

	c.snapshot.NetworkStat = *networkStat
}

func (c *DwmBarMetricsCollector) updatePowerState() {
	if c.config.NoPowerState {
		return
	}

	state, err := battery_state.Get()
	if c.checker.ErrorFound(err) {
		return
	}

	c.snapshot.BatteryState = *state
}

func (c *DwmBarMetricsCollector) updateNowDateTime() {
	c.snapshot.NowDateTime = time.Now()
}

func (c *DwmBarMetricsCollector) updateTemp() {
	if c.config.NoTemp {
		return
	}

	cpuTemp, err := cpu_temp.Get()
	if c.checker.ErrorFound(err) {
		return
	}

	c.snapshot.CpuTemp = *cpuTemp
}

func (c *DwmBarMetricsCollector) updateKeyboardLayout() {
	if c.config.NoKeyboardLayout {
		return
	}

	keyboardLayout, err := keyboard_layout.Get()
	if c.checker.ErrorFound(err) {
		return
	}

	c.snapshot.KeyboardLayout = *keyboardLayout
}

func (c *DwmBarMetricsCollector) updateVolumeState() {
	if c.config.NoVolume {
		return
	}

	volumeState, err := volume_state.Get()
	if c.checker.ErrorFound(err) {
		return
	}

	c.snapshot.VolumeState = *volumeState
}

func (c *DwmBarMetricsCollector) updateBrightnessState() {
	if c.config.NoBrightness {
		return
	}

	brightnessState, err := brightness_state.Get()
	if c.checker.ErrorFound(err) {
		return
	}

	c.snapshot.BrightnessState = *brightnessState
}
