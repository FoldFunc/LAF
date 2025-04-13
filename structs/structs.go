package structs

type ReqestAddAndLostHandler struct {
	Name     string `json:"name"`
	WhoFound string `json:"whoFound"`
	Where    string `json:"where"`
	Notes    string `json:"notes"`
}

type ReqestGetAndLostHandler struct {
	Name        string `json:"name"`
	WhoFound    string `json:"whoFound"`
	Where       string `json:"where"`
	Notes       string `json:"notes"`
	DateAdded   string `json:"dateAdded"`
	DateClaimed string `json:"dateClaimed"`
}
type LostAndFound struct {
	Name        string `json:"name"`
	WhoFound    string `json:"whoFound"`
	Where       string `json:"where"`
	Notes       string `json:"notes"`
	DateAdded   string `json:"dateAdded"`
	DateClaimed string `json:"dateClaimed"`
}

type ReqestClaimLostAndFound struct {
	Name         string `json:"name"`
	ClaimingName string `json:"claimingName"`
	DateClaimed  string `json:"dateClaimed"`
}
