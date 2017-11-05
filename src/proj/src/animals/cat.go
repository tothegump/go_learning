package animals

func CatCall(count int) string {
	sound := ""
	for i := 0; i < count; i ++ {
		sound += "miao.. "
	}
	return sound
}
