package random

import (
    "testing"
)

func TestRandom(t *testing.T) {
    if (Random(5, 2, 3) == Random(5, 3, 2) && Random(4, 1, 3) == Random(4, 3, 1)) {
        t.Error("Random seems to not care about arguments order")
    }
    if (Random(5, 1, 4) == Random(5, 2, 3) && Random(6, 7, 9) == Random(6, 6, 10)) {
        t.Error("Random seems to only care about sum of arguments")
    }
    if (Random(5, 1, 4) == Random(4, 1, 4)) {
        t.Error("Random seems to not care about the seed")
    }
    if (Random(5, 1, 4) != Random(5, 1, 4)) {
        t.Error("Random is not reproducible")
    }
}

func BenchmarkRandom(b *testing.B) {
    var i uint
    var n uint
    n = uint(b.N)
    for i=0; i < n; i++ {
        Random(45, i, 6, i)
    }
}
