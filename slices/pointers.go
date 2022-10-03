package slices

type MyStruct struct {
	V int
}

func AppendSlicePointers(n int) []*MyStruct {
	slice := make([]*MyStruct, 0, 100)
	for j := 0; j < 100; j++ {
		slice = append(slice, &MyStruct{V: j})
	}
	return slice
}

func AppendSliceNoPointers(n int) []MyStruct {
	slice := make([]MyStruct, 0, 100)
	for j := 0; j < 100; j++ {
		slice = append(slice, MyStruct{V: j})
	}
	return slice
}

// If you "have to" return a slice of pointers
func AppendSliceHybrid(n int) []*MyStruct {
	slice := make([]MyStruct, 0, 100)
	for j := 0; j < 100; j++ {
		slice = append(slice, MyStruct{V: j})
	}

	slicep := make([]*MyStruct, len(slice))
	for j := range slice {
		slicep[j] = &slice[j]
	}
	return slicep
}
