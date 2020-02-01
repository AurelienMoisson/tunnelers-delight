package blocks

type Block interface {
    getWeight() int
    getSolidity() int
}
