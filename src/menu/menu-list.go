package menu

import (
	"strconv"

	"github.com/nsf/termbox-go"
)

var (
	offlinePlayersConfigElements = []opt{}

	supportingColors = []termbox.Attribute{termbox.ColorBlue, termbox.ColorRed, termbox.ColorGreen, termbox.ColorYellow}
	colors           []termbox.Attribute
	maxBotCount      = 0

	mainMenu = menu{
		opts: []opt{
			{
				label: "Start",
				onSelect: func(mpt *menuPagesType) (bool, callback) {
					mpt.changeMenuPage(1)
					return false, nil
				},
			},
			{
				label: "Quit",
				onSelect: func(mpt *menuPagesType) (bool, callback) {
					return true, nil
				},
			},
		},
	}

	startMenu = menu{
		opts: []opt{
			{
				label: "Offline",
				onSelect: func(mpt *menuPagesType) (bool, callback) {
					mpt.changeMenuPage(2)
					return false, nil
				},
			},
			{
				label: "Online",
				onSelect: func(mpt *menuPagesType) (bool, callback) {
					return false, nil
				},
			},
			{
				label: "Back",
				onSelect: func(mpt *menuPagesType) (bool, callback) {
					mpt.changeMenuPage(0)
					return false, nil
				},
			},
		},
	}

	offlineMenu = menu{
		opts: []opt{
			{
				label:   "No of Players: ",
				subOpts: []interface{}{"2", "3", "4"},
				onSelect: func(mpt *menuPagesType) (bool, callback) {
					setAvailableColors(mpt.menus[mpt.curIdx].opts[0].curIdx + 2)
					offlinePlayersConfigElements = getPlayersOptionsElements(mpt.menus[mpt.curIdx].opts[0].curIdx + 2)
					offlinePlayersConfig.opts = append(offlinePlayersConfigElements, offlinePlayersConfig.opts[len(offlinePlayersConfig.opts)-1])
					mpt.changeMenuPage(3)
					return false, nil
				},
			},
			{
				label: "Back",
				onSelect: func(mpt *menuPagesType) (bool, callback) {
					mpt.changeMenuPage(1)
					return false, nil
				},
			},
		},
	}

	offlinePlayersConfig = menu{
		opts: []opt{
			{
				label: "Back",
				onSelect: func(mpt *menuPagesType) (bool, callback) {
					mpt.changeMenuPage(1)
					return false, nil
				},
			},
		},
	}
)

func setAvailableColors(count int) {
	maxBotCount = count - 1
	colors = supportingColors[0:count]
}

func getPlayersOptionsElements(count int) []opt {
	optsElements := []opt{}
	for i := 1; i <= count; i++ {
		optsElements = append(optsElements, opt{
			label:   "Player " + strconv.Itoa(i),
			subOpts: []interface{}{"bot"},
		})
	}

	return optsElements
}
