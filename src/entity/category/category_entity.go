package category

type Category struct {
	id   CategoryID
	name CategoryName
}

func NewCategory(categoryID CategoryID, categoryName CategoryName) (*Category, error) {

	return &Category{
		id:   categoryID,
		name: categoryName,
	}, nil

}
