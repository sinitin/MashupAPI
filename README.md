# MashupAPI
REST API which offers a  mashup of other API's

##Instructions:

###Prerequisites
[Go must be installed of your system](https://golang.org/doc/install)

Clone this repository

Update your GOPATH to point to the MashupAPI workspace:
>$ export GOPATH=$HOME/MashupAPI

Change into the recently cloned directory 

Download the dependencies:
>$ go get ./...

Build the api with the following command, a file with the name myapi will now reside in the current directory afterwards:
>go build

Run the unit tests:
>go test

Start the webserver, the webserver will serve at http://localhost:8080:
>$ ./myapi

Run the function tests:
>$ cd $HOME/MashupAPI/test/
>go test

Try it out with curl:
>curl -i -H "Accept: application/json" http://localhost:8080/musicinfo/45a663b5-b1cb-4a91-bff6-2bef7bbfdd76
