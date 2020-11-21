package player

// Avatar
type Avatar struct {

	// URL avatar du joueur
	Url string
}

// Jouer
// Example : Player 1: {Bob 10 {} }
type Player struct {
	Name     string
	Age      int
	Avatar   Avatar
	password string
}

// Create a new player
func New(name string) Player {
	return Player{
		Name:     name,
		password: "defaultpassword",
	}
}
