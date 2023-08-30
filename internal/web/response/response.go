package response

type SlugsResp struct {
	ID    int32    `json:"id"`
	Slugs []string `json:"slugs"`
}
