package keyspace

type Keyspace struct{
	space []byte
	length int
	current []int
}

func (k * Keyspace) Next() []byte {
	k.incrementString(k.length-1)
	x := make([]byte, k.length)
	for idx, value := range k.current {
      x[idx]=k.space[value]
    }
    return x
}

func (k * Keyspace) incrementString(offset int) {
	x := k.current
	if(x[offset]==len(k.space)-1) {
		x[offset]=0
		k.incrementString(offset-1)
	} else {
		x[offset]++
	}
	k.current = x
}

func New(space []byte, length int) Keyspace {
	var k Keyspace
    k.space = space
    k.length = length

    k.current = make([]int, length)
    for i := range k.current  { k.current[i] = 0 }
    return k
}
