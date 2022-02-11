package twofer

// returns a two-fer string with a name else two-fer you/me string.
func ShareWith(name string) string {
	fer := "One for you, one for me."
	if (name!="") {
		fer = "One for "+name+ ", one for me."
	}
	return fer
}
