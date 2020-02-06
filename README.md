# trello-daily-logs

[![GoDoc](https://godoc.org/l-lin/trello-daily-logs?status.svg)](https://godoc.org/l-lin/trello-daily-logs)

I want to track what I did during the day, and as I'm using Trello to manage my daily tasks, I can
use the [Trello APIs](https://developers.trello.com/docs/api-introduction) to fetch the card
information, especially the ones from my `DONE` list.

This CLI is used to:

- get the Trello card information
- append/prepend today's content into a file `/path/to/daily-logs/2020/02.md`:

```md
## Thursday 06

- ABANDONED
  - api@spec: write OpenAPI specifications
- PERSO
  - career@interview: find good questions for interview
- WORK
  - api@gravitee: install in prod environment
  - api@spec: write OpenAPI specifications

```

## Getting started

```bash
# Build
make compile
```

## Usage

```bash
# Run binary
./bin/trello-daily-logs -h
# Or directly using go
go run .
# Install
go install
# Don't forget to run at least once to initialize the config file
trello-daily-logs
```

Add to the crontab:

```bash
# If you want to directly write into a file
55 17 * * 1-4 /home/llin/go/bin/trello-daily-logs -f file 1> /dev/null 2> /var/log/trello-daily-logs.log
55 16 * * 5 /home/llin/go/bin/trello-daily-logs -f file 1> /dev/null 2> /var/log/trello-daily-logs.log
# If you want to use shell command and prepend to a file
55 17 * * 1-4 local folder="/home/llin/perso/daily-logs/$(date +%Y)" && mkdir -p "${folder}" && touch "${folder}/$(date +%m).md" && printf '%s\n\n%s' "$(/home/llin/go/bin/trello-daily-logs)" "$(cat ${folder}/$(date +%m).md)" > ${folder}/$(date +%m).md
55 16 * * 5 local folder="/home/llin/perso/daily-logs/$(date +%Y)" && mkdir -p "${folder}" && touch "${folder}/$(date +%m).md" && printf '%s\n\n%s' "$(/home/llin/go/bin/trello-daily-logs)" "$(cat ${folder}/$(date +%m).md)" > ${folder}/$(date +%m).md
```

