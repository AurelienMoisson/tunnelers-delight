package noise

import (
    "testing"
    "log"
)

func TestGetDensity(t *testing.T) {
    distr := getDistrib()
    log.Println(distr.GetDensity(2,2,3))
}

func TestGetZoneDensity(t *testing.T) {
    distr := getDistrib()
    zoneDensity := distr.GetZoneDensity(1, 5, 1, 5, 1, 5)
    log.Println(zoneDensity)
    log.Println(zoneDensity[25+5+2])
}

func BenchmarkGetDensity(b *testing.B) {
    distr := getDistrib()
    n := uint(b.N)
    b.ResetTimer()
    var i uint
    for i=0; i < n; i++ {
        distr.GetDensity(i, i+5, 17)
    }
}

func BenchmarkGetZoneDensity(b *testing.B) {
    distr := getDistrib()
    b.ResetTimer()
    for i:=0; i < b.N; i++ {
        distr.GetZoneDensity(0, 128, 0, 128, 0, 128)
    }
}

func getDistrib() Distribution {
    density := make([]uint, 256)
    for i:=0; i < 256; i++ {
        density[i] = 1000-uint(i)
    }
    periods := [][2]uint{
        [2]uint{1,2},
        [2]uint{3,6},
        [2]uint{11, 110},
        [2]uint{17, 71},
        [2]uint{376, 151},
    }
    return NewDistribution(density, periods, 3)
}
