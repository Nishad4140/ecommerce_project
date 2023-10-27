package response

type Category struct {
	Id           int
	CategoryName string
}

type Brand struct {
	Id           int `json:",omitempty"`
	Name         string
	Description  string
	CategoryName string
}

type Model struct {
	Id           uint
	ModelName    string
	Description  string
	Brand        string
	CategoryName string
	Sku          string
	QtyInStock   int
	Color        string
	Ram          int
	Battery      int
	ScreenSize   float64
	Storage      int
	Camera       int
	Price        int
	Image        []string
}
