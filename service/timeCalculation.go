package service

import (
	"fmt"
	"math"
	"sort"

	"example.com/courier-service/model"
)

func CalculateDeliveryTime(deliveryInfo *model.DeliveryInfo) {
	if deliveryInfo.CalculationMode != model.DeliveryCostAndTime {
		fmt.Println("Skipping delivery time calculation...")
		return
	}

	fmt.Print("\nCalculating Delivery Time...\n")
	//Build all the possible package combination
	allPkgCombination := []model.PackageCombination{}
	allPkgCombination = loopPkgCombination(deliveryInfo, 0, 0, allPkgCombination, 0, 0, []string{})

	//Sort combination by highest totalWeight, if same weight, sort by lowest totalDistance
	sort.Slice(allPkgCombination, func(a, b int) bool {
		if allPkgCombination[a].TotalWeight == allPkgCombination[b].TotalWeight {
			return allPkgCombination[a].TotalDistance < allPkgCombination[b].TotalDistance
		} else {
			return allPkgCombination[a].TotalWeight > allPkgCombination[b].TotalWeight
		}
	})

	var vehicleList = []*model.Vehicle{}
	for i := 0; i < deliveryInfo.NoOfVehicle; i++ {
		vehicleList = append(vehicleList, &model.Vehicle{IsAvailable: false, DeliveryStartTime: 0})
	}

	//Loop through all the possible combination
	fmt.Println("Possible Package Combinations:")
	fmt.Println("[Total Weight] [Total Distance] [Packages]")
	var calculatedPackages = []string{}

	for _, pkgComb := range allPkgCombination {
		fmt.Println(pkgComb.TotalWeight, pkgComb.TotalDistance, pkgComb.PackageIDs)

		//Skip the combination if the package in the combination already calculated
		if len(calculatedPackages) > 0 && isPackageCalculated(pkgComb.PackageIDs, calculatedPackages) {
			continue
		}

		vehicleIdx := getNextAvailableVehicle(vehicleList) //Find out the first available vehicle and later get it's start time
		deliveryStartTime := vehicleList[vehicleIdx].DeliveryStartTime
		var longestDeliveryTime float64
		for _, pkgId := range pkgComb.PackageIDs {
			deliveryTime := setDeliveryTime(deliveryInfo, pkgId, deliveryStartTime)
			if deliveryTime > longestDeliveryTime {
				longestDeliveryTime = deliveryTime
			}
		}
		vehicleList[vehicleIdx].DeliveryStartTime = longestDeliveryTime * 2 //Set the vehicle next available time
		calculatedPackages = append(calculatedPackages, pkgComb.PackageIDs...)
		if len(calculatedPackages) == deliveryInfo.NoOfPackages {
			break
		}
	}

	if deliveryInfo.CalculationMode == model.DeliveryCostAndTime {
		fmt.Print("\n*****Final Result*****\n")
		for _, pkg := range deliveryInfo.PackageDetailList {
			fmt.Println(pkg.PkgId, pkg.Discount, pkg.DeliveryCost, pkg.DeliveryTime)
		}
	}
}

func loopPkgCombination(deliveryInfo *model.DeliveryInfo, loopIdx int, previousPkgIdx int, allPkgCombination []model.PackageCombination, sumOfWeight int, sumOfDistance int, pkgArray []string) []model.PackageCombination {
	pkgList := deliveryInfo.PackageDetailList[loopIdx:] //Only loop
	for pkgIdx, pkg := range pkgList {
		if loopIdx > 0 && (pkgIdx+loopIdx) <= previousPkgIdx { //So that we wont add the same PackageId or package combination in different order
			continue
		}
		var newSumOfWeight int = sumOfWeight + pkg.PkgWeight
		var newSumOfDistance int = sumOfDistance + pkg.Distance
		var newPkgArray []string = append(pkgArray, pkg.PkgId)
		if newSumOfWeight <= deliveryInfo.MaxCarryWeight {
			pkgCombination := model.PackageCombination{TotalWeight: newSumOfWeight, TotalDistance: newSumOfDistance, PackageIDs: newPkgArray}
			allPkgCombination = append(allPkgCombination, pkgCombination)
		}

		if loopIdx < (deliveryInfo.NoOfPackages / deliveryInfo.NoOfVehicle) {
			allPkgCombination = loopPkgCombination(deliveryInfo, (loopIdx + 1), (pkgIdx + loopIdx), allPkgCombination, newSumOfWeight, newSumOfDistance, newPkgArray)
		}
	}
	return allPkgCombination
}

func isPackageCalculated(currentPackages []string, calculatedPackages []string) bool {
	for _, calPkgId := range calculatedPackages {
		for _, curPkgId := range currentPackages {
			if curPkgId == calPkgId {
				return true
			}
		}
	}
	return false
}

func getNextAvailableVehicle(vehicleList []*model.Vehicle) int { //Return the index of the available vehicle
	var nearestDeliveryTime float64 = 0
	var nearestIdx int
	for idx, vehicle := range vehicleList {
		if vehicle.DeliveryStartTime == 0 { //If 0 means vehicle is available straight away
			return idx
		}
		if nearestDeliveryTime == 0 || vehicle.DeliveryStartTime < nearestDeliveryTime {
			nearestDeliveryTime = vehicle.DeliveryStartTime
			nearestIdx = idx
		}
	}
	return nearestIdx
}

func setDeliveryTime(deliveryInfo *model.DeliveryInfo, pkgId string, deliveryStartTime float64) float64 {
	for idx, pkgDetail := range deliveryInfo.PackageDetailList {
		if pkgDetail.PkgId == pkgId {
			var deliveryTime float64 = deliveryStartTime + (float64(pkgDetail.Distance) / float64(deliveryInfo.MaxSpeed))
			deliveryTime = math.Floor(deliveryTime*100) / 100
			pkgDetail.DeliveryTime = deliveryTime
			deliveryInfo.PackageDetailList[idx] = pkgDetail

			return deliveryTime
		}
	}
	return 0
}
