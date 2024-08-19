package network_stat

import (
	"errors"
	"github.com/mackerelio/go-osstat/network"
	"time"
)

func Get(currentInterfaceName string) (*Stats, error) {
	stats, err := network.Get()
	time.Sleep(time.Second)
	stats2, err := network.Get()
	if err != nil {
		return nil, err
	}

	currentStats := findByName(stats, currentInterfaceName)
	if currentStats == nil {
		return nil, errors.New(
			"can't get stats by provided interface(currentStats)",
		)
	}

	currentStats2 := findByName(stats2, currentInterfaceName)
	if currentStats2 == nil {
		return nil, errors.New(
			"can't get stats by provided interface(currentStats2)",
		)
	}

	return &Stats{
		RxBytes: currentStats2.RxBytes - currentStats.RxBytes,
		TxBytes: currentStats2.TxBytes - currentStats.TxBytes,
	}, nil
}

type Stats struct {
	RxBytes, TxBytes uint64
}

func findByName(stats []network.Stats, name string) *network.Stats {
	for i := range stats {
		if stats[i].Name == name {
			return &stats[i]
		}
	}

	return nil
}
