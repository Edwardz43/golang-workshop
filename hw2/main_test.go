package hw2

import "testing"

func TestAdd100AndNegative1(t *testing.T) {
	_, e := Add(100, -1)
	if e == nil {
		t.Errorf("Add(100, -1) should returned error")
	}
	t.Logf("Add(100, -1) returned error: %v", e)
}

func TestAdd100And100(t *testing.T) {
	r, e := Add(100, 100)
	if e != nil {
		t.Errorf("Add(100, 100) should not returned error")
	}
	if r != 200 {
		t.Errorf("Add(100, 100) should returned 200, but returned %v", r)
	}
	t.Logf("Add(100, 100) returned %v", r)
}

func TestSum100(t *testing.T) {
	r, e := Sum(100)
	if e != nil {
		t.Errorf("Sum(100) should not returned error")
	}
	if r != 5050 {
		t.Errorf("Sum(100) should returned 5050, but returned %v", r)
	}
	t.Logf("Sum(100) returned %v", r)
}

func TestSumNegative100(t *testing.T) {
	_, e := Sum(-100)
	if e == nil {
		t.Errorf("Sum(-100) should returned error")
	}
	t.Logf("Sum(-100) returned error: %v", e)
}
