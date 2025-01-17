package coverage

import (
	"os"
	"reflect"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

func TestLen(t *testing.T) {
	p := People(make([]Person, 4))
	p = People{
		{firstName: "Akmal", lastName: "Kadirov", birthDay: time.Date(1998, 10, 23, 0, 0, 0, 0, time.UTC)},
		{firstName: "Batir", lastName: "Turgunov", birthDay: time.Date(1998, 10, 23, 23, 2, 4, 5, time.UTC)},
		{firstName: "Tair", lastName: "Sultanov", birthDay: time.Date(1998, 4, 5, 5, 6, 7, 7, time.UTC)},
		{firstName: "Shakir", lastName: "Smith", birthDay: time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)},
	}
	expected := len(p)
	if output := p.Len(); output != expected {
		t.Errorf("Expected: %v, Got: %v", expected, output)
	}
}

func TestLess(t *testing.T) {
	p := People(make([]Person, 3))

	p = People{
		{firstName: "Aadirov", lastName: "Akmal", birthDay: time.Date(2000, 10, 23, 0, 0, 0, 0, time.UTC)},
		{firstName: "Turgunov", lastName: "Batir", birthDay: time.Date(1998, 10, 23, 23, 2, 4, 5, time.UTC)},
		{firstName: "Bair", lastName: "Aultanov", birthDay: time.Date(2000, 10, 23, 0, 0, 0, 0, time.UTC)},
		{firstName: "Bair", lastName: "Bakhimov", birthDay: time.Date(2000, 10, 23, 0, 0, 0, 0, time.UTC)},
	}

	trueTestCases := [3][2]int{
		{0, 1},
		{0, 2},
		{2, 3},
	}

	for _, testCase := range trueTestCases {
		if output := p.Less(testCase[0], testCase[1]); output != true {
			t.Errorf("TestCase: [%v, %v] Expected: %v; Got: %v", testCase[0], testCase[1], true, output)
		}
	}
}

func TestSwap(t *testing.T) {
	p := People(make([]Person, 3))
	p = People{
		{firstName: "Kadirov", lastName: "Akmal", birthDay: time.Date(1998, 10, 23, 0, 0, 0, 0, time.UTC)},
		{firstName: "Turgunov", lastName: "Batir", birthDay: time.Date(1998, 10, 23, 23, 2, 4, 5, time.UTC)},
		{firstName: "Tair", lastName: "Sultanov", birthDay: time.Date(1998, 4, 5, 5, 6, 7, 7, time.UTC)},
		//{firstName: "Shakir", lastName: "Smith", birthDay: time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)},
	}
	i := 0
	j := 2
	x := p[i]
	y := p[j]
	p.Swap(i, j)
	if p[i] != y || p[j] != x {
		t.Errorf("The test failed as the swap function did not work.")
	}
}

func TestNew(t *testing.T) {
	str1 := "1 2 3\n4 5 6\n7 8 9"
	m := Matrix{
		rows: 3,
		cols: 3,
		data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}

	out, err := New(str1)
	if err != nil {
		t.Errorf("Error while parsing")
	}

	if out.rows != m.rows || out.cols != m.cols {
		t.Errorf("Failed: Expected rows, cols, data: %v, %v, %v Got rows: %v, cols: %v, data: %v", m.rows, m.cols, m.data, out.rows, out.cols, out.data)
	}
}

func TestRows(t *testing.T) {
	m := Matrix{
		rows: 3,
		cols: 3,
		data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}
	expected := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	out := m.Rows()
	if !reflect.DeepEqual(out, expected) {
		t.Errorf("Test Fails: Expected - %v, Got - %v", expected, out)
	}
}

func TestCols(t *testing.T) {
	m := Matrix{
		rows: 2,
		cols: 3,
		data: []int{1, 2, 3, 4, 5, 6},
	}
	expected := [][]int{
		{1, 4},
		{2, 5},
		{3, 6},
	}
	out := m.Cols()
	if !reflect.DeepEqual(out, expected) {
		t.Errorf("Test fails: Expected - %v, Got - %v", expected, out)
	}
}

func TestSet(t *testing.T) {
	expected := true
	m := Matrix{
		rows: 3,
		cols: 2,
		data: []int{1, 2, 3, 4, 5, 6},
	}
	row := 2
	col := 1
	val := 10
	out := m.Set(row, col, val)
	if out != expected {
		t.Errorf("Test Fails: Expected - %v, Got - %v", expected, out)
	}

}
