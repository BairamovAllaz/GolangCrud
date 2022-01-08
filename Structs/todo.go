package structs

type Todolist struct {
	Id           int    `json:"id"`
	Title        string `json:"title" binding:"required"`
	Desctription string `json:"desctription"`
}
type Userlist struct {
	Id     int
	Userid string
	Listid string
}
type Todoitem struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Desctription string `json:"desctription"`
	Done         bool   `json:"done"`
}
type Listitem struct {
	Id     int
	ListId int
	ItemId int
}

type UpdateListItem struct{ 
	Title *string `json:"title"`
	Desctription *string `json:"desctription"`
}
