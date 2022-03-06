package model

const (
	DRAFT = iota
	DELETED
	PUBLISH
)

func StatusString(status int) string {
	switch status {
	case DRAFT:
		return "draft"
	case DELETED:
		return "deleted"
	case PUBLISH:
		return "publish"
	}

	return ""
}

func StatusDict() map[string]bool {
	return map[string]bool{
		"draft":   true,
		"deleted": true,
		"publish": true,
	}
}
