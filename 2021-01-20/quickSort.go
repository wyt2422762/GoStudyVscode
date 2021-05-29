package maximumproduct

func qs() []int {
	nums := []int{6, 1, 2, 7, 9, 11, 4, 5, 10, 8}	
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, low, high int) {

	if(low > high){
		return;
	}

	left, right := low, high

	temp := nums[left]

	for left < right{
		for left < right && nums[right] >= temp{
			right--
		}
		for left < right && nums[left] <= temp{
			left++
		}
	
		if left < right {
			//交换
			nums[left], nums[right] = nums[right], nums[left]
		}

	}

	nums[low], nums[left] = nums[left], nums[low]

	 //递归调用左半数组
	 quickSort(nums, low, right-1);
	 //递归调用右半数组
	 quickSort(nums, right+1, high);

}