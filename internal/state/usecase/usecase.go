package usecase

import (
	"log"
)

type StateUseCase struct {
	UserState map[int][]string
}

func NewStateUseCase() StateUseCase {
	return StateUseCase{
		UserState: make(map[int][]string),
	}
}

func (u *StateUseCase) RememberUserState(userId int, newState string) {
	countUserStates := len(u.UserState[userId])
	if countUserStates == 0 {
		u.UserState[userId] = append(u.UserState[userId], newState)
		return
	}
	if u.UserState[userId][countUserStates-1] != newState {
		u.UserState[userId] = append(u.UserState[userId], newState)
		u.CleanUserState(userId)
	}
}

//	if len(u.UserState[userId]) == 0 {
//		return []string{}, fmt.Errorf("user state empty")
//	}
func (u *StateUseCase) GetUserState(userId int) []string {
	log.Println(" ---- user state in getting state -", userId)
	if _, ok := u.UserState[userId]; !ok {
		return []string{}
	}
	return u.UserState[userId]
}

func (u *StateUseCase) CleanUserState(userId int) {
	if len(u.UserState[userId]) > 10 {
		u.UserState[userId] = u.UserState[userId][len(u.UserState[userId])-10:]
	}
}
