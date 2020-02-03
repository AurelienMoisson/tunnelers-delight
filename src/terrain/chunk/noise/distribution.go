package noise

type Distribution struct {
    density []uint
    periods [][2]uint
    offsets [][3]uint
    seed uint
}

func newDistribution(density []uint, periods [][2]uint, seed uint) (Distribution) {
    offsets := make([][3]uint, len(periods))
    for i:=0; i <len(periods); i++ {
        offsets[i] = [3]uint{random(seed, 1), random(seed, 2), random(seed, 3)}
    }
    return Distribution{density, periods, offsets, seed}
}

func (d Distribution) GetZoneDensity(xMin, xMax, yMin, yMax, zMin, zMax uint) []uint {
    totalSize := (xMax-xMin) * (yMax-yMin) * (zMax-zMin)
    dens := make([]uint, totalSize)
    // for each period
    for i:=0; i<len(d.periods); i++ {
        xStart := xMin
        yStart := yMin
        zStart := zMin
        lowerZ := zMin/d.periods[i][0]
        upperZ := lowerZ + d.periods[i][0]
        // iterate over blocks of size period
        for lowerZ<zMax {
            lowerY := yMin/d.periods[i][0]
            upperY := lowerY + d.periods[i][0]
            for lowerY<yMax {
                lowerX := xMin/d.periods[i][0]
                upperX := lowerX + d.periods[i][0]
                for lowerX<xMax {
                    z := zStart
                    corner1 := float32(random(d.seed, lowerX, lowerY, lowerZ, d.periods[i][0]))
                    corner2 := float32(random(d.seed, lowerX, lowerY, upperZ, d.periods[i][0]))
                    corner3 := float32(random(d.seed, lowerX, upperY, lowerZ, d.periods[i][0]))
                    corner4 := float32(random(d.seed, lowerX, upperY, upperZ, d.periods[i][0]))
                    corner5 := float32(random(d.seed, upperX, lowerY, lowerZ, d.periods[i][0]))
                    corner6 := float32(random(d.seed, upperX, lowerY, upperZ, d.periods[i][0]))
                    corner7 := float32(random(d.seed, upperX, upperY, lowerZ, d.periods[i][0]))
                    corner8 := float32(random(d.seed, upperX, upperY, upperZ, d.periods[i][0]))
                    // iterate inside block
                    for z<upperZ && z<zMax {
                        zProp := float32(z%d.periods[i][0])/float32(d.periods[i][0])
                        edge1 := smoothStep(zProp, corner1, corner2)
                        edge2 := smoothStep(zProp, corner3, corner4)
                        edge3 := smoothStep(zProp, corner5, corner6)
                        edge4 := smoothStep(zProp, corner7, corner8)
                        y := yStart
                        for y<upperY && y<yMax {
                            yProp := float32(y%d.periods[i][0])/float32(d.periods[i][0])
                            face1 := smoothStep(yProp, edge1, edge2)
                            face2 := smoothStep(yProp, edge3, edge4)
                            x := xStart
                            for x<upperX && x<xMax {
                                xProp := float32(x%d.periods[i][0])/float32(d.periods[i][0])
                                dens[x+y*(xMax-xMin + z*(yMax-yMin))] += uint(smoothStep(xProp, face1, face2))
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
    var dens float32
    dens = 0
    maxUint := ^uint(0)
    for i:=0; i<len(d.periods); i++ {
        dens += float32(d.periods[i][1])/float32(maxUint) * d.oneWavePeriod(x,y,z, i)
    }
    return uint(dens * float32(d.density[z])/float32(maxUint))
}

func (d Distribution) oneWavePeriod(x,y,z uint, i int) float32 {
    lowerX := x/d.periods[i][0]
    lowerY := y/d.periods[i][0]
    lowerZ := z/d.periods[i][0]
    upperX := lowerX + d.periods[i][0]
    upperY := lowerY + d.periods[i][0]
    upperZ := lowerZ + d.periods[i][0]
    xProp := float32(x-lowerX) / float32(d.periods[i][0])
    yProp := float32(y-lowerY) / float32(d.periods[i][0])
    zProp := float32(z-lowerZ) / float32(d.periods[i][0])
    corner1 := float32(random(d.seed, lowerX, lowerY, lowerZ, 4))
    corner2 := float32(random(d.seed, lowerX, lowerY, upperZ, 4))
    corner3 := float32(random(d.seed, lowerX, upperY, lowerZ, 4))
    corner4 := float32(random(d.seed, lowerX, upperY, upperZ, 4))
    corner5 := float32(random(d.seed, upperX, lowerY, lowerZ, 4))
    corner6 := float32(random(d.seed, upperX, lowerY, upperZ, 4))
    corner7 := float32(random(d.seed, upperX, upperY, lowerZ, 4))
    corner8 := float32(random(d.seed, upperX, upperY, upperZ, 4))
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
    return float32(v1 + (v2-v1)) * (3*squared + 2*squared*x)
}

func random(seed uint, vals ...uint) uint {
    var r uint
    r = seed
    for _, v := range vals {
        r = hash(r^v)
    }
    return r
}

func hash(x uint) uint {
    x = ((x>>16) ^ x) * 0x45d9f3b
    x = ((x>>16) ^ x) * 0x45d9f3b
    x = ((x>>16) * x)
    return x
}
