package category

type Category struct {
	id   CategoryID
	name CategoryName
}

func NewCategory(categoryID CategoryID, categoryName CategoryName) (*Category, error) {

	err := NewCategoryID(categoryID)
	if err != nil {
		return nil, err
	}

	err = NewCategoryName(categoryName)
	if err != nil {
		return nil, err
	}

	return &Category{
		id:   categoryID,
		name: categoryName,
	}, nil

}
