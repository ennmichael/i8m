package engine

import "encoding/json"

// Player represents a player in the game.
type Player struct {
	Position  Vector
	Direction Vector
	firedRail int // XXX Incomplete
}

// move moves the player for distance d.
func (p *Player) move(d float64) {
	p.Position = p.Position.Translate(p.Direction.Scale(d))
}

// FireRail causes the player to fire a rail.
func (p *Player) FireRail() {
}

// Die kills the player.
func (p *Player) Die() {
}

// Engine is the game engine. It holds the gameplay settings and the current game state.
type Engine struct {
	Timestep    float64 // Engine timestep in milliseconds.
	PlayerSpeed float64 // Player speed in pixels per millisecond.
	players     map[string]*Player
}

// NewEngine creates a new engine.
func NewEngine(Timestep, PlayerSpeed float64) *Engine {
	return &Engine{
		Timestep:    Timestep,
		PlayerSpeed: PlayerSpeed,
		players:     make(map[string]*Player),
	}
}

// Update advances the game engine state. deltaTime should be the time elapsed
// since the previous call to update. This is required because the engine operates
// within fixed timesteps. The remaining time that doesn't fall within the timestep is
// returned. This value will always be smaller than e.Timestep.
func (e *Engine) Update(deltaTime float64) (remainingTime float64) {
	for deltaTime >= e.Timestep {
		deltaTime -= e.Timestep
		for _, player := range e.players {
			player.move(e.PlayerSpeed * e.Timestep)
		}
	}
	return deltaTime
}

// AddPlayer adds a new player to the simulation.
func (e *Engine) AddPlayer(id string, p *Player) {
	e.players[id] = p
}

// RemovePlayer removes a player from the simulation.
func (e *Engine) RemovePlayer(id string) {
	delete(e.players, id)
}

// StateJSON returns a JSON dump of engine state.
func (e *Engine) StateJSON() (state []byte, err error) {
	state, err = json.Marshal(e.players)
	return
}
