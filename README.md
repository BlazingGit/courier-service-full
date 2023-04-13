# courier-service-full

<b>--- Change Log ---</b>
1. Removed global variables and introduced a main DeliveryInfo object to store delivery parameters.
2. Restructured the solution and separated the logic to multiple files by adding service layer.
3. Combined the solution for problem 1 and 2 by giving user a choice to calculate cost only or calculate both cost and time.
4. Added option to insert new coupons.
5. Added more test cases to cover more scenario for both problem.
<b>------------------</b>

<b>----- Note -----</b>
For the delivery time calculation of below example input. It should produce the below output (similar to my previous solution).
As the requirement assumed the destinations are on a single route. Hence in the same trip, packages that are at the same distance should arrived at the same time.

Input:  
100 3  
PKG1 75 5 OFR001  
PKG2 15 5 OFR002  
PKG3 10 100 OFR003  
1 40 100  

Output:  
PKG1 87.5 787.5 0.12  
PKG2 0 275 0.12  
PKG3 35 665 2.5  

<b>----------------</b>

<b>-- Introduction --</b>
The project is written with Go.

You can run the program by opening the executable main.exe.

If Go is installed in your machine, you can run below command to execute the program:
go run ./main.go

At the main_test.go file, function getTestDataList() is where I put the test input and output list.
Run 'go test -v' to execute the test function.

I also attached few sample screenshots of the program in the screenshot folder.

Sample Input:  
100 5  
PKG1 50 30 OFR001  
PKG2 75 125 OFR008  
PKG3 175 100 OFR003  
PKG4 110 60 OFR002  
PKG5 155 95 NA  
2 70 200  

Sample Output:  
PKG1 0 750 3.98  
PKG2 0 1475 1.78  
PKG3 0 2350 1.42  
PKG4 105 1395 0.85  
PKG5 0 2125 4.19  


<b>-- Thought Process --</b>
The first problem require us to calculate the delivery cost based on the given weight, distance, and discount coupon. Weight and distance are straight forward. To calculate correctly based on the coupon offerId, I store the information of the coupon into a Map so I can retrieve it easily based on the user's input. Then I check if the given weight/distance matched the coupon's criteria, if yes then I applied the discount to the delivery cost calculation.

For the second problem, vehicle, maximum speed, and maximum carriable weight are introduced to calculate the time packages would be delivered. In addition to the problem 1 calculation, I first find out what could be the possible combination of packages that falls within the max carriable weight. I've decided to set the optimum number of packages per combination/trip is (Number of Packages/Number of vehicle). Example if there are 9 packages and 2 vehicle available, at most 1 vehicle can carry 5 packages, given that the weight is within the max carriable weight.

After that, I will sort the combination by weight in descending order, if the weights are the same, then I will sort them by distance in ascending order. Then I will loop the package combination to calculate the delivery time. Need to take note that the time taken for the vehicle to be available is getting the furthest destination devide by the max speed then multiply by 2. If 2 packages located at the same distance, I will assume they are delivered at the same time.

Thank you~
