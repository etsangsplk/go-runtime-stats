package goruntimestats

import (
	"os"
	"time"

	"github.com/jtaczanowski/go-graphite-client"
	"github.com/jtaczanowski/go-runtime-stats/api"
	"github.com/jtaczanowski/go-runtime-stats/collector"
	"github.com/jtaczanowski/go-runtime-stats/publisher"
	"github.com/jtaczanowski/go-scheduler"
)

type Config struct {
	GraphiteHost     string
	GraphitePort     int
	GraphiteProtocol string
	GraphitePrefix   string
	Interval         time.Duration
	HTTPOn           bool
	HTTPPort         int
}

// Start - starts goruntimestats in background
func Start(config Config) {
	collector := collector.NewCollector()

	graphite := graphite.NewClient(config.GraphiteHost, config.GraphitePort, config.GraphitePrefix, config.GraphiteProtocol)
	publisher := publisher.NewPublisher(collector, graphite)

	scheduler.RunTaskAtInterval(collector.CollectStats, config.Interval, time.Second*0)

	/* Run PublishToGraphite function every config.Interval, delay the first start by one second.
	 * This time delay allows to runing PublishToGraphite function after CollectStats function.
	 */
	scheduler.RunTaskAtInterval(publisher.PublishToGraphite, config.Interval, time.Second*1)

	if config.HTTPOn {
		api := api.NewHttpServer(config.HTTPPort, collector)
		err := api.Start()
		if err != nil {
			os.Exit(1)
		}
	}
}
