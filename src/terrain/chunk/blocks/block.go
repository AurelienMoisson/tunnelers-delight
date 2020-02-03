package blocks

type Block interface {
    GetWeight() int
    GetFragility() int
    IsSolid() bool
}
