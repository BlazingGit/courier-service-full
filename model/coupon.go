package model

type Coupon struct {
	OfferCode    string
	DiscountPerc int
	MinDistance  int
	MaxDistance  int
	MinWeight    int
	MaxWeight    int
}

func InitializeCouponMap(deliveryInfo *DeliveryInfo) {

	deliveryInfo.CouponMap = make(map[string]Coupon)

	deliveryInfo.CouponMap["OFR001"] = Coupon{OfferCode: "OFR001", DiscountPerc: 10, MinDistance: 0, MaxDistance: 199, MinWeight: 70, MaxWeight: 200}
	deliveryInfo.CouponMap["OFR002"] = Coupon{OfferCode: "OFR002", DiscountPerc: 7, MinDistance: 50, MaxDistance: 150, MinWeight: 100, MaxWeight: 250}
	deliveryInfo.CouponMap["OFR003"] = Coupon{OfferCode: "OFR003", DiscountPerc: 5, MinDistance: 50, MaxDistance: 250, MinWeight: 10, MaxWeight: 150}
}
