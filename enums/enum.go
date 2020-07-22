package enums

import (
	"sort"
)

const Rate float32 = 4

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

// 附件类型
const (
	EnclosureTypeZero uint32 = iota
	// 整备单据
	EnclosureTypeRecondition
)

func FuelTypeText(code int) string {
	return fuelTypeText[code]
}

func FuelTypeList() []int {
	return getKeys(fuelTypeText)
}

func getKeys(maps map[int]string) []int {
	keys := make([]int, 0, len(maps))
	for mapKey, _ := range maps {
		keys = append(keys, mapKey)
	}
	sort.Ints(keys)

	return keys
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
func ColorList() []int {
	return getKeys(colorText)
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

func TransmissionList() []int {
	return getKeys(transmissionText)
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

func RegistrationTypeList() []int {
	return getKeys(registrationTypeText)
}

var registrationTypeCode = map[string]int{
	"Private": RegistrationTypePrivate,
	"Company": RegistrationTypeCompany,
}

func RegistrationTypeCode(text string) int {
	return registrationTypeCode[text]
}

// 库存状态：1在库，2离库
const (
	InventoryStatusIn  = 1
	InventoryStatusNot = 2
)

// 车辆状态：100 待整备，110整备中，120待核查，200待交易，210交易中，220已出售
const (
	CarStateToDoRecondition = 100
	CarStateReconditioning  = 110
	CarStateUnchecked       = 120
	CarStateWaitingTrade    = 200
	CarStateTrading         = 210
	CarStateTraded          = 220
)

// 上架状态：1未上架，2已上架
const (
	SaleStatusOffline = 1
	SaleStatusOnline  = 2
)

// 经销商类型：2 Carsome，3 Partner
const (
	// 经销商类型：0自营
	DealerTypeCarsome = 2
	// 经销商类型：1第三方
	DealerTypePartner = 3
)

// 枚举类型名称，对应car_enums_config表中的type字段
const (
	// 车体类型
	EnumNameBodyType = "BodyType"
	// 颜色
	EnumNameColor = "Color"
	// 变速箱类型：手动、自动
	EnumNameTransmission = "Transmission"
	// 几座车
	EnumNameSeat = "Seat"
	// 注册类型
	EnumNameRegistrationType = "RegistrationType"
	// 车辆状态
	EnumNameCarState = "CarState"
	// 库存状态
	EnumNameInventoryStatus = "InventoryStatus"
	// 检测点状态类型
	EnumNameReconditionPointType = "ReconditionPointType"
	// 上架状态
	EnumNameSaleStatus = "SaleStatus"
	// 经销商类型
	EnumNameCarDealerType = "CarDealerType"
)

// 车辆资源类型枚举
const (
	// 车辆资源：内部照片
	CarRes uint32 = iota
	CarResImgInner
	// 车辆资源：外部照片
	CarResImgOuter
	// 整备单据
	CarResReconditionFiles = 10
)
