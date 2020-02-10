package terrain

import (
    "testing"
    "./chunk"
    "./chunk/blocks"
    "./chunk/noise"
    "./chunk/noise/random"
    "log"
)

func getDistributions() ([]chunk.BlockDistribution) {
    stoneConst := make([]uint, chunkHeight)
    deepMetalConst := make([]uint, chunkHeight)
    shallowMetalConst := make([]uint, chunkHeight)
    airConst := make([]uint, chunkHeight)
    var z uint
    for z=0; z<chunkHeight; z++ {
        stoneConst[z] = 100000
        deepMetalConst[z] = 92000 - 64*z
        shallowMetalConst[z] = 81000 + 3*z*(chunkHeight-z)
        airConst[z] = 85000+128*z
    }
    stone := blocks.NewSolidBlock(10, 10)
    platinum := blocks.NewSolidBlock(12, 8)
    gold := blocks.NewSolidBlock(12, 6)
    air := blocks.NewAirBlock()
    randomSeed := uint(0);
    stoneDistrib := noise.NewDistribution(stoneConst, [][2]uint{}, random.Random(randomSeed, 1))
    platinumDistrib := noise.NewDistribution(deepMetalConst, [][2]uint{
        [2]uint{1, 600},
        [2]uint{3, 10000},
        [2]uint{11, 1000},
        [2]uint{92, 1000},
    }, random.Random(randomSeed, 2))
    goldDistrib := noise.NewDistribution(shallowMetalConst, [][2]uint{
        [2]uint{1, 300},
        [2]uint{3, 10000},
        [2]uint{11, 7000},
        [2]uint{92, 2000},
    }, random.Random(randomSeed, 3))
    airDistrib := noise.NewDistribution(airConst, [][2]uint{
        [2]uint{1, 9000},
        [2]uint{3, 6900},
        [2]uint{5, 6300},
        [2]uint{17, 19300},
        [2]uint{31, 22300},
    }, random.Random(randomSeed, 4))
    return []chunk.BlockDistribution{
        chunk.NewBlockDistribution(stone, stoneDistrib),
        chunk.NewBlockDistribution(platinum, platinumDistrib),
        chunk.NewBlockDistribution(gold, goldDistrib),
        chunk.NewBlockDistribution(air, airDistrib),
    }
}

func TestTerrain(t *testing.T) {
    m := NewTerrain(10, 10, getDistributions())
    c := m.getChunk(0, 0)
    countStone := 0
    countPlatinum := 0
    countGold := 0
    countAir := 0
    for z:=uint(1); z<chunkHeight; z++ {
        for y:=uint(1); y<chunkSize; y++ {
            line := ""
            for x:=uint(1); x<chunkSize; x++ {
                switch c.GetBlock(x,y,z).GetFragility() {
                case 10:
                    line += "#"
                    countStone += 1
                case 8:
                    line += "+"
                    countPlatinum += 1
                case 6:
                    line += "*"
                    countGold += 1
                default:
                    line += " "
                    countAir += 1
                }
            }
            log.Println(line)
        }
        log.Println("----")
    }

    log.Println("stone", countStone)
    log.Println("platinum", countPlatinum)
    log.Println("gold", countGold)
    log.Println("air", countAir)
}
