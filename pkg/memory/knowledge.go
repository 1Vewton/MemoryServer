package memory

// Knowledge defines single knowledge for agent
type Knowledge struct {
	Content   string
	Field     string
	SourceURL string
	Uploader  string
}

// NewKnowledge creates new knowledge
func NewKnowledge(
	content string,
	field string,
	sourceURL string,
	uploader string,
) *Knowledge {
	return &Knowledge{
		Content:   content,
		Field:     field,
		SourceURL: sourceURL,
		Uploader:  uploader,
	}
}
