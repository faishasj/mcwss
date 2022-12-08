package event

import "github.com/sandertv/mcwss/mctype"

const (
	_ = iota
	AcquisitionPickedUp
	AcquisitionCrafted
	AcquisitionTakenFromContainer
	_
	_
	AcquisitionForged
	AcquisitionSmelted
	AcquisitionBrewed
	AcquisitionFilledBottle
	AcquistionTraded
	AcquisitionFished
)

// ItemAcquired is sent by the client when it acquires an item. This is most notably done when a player picks
// an item up from the ground, but also when it crafts a new item, gets one from a chest etc.
type ItemAcquired struct {
	AcquisitionMethodId int
	Count               int
	Item                mctype.Block
	Player              mctype.Player
}
