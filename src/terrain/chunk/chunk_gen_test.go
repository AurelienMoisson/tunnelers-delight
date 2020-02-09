package chunk

import (
    "./noise"
    "./blocks"
    "./noise/random"
    "log"
    "testing"
)

func createBlockDistributions(seed uint) []BlockDistribution {
    blockDistribs := make([]BlockDistribution, 3)
    var block blocks.Block
    block = blocks.NewSolidBlock(10,10)
    stoneDensity := []uint{
        1000,
        1000,
        1000,
        1000,
        1000,
        1000,
        1000,
        1000,
    }
    distrib := noise.NewDistribution(stoneDensity, [][2]uint{}, random.Random(seed, 1))
    blockDistribs[0] = NewBlockDistribution(block, distrib)
    block = blocks.NewSolidBlock(11, 9)
    zeroDensity := []uint{
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
    }
    periods := [][2]uint{
        [2]uint{7, 50},
        [2]uint{17, 250},
        [2]uint{128, 1500},
    }
    distrib = noise.NewDistribution(zeroDensity, periods, random.Random(seed,2))
    blockDistribs[1] = NewBlockDistribution(block, distrib)
    block = blocks.NewSolidBlock(12, 8)
    periods = [][2]uint{
        [2]uint{1, 5},
        [2]uint{3, 200},
        [2]uint{10, 750},
        [2]uint{128, 500},
    }
    distrib = noise.NewDistribution(zeroDensity, periods, random.Random(seed,3))
    blockDistribs[2] = NewBlockDistribution(block, distrib)
    return blockDistribs
}

func TestPopulateChunk(t *testing.T) {
    blockDistribs := createBlockDistributions(3)
    yWidth := uint(12)
    xWidth := uint(32)
    c := NewChunk(1,1,1,xWidth,yWidth,8)
    c.PopulateChunk(blockDistribs)
    for z:=uint(1); z<9; z++ {
        for y:=uint(1); y<yWidth+1; y++ {
            line := ""
            for x:=uint(1); x<xWidth+1; x++ {
                switch c.GetBlock(x,y,z).GetFragility() {
                case 10:
                    line += "#"
                case 9:
                    line += "%"
                case 8:
                    line += "+"
                default:
                    line += "/"
                }
            }
            log.Println(line)
        }
        log.Println()
    }
}
