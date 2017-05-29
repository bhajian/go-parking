# go-parking

- export GOPATH=$(go env GOPATH)
- mkdirp ~/go/src/github.com/user/go-parking
- git clone git@github.com:bhajian/go-parking.git
- cd ~/go/src/github.com/user/go-parking
- go install ./cmd/parking-app-server/
- $GOPATH/bin/parking-app-server --scheme http

# API Usage:

- To create a new lot:
	- In your postman paste the following url: http://127.0.0.1:PORT/api/lot
	- Place the port with the port you get from last command printed in the screen
	- Change the protocol to POST in postman
	- Paste the following body in the body of the request:
	
	```
	{
		"name": "Lot2",
		"lotType": "DAILY",
		"address":"",
		"smallCarSize": 40,
		"mediumCarSize": 40
	}
	``` 
	
- To get list of Lots
	- curl http://127.0.0.1:PORT/api/lot or paste the URL in the browser or postman with get protocol

- To get a ticket for a car on the entry to a lot send a request with the following URL to server with GET:
	- http://127.0.0.1:PORT/api/ticket/getTicket/lot/1/carSize/SMALL
	you will receive the following string as a result of get ticket which shows your entry time
	```
	{
	  "carSize": "SMALL",
	  "entryTime": "2017-05-29T18:18:12-04:00",
	  "id": 1,
	  "lotId": 1
	}
	```
- When the car leaves send a request to the following URL: http://127.0.0.1:51476/api/lot/1/ticket/1/carLeaves with POST
```
{
	
}
```
you will receive the following response:
```
{
  "amount": 5
}
```

- To update a lot use put 
- please see swagger for more info about the API

