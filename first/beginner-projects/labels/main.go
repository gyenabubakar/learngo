package main

func main() {
	println(1)
	println(2)
	goto End

Return:
	println(3)
	return

End:
	println(4)
	goto Return
}
