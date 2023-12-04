package dictionary

type UseCase interface {
	AddNewWord(word string) error
	GetContextBySentence(sentence string) ([]string, error)
}
