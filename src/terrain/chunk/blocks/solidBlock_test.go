package blocks

import (
    "testing"
)

func TestStone(t *testing.T) {
    var stone SolidBlock = SolidBlock{1, 3}
    if stone.getWeight() != 1 {
        t.Error("expected weight of 1 from getWeight()")
    }
    if stone.getSolidity() != 3 {
        t.Error("expected solidity of 3 from getSolidity()")
    }
}
