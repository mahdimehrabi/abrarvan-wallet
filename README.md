## Please Concern

I implement this project **without** using any framework and ORM
for best performance , scalability.<br>
for the purpose of this project I didn't implement all needed features
so interfaces don't contain all needed methods ,
and I didn't implement UX related stuff like validation messages and response messages ,
also I didn't implement docker-compose version for production.
I know you don't needed test for this project , but I wrote unit test and mock for this project.
<br>

## More Scalability

Used standard library for routing not any framework and no ORM used. <br>
No external packages that make the performance lower used.<br>
Used [pgx](https://github.com/jackc/pgx) as database interface because
we only use PostgreSQL as database and pgx is faster than `database/sql` package in standard library.

## Architecture , Design

The architecture of this project is clean architecture,
I created an image to help you understand architecture of this project better.<br>
![clean architecture](https://raw.githubusercontent.com/mahdimehrabi/go-challenge/main/clean.png)

Used interface for getting tools like logger, db, memoryDB so using another tool for example another
logger or db don't force you to edit all codes of different layers.
<br><br>
Used [uber fx](https://github.com/uber-go/fx) as dependency injection system
to increase readability and save more memory.

## Solution

I used redis as in memory db and postgreSQL as database.<br>
when we create a discount/charge code this object will be created in both postgres and redis.<br>
we defined remaining of usage field as `consumer_count`.<br>
when user use charge/discount code we decrease count of consumer_count in redis,
when amount of remaining usages become 0 in redis we set rconsumer_count column in postgre
to 0 too and we set consumer_count in redis to `-1`.
and after this any requests that come to our server because consumer_count in
redis is -1 we just throw 404 error to them.<br>
To prevent losing data from redis we run a periodic task
every hour and update consumer_count of **active** codes in postgres.
in decreasing action if anything goes wrong we cancel other actions ,
for example if adding charge to user be successful,
but decreasing consumer_count fail we will not add charge to user.

## Security

prevented race condition in redis by using [redis transaction](https://redis.io/docs/manual/transactions/).<br>
prevented users to use charge more than once by defining `recevied_charge` column in users table.<br>

## Getting started

`git clone https://github.com/mahdimehrabi/go-challenge.git` <br>
copy env file <br>
`cd gin-gorm-boilerplate`<br>
`cp env.example .env` <br>

create docker volume and start <br>
`docker volume create psql_data` <br>
`docker-compose up -d ` <br>
create database <br>
`docker-compose exec database psql -U root`<br>
`CREATE DATABASE challenge;`<br>
run migrations <br>
`make migrate-up` <br>

run tests <br>
`make test-all`

generate some users (seed), this command create between 1 and 1000 random users
,so feel free to use this command as many times as you want to create more users.<br>
`make seed`

now please edit a file(just add new line or tab or space is enought) and save (to restart delve server)
or restart docker-compose `docker-compose down && docker-compsoe up -d` <br>

#### Now you can send your requests to Endpoints

POST `localhost:8000/users`  Create new user send <br>
example **request** data

```
{
    "ID":"fsas42aa3af",
    "segment":"1aa1x",
    "expiredSegment":1655209411
}
```

GET `localhost:8000/segments/count` get segments and it user counts

example **response** data

```
[
{"title":"RU","usersCount":50052132},
{"title":"z12XTowIA2","usersCount":45235009}
]
```
