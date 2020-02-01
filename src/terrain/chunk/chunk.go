package chunk

import (
    "./blocks"
)

type Chunk struct {
    xWidth int
    yWidth int
    height int
    blockSlice []blocks.Block
}

func NewChunk(xWidth, yWidth, height int) (c Chunk) {
    var blockSlice []blocks.Block
    blockSlice = make([]blocks.Block, xWidth * yWidth * height)
    c = Chunk{xWidth, yWidth, height, blockSlice}
    return c
}

func (c Chunk) getBlock(x,y,z int) (blocks.Block) {
    return c.blockSlice[x + (y + z * c.yWidth) * c.xWidth]
}

func (c Chunk) setBlock(x,y,z int, b blocks.Block) {
    c.blockSlice[x + (y + z * c.yWidth) * c.xWidth] = b
}
