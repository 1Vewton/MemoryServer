package requester

// MediaWikiRequester defines the requester to media wiki
type MediaWikiRequester struct {
	url string
}

// NewMediaWikiRequester creates new requester for media wiki
func NewMediaWikiRequester(
	url string,
) *MediaWikiRequester {
	return &MediaWikiRequester{
		url: url,
	}
}
