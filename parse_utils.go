package main

// knownUnit returns true if the string passed is a known unit type like, "cum", "metre"
func knownUnit(un string) bool {
	knownUnits := []string{"cum", "each", "metre", "meter", "sqm"}
	for _, i := range knownUnits {
		if i == un {
			return true
		}
	}
	return false
}
