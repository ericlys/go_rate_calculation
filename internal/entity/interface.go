package entity

type OrderRepositoryInterface interface { //qualquer package que tenhas essas funcoes já extende altomaticamente essa interface
	Save(order *Order) error
	GetTotal() (float64, error)
}
