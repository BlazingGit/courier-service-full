package service

import (
	"fmt"
	"log"

	"example.com/courier-service/model"
)

func GetCalculationMode(deliveryInfo *model.DeliveryInfo, mode model.CalculationMode) {
	if mode != model.Undefined {
		deliveryInfo.CalculationMode = mode

	} else {
		var modeInput model.CalculationMode
		fmt.Println("Press 1 and enter to calculate delivery cost only.")
		fmt.Println("Press 2 and enter to calculate delivery cost and delivery time.")
		_, err := fmt.Scan(&modeInput)

		if err != nil || (modeInput != model.DeliveryCost && modeInput != model.DeliveryCostAndTime) {
			fmt.Print("Invalid input...\n\n")
			err = nil
			GetCalculationMode(deliveryInfo, model.Undefined)
		} else {
			deliveryInfo.CalculationMode = modeInput
		}
	}
}

func AddCoupon(deliveryInfo *model.DeliveryInfo, testCoupon *model.Coupon) {
	if testCoupon != nil {
		_, couponExist := deliveryInfo.CouponMap[testCoupon.OfferCode]
		if couponExist {
			log.Fatal("Coupon", testCoupon.OfferCode, "already exist!")
		} else {
			deliveryInfo.CouponMap[testCoupon.OfferCode] = *testCoupon
		}
	} else {
		var addCoupon int
		var offerCodes string

		for _, value := range deliveryInfo.CouponMap {
			offerCodes = offerCodes + " " + value.OfferCode
		}

		fmt.Print("\nThere are ", len(deliveryInfo.CouponMap), " coupon currently ("+offerCodes+" )\nPress 1 and enter to insert another coupon. Enter other key to skip.\n")
		_, err := fmt.Scan(&addCoupon)

		if err != nil || addCoupon != 1 {
			fmt.Print("No more coupon to add...\n\n")
		} else {
			var offerCode string
			var discPerc, minDist, maxDist, minWeight, maxWeight int
			fmt.Println("Please enter coupon detail in 'OfferCode DiscountPercentage MinDistance MaxDistance MinWeight MaxWeight' format.")
			fmt.Println("Example: OFR004 20 251 300 200 400")

			_, couponErr := fmt.Scan(&offerCode, &discPerc, &minDist, &maxDist, &minWeight, &maxWeight)

			if couponErr != nil {
				fmt.Print("Coupon input is invalid...\n\n")
				couponErr = nil
				AddCoupon(deliveryInfo, nil)
			} else {
				_, couponExist := deliveryInfo.CouponMap[offerCode]
				if couponExist {
					fmt.Print("Coupon OfferCode " + offerCode + " already exist, please enter different OfferCode...\n\n")
					couponErr = nil
					AddCoupon(deliveryInfo, nil)
				} else {
					deliveryInfo.CouponMap[offerCode] = model.Coupon{
						OfferCode:    offerCode,
						DiscountPerc: discPerc,
						MinDistance:  minDist,
						MaxDistance:  maxDist,
						MinWeight:    minWeight,
						MaxWeight:    maxWeight,
					}
					couponErr = nil
					AddCoupon(deliveryInfo, nil)
				}
			}
		}
	}
}

func GetInitialInput(deliveryInfo *model.DeliveryInfo) {
	var baseDeliveryCost, noOfPackages int
	fmt.Print("\nPlease enter Base Delivery Cost and Number of Package separated by space:\n")
	_, err := fmt.Scan(&baseDeliveryCost, &noOfPackages)

	if err != nil {
		fmt.Print("Base Delivery Cost and Number of Package must be a number...\n\n")
		err = nil
		GetInitialInput(deliveryInfo)

	} else {
		deliveryInfo.BaseDeliveryCost = baseDeliveryCost
		deliveryInfo.NoOfPackages = noOfPackages
	}
}

func GetPkgListInput(deliveryInfo *model.DeliveryInfo) {
	var pkgDetailList = []model.PackageDetail{}
	var pkgId, offerCode string
	var pkgWeight, distance int
	fmt.Print("\nPlease enter ", deliveryInfo.NoOfPackages, " lines of package details:\n")

	for i := 0; i < deliveryInfo.NoOfPackages; i++ {
		_, err := fmt.Scan(&pkgId, &pkgWeight, &distance, &offerCode)

		if err != nil {
			fmt.Print("Package detail not in \"string int int string\" format, please enter the list again...\n\n")
			pkgDetailList = []model.PackageDetail{}
			i = -1
			err = nil
		} else {
			pkgDetailList = append(pkgDetailList, model.PackageDetail{PkgId: pkgId, PkgWeight: pkgWeight, Distance: distance, OfferCode: offerCode})
		}
	}
	deliveryInfo.PackageDetailList = pkgDetailList
}

func GetFinalInput(deliveryInfo *model.DeliveryInfo) {
	var noOfVehicle, maxSpeed, maxCarryWeight int
	fmt.Print("\nPlease enter Number of Vehicle, Max Speed, and Max Carriable Weight separated by space:\n")
	_, err := fmt.Scan(&noOfVehicle, &maxSpeed, &maxCarryWeight)

	if err != nil {
		fmt.Print("Number of Vehicle, Max Speed, and Max Carriable Weight must be a number...\n\n")
		err = nil
		GetFinalInput(deliveryInfo)
	} else {
		deliveryInfo.NoOfVehicle = noOfVehicle
		deliveryInfo.MaxSpeed = maxSpeed
		deliveryInfo.MaxCarryWeight = maxCarryWeight
	}
}
