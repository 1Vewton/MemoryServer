package knowledge

// Piece defines single piece of knowledge for agent
type Piece struct {
	Content   string
	BaseID    string
	SourceURL string
	Uploader  string
}

// NewKnowledgePiece creates new knowledge
func NewKnowledgePiece(
	content string,
	baseID string,
	sourceURL string,
	uploader string,
) *Piece {
	return &Piece{
		Content:   content,
		BaseID:    baseID,
		SourceURL: sourceURL,
		Uploader:  uploader,
	}
}
