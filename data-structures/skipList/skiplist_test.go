package skipList

import (
	"fmt"
	"testing"
)

func TestSkipList(t *testing.T) {
	newSkipList := Constructor()
	newSkipList.Add(1)
	newSkipList.Add(2)
	newSkipList.Add(3)
	newSkipList.Add(4)
	newSkipList.Add(5)
	newSkipList.Add(6)
	newSkipList.Add(7)
	newSkipList.Add(8)
	newSkipList.Add(9)
	newSkipList.Add(10)
	fmt.Println(newSkipList.Search(11))
}
