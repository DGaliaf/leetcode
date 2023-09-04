package medium

func searchRange(nums []int, target int) []int {
    res := []int{-1,-1}

    start, end := 0, len(nums)-1
    for start <= end {
        mid := start+(end-start)/2

        if target > nums[mid] {
            start = mid+1;
        } else if target < nums[mid] {
            end = mid-1;
        } else {
            res[0] = mid;
            end = mid-1;
        }
    }

    start, end = 0, len(nums)-1
    for start <= end {
        mid := start+(end-start)/2

        if target > nums[mid] {
            start = mid+1;
        } else if target < nums[mid] {
            end = mid-1;
        } else {
            res[1] = mid;
            start = mid+1;
        }
    }

    return res
}