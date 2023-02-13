package event

import "github.com/sandertv/mcwss/mctype"

const (
	PlacementDefault = iota
)

// BlockPlaced is called by a client when it places a block.
type BlockPlaced struct {
	Block            mctype.Block
	Count            int
	PlacedUnderWater bool
	PlacementMethod  int
	Player           mctype.Player
	Tool             mctype.Tool
}
