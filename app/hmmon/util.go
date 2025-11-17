package hmmon

import (
	"net"

	"github.com/rs/zerolog/log"
)

func pingInternet(address string) bool {
	conn, err := net.Dial("tcp", address+":80")
	if err != nil {
		log.Error().Err(err).Msg("pingInternet failed")
		return false
	}
	defer conn.Close()
	log.Info().Msg("pingInternet success")
	return true
}

// func
