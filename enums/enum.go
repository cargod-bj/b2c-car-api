package enums

const (
	FuelTypePetrol         = 1
	FuelTypePetrolWithNGV  = 2
	FuelTypeDiesel         = 3
	FuelTypeDieselWithNGV  = 4
	FuelTypeHybridElectric = 5
)

var fuelTypeText = map[int]string{
	FuelTypePetrol:         "Petrol",
	FuelTypePetrolWithNGV:  "Petrol with NGV",
	FuelTypeDiesel:         "Diesel",
	FuelTypeDieselWithNGV:  "Diesel with NGV",
	FuelTypeHybridElectric: "Hybrid/Electric",
}

func FuelTypeText(code int) string {
	return fuelTypeText[code]
}

var fuelTypeCode = map[string]int{
	"Petrol":          FuelTypePetrol,
	"Petrol with NGV": FuelTypePetrolWithNGV,
	"Diesel":          FuelTypeDiesel,
	"Diesel with NGV": FuelTypeDieselWithNGV,
	"Hybrid/Electric": FuelTypeHybridElectric,
}

func FuelTypeCode(text string) int {
	return fuelTypeCode[text]
}

const (
	ColorBlack  = 1
	ColorBlue   = 2
	ColorBrown  = 3
	ColorGold   = 4
	ColorGray   = 5
	ColorGreen  = 6
	ColorOrange = 7
	ColorPurple = 8
	ColorRed    = 9
	ColorSliver = 10
	ColorTan    = 11
	ColorWhite  = 12
)

var colorText = map[int]string{
	ColorBlack:  "Black",
	ColorBlue:   "Blue",
	ColorBrown:  "Brown",
	ColorGold:   "Gold",
	ColorGray:   "Gray",
	ColorGreen:  "Green",
	ColorOrange: "Orange",
	ColorPurple: "Purple",
	ColorRed:    "Red",
	ColorSliver: "Sliver",
	ColorTan:    "Tan",
	ColorWhite:  "White",
}

func ColorText(code int) string {
	return colorText[code]
}

var colorCode = map[string]int{
	"Black":  ColorBlack,
	"Blue":   ColorBlue,
	"Brown":  ColorBrown,
	"Gold":   ColorGold,
	"Gray":   ColorGray,
	"Green":  ColorGreen,
	"Orange": ColorOrange,
	"Purple": ColorPurple,
	"Red":    ColorRed,
	"Sliver": ColorSliver,
	"Tan":    ColorTan,
	"White":  ColorWhite,
}

func ColorCode(text string) int {
	return colorCode[text]
}

const (
	TransmissionManual = 1
	TransmissionAuto   = 2
)

var transmissionText = map[int]string{
	TransmissionManual: "Manual",
	TransmissionAuto:   "Auto",
}

func TransmissionText(code int) string {
	return transmissionText[code]
}

var transmissionCode = map[string]int{
	"Manual": TransmissionManual,
	"Auto":   TransmissionAuto,
}

func TransmissionCode(text string) int {
	return transmissionCode[text]
}

const (
	RegistrationTypePrivate = 1
	RegistrationTypeCompany = 2
)

var registrationTypeText = map[int]string{
	RegistrationTypePrivate: "Private",
	RegistrationTypeCompany: "Company",
}

func RegistrationTypeText(code int) string {
	return registrationTypeText[code]
}

var registrationTypeCode = map[string]int{
	"Private": RegistrationTypePrivate,
	"Company": RegistrationTypeCompany,
}

func RegistrationTypeCode(text string) int {
	return registrationTypeCode[text]
}