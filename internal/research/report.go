package research

type Report struct {
	Topic     string    `json:"topic"`
	Summary   string    `json:"summary"`
	Findings  []Finding `json:"findings"`
	Conflicts []string  `json:"conflicts"`
	Angles    []string  `json:"angles"`
	Sources   []Source  `json:"sources"`
}

type Finding struct {
	Claim     string   `json:"claim"`
	SourceIDs []string `json:"source_ids"`
}

type Source struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
}
