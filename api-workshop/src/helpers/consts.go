package helpers

// LayoutISO ...
const LayoutISO = "2006-01-02"

// LayoutDateTime ...
const LayoutDateTime = "2006-01-02 15:04:05"

// GetCultivarEnums ...
func GetCultivarEnums() map[string]int {
	cultivarE := map[string]int{"soja": 1, "milho": 2, "trigo": 3, "tomate": 4}

	return cultivarE
}

// GetAreaEnums ...
func GetAreaEnums() map[string]int {
	areaE := map[string]int{"pivot": 1, "sequeiro": 2, "calcinha": 3}

	return areaE
}
