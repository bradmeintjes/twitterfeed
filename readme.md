## Description

This application is a simplified twitter feed application which makes use of two input files, supplied as command line parameters, specifying the users and tweets from each user.

See `/res/Retail IT Coding Assignment Version 4.pdf` for the full application requirements specification

## Usage

````
feed <usersfile> <tweetsfile>
````

## Building

Application makes use of go modules (requires go 1.11 or later). To build it manually, run 

````
go build -o feed
````

### Docker

This application can be built and containerised as a docker container using
````
docker build -t twitter-feed .
````

which can be executed using 
````
docker run -v `pwd`/res/:/var/ twitter_feed:latest /var/user.txt ~/var/tweet.txt
````

## Testing

Execute the unit test suite with

````
go test ./...
````
