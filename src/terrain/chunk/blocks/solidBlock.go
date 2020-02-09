package blocks

type SolidBlock struct {
    weight int
    fragility int
}

func NewSolidBlock(weight, fragility int) (SolidBlock) {
    return SolidBlock{weight, fragility}
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
