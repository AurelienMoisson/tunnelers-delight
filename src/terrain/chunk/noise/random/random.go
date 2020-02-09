package random

func Random(seed uint, vals ...uint) uint {
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
