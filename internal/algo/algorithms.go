package algo

// Bubble sorting algorithm
func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
        for j := 0; j < len(arr)-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
	return arr
}


// Selection sorting algorithm
func SelectionSort(arr []int) []int{
	 for i := 0; i < len(arr)-1; i++ {
        minIndex := i
        for j := i + 1; j < len(arr); j++ {
            if arr[j] < arr[minIndex] {
                arr[j], arr[minIndex] = arr[minIndex], arr[j]
            }
        }
    }
	return arr
}

func InsertionSort(arr []int) []int{
	for i := 1; i < len(arr); i++ {
        for j := 0; j < i; j++ {
            if arr[j] > arr[i] {
                arr[j], arr[i] = arr[i], arr[j]
            }
        }
    }
	return arr
}