package main

type logWriter struct{}

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://canva.com",
		"http://amazon.com",
		"http://golang.com",
	}

	for _, l := range links {
		println(l)
	}
}
