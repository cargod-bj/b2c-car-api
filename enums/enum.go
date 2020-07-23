package enums

const Rate float32 = 0.04

const (
	color uint32 = iota
	ColorBlack
	ColorWhite
	ColorGray
	ColorSilver
	ColorRed
	ColorBlue
	ColorBrown
	ColorGold
	ColorGreen
	ColorOrange
	ColorBeige
	ColorPurple
	ColorBronze
	ColorOther
)

// v 是否是合法的transmission
func IsColor(v uint32) bool {
	return v == ColorBlack || v == ColorWhite ||
		v == ColorGray || v == ColorSilver ||
		v == ColorRed || v == ColorBlue ||
		v == ColorBrown || v == ColorGold ||
		v == ColorGreen || v == ColorOrange ||
		v == ColorBeige || v == ColorPurple ||
		v == ColorBronze || v == ColorOther
}

// 变速箱
const (
	// 变速箱：手动
	TransmissionManual = 1
	// 变速箱：自动
	TransmissionAuto = 2
)

// v 是否是合法的transmission
func IsTransmission(v uint32) bool {
	return v == TransmissionManual || v == TransmissionAuto
}

// 注册类型
const (
	// 注册类型: 私家
	RegistrationTypePrivate = 1
	// 注册类型: 公司
	RegistrationTypeCompany = 2
)

// v 是否是合法的RegistrationType
func IsRegistrationType(v uint32) bool {
	return v == RegistrationTypePrivate || v == RegistrationTypeCompany
}

// 库存状态：1在库，2离库
const (
	// 库存状态：在库
	InventoryStatusIn = 1
	// 库存状态：离库
	InventoryStatusNot = 2
)

// v 是否是合法的InventoryStatus
func IsInventoryStatus(v uint32) bool {
	return v == InventoryStatusIn || v == InventoryStatusNot
}

// 车辆状态：100 待整备，110整备中，120待核查，200待交易，210交易中，220已出售
const (
	// 车辆状态: 待整备
	CarStateToDoRecondition = 100
	// 车辆状态: 整备中
	CarStateReconditioning = 110
	// 车辆状态: 待核查
	CarStateUnchecked = 120
	// 车辆状态: 待交易
	CarStateWaitingTrade = 200
	// 车辆状态: 交易中
	CarStateTrading = 210
	// 车辆状态: 已出售
	CarStateTraded = 220
)

// v 是否是合法的CarState
func IsCarState(v uint32) bool {
	return v == CarStateToDoRecondition || v == CarStateReconditioning ||
		v == CarStateUnchecked || v == CarStateWaitingTrade ||
		v == CarStateTrading || v == CarStateTraded
}

// 上架状态：1未上架，2已上架
const (
	// 上架状态: 未上架
	SaleStatusOffline = 1
	// 上架状态: 已上架
	SaleStatusOnline = 2
)

// v 是否是合法的SaleStatus
func IsSaleStatus(v uint32) bool {
	return v == SaleStatusOffline || v == SaleStatusOnline
}

// 经销商类型：2 Carsome，3 Partner
const (
	// 经销商类型：自营
	DealerTypeCarsome = 2
	// 经销商类型：第三方
	DealerTypePartner = 3
)

// v 是否是合法的DealerType
func IsDealerType(v uint32) bool {
	return v == DealerTypeCarsome || v == DealerTypePartner
}

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
	carRes uint32 = iota
	// 车辆资源：内部照片
	CarResImgInner
	// 车辆资源：外部照片
	CarResImgOuter
	// 整备单据
	CarResReconditionFiles = 10
)

// v 是否是合法的CarRes
func IsCarRes(v uint32) bool {
	return v == CarResImgInner || v == CarResImgOuter || v == CarResReconditionFiles
}
