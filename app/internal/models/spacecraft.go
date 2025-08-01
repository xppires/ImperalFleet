package models

type Armament struct {
	ID       int    `json:"id"`
	CraftID  int    `json:"spaceship_id"`
	Title    string `json:"title"`
	Quantity string `json:"quantity" db:"qty"`
}


// Spacecraft is the model representation of a spacecraft that is stored in the database.
type Spacecraft struct {
	ID       int        `json:"id"`
	Name     string     `json:"name"`
	Class    string     `json:"class"`
	Status   string     `json:"status"`
	Armament []Armament `json:"armament"`
	Image    string     `json:"image"`
	Crew     int        `json:"crew"`
	Value    float32      `json:"value"`
}


// SpacecraftRequest is a model that contains the minimum information required to register a new spacecraft.
type SpacecraftRequest struct {
	Name     string     `json:"name"`
	Class    string     `json:"class"`
	Status   string     `json:"status"`
	Armament []Armament `json:"armament"`
	Image    string     `json:"image"`
	Crew     int        `json:"crew"`
	Value    int        `json:"value"`
}

func (r SpacecraftRequest) Valid() bool {
	return r.Name != "" && r.Class != "" && r.Status != "" && len(r.Armament) != 0 && r.Image != "" && r.Crew != 0
}