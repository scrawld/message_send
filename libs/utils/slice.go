package utils

import (
	"reflect"
)

func InSlice(val interface{}, slice interface{}) (exists bool, index int) {
	exists = false
	index = -1

	if reflect.TypeOf(slice).Kind() != reflect.Slice {
		return
	}
	s := reflect.ValueOf(slice)
	for i := 0; i < s.Len(); i++ {
		if reflect.DeepEqual(val, s.Index(i).Interface()) == false {
			continue
		}
		index = i
		exists = true
		return
	}
	return
}

func SliceDiff(slice1, slice2 interface{}) (r []interface{}) {
	if reflect.TypeOf(slice1).Kind() != reflect.Slice || reflect.TypeOf(slice2).Kind() != reflect.Slice {
		return
	}
	s := reflect.ValueOf(slice1)
	for i := 0; i < s.Len(); i++ {
		if exists, _ := InSlice(s.Index(i).Interface(), slice2); exists {
			continue
		}
		r = append(r, s.Index(i).Interface())
	}
	s = reflect.ValueOf(slice2)
	for i := 0; i < s.Len(); i++ {
		if exists, _ := InSlice(s.Index(i).Interface(), slice1); exists {
			continue
		}
		r = append(r, s.Index(i).Interface())
	}
	return
}

func SliceUnique(slice interface{}) (r []interface{}) {
	if reflect.TypeOf(slice).Kind() != reflect.Slice {
		return
	}
	s := reflect.ValueOf(slice)
	for i := 0; i < s.Len(); i++ {
		if exists, _ := InSlice(s.Index(i).Interface(), r); exists {
			continue
		}
		r = append(r, s.Index(i).Interface())
	}
	return
}

func SliceUniqueInt(slice []int) (r []int) {
	for _, v := range slice {
		if exists, _ := InSlice(v, r); exists {
			continue
		}
		r = append(r, v)
	}
	return
}

func SliceUniqueInt32(slice []int32) (r []int32) {
	for _, v := range slice {
		if exists, _ := InSlice(v, r); exists {
			continue
		}
		r = append(r, v)
	}
	return
}
