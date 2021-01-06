package main
import ("net/http"; "io")
func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", nil)
}
func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<doctype html>
		<html>
		<head>
		<title>Hello Vijay</title>
		</head>
		<body>
		Hello Vijay this life is empty
		</body>
		</html>`,
	)
	http.Handle(
		"/assets/",
		http.StripPrefix(
			"/assets/",
			http.FileServer(http.Dir("assets")),
		),
	)
}

