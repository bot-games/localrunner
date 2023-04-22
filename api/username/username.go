package username

import "fmt"

func UserName(uid uint32) string {
	if uid == 0 {
		return "SmartGuy"
	}

	return fmt.Sprintf("User %d", uid)
}
