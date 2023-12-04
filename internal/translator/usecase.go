package translator

type UseCase interface {
	Translate(data string) (string, error)
}
