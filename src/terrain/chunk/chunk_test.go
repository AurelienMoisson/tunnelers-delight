package chunk

import (
    "testing"
    "./blocks"
    "log"
)

func TestNewChunk(t *testing.T) {
    var chunk Chunk
    chunk = NewChunk(1, 1, 1, 4, 4, 4)
    var block1 blocks.Block
    block1 = blocks.NewSolidBlock(2, 20)
    chunk.SetBlock(2, 3, 3, block1)
    if chunk.GetBlock(2, 3, 3).GetWeight() != 2 {
        t.Error("didn't get the right block")
    }
}

func TestStability(t *testing.T) {
    var chunk Chunk
    chunk = NewChunk(1, 1, 1, 4, 4, 4)
    var block1 blocks.Block
    var air blocks.Block
    block1 = blocks.NewSolidBlock(2, 20)
    air = blocks.NewAirBlock()
    if chunk.GetStability(2, 3, 3) > 0 {
        log.Print(chunk.GetStability(2,3, 3))
        t.Error("didn't get the right stability")
    }
    chunk.SetStability(0, 2, 2, MaxStability)
    if chunk.GetStability(0, 2, 2) != MaxStability {
        t.Error("SetStability didn't work")
    }
    chunk.SetBlock(1, 2, 2, block1)
    if chunk.GetStability(1, 2, 2) != MaxStability-20 {
        log.Print("found stability of ", chunk.GetStability(1, 2, 2), " instead of ", MaxStability-20)
        t.Error("stability did not propagate correctly horrizontally")
    }
    chunk.SetBlock(1, 2, 2, air)
    chunk.SetBlock(2,2,2, block1)
    if chunk.GetStability(2,2,2) > 0 {
        t.Error("positive stability for floating block")
    }
    chunk.SetBlock(3,2,2, block1)
    chunk.SetBlock(1,2,2, block1)
    if chunk.GetStability(3,2,2) != MaxStability-60 {
        log.Print("found stability of ", chunk.GetStability(3, 2, 2), " instead of ", MaxStability-60)
        t.Error("stability didn't propagate far successfully")
    }
}

func BenchmarkStability(b *testing.B) {
    var chunk Chunk
    chunk = NewChunk(1, 1, 1, 10, 10, 10)
    var block1 blocks.Block
    var air blocks.Block
    block1 = blocks.NewSolidBlock(2, 20)
    air = blocks.NewAirBlock()
    var x uint
    for x= 2; x<9; x++ {
        var y uint
        for y= 2; y<9; y++ {
            chunk.SetBlock(x,y,2, block1)
        }
    }
    chunk.SetBlock(1,1,2, block1)
    chunk.SetBlock(7,8,1, block1)
    b.ResetTimer()
    for i:=0; i<b.N; i++ {
        chunk.SetBlock(7,8,1, air)
        chunk.SetBlock(7,8,1, block1)
    }
}
