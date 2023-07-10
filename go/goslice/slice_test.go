package goslice

import "testing"

func TestIntersect(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5, 6, 7}
	s2 := []int{5, 7, 9, 10}
	s3 := Intersect[int](s1, s2)
	t.Log(s3, len(s3))
}

func TestUnion(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5, 6, 7}
	s2 := []int{5, 7, 9, 10}
	s3 := Union[int](s1, s2)
	t.Log(s3, len(s3))
}

func TestExcept(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5, 6, 7}
	s2 := []int{5, 7, 9, 10}
	s3 := Except[int](s1, s2)
	t.Log(s3, len(s3))
}
