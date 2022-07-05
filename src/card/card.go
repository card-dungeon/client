package card

type CardType int

const (
	CHARACTER CardType = iota + 1
	PASSIVE
	MAGIC
	EQUIP
)

type Status struct {
	name string
	desc string
	atk  int
	hp   int
	sd   int
}

func Init(cardType CardType, data []Status) {

}
