package chunk

import (
    "./blocks"
    "log"
)
var MaxStability int = 100

type Chunk struct {
    xWidth int
    yWidth int
    height int
    blockSlice []blocks.Block
    stability []int
}

func NewChunk(xWidth, yWidth, height int) (c Chunk) {
    var blockSlice []blocks.Block
    blockSlice = make([]blocks.Block, xWidth * yWidth * height)
    for i:=0; i < xWidth * yWidth * height; i++ {
        blockSlice[i] = blocks.NewAirBlock()
    }
    var stability []int
    stability = make([]int, (xWidth + 2) * (yWidth + 2) * height)
    c = Chunk{xWidth, yWidth, height, blockSlice, stability}
    return c
}

func (c Chunk) getAddress(x,y,z int) int {
    return x + (y + z * c.yWidth) * c.xWidth
}

func (c Chunk) getLargeAddress(x,y,z int) int {
    return x+1 + (y+1 + z * (c.yWidth+1)) * (c.xWidth+1)
}

func (c Chunk) GetBlock(x,y,z int) (blocks.Block) {
    return c.blockSlice[c.getAddress(x,y,z)]
}

func (c Chunk) SetBlock(x,y,z int, b blocks.Block) {
    c.blockSlice[c.getAddress(x,y,z)] = b
    c.updateStability(x,y,z)
    if c.GetBlock(x,y,z).IsSolid() != b.IsSolid(){
        c.updateStability(x,y,z+1)
    }
}

func (c Chunk) GetStability(x,y,z int) int {
    return c.stability[c.getLargeAddress(x,y,z)]
}

func (c Chunk) setStability(x,y,z int, stability int) {
    c.stability[c.getLargeAddress(x,y,z)] = stability
}

func (c Chunk) SetStability(x,y,z int, stability int) {
    if (x>=0 && x<c.xWidth && y>=0 && y<c.yWidth) {
        log.Fatal("shouldn't SetStability inside chunk borders")
    }
    c.setStability(x,y,z, stability)
    if (x==-1) {
        c.updateStability(x+1,y,z)
    }
    if (x==c.xWidth) {
        c.updateStability(x-1,y,z)
    }
    if (y==-1) {
        c.updateStability(x,y+1,z)
    }
    if (y==c.yWidth) {
        c.updateStability(x,y-1,z)
    }
}

func (c Chunk) updateStability(x,y,z int) {
    if (x<0 || y<0 || x>=c.xWidth || y>=c.yWidth) {
        return
    }
    if (!c.GetBlock(x,y,z).IsSolid()) {
        if (c.GetStability(x,y,z) == 0) {
            return
        }
        c.setStability(x,y,z, 0)
    } else if (z==0 || c.GetBlock(x,y,z-1).IsSolid()) {
        c.setStability(x,y,z, MaxStability)
    } else {
        var stability int
        stability = c.GetStability(x-1,y,z)
        if (stability < c.GetStability(x,y-1,z)) {
            stability = c.GetStability(x,y-1,z)
        }
        if (stability < c.GetStability(x+1,y,z)) {
            stability = c.GetStability(x+1,y,z)
        }
        if (stability < c.GetStability(x,y+1,z)) {
            stability = c.GetStability(x,y+1,z)
        }
        stability -= c.GetBlock(x,y,z).GetFragility()
        if stability == c.GetStability(x,y,z) {
            return
        }
        c.setStability(x,y,z, stability)
    }
    c.updateStability(x+1,y,z)
    c.updateStability(x-1,y,z)
    c.updateStability(x,y+1,z)
    c.updateStability(x,y-1,z)
}

func (c Chunk) shouldFall(x,y,z int) bool {
    return c.GetStability(x,y,z)>0
}
