package keyspace

type Keyspace struct{
	space []byte
	length int
	current []int
}

func New(space []byte, length int) Keyspace {
	var k Keyspace
    k.space = space
    k.length = length

    k.current = make([]int, length)
    for i := range k.current  { k.current[i] = 0 }
    return k
}
