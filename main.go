package main

import (
	"go.uber.org/dig"
	"log"
	"main/drawer"
	"main/metrics_collector"
	"main/snapshot"
	"main/themes"
	"main/util"
	"sync"
)

func f(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	c := dig.New()
	var wg sync.WaitGroup

	f(c.Provide(util.NewConfig))

	// Inject logger singleton
	f(c.Provide(util.NewLogger))

	// Inject error checker singleton
	f(c.Provide(util.NewErrorChecker))

	// Inject theme singleton(you can use another theme, see themes namespace)
	f(c.Provide(themes.CreateThemeTomorrowNight))

	// Inject snapshot(used for store metrics) singleton
	f(c.Provide(snapshot.NewDwmBarStatsSnapshot))

	// Inject metrics collector singleton
	f(c.Provide(metrics_collector.NewDwmBarMetricsCollector))

	// Inject drawer(required snapshot) singleton
	f(c.Provide(drawer.NewDwmBarDrawer))

	// Run metrics collector
	f(c.Invoke(func(collector *metrics_collector.DwmBarMetricsCollector) {
		collector.FirstCollect()

		wg.Add(1)
		go func() {
			defer wg.Done()
			collector.Run()
		}()
	}))

	// Run drawer
	f(c.Invoke(func(drawer *drawer.Drawer) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			drawer.Run()
		}()
	}))

	wg.Wait()
}
