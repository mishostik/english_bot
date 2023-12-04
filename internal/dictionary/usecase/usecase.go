package usecase

type DictionaryUseCase struct {
}

func NewDictionaryUsecase() DictionaryUseCase {
	return DictionaryUseCase{}
}

func (u *DictionaryUseCase) AddNewWord(word string) error {
	return nil
}

func (u *DictionaryUseCase) GetContextBySentence(sentence string) ([]string, error) {
	return []string{}, nil
}
