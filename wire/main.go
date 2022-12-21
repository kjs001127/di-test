package main

func main() {
	err := NewServer().ListenAndServe()
	if err != nil {
		return
	}
}
