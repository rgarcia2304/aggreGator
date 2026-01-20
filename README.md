# aggreGator- Blog Aggregator

Welcome to aggregator this RSS feed aggregator I built. with this you can get RSS feeds from accross the web, store them in your postgres database, You can follow other users feeds, And view summaries of those posts all in your terminal!

## Requirements 
You will need to have an up to date version of golang and I used postgres 15 to run this CLI application. 

## Installation 

`go install github.com/rgarcia2304/aggreGator`

## Set Up 

To start you will want to create a config file. Host this config file at the root of your computer. Name it .gatorconfig.json Inside it will contain a json file that will contain the current user and database url. It should look like this. 

`
{
  "db_url": "connection_string_goes_here",
  "current_user_name": "username_goes_here"
}
`

Next open up your postgres shell and create a database called gator.
`CREATE DATABASE`
Connect to the database.
`\c gator`
Get your connection string: Will look something like this. 
`protocol://bob:@host:5432/gator` or generally like
`protocol://username:password@host:port/database`

To get all the updated tables you will need to run in sql/schemas folder inside the project 
`goose postgres <connection_string> up `
Do this until, you are at the highest migration. 

Once that is all set up you can now use the application

## Commands 
All commands are run at the main package directory level

Some general commands to run are :
`register` - Allows you to register a user 
Usage - `go run . register {name of user}`

`register` - Allows you to make user active
Usage - `go run . login {name of registered user}`

`addfeed` - Allows you to add feed for the user
Usage - `go run . addfeed {Name of feed} {url of feed}`

`agg` - Allows you to aggregate results of the feed 
Usage - `go run . agg {time_interval eg: 10s, 1min, 1hr}`

`browse` - Allows you to view different feeds you have aggregated 
Usage - `go run . browse {Number of feeds}`

All other commands are in the main.go file

Enjoy!






