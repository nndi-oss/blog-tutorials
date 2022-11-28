package main

import (
	"fmt"
	"log"
	"net/http"
)

const loginTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Nginx CPanel Login example</title>
</head>
<body>
    <h1>Dummy CPanel</h1>
    <form method="post" action="/login" autocomplete="off">
        <input type="hidden" name="csrf" value="__some_value__" />
        <div>
            <label>Username</label><br/>
            <input type="text" name="username" />
        </div>
        <div>
            <label>Password</label><br/>
            <input type="password" name="password" />
        </div>
        <button>Login</button>
    </form>
</body>
</html>
`

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html;charset=utf-8")
		w.Write([]byte(loginTemplate))
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Logged in to CPanel served by Go! ;)"))
	})

	fmt.Println("Starting CPanel server at :2083")
	err := http.ListenAndServe(":2083", nil)
	if err != nil {
		log.Fatalf("failed to start %v", err)
	}
}
