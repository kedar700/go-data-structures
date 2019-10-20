package queues

type Luggage struct {
	Weight    int
	Passenger string
}

func NewLuggage(weight int, passenger string) *Luggage {
	return &Luggage{Weight: weight, Passenger: passenger}
}

type Belt []*Luggage

func (b *Belt) Add(newLuggage *Luggage) {
	*b = append(*b, newLuggage)
}

func (b *Belt) Remove() *Luggage {
	first, rest := (*b)[0], (*b)[1:]
	*b = rest
	return first
}
