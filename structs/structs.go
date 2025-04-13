package structs

type ReqestGetAndLostHandler struct {
	Name     string `json:"name"`
	WhoFound string `json:"whoFound"`
	Where    string `json:"where"`
	Notes    string `json:"notes"`
}
