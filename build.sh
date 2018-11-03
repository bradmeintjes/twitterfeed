#!/bin/bash

docker build --tag twitter_feed .
docker run -it -v `pwd`/res/:/var/ twitter_feed:latest /var/user.txt /var/tweet.txt