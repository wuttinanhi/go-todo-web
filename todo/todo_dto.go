package todo

// create todo dto
type TodoAddDTO struct {
	Name       string `json:"name" binding:"required,max=255"`
	CategoryId uint   `json:"category_id" binding:"required,gte=1,number"`
}

// update todo dto
type TodoUpdteDTO struct {
	Id         int    `json:"id" binding:"required,gte=1,number"`
	Name       string `json:"name" binding:"required,max=255"`
	CategoryId uint   `json:"category_id" binding:"required,gte=1,number"`
}

// delete todo dto
type TodoDeleteDTO struct {
	Id int `json:"id" binding:"required,gte=1,number"`
}
