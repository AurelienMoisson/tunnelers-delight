package blocks

type SolidBlock struct {
    weight int
    fragility int
}

func NewSolidBlock(weight, solidity int) (SolidBlock) {
    return SolidBlock{weight, solidity}
}

func (b SolidBlock) GetWeight() (int) {
    return b.weight
}

func (b SolidBlock) GetFragility() (int) {
    return b.fragility
}

func (b SolidBlock) IsSolid() (bool) {
    return true
}
