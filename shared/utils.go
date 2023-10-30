package shared

func SliceStringToInterfaces(slices []string) []interface{} {
	results := []interface{}{}
	for _, s := range slices {
		results = append(results, s)
	}
	return results
}
