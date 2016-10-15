package dawg

import "testing"

func TestCreateDAWG(t *testing.T) {
	dawg := CreateDAWG([][]int{{1, 2, 3, 1}, {4, 2, 3, 1}, {5, 2, 3, 1}, {5, 6, 1, 2}})
	if dawg.nodesCount != 8 {
		t.Error("Creation failed")
	}
}

func TestSearch(t *testing.T) {
	dawg := CreateDAWG([][]int{{1, 2, 3, 1}, {4, 2, 3, 1}, {5, 2, 3, 1}, {5, 6, 1, 2}})

	test, err := dawg.Search([]int{1, 2, 3, 1}, 0, 1, false, false)
	if err != nil || len(test) != 1 || test[0] != []int{1, 2, 3, 1} {
		t.Error("Search failed")
	}
}

func TestSaveDawgFile(t *testing.T) {
	dawg := CreateDAWG([][]int{{1, 2, 3, 1}, {4, 2, 3, 1}, {5, 2, 3, 1}, {5, 6, 1, 2}})
	
	err := dawg.SaveToFile("Test.dawg")
	if err != nil {
		t.Error("Saving Failed")
	}
}

// func TestCreateBigDAWGfromFile(t *testing.T) {
// 	dawg, err := CreateDAWGFromFile("TestDawg.txt")
// 	if err != nil {
// 		t.Error("Creation From File Failed")
// 	}
// 	dawg.SaveToFile("Test.dawg")
// }
