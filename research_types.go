package main

type ResearchReport struct {
	Topic     string
	Summary   string
	Findings  []Finding
	Conflicts []string
	Angles    []string
	Sources   []Source
}

type Finding struct {
	Claim     string
	SourceIDs []string
}

type Source struct {
	ID    string
	Title string
	URL   string
}
