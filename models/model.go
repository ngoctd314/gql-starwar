package models

// CharacterFields ...
type CharacterFields struct {
	ID        string
	Name      string
	Fields    []string
	AppearsIn []Episode
}

// Human ...
type Human struct {
	CharacterFields
	StarshipIds  []string
	HeightMeters float64
	Mass         float64
}

// func (h *Human) Height(unit Len )

// IsCharacter ...
func (Human) IsCharacter()    {}
func (Human) IsSearchResult() {}

// Review ...
type Review struct{}

// Droid ...
type Droid struct{}

// IsCharacter ...
func (Droid) IsCharacter() {}

// IsSearchResult ...
func (Droid) IsSearchResult() {}

// FriendsConnection ...
type FriendsConnection struct {
	Ids  []string
	From int
	To   int
}
