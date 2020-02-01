package blocks

type SolidBlock struct {
    weight int
    solidity int
}

func (b *SolidBlock) getWeight() (int) {
    return b.weight
}

func (b *SolidBlock) getSolidity() (int) {
    return b.solidity
}
