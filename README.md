# Description

FizzBuzz LeBonCoin is an API web service written in go and postgresql that do a fizzbuzz exercise and compt the maximum requests of the same sent parameters

Features :

- Init API 
- Add an endpoint that send 5 parameters and return a formatted string as fizzbuzz exercise
- Add an endpoint that return the most requested parameters for the fizzbuzz endpoint

# Getting started

You must have docker-compose 2.5 installed
.env contains the database credentials, to copy .env.template :

    cp .env.template .env

then

    docker-compose up


Default API URL is http://localhost:4747

You can see documentation once the api started on http://localhost:4747/swagger/index.html