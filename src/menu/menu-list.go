package menu

import (
	"ludo/src/game"
	gameUtils "ludo/src/game-utils"
	"ludo/src/network"

	"github.com/nsf/termbox-go"
)

var (
	mainMenu = menu{
		opts: []opt{
			{
				label: "Start",
				onSelect: func(mpt *menuPagesType) (bool, callback) {
					mpt.changeMenuPage(START_MENU)
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
					mpt.changeMenuPage(OFFLINE_MENU)
					return false, nil
				},
			},
			{
				label: "Online",
				onSelect: func(mpt *menuPagesType) (bool, callback) {
					mpt.changeMenuPage(ONLINE_MENU)
					return false, nil
				},
			},
			{
				label: "Back",
				onSelect: func(mpt *menuPagesType) (bool, callback) {
					mpt.changeMenuPage(MAIN_MENU)
					return false, nil
				},
			},
		},
	}

	playerSubOpts = []string{"Player", "Bot", "-"}

	onlineMenu = menu{
		opts: []opt{
			{
				label: "Host",
				onSelect: func(mpt *menuPagesType) (bool, callback) {
					mpt.changeMenuPage(HOST_MENU)
					return false, nil
				},
			},
			{
				label: "Join",
				onSelect: func(mpt *menuPagesType) (bool, callback) {
					return false, func() {
						game.StartGameOnline()
					}
				},
			},
			{
				label: "Back",
				onSelect: func(mpt *menuPagesType) (bool, callback) {
					mpt.changeMenuPage(START_MENU)
					return false, nil
				},
			},
		},
	}

	hostMenu = menu{
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
					players := []gameUtils.PlayerData{}
					curMenuOpts := mpt.menus[mpt.curIdx].opts
					curMenuOpts = curMenuOpts[:len(curMenuOpts)-2]

					for _, opt := range curMenuOpts {
						players = append(players, gameUtils.PlayerData{Color: opt.label.(termbox.Attribute), Type: opt.subOpts[opt.curIdx]})
					}

					return true, func() {
						termbox.Close()
						termbox.Init()
						go network.Host(players)
						game.StartGameOnline()
					}
				},
			},
			{
				label: "Back",
				onSelect: func(mpt *menuPagesType) (bool, callback) {
					mpt.changeMenuPage(ONLINE_MENU)
					return false, nil
				},
			},
		},
	}

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
					players := []gameUtils.PlayerData{}
					curMenuOpts := mpt.menus[mpt.curIdx].opts
					curMenuOpts = curMenuOpts[:len(curMenuOpts)-2]

					for _, opt := range curMenuOpts {
						players = append(players, gameUtils.PlayerData{Color: opt.label.(termbox.Attribute), Type: opt.subOpts[opt.curIdx]})
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
					mpt.changeMenuPage(START_MENU)
					return false, nil
				},
			},
		},
	}
)
