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

func StatusDict() map[string]int {
	return map[string]int{
		"draft":   DRAFT,
		"deleted": DELETED,
		"publish": PUBLISH,
	}
}
