package noise

import (
    "testing"
)

func TestRandom(t *testing.T) {
    if (random(5, 2, 3) == random(5, 3, 2) && random(4, 1, 3) == random(4, 3, 1)) {
        t.Error("random seems to not care about arguments order")
    }
    if (random(5, 1, 4) == random(5, 2, 3) && random(6, 7, 9) == random(6, 6, 10)) {
        t.Error("random seems to only care about sum of arguments")
    }
    if (random(5, 1, 4) == random(4, 1, 4)) {
        t.Error("random seems to not care about the seed")
    }
    if (random(5, 1, 4) != random(5, 1, 4)) {
        t.Error("random is not reproducible")
    }
}

func BenchmarkRandom(b *testing.B) {
    var i uint
    var n uint
    n = uint(b.N)
    for i=0; i < n; i++ {
        random(45, i, 6, i)
    }
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
        density[i] = 10000-uint(i)
    }
    periods := [][2]uint{
        [2]uint{1,2},
        [2]uint{3,6},
        [2]uint{11, 110},
        [2]uint{17, 71},
        [2]uint{376, 150},
    }
    return newDistribution(density, periods, 7)
}
