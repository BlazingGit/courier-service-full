package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"example.com/courier-service/model"
	"example.com/courier-service/service"
)

func TestCalculateDeliveryCostAndTime(t *testing.T) {
	testDataList := getTestDataList()

	for i, testData := range testDataList {
		fmt.Print("\n### Processing Test Case ", (i + 1), " ###\n")
		processTestData(&testData)

		service.CalculateDeliveryCost(&testData.DeliveryInfo)
		service.CalculateDeliveryTime(&testData.DeliveryInfo)

		fmt.Print("\n***Comparing output and expected***\n")
		for j, expected := range testData.ExpectedOutput {
			output := testData.DeliveryInfo.PackageDetailList[j].PkgId +
				" " +
				strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.2f", testData.DeliveryInfo.PackageDetailList[j].Discount), "0"), ".") +
				" " +
				strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.2f", testData.DeliveryInfo.PackageDetailList[j].DeliveryCost), "0"), ".")

			if testData.DeliveryInfo.CalculationMode == model.DeliveryCostAndTime {
				output = output +
					" " +
					strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.2f", testData.DeliveryInfo.PackageDetailList[j].DeliveryTime), "0"), ".")
			}

			if expected != output {
				t.Errorf("Dataset %v output %v not equal to expected %v", i, output, expected)
			} else {
				fmt.Println(output, "passed.")
			}
		}
	}
}

func processTestData(testData *model.TestData) {
	pkgDetailList := []model.PackageDetail{}

	for i, input := range testData.Input {
		s := strings.Split(input, " ")
		if i == 0 {
			base, _ := strconv.Atoi(s[0])
			testData.DeliveryInfo.BaseDeliveryCost = base
			noOfPkg, _ := strconv.Atoi(s[1])
			testData.DeliveryInfo.NoOfPackages = noOfPkg

		} else if testData.DeliveryInfo.CalculationMode == model.DeliveryCostAndTime && i == (len(testData.Input)-1) {
			vehicleNo, _ := strconv.Atoi(s[0])
			testData.DeliveryInfo.NoOfVehicle = vehicleNo
			speed, _ := strconv.Atoi(s[1])
			testData.DeliveryInfo.MaxSpeed = speed
			maxWeight, _ := strconv.Atoi(s[2])
			testData.DeliveryInfo.MaxCarryWeight = maxWeight

		} else {
			weight, _ := strconv.Atoi(s[1])
			distance, _ := strconv.Atoi(s[2])
			pkgDetailList = append(pkgDetailList, model.PackageDetail{PkgId: s[0], PkgWeight: weight, Distance: distance, OfferCode: s[3]})
		}
	}
	testData.DeliveryInfo.PackageDetailList = pkgDetailList
}

func getTestDataList() []model.TestData {
	testDataList := []model.TestData{}
	var deliveryInfo *model.DeliveryInfo
	var input, expectedOutput []string

	//TestCase1: Following the given Delivery Cost sample
	deliveryInfo = new(model.DeliveryInfo)
	deliveryInfo.CalculationMode = model.DeliveryCost
	model.InitializeCouponMap(deliveryInfo)
	input = []string{
		"100 3",
		"PKG1 5 5 OFR001",
		"PKG2 15 5 OFR002",
		"PKG3 10 100 OFR003",
	}
	expectedOutput = []string{
		"PKG1 0 175",
		"PKG2 0 275",
		"PKG3 35 665",
	}
	testDataList = append(testDataList, model.TestData{DeliveryInfo: *deliveryInfo, Input: input, ExpectedOutput: expectedOutput})

	//TestCase2: Following the given Delivery Time sample
	deliveryInfo = new(model.DeliveryInfo)
	deliveryInfo.CalculationMode = model.DeliveryCostAndTime
	model.InitializeCouponMap(deliveryInfo)
	input = []string{
		"100 5",
		"PKG1 50 30 OFR001",
		"PKG2 75 125 OFR008",
		"PKG3 175 100 OFR003",
		"PKG4 110 60 OFR002",
		"PKG5 155 95 NA",
		"2 70 200",
	}
	expectedOutput = []string{
		"PKG1 0 750 3.98",
		"PKG2 0 1475 1.78",
		"PKG3 0 2350 1.42",
		"PKG4 105 1395 0.85",
		"PKG5 0 2125 4.19",
	}
	testDataList = append(testDataList, model.TestData{DeliveryInfo: *deliveryInfo, Input: input, ExpectedOutput: expectedOutput})

	//TestCase3: Test with more packages
	deliveryInfo = new(model.DeliveryInfo)
	deliveryInfo.CalculationMode = model.DeliveryCostAndTime
	model.InitializeCouponMap(deliveryInfo)
	input = []string{
		"100 8",
		"PKG1 50 30 OFR001",
		"PKG2 75 125 OFR008",
		"PKG3 175 100 OFR003",
		"PKG4 110 60 OFR002",
		"PKG5 155 95 NA",
		"PKG6 50 20 NA",
		"PKG7 50 30 NA",
		"PKG8 50 40 NA",
		"2 70 200",
	}
	expectedOutput = []string{
		"PKG1 0 750 0.42",
		"PKG2 0 1475 1.78",
		"PKG3 0 2350 2.56",
		"PKG4 105 1395 0.85",
		"PKG5 0 2125 4.91",
		"PKG6 0 700 0.28",
		"PKG7 0 750 0.42",
		"PKG8 0 800 0.57",
	}
	testDataList = append(testDataList, model.TestData{DeliveryInfo: *deliveryInfo, Input: input, ExpectedOutput: expectedOutput})

	//TestCase4: Test single vehicle
	deliveryInfo = new(model.DeliveryInfo)
	deliveryInfo.CalculationMode = model.DeliveryCostAndTime
	model.InitializeCouponMap(deliveryInfo)
	input = []string{
		"100 3",
		"PKG1 75 5 OFR001",
		"PKG2 15 5 OFR002",
		"PKG3 10 100 OFR003",
		"1 40 100",
	}
	expectedOutput = []string{
		"PKG1 87.5 787.5 0.12",
		"PKG2 0 275 0.12",
		"PKG3 35 665 2.5",
	}
	testDataList = append(testDataList, model.TestData{DeliveryInfo: *deliveryInfo, Input: input, ExpectedOutput: expectedOutput})

	//TestCase5: Test with new Coupon
	deliveryInfo = new(model.DeliveryInfo)
	deliveryInfo.CalculationMode = model.DeliveryCostAndTime
	model.InitializeCouponMap(deliveryInfo)
	service.AddCoupon(deliveryInfo, &model.Coupon{
		OfferCode:    "OFR004",
		DiscountPerc: 50,
		MinDistance:  100,
		MaxDistance:  200,
		MinWeight:    100,
		MaxWeight:    200,
	})
	input = []string{
		"100 4",
		"PKG1 75 5 OFR001",
		"PKG2 15 5 OFR002",
		"PKG3 10 100 OFR003",
		"PKG4 100 100 OFR004",
		"2 40 100",
	}
	expectedOutput = []string{
		"PKG1 87.5 787.5 0.12",
		"PKG2 0 275 0.12",
		"PKG3 35 665 2.5",
		"PKG4 800 800 2.5",
	}
	testDataList = append(testDataList, model.TestData{DeliveryInfo: *deliveryInfo, Input: input, ExpectedOutput: expectedOutput})

	return testDataList
}
