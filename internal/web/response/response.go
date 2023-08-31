package response

type Slugs struct {
	ID    int32    `json:"id"`
	Slugs []string `json:"slugs"`
}

type UserChanges struct {
	ID      int32     `json:"id"`
	Removed *[]string `json:"removed"`
	Added   *[]string `json:"added"`
}

type Records struct {
	Link string `json:"link"`
}
