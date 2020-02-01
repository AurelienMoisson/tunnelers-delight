package blocks

type SolidBlock struct {
    weight int
    solidity int
}

func NewSolidBlock(weight, solidity int) (SolidBlock) {
    return SolidBlock{weight, solidity}
}

func (b SolidBlock) GetWeight() (int) {
    return b.weight
}

func (b SolidBlock) GetSolidity() (int) {
    return b.solidity
}
