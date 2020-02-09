package noise

import (
    "./random"
)

type Distribution struct {
    density []uint
    periods [][2]uint
    seed uint
}

func NewDistribution(density []uint, periods [][2]uint, seed uint) (Distribution) {
    for i:=0; i <len(periods); i++ {
        if (periods[i][0] == 0) {
            panic("can't handle period of 0; please use density")
        }
    }
    return Distribution{density, periods, seed}
}

func (d Distribution) GetZoneDensity(xMin, xMax, yMin, yMax, zMin, zMax uint) []uint {
    totalSize := (xMax-xMin) * (yMax-yMin) * (zMax-zMin)
    dens := make([]uint, totalSize)
    var z uint
    for z=0; z<(zMax-zMin); z++ {
        var j uint
        for j=0; j<(xMax-xMin)*(yMax-yMin); j++ {
            dens[z*(yMax-yMin)*(xMax-xMin) + j] += d.density[z]
        }
    }
    // for each period
    for i:=0; i<len(d.periods); i++ {
        xStart := xMin
        yStart := yMin
        zStart := zMin
        lowerZ := zMin/d.periods[i][0] * d.periods[i][0]
        upperZ := lowerZ + d.periods[i][0]
        // iterate over blocks of size period
        for lowerZ<zMax {
            lowerY := yMin/d.periods[i][0] * d.periods[i][0]
            upperY := lowerY + d.periods[i][0]
            for lowerY<yMax {
                lowerX := xMin/d.periods[i][0] * d.periods[i][0]
                upperX := lowerX + d.periods[i][0]
                for lowerX<xMax {
                    z := zStart
                    corner1 := float32(random.Random(d.seed, lowerX, lowerY, lowerZ, d.periods[i][0])%d.periods[i][1])
                    corner2 := float32(random.Random(d.seed, lowerX, lowerY, upperZ, d.periods[i][0])%d.periods[i][1])
                    corner3 := float32(random.Random(d.seed, lowerX, upperY, lowerZ, d.periods[i][0])%d.periods[i][1])
                    corner4 := float32(random.Random(d.seed, lowerX, upperY, upperZ, d.periods[i][0])%d.periods[i][1])
                    corner5 := float32(random.Random(d.seed, upperX, lowerY, lowerZ, d.periods[i][0])%d.periods[i][1])
                    corner6 := float32(random.Random(d.seed, upperX, lowerY, upperZ, d.periods[i][0])%d.periods[i][1])
                    corner7 := float32(random.Random(d.seed, upperX, upperY, lowerZ, d.periods[i][0])%d.periods[i][1])
                    corner8 := float32(random.Random(d.seed, upperX, upperY, upperZ, d.periods[i][0])%d.periods[i][1])
                    // iterate inside block
                    for z<upperZ && z<zMax {
                        zProp := float32(z-lowerZ)/float32(d.periods[i][0])
                        edge1 := smoothStep(zProp, corner1, corner2)
                        edge2 := smoothStep(zProp, corner3, corner4)
                        edge3 := smoothStep(zProp, corner5, corner6)
                        edge4 := smoothStep(zProp, corner7, corner8)
                        y := yStart
                        for y<upperY && y<yMax {
                            yProp := float32(y-lowerY)/float32(d.periods[i][0])
                            face1 := smoothStep(yProp, edge1, edge2)
                            face2 := smoothStep(yProp, edge3, edge4)
                            x := xStart
                            for x<upperX && x<xMax {
                                xProp := float32(x-lowerX)/float32(d.periods[i][0])
                                dens[x-xStart+(y-yStart)*(xMax-xMin + (z-zStart)*(yMax-yMin))] += uint(smoothStep(xProp, face1, face2))
                                x++
                            }
                            y++
                        }
                        z++
                    }
                    lowerX = upperX
                    xStart = lowerX
                    upperX += d.periods[i][0]
                }
                lowerY = upperY
                yStart = lowerY
                upperY += d.periods[i][0]
            }
            lowerZ = upperZ
            zStart = lowerZ
            upperZ += d.periods[i][0]
        }
    }
    return dens
}

func (d Distribution) GetDensity(x,y,z uint) uint {
    if d.density[z] == 0 {
        return 0
    }
    var dens uint
    dens = 0
    dens+=d.density[z]
    for i:=0; i<len(d.periods); i++ {
        dens += uint(d.oneWavePeriod(x,y,z, i))
    }
    return uint(dens)
}

func (d Distribution) oneWavePeriod(x,y,z uint, i int) float32 {
    lowerX := x/d.periods[i][0] * d.periods[i][0]
    lowerY := y/d.periods[i][0] * d.periods[i][0]
    lowerZ := z/d.periods[i][0] * d.periods[i][0]
    upperX := lowerX + d.periods[i][0]
    upperY := lowerY + d.periods[i][0]
    upperZ := lowerZ + d.periods[i][0]
    for upperX<x {
        lowerX = upperX
        upperX += d.periods[i][0]
    }
    for upperY<y {
        lowerY = upperY
        upperY += d.periods[i][0]
    }
    for upperZ<z {
        lowerZ = upperZ
        upperZ += d.periods[i][0]
    }
    xProp := float32(x-lowerX) / float32(d.periods[i][0])
    yProp := float32(y-lowerY) / float32(d.periods[i][0])
    zProp := float32(z-lowerZ) / float32(d.periods[i][0])
    corner1 := float32(random.Random(d.seed, lowerX, lowerY, lowerZ, d.periods[i][0])%d.periods[i][1])
    corner2 := float32(random.Random(d.seed, lowerX, lowerY, upperZ, d.periods[i][0])%d.periods[i][1])
    corner3 := float32(random.Random(d.seed, lowerX, upperY, lowerZ, d.periods[i][0])%d.periods[i][1])
    corner4 := float32(random.Random(d.seed, lowerX, upperY, upperZ, d.periods[i][0])%d.periods[i][1])
    corner5 := float32(random.Random(d.seed, upperX, lowerY, lowerZ, d.periods[i][0])%d.periods[i][1])
    corner6 := float32(random.Random(d.seed, upperX, lowerY, upperZ, d.periods[i][0])%d.periods[i][1])
    corner7 := float32(random.Random(d.seed, upperX, upperY, lowerZ, d.periods[i][0])%d.periods[i][1])
    corner8 := float32(random.Random(d.seed, upperX, upperY, upperZ, d.periods[i][0])%d.periods[i][1])
    edge1 := smoothStep(zProp, corner1, corner2)
    edge2 := smoothStep(zProp, corner3, corner4)
    edge3 := smoothStep(zProp, corner5, corner6)
    edge4 := smoothStep(zProp, corner7, corner8)
    face1 := smoothStep(yProp, edge1, edge2)
    face2 := smoothStep(yProp, edge3, edge4)
    return smoothStep(xProp, face1, face2)
}

func smoothStep(x float32,v1,v2 float32) float32 {
    squared := x*x
    return v1 + (v2-v1) * (3*squared + 2*squared*x)
}
