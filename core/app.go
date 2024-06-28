package core

import (
	"github.com/lkmproject/bigtwo/helper"
)

const (
	MaxPrivateRoomSize = 100
	NameMaxSize        = 8
	NameSmallSize      = 2
)

type App struct {
	room map[int32]int8 //roomid , player count
}

func CreateRoom(m map[int32]int8, size int) {
	count := 0
	for {
		dig := helper.Random5DigitNumber()
		if _, ok := m[dig]; !ok {
			m[dig] = 0
			count++
		}

		if count == size {
			break
		}
	}

}

func NewApp() *App {
	app := &App{
		room: make(map[int32]int8),
	}

	CreateRoom(app.room, MaxPrivateRoomSize)

	return app
}
