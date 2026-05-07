func twoSum(nums []int, target int) []int {
    var indexes []int
    for i, value1 := range nums{
        for j, value2 := range nums{
            if i == j {
                continue
            }
            sum := value1 + value2
            if sum == target {
                indexes = append(indexes, i)
                indexes = append(indexes, j)
                return indexes
            }
            
        }
    }
    return indexes
}
