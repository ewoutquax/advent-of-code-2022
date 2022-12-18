package puzzle

import ()

func isValid(left Node, right Node) bool {
	if !left.isInteger && !right.isInteger {
		return isValid(*(left.childNode), *(right.childNode))
	} else if left.isInteger && right.isInteger && left.value == right.value {
		return isValid(*(left.nextNode), *(right.nextNode))
	} else if left.isInteger && right.isInteger && left.value < right.value {
		return true
	}

	return false
}
