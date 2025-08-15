package entities

type Vendor struct {
	UserID          int
	Name            string
	BusinessEmail   string
	BusinessAddress string
	BusinessPhone   string
	BannerURL       string
	LogoURL         string
	Description     string
	Status          string
}

func NewVendor(
	userID int,
	name,
	businessEmail,
	businessAddress,
	businessPhone,
	description string,
) *Vendor {
	return &Vendor{
		UserID:          userID,
		Name:            name,
		BusinessEmail:   businessEmail,
		BusinessAddress: businessAddress,
		BusinessPhone:   businessPhone,
		Description:     description,
		Status:          "pending",
	}
}
