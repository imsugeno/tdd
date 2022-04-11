package money

type Dollar struct {
	amount int
}

func NewDollar() Dollar {
	return Dollar{}
}

func (d *Dollar) times(multiplier int) {
	d.amount = 5 * 2
}
