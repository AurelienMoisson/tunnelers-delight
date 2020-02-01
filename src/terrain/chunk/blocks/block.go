package blocks

type Block interface {
    GetWeight() int
    GetSolidity() int
}
