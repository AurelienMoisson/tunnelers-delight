package terrain

import (
    "./chunk"
)

type Terrain struct {
    xWidth uint
    yWidth uint
    height uint
    chunkXLen, chunkYLen uint
    chunks [][]chunk.Chunk
    blockDistribs []chunk.BlockDistribution
}

func NewTerrain(xWidth, yWidth uint, blockDistribs []chunk.BlockDistribution) (Terrain) {
    chunks := make([][]chunk.Chunk, yWidth)
    var y uint
    for y=0; y<yWidth; y++ {
        chunks[y] = make([]chunk.Chunk, xWidth)
        var x uint
        for x=0; x<xWidth; x++ {
            chunks[y][x] = chunk.NewChunk(
                x*chunkSize+1,
                y*chunkSize+1,
                1,
                chunkSize,
                chunkSize,
                chunkHeight,
            )
        }
    }
    return Terrain{
        xWidth*chunkSize,
        yWidth*chunkSize,
        chunkHeight,
        xWidth,
        yWidth,
        chunks,
        blockDistribs,
    }
}

func (m *Terrain) getChunk(x,y uint) (*chunk.Chunk) {
    c := &m.chunks[y][x]
    // TODO : check if chunk is already generated
    c.PopulateChunk(m.blockDistribs)
    return c
}
