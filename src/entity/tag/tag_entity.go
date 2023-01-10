package tag

type Tag struct {
	id   TagID
	name TagName
}

func NewTag(id TagID, name TagName) (*Tag, error) {

	return &Tag{
		id:   id,
		name: name,
	}, nil

}
