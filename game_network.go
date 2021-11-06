package go_boardgame_networking

import (
	"fmt"
	"github.com/quibbble/go-boardgame/pkg/bgn"
	"strings"
)

type GameNetwork struct {
	hubs map[string]*gameHub // mapping from game key to game hub
}

func NewGameNetwork(options GameNetworkOptions) *GameNetwork {
	hubs := make(map[string]*gameHub)
	for _, builder := range options.Games {
		hub := newGameHub(builder, options.GameExpiry, options.Adapters)
		go hub.Start()
		hubs[builder.Key()] = hub
	}
	return &GameNetwork{
		hubs: hubs,
	}
}

func (n *GameNetwork) CreateGame(options CreateGameOptions) error {
	hub, ok := n.hubs[options.NetworkOptions.GameKey]
	if !ok {
		return fmt.Errorf("game key '%s' does not exist", options.NetworkOptions.GameKey)
	}
	if len(options.NetworkOptions.Players) > 0 && len(options.NetworkOptions.Players) != len(options.GameOptions.Teams) {
		return fmt.Errorf("number of teams are inconsistent")
	}
	return hub.Create(options)
}

func (n *GameNetwork) LoadGame(options LoadGameOptions) error {
	hub, ok := n.hubs[options.NetworkOptions.GameKey]
	if !ok {
		return fmt.Errorf("game key '%s' does not exist", options.NetworkOptions.GameKey)
	}
	if len(options.NetworkOptions.Players) > 0 && len(options.NetworkOptions.Players) != len(strings.Split(options.BGN.Tags["Teams"], ", ")) {
		return fmt.Errorf("number of teams are inconsistent")
	}
	return hub.Load(options)
}

func (n *GameNetwork) JoinGame(options JoinGameOptions) error {
	hub, ok := n.hubs[options.GameKey]
	if !ok {
		return fmt.Errorf("game key '%s' does not exist", options.GameKey)
	}
	return hub.Join(options)
}

func (n *GameNetwork) GetBGN(gameKey, gameID string) (*bgn.Game, error) {
	hub, ok := n.hubs[gameKey]
	if !ok {
		return nil, fmt.Errorf("game key '%s' does not exist", gameKey)
	}
	game, ok := hub.games[gameID]
	if !ok {
		return nil, fmt.Errorf("game id '%s' already exists", gameID)
	}
	return game.game.GetBGN(), nil
}
