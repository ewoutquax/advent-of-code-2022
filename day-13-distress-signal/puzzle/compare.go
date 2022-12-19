package puzzle

import ()

type compareResult struct {
}

func isValid(left Node, right Node) bool {
	if !left.isInteger && !right.isInteger {
		tempResult := isValid(*(left.childNode), *(right.childNode))
		if tempResult == nil {
			return isValid(*(left.nextNode), *(right.nextNode))
		} else {
			return tempResult
		}
	} else if left.isInteger && right.isInteger && left.value == right.value {
		return isValid(*(left.nextNode), *(right.nextNode))
	} else if left.isInteger && right.isInteger && left.value < right.value {
		return true
	} else if left.nextNode == nil && right.nextNode == nil {
		return nil
	}

	return false
}
