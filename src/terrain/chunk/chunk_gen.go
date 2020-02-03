package chunk

import (
    "./blocks"
    "./noise"
)

type BlockDistribution struct {
    block blocks.Block
    distribution noise.Distribution
}

func NewBlockDistribution(block blocks.Block, distribution noise.Distribution) BlockDistribution {
    return BlockDistribution{block, distribution}
}

func (c Chunk) PopulateChunk(blockDistribs []BlockDistribution) {
    var bestScore []uint
    totalSize := c.xWidth * c.yWidth * c.height
    bestScore = make([]uint, totalSize)
    for _, blockDistr := range blockDistribs {
        newScore := blockDistr.distribution.GetZoneDensity(c.xStart, c.xEnd, c.yStart, c.yEnd, c.zStart, c.zEnd)
        var x uint
        for x=0; x<c.xWidth; x++ {
            var y uint
            for y=0; y<c.yWidth; y++ {
                var z uint
                for z=0; z<c.height; z++ {
                    i:= x + (y + z*c.yWidth)*c.xWidth
                    if newScore[i] > bestScore[i] {
                        c.SetBlock(x+c.xStart, y+c.yStart, z+c.zStart, blockDistr.block)
                        bestScore[i] = newScore[i]
                    }
                }
            }
        }
    }
}
