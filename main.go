package main

func main() {
	server := NewServer(":8081")
	server.Handle("GET", "/", HandleRoot)
	server.Handle("POST", "/api", server.AddMiddleware(HandleHome, CheckAuth(), Logging()))
	server.Handle("POST", "/create", PostRequest)
	server.Handle("POST", "/user", UserPostRequest)
	server.Listen()
}
