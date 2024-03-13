package mapper

type Mapper struct {
	OrderMapper OrderMapping
}

func NewMapper() *Mapper {
	return &Mapper{
		OrderMapper: NewOrderMapper(),
	}
}
