package state

type UseCase interface {
	RememberUserState(userId int, newState string)
	GetUserState(userId int) ([]string, error)
	CleanUserState(userId int)
}
