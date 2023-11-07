package crew

type CrewItem struct {
	Id        uint64 `json:"id"`
	Name      string `json:"name"`
	Birthdate string `json:"birth_date"`
	Photo     string `json:"photo"`
}

type Character struct {
	IdActor       uint64 `json:"actor_id"`
	ActorPhoto    string `json:"actor_photo"`
	NameActor     string `json:"actor_name"`
	NameCharacter string `json:"character_name"`
}
