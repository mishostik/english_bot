package state

type UseCase interface {
	RememberUserState(userId int, newState string)
	GetUserState(userId int) []string
	CleanUserState(userId int)
}
