package calc

// Compare numbers
func Compare(fn int, ln int) (r bool, err error) {
	switch true {
	case fn > ln:
		return true, err
	default:
		return false, err
	}
}
