/*
 * https://leetcode.com/problems/maximum-units-on-a-truck/
 *
 * 1710. Maximum Units on a Truck
 * Easy
 *
 * Runtime: 40 ms
 * Memory Usage: 6.5 MB
 *
 * You are assigned to put some amount of boxes onto one truck.
 * You are given a 2D array boxTypes, where boxTypes[i] = [numberOfBoxesi, numberOfUnitsPerBoxi]
 *
 * numberOfBoxesi is the number of boxes of type i
 * numberOfUnitsPerBoxi is the number of units in each box of the type i
 *
 * You are also given an integer truckSize, which is the maximum number of boxes that can be put on the truck. You can choose any boxes to put on the truck as long as the number of boxes does not exceed truckSize.
 *
 * Return the maximum total number of units that can be put on the truck.
 *
 * Example 1:
 *
 * Input: boxTypes = [[1,3],[2,2],[3,1]], truckSize = 4
 * Output: 8
 *
**/
func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.SliceStable(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})

	count := 0
	for i := 0; i < len(boxTypes); i++ {
		if truckSize <= 0 {
			break
		}
		if truckSize > boxTypes[i][0] {
			count += boxTypes[i][0] * boxTypes[i][1]
			truckSize -= boxTypes[i][0]
		} else {
			count += truckSize * boxTypes[i][1]
			truckSize -= truckSize
		}
	}

	return count
}