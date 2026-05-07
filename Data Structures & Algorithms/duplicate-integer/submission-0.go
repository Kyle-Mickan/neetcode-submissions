func hasDuplicate(nums []int) bool {
    dupes := false

    for i := 0; i < len(nums); i++{
        for j := 0; j < len(nums); j++{
            if i == j {
                continue
            }
            if nums[i] == nums[j] {
                dupes = true
            }
        }
    }

    return dupes
}
