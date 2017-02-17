# MashupAPI
REST API which offers a mashup of other API's.

##Features
The api will return with a 200 ok as long as there is a match in the music brainz database for the artist or band, even if there was no wikipedia information found or no cover art found. The response will in that case only include the album title and the album id.

##Instructions:

###Prerequisites
[Go must be installed of your system](https://golang.org/doc/install)

Clone this repository

Update your GOPATH to point to the MashupAPI workspace:
>$ export GOPATH=$HOME/MashupAPI

Change into the recently cloned directory and download the dependencies:
>$ go get ./...

Build the api, a file with the name myapi will now reside in the current directory afterwards:
>go build

Run the unit tests:
>go test

Start the webserver, it will serve at http://localhost:8080
>$ ./myapi

Run the function tests:
>$ cd $HOME/MashupAPI/test/
>go test

Try it out with curl:
>curl -i -H "Accept: application/json" http://localhost:8080/musicinfo/45a663b5-b1cb-4a91-bff6-2bef7bbfdd76

##Limitations/left to do:

Logging:
-log levels

Testing and testability:
-unit tests should cover corner cases
-load testing
-some refactoring to make unit testing easier
