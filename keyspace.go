package keyspace

import (
	"errors"
)

type Keyspace struct{
	space []byte
	length int
	current []int
}

// Should increment to the next entry in the keyspace, and return that entry.
// Will return an error when the keyspace is exhausted
func (k * Keyspace) Next() ([]byte, error) {
	err := k.incrementString(k.length-1)
	return k.toBytes(), err
}

// Should decrement to the previous entry in the keyspace, and return that entry.
// Will return an error when the keyspace is exhausted
func (k * Keyspace) Previous() ([]byte, error) {
	err := k.decrementString(k.length-1)
	return k.toBytes(), err
}

func (k * Keyspace) Length() int {
	return k.length
}

func (k * Keyspace) toBytes() []byte {
	x := make([]byte, k.length)
	for idx, value := range k.current {
      x[idx]=k.space[value]
    }
    return x
}

func (k * Keyspace) incrementString(offset int) (err error) {
	x := k.current

	if(offset == 0 && x[offset] == len(k.space)-1) {
	  return errors.New("Keyspace: Maximum value reached")
	}
	if(x[offset]==len(k.space)-1) {
		x[offset]=0
		if err := k.incrementString(offset-1); err != nil {
			return err
		}
	} else {
		x[offset]++
	}
	k.current = x
	return nil
}

func (k * Keyspace) decrementString(offset int) (err error) {
	x := k.current

  // if(offset == 0 && x[offset] == len(k.space)-1) {
  //   return errors.New("Keyspace: Maximum value reached")
  // }
  hasPositiveNumber := false
  for num := range x {
    if num > 0 {
      hasPositiveNumber = true
      break
    }
  }
  if(!hasPositiveNumber) {
    return errors.New("Keyspace: Minimum value reached")
  }
	if(x[offset]==0) {
		x[offset]=len(k.space)-1
		if err := k.decrementString(offset-1); err != nil {
			return err
		}
	} else {
		x[offset]--
	}
	k.current = x
	return nil
}


func New(space []byte, length int) Keyspace {
	var k Keyspace
    k.space = space
    k.length = length

    k.current = make([]int, length)
    for i := range k.current  { k.current[i] = 0 }
    k.current[length-1]=-1 // otherwise we miss the first value
    return k
}
