package signal

type Signal struct {
	Turn bool
}

func New(_switch bool) Signal {
	return Signal{
		Turn: _switch,
	}
}