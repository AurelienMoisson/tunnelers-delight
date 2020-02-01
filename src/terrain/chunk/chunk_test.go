package chunk

import (
    "testing"
    "./blocks"
)

func TestNewChunk(t *testing.T) {
    var chunk Chunk
    chunk = NewChunk(16, 16, 16)
    var block1 blocks.Block
    block1 = blocks.NewSolidBlock(2, 3)
    chunk.setBlock(2, 3, 4, block1)
    if chunk.getBlock(2, 3, 4).GetWeight() != 2 {
        t.Error("didn't get the right block")
    }
}
