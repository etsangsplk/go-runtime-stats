package goruntimestats

import (
	"appstats/goruntimestats/scheduler"
	"os"
	"time"

	"github.com/jtaczanowski/go-runtime-stats/api"
	"github.com/jtaczanowski/go-runtime-stats/collector"
	"github.com/jtaczanowski/go-runtime-stats/publisher"
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

	graphite := graphite.NewGraphiteClient(config.GraphiteHost, config.GraphitePort, config.GraphitePrefix, config.GraphiteProtocol)
	publisher := publisher.NewPublisher(collector, graphite)

	scheduler.RunTaskAtInterval(collector.CollectStats, config.Interval, time.Second*0)

	/* Run PublishToGraphite function every 10s, delay the first start by one second.
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
