package main

//go:noinline
func escape() *int {
	i := 1
	return &i
}

func main() {
	a := escape()
	*a++
}

// go build -gcflags '-m' main.go
