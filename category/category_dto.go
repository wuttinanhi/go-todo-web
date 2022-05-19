package category

// create category dto
type CategoryAddDTO struct {
	Name string `json:"name" binding:"required,max=255"`
}

// update category dto
type CategoryUpdateDTO struct {
	Id   int    `json:"id" binding:"required,gte=1,number"`
	Name string `json:"name" binding:"required,max=255"`
}

// delete category dto
type CategoryDeleteDTO struct {
	Id int `json:"id" binding:"required,gte=1,number"`
}
