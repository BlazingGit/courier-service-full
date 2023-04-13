package model

type CalculationMode int64

const (
	Undefined CalculationMode = iota
	DeliveryCost
	DeliveryCostAndTime
)

type DeliveryInfo struct {
	CalculationMode   CalculationMode
	BaseDeliveryCost  int
	NoOfPackages      int
	NoOfVehicle       int
	MaxSpeed          int
	MaxCarryWeight    int
	CouponMap         map[string]Coupon
	PackageDetailList []PackageDetail
}
