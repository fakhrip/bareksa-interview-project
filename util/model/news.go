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
