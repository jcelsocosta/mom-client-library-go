package main

func main() {
	middleware := NewMiddlewareIntegration()

	middleware.Bind("localhost:8080")
}
