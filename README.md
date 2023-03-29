# ToDo Sport Bot

![](https://img.shields.io/github/go-mod/go-version/shipaaa/ToDoSportBot?style=plastic)

🤖 A telegram bot that sends exercises to a muscle group that the user chooses

## What can I do with this bot

- View exercises for a specific muscle group
- View the training program for a specific day
- Suggest your idea or add a new exercise

## Example
Usage example of a [Telegram Bot](https://t.me/todoSportBot) on Russian

<img height="450" src="video-example.gif" width="400"/>

### Description of bot commands

- `/start` — welcome message
- `/help` — help bot
- `/allexercises` — sends all exercises currently available in the database
- `/training` — sends exercises for a specific training day.

## How to start

Clone project and change the configuration file by substituting your values

```bash
git clone https://github.com/shipaaa/ToDoSportBot.git
cp .env.example .env
```

And then depending on your preferences run the program

#### On your local machine

1. Get install [Postgresql](https://www.postgresql.org/download/)
2. Create Database and run the following queries in your terminal

```bash
psql -U username -d DataBase -a -f pkg/models/tablesCreation.sql
psql -U username -d DataBase -a -f pkg/models/tablesFilling.sql
```

3. Launch the application `make run`

#### Using Docker

1. Make sure [Docker](https://docs.docker.com/engine/install/) is installed on your computer
2. Build and up container with telegram bot `make up`

## Ideas

- Add CI/CD
- Increase the exercise database
- Add exercises for women
