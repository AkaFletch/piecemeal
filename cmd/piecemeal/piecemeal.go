package main

import (
	"strings"

	log "github.com/rs/zerolog/log"
	"threlfall.dev/piecemeal/v2/internal/fen"
)

func main() {
	log.Info().Msg("Starting Piecemeal")
	testReader := strings.NewReader("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	fen.ReadFen(testReader)
}
