package menu

import (
	"ludo/src/game"

	"github.com/nsf/termbox-go"
)

var (
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

	playerSubOpts = []string{"Player", "Bot", "-"}

	offlineMenu = menu{
		opts: []opt{
			{
				label:   termbox.ColorBlue,
				subOpts: playerSubOpts,
			},
			{
				label:   termbox.ColorRed,
				subOpts: playerSubOpts,
			},
			{
				label:   termbox.ColorYellow,
				subOpts: playerSubOpts,
			},
			{
				label:   termbox.ColorGreen,
				subOpts: playerSubOpts,
			},
			{
				label: "Done",
				onSelect: func(mpt *menuPagesType) (bool, callback) {
					players := []game.PlayerData{}
					curMenuOpts := mpt.menus[mpt.curIdx].opts
					curMenuOpts = curMenuOpts[:len(curMenuOpts)-2]

					for _, opt := range curMenuOpts {
						players = append(players, game.PlayerData{Color: opt.label.(termbox.Attribute), Type: opt.subOpts[opt.curIdx]})
					}
					return true, func() {
						termbox.Close()
						termbox.Init()
						game.StartGameOffline(players)
					}
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
)
