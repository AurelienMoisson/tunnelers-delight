package chunk

import (
    "./blocks"
    "log"
)
var MaxStability int = 100

type Chunk struct {
    xStart uint
    yStart uint
    zStart uint
    xWidth uint
    yWidth uint
    height uint
    xEnd uint
    yEnd uint
    zEnd uint
    blockSlice []blocks.Block
    stability []int
}

func NewChunk(xStart, yStart, zStart, xWidth, yWidth, height uint) (c Chunk) {
    var blockSlice []blocks.Block
    blockSlice = make([]blocks.Block, xWidth * yWidth * height)
    var i uint
    for i=0; i < xWidth * yWidth * height; i++ {
        blockSlice[i] = blocks.NewAirBlock()
    }
    var stability []int
    stability = make([]int, (xWidth + 2) * (yWidth + 2) * height)
    c = Chunk{xStart, yStart, zStart, xWidth, yWidth, height, xStart+xWidth, yStart+yWidth, zStart+height, blockSlice, stability}
    return c
}

func (c Chunk) getAddress(x,y,z uint) uint {
    if (x==0) {
        panic("out of bounds address")
    }
    return x-c.xStart + (y-c.yStart + (z-c.zStart) * c.yWidth) * c.xWidth
}

func (c Chunk) getLargeAddress(x,y,z uint) uint {
    return x+1-c.xStart + (y+1-c.yStart + (z-c.zStart) * (c.yWidth+2)) * (c.xWidth+2)
}

func (c Chunk) GetBlock(x,y,z uint) (blocks.Block) {
    return c.blockSlice[c.getAddress(x,y,z)]
}

func (c Chunk) SetBlock(x,y,z uint, b blocks.Block) {
    c.blockSlice[c.getAddress(x,y,z)] = b
    c.updateStability(x,y,z)
    if c.GetBlock(x,y,z).IsSolid() != b.IsSolid(){
        c.updateStability(x,y,z+1)
    }
}

func (c Chunk) GetStability(x,y,z uint) int {
    return c.stability[c.getLargeAddress(x,y,z)]
}

func (c Chunk) setStability(x,y,z uint, stability int) {
    c.stability[c.getLargeAddress(x,y,z)] = stability
}

func (c Chunk) SetStability(x,y,z uint, stability int) {
    if (x>=c.xStart && x<c.xEnd && y>=c.yStart && y<c.yEnd) {
        log.Fatal("shouldn't SetStability inside chunk borders")
    }
    c.setStability(x,y,z, stability)
    if (x==c.xStart-1) {
        c.updateStability(x+1,y,z)
    }
    if (x==c.xEnd) {
        c.updateStability(x-1,y,z)
    }
    if (y==c.yStart-1) {
        c.updateStability(x,y+1,z)
    }
    if (y==c.yEnd) {
        c.updateStability(x,y-1,z)
    }
}

func (c Chunk) updateStability(x,y,z uint) {
    if (x<c.xStart || y<c.yStart || x>=c.xEnd || y>=c.yEnd) {
        return
    }
    if (!c.GetBlock(x,y,z).IsSolid()) {
        if (c.GetStability(x,y,z) == 0) {
            return
        }
        c.setStability(x,y,z, 0)
    } else if (z==c.zStart || c.GetBlock(x,y,z-1).IsSolid()) {
        if (c.GetStability(x,y,z) == MaxStability) {
            return
        }
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

func (c Chunk) shouldFall(x,y,z uint) bool {
    return c.GetStability(x,y,z)>0
}
