package enums

const Rate float32 = 0.035

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
	// 燃油类型
	EnumNameFuelType = "FuelType"
	// 是否有某项配置
	EnumNameWithOrWithout = "WithOrWithout"
	// 车检类型
	EnumNamePuspakomItem = "PuspakomItem"
	// 活动状态
	EnumNameCampaignStatus = "CarCampaignStatus"
	// 整备类型
	EnumNameCarReconditionCategory = "CarReconditionCategory"
	// 保修是否有最大里程限制
	EnumNameWarrantyMaxMileageLimit = "WarrantyMaxMileageLimit"
)

// 燃油类型：fuelType
const (
	fuelType uint32 = iota
	FuelTypePetrol
	FuelTypePetrolWithNGV
	FuelTypeDiesel
	FuelTypeDieselWithNGV
	FuelTypeHybridElectric
)

// 车辆颜色
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

// 变速箱
const (
	// 变速箱：手动
	TransmissionManual = 1
	// 变速箱：自动
	TransmissionAuto = 2
)

// 注册类型
const (
	// 注册类型: 私家
	RegistrationTypePrivate = 1
	// 注册类型: 公司
	RegistrationTypeCompany = 2
)

// 库存状态：1在库，2离库
const (
	// 库存状态：在库
	InventoryStatusIn = 1
	// 库存状态：离库
	InventoryStatusNot = 2
)

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

// 上架状态：1未上架，2已上架
const (
	// 上架状态: 未上架
	SaleStatusOffline = 1
	// 上架状态: 已上架
	SaleStatusOnline = 2
)

// 经销商类型：2 Carsome，3 Partner
const (
	// 经销商类型：自营
	DealerTypeCarsome = 2
	// 经销商类型：第三方
	DealerTypePartner = 3
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

// 车辆 '是否有' 枚举
const (
	carWithOrWithout uint32 = iota
	// 车辆 有 某项资源
	CarWith
	// 车辆 没有 某项资源
	CarWithout
)

// 车检类型
const (
	carPuspakomItem uint32 = iota
	// 车检类型：B5
	PuspakomItemB5
	// 车检类型：B7
	PuspakomItemB7
)

// 车辆费用类型
const (
	carCostType uint32 = iota
	// 车辆费用类型: 整备
	CarCostTypeRecondition
	// 车辆费用类型: 保养
	CarCostTypeMaintenance
	// 车辆费用类型: 车检
	CarCostTypePuspakom
	//政府保险费用
	CarCostTypeGovernmentInsurance
	//道路税费用
	CarCostTypeRoadTax
)

// 车辆活动类型
const (
	carCampaignStatus uint32 = iota
	// 活动类型：未激活
	CarCampaignStatusInactive
	// 活动类型：激活
	CarCampaignStatusActive
	// 活动类型：无效
	CarCampaignStatusExpired
)

// 导入活动车辆失败
const (
	carCampaignImportFailed uint32 = iota
	// 导入失败：carNo错误
	CarCampaignImportFailedNo
	// 导入失败：carNo和LicensePlate不匹配
	CarCampaignImportFailedNoAndPlateNoMatch
	// 导入失败：车辆状态不支持活动
	CarCampaignImportFailedCarStateError
	// 导入失败：车辆价格为空
	CarCampaignImportFailedCarPriceZero
)

// 整备类型
const (
	carReconditionCategory uint32 = iota
	// 整备类型：Car service and repair
	CarReconditionCategoryServiceAndRepair
	// 整备类型：Cleaning/Auto spa
	CarReconditionCategoryCleanSpa
	// 整备类型：Painting
	CarReconditionCategoryPainting
)

// 保修是否有最大里程限制
const (
	carWarrantyMaxMileageLimit uint32 = iota
	// 保修是否有最大里程限制：有限制
	CarWarrantyMaxMileageLimited
	// 保修是否有最大里程限制：无限制
	CarWarrantyMaxMileageUnlimited
)
