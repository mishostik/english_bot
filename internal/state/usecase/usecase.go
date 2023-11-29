package usecase

import "fmt"

type StateUseCase struct {
	UserState map[int][]string
}

func NewStateUseCase() StateUseCase {
	return StateUseCase{
		UserState: make(map[int][]string),
	}
}

func (u *StateUseCase) RememberUserState(userId int, newState string) {
	u.UserState[userId] = append(u.UserState[userId], newState)
	u.CleanUserState(userId)
}

func (u *StateUseCase) GetUserState(userId int) ([]string, error) {
	if len(u.UserState[userId]) == 0 {
		return []string{}, fmt.Errorf("user state empty")
	}
	return u.UserState[userId], nil
}

func (u *StateUseCase) CleanUserState(userId int) {
	if len(u.UserState[userId]) > 10 {
		u.UserState[userId] = u.UserState[userId][len(u.UserState[userId])-10:]
	}
}
