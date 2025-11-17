package hmmon

import (
	"context"
	"time"

	probing "github.com/prometheus-community/pro-bing"
	"github.com/rs/zerolog/log"
)

func pingInternet(address string) bool {
	pinger, err := probing.NewPinger(address)
	if err != nil {
		panic(err)
	}
	pinger.Count = 3

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	c := make(chan struct{})

	go func() {
		_ = pinger.Run()
		close(c)
	}()

	select {
	case <-ctx.Done():
	case <-c:
	}

	pinger.Stop()

	cancel()

	stats := pinger.Statistics()

	log.Debug().Msgf("Ping stats to %s: %+v", address, stats)
	log.Debug().Msgf("Ping packets loss to %s: %.2f%%", address, stats.PacketLoss)
	log.Debug().Msgf("Ping avg rtt to %s: %s", address, stats.AvgRtt.String())
	log.Debug().Msgf("Ping min rtt to %s: %s", address, stats.MinRtt.String())

	ok := stats.PacketsRecv > 0
	if !ok {
		log.Warn().Msgf("Internet ping to %s failed", address)
	} else {
		log.Info().Msgf("Internet ping to %s successful", address)
	}

	return ok
}

// func
