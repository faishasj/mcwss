package event

import "github.com/sandertv/mcwss/mctype"

const (
	DestructionBreak = iota
)

// BlockBroken is sent by the client when it breaks a block.
type BlockBroken struct {
	Block             mctype.Block
	Count             int
	DestructionMethod int
	Player            mctype.Player
	Tool              mctype.Tool
	Variant           int
}
