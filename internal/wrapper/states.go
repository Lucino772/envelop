package wrapper

type ServerState struct {
	Status  ServerState_Status  `json:"status"`
	Players ServerState_Players `json:"players"`
}

type ServerState_Status struct {
	Description string `json:"description"`
}

type ServerState_Players struct {
	Count int                  `json:"count"`
	Max   int                  `json:"max"`
	List  []ServerState_Player `json:"list"`
}

type ServerState_Player struct {
	Id         string         `json:"id"`
	Attributes map[string]any `json:"attributes"`
}
