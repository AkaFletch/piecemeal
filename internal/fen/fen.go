package fen

import (
	"io"
	"io/ioutil"

	log "github.com/rs/zerolog/log"

	"threlfall.dev/piecemeal/v2/internal/board"
)

func ReadFen(reader io.Reader) board.BoardRepresentation {
	buf, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal().Err(err).Msgf("Error reading FEN")
	}
	fen := string(buf[:])
	log.Debug().Str("FEN Input", fen).Msg("Reading FEN")
	board := board.BoardRepresentation{}

	return board
}
