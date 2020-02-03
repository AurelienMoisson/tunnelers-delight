package blocks

type AirBlock struct {}

func NewAirBlock() AirBlock {
    return AirBlock{}
}

func (b AirBlock) IsSolid() bool {
    return false
}

func (b AirBlock) GetWeight() int {
    return 0
}

func (b AirBlock) GetFragility() int {
    return 100
}
