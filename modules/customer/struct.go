package customer

type BindingDataCustomer struct {
	ID          int      `validate:"required" binding:"required" form:"id" json:"id"`
	Subdomain   string   `validate:"required" binding:"required" form:"subdomain" json:"subdomain"`
	Fullname    string   `validate:"required" binding:"required" form:"fullname" json:"fullname"`
	Mobile      string   `validate:"required" binding:"required" form:"mobile" json:"mobile"`
	Email       string   `form:"email" json:"email"`
	Gender      string   `form:"gender" json:"gender"`
	IdNumber    string   `form:"id_number" json:"id_number"`
	Location    Location `form:"location" json:"location"`
	Address     string   `form:"address" json:"address"`
	DateOfBirth string   `form:"date_of_birth" json:"date_of_birth"`
	Note        string   `form:"note" json:"note"`
	ImageFile   string   `validate:"required" binding:"required" form:"image_file" json:"image_file"`
	Images      Images   `validate:"required" binding:"required" form:"images" json:"images"`
	IsDeleted   bool     `form:"is_deleted" json:"is_deleted"`
}

type Images struct {
	Large  string `validate:"required" binding:"required" form:"large" json:"large"`
	Medium string `validate:"required" binding:"required" form:"medium" json:"medium"`
	Small  string `validate:"required" binding:"required" form:"small" json:"small"`
}

type Location struct {
	ID   int    `validate:"required" binding:"required" form:"id" json:"id"`
	Name string `validate:"required" binding:"required" form:"name" json:"name"`
}

type RequestDataCustomer struct {
	ID          int
	Subdomain   string
	Fullname    string
	Email       string
	Mobile      string
	Age         string
	Gender      string
	IdNumber    string
	Location    Location
	Address     string
	DateOfBirth string
	Note        string
	ImageFile   string
	Images      Images
	IsDeleted   bool
}
