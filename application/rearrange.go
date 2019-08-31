package application

import "fmt"

type Case struct {
  Order int
  Name  string
}

var cases = []Case{
  {Order: 1, Name: "one"},
  {Order: 2, Name: "two"},
  {Order: 3, Name: "tri"},
  {Order: 4, Name: "fou"},
  {Order: 5, Name: "fiv"},
}

func rearrange(arr []Case, needle, gold int) []Case {
	temp_arr := make([]Case, len(arr))
	copy(temp_arr, arr)

	for v, _ := range temp_arr {
		if temp_arr[v].Order == needle{
			temp_arr[v].Name = arr[gold-1].Name
		}
		if temp_arr[v].Order == gold{
			temp_arr[v].Name = arr[needle-1].Name
		}
	}
	return temp_arr
}

func GetRearrange()  {
	cases = rearrange(cases, 4,2)
	cases = rearrange(cases, 1,2)
	fmt.Printf("Arr now = %v \n", cases)
}