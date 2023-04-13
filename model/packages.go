package model

type PackageDetail struct {
	PkgId        string
	PkgWeight    int
	Distance     int
	OfferCode    string
	DeliveryCost float64
	DeliveryTime float64
	Discount     float64
}

type PackageCombination struct {
	TotalWeight   int
	TotalDistance int
	PackageIDs    []string
}
