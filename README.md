# goscore
go exercise. Hex architecture, CQRS, no framework, in memory database, testing.<br>
No other libraries apart from the standard library have been used.<br>
This has been a very good exercise. Go is a fun language.<br>
<br>
In order to execute this program, just go to the root dir and type:<br>
_$ go run main.go_<br>
Alternatively, build the project first with 'go build' and simply execute<br>
_$ Goscore.exe_ in windows or<br>
_$ ./Goscore_ in Unix systems<br>
<br>
The server will populate the in memory DB with a 1000 user scores and will listen to connections to localhost:8080<br>
You can either throw the tests in Infrastructure/Tests with<br>
_$ go test_<br>
or you can do requests with your browser, curl, Postman or the like.<br>
<br>
The endpoints are:<br>

- GET "/" for a simple Hello World
- GET "/fetch_absolute" to get top scores. Params: rank(int) 
- GET "/fetch_relative" to get relative scores. Params: rank(int), n_relatives(int)
- POST "/new_total". Params: (json) {"user": int, "total": int}
- POST "/new_diff". Params: (json) {"user": int, "score": string}