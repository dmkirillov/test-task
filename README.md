Build a middleware using echo framework
First of all, you should create a handler which sends how many days left until 1 Jan 2024 
and response with HTTP 200 OK status code.

Secondly, build a middleware, which checks HTTP header User-Role presents and contains admin and prints red button user detected to the console.
<hr/>
<b>Run</b>

go mod tidy

make run
<hr/>
<b>Test<b>

curl -v  '127.0.0.1:8080/status'

curl --location --request GET '127.0.0.1:8080/status' --header 'User-Role: admin'