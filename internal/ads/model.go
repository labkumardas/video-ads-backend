package ads

type Ad struct {
	ID        int    `json:"id"`
	ImageURL  string `json:"image_url"`
	TargetURL string `json:"target_url"`
}

type Click struct {
	AdID         int    `json:"ad_id"`
	IP           string `json:"ip"`
	PlaybackTime int    `json:"playback_time"`
}
