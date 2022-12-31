package category

type Category struct {
	id           CategoryID
	categoryName CategoryName
}

func NewCategory(id, name string) (*Category, error) {

	cID, err := CategoryIDValidate(id)
	if err != nil {
		return nil, err
	}

	cName, err := CategoryNameValidate(name)
	if err != nil {
		return nil, err
	}

	return &Category{
		id:           *cID,
		categoryName: *cName,
	}, nil

}
