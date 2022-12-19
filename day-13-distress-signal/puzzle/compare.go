package puzzle

import "fmt"

func isValid(left Node, right Node) string {
	fmt.Println("Start comparison loop")

	if left.isInteger {
		fmt.Print("left is integer '", left.value, "'")
	} else {
		fmt.Print("left is list")
	}

	if right.isInteger {
		fmt.Print("; right is integer '", right.value, "'")
	} else {
		fmt.Print("; right is list")
	}
	fmt.Print("\n")

	if !left.isInteger && !right.isInteger {
		fmt.Println("Found 2 lists, go into subloop")
		tempResult := isValid(*(left.childNode), *(right.childNode))
		if tempResult == "unknown" {
			fmt.Println("Returned from undecisive subloop; continue with nextNode")
		} else {
			fmt.Println("returned from decisive subloop; return result")
			return tempResult
		}
	} else if left.isInteger && right.isInteger && left.value > right.value {
		fmt.Println("Left integer is higher (", left.value, "vs", right.value, "); return false")
		return "false"
	} else if left.isInteger && right.isInteger && left.value < right.value {
		fmt.Println("Left integer is lower (", left.value, "vs", right.value, "); return true")
		return "true"
	} else if !left.isInteger && right.isInteger && right.nextNode == nil {
		fmt.Println("Left is array and right is last digit: convert")
		return isValid(*(left.nextNode), right)
	}

	if left.nextNode == nil && right.nextNode == nil {
		fmt.Println("end of both lists; go to upper loop without decisive answer")
		return "unknown"
	} else if left.nextNode == nil {
		fmt.Println("left side is out of items: return true")
		return "true"
	} else if right.nextNode == nil {
		fmt.Println("right side is out of items: return false")
		return "false"
	} else {
		fmt.Println("Continue with nextNode from both sides")
		return isValid(*(left.nextNode), *(right.nextNode))
	}
}
