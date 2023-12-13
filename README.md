# concurrency_go

- It shows how a client can launch several Goroutines to a server and wait for the result using channels.
- We can set a timeout on client side to handle slow requests on the server side.
- The server might send an HTTP 500 anytime..

Example:

```
# request client side

go run main.go # by default launch 10 Goroutines 

{"level":"info","message":"Sending clientId %!n(int=5)"}
{"level":"info","message":"Sending clientId %!n(int=2)"}
{"level":"info","message":"Sending clientId %!n(int=8)"}
{"level":"info","message":"Sending clientId %!n(int=9)"}
{"level":"info","message":"Sending clientId %!n(int=0)"}
{"level":"info","message":"Sending clientId %!n(int=4)"}
{"level":"info","message":"Sending clientId %!n(int=6)"}
{"level":"info","message":"Sending clientId %!n(int=1)"}
{"level":"info","message":"Sending clientId %!n(int=7)"}
{"level":"info","message":"Sending clientId %!n(int=3)"}
```

```
# response client side

{SERVER GOT clientId 0 - DONE PROCESSING AFTER 1s
 <nil>}
{SERVER GOT clientId 1 - DONE PROCESSING AFTER 2s
 <nil>}
{SERVER GOT clientId 2 - DONE PROCESSING AFTER 2s
 <nil>}
{SERVER GOT clientId 3 - DONE PROCESSING AFTER 2s
 <nil>}
{SERVER GOT clientId 6 - DONE PROCESSING AFTER 2s
 <nil>}
{SERVER GOT clientId 9 - DONE PROCESSING AFTER 3s
 <nil>}
{SERVER GOT clientId 8 - DONE PROCESSING AFTER 4s
 <nil>}
{request has failed for clientId 5 # <- a request might fail on server side, use clientId to track it
 <nil>}
{SERVER GOT clientId 4 - DONE PROCESSING AFTER 5s
 <nil>}
{SERVER GOT clientId 7 - DONE PROCESSING AFTER 5s
 <nil>}
```
