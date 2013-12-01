package keyspace

type Keyspace struct{
	space []byte
	length int
}

func New(space []byte, length int) Keyspace {
	var k Keyspace
    k.space = space
    k.length = length
    return k
}
