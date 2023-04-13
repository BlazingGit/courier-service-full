package service

import (
	"fmt"
	"log"

	"example.com/courier-service/model"
)

func CalculateDeliveryCost(deliveryInfo *model.DeliveryInfo) {
	if deliveryInfo.CalculationMode == model.Undefined {
		log.Fatal("Calculation mode is Undefined...")
		return
	}

	fmt.Print("\nCalculating Delivery Cost...\n\n")
	for idx, pkg := range deliveryInfo.PackageDetailList {
		var discount, deliveryCost float64
		deliveryCost = float64(deliveryInfo.BaseDeliveryCost) + (float64(pkg.PkgWeight) * 10) + (float64(pkg.Distance) * 5)
		discount = calculateDiscount(deliveryCost, &pkg, deliveryInfo.CouponMap)
		deliveryCost -= discount
		pkg.DeliveryCost = deliveryCost
		pkg.Discount = discount
		deliveryInfo.PackageDetailList[idx] = pkg
	}

	if deliveryInfo.CalculationMode == model.DeliveryCost {
		fmt.Print("\n*****Final Result*****\n")
		for _, pkg := range deliveryInfo.PackageDetailList {
			fmt.Println(pkg.PkgId, pkg.Discount, pkg.DeliveryCost)
		}
	}
}

func calculateDiscount(deliveryCost float64, pkgDetail *model.PackageDetail, couponMap map[string]model.Coupon) (result float64) {
	coupon, couponExist := couponMap[pkgDetail.OfferCode]
	if couponExist {
		if pkgDetail.PkgWeight <= coupon.MaxWeight && pkgDetail.PkgWeight >= coupon.MinWeight && pkgDetail.Distance <= coupon.MaxDistance && pkgDetail.Distance >= coupon.MinDistance {
			result = deliveryCost * float64(coupon.DiscountPerc) / 100
			fmt.Printf("%v : Calculated discount is %v.\n", pkgDetail.PkgId, result)
		} else {
			fmt.Println(pkgDetail.PkgId, ": Weight or distance does not meet coupon", pkgDetail.OfferCode, "criteria.")
		}
	} else {
		fmt.Println(pkgDetail.PkgId, ": Coupon with offer code", pkgDetail.OfferCode, "does not exist.")
	}
	return result
}
