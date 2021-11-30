# trello-daily-logs

[![GoDoc](https://godoc.org/l-lin/trello-daily-logs?status.svg)](https://godoc.org/l-lin/trello-daily-logs)

I want to track what I did during the day, and as I'm using Trello to manage my daily tasks, I can
use the [Trello APIs](https://developers.trello.com/docs/api-introduction) to fetch the card
information, especially the ones from my `DONE` list.

This CLI is used to:

- get the Trello card information
- append/prepend today's content into a file `/path/to/daily-logs/2020/02.md`:

```md
## Sunday 09

**ABANDONED**

&nbsp;&nbsp;&nbsp; api@spec: write OpenAPI specifications

**PERSO**

&nbsp;&nbsp;&nbsp; career@interview: find good questions for interview

&nbsp;&nbsp;&nbsp; career: find good path

**WORK**

<details>
<summary>api@gravitee: install in prod environment</summary>

# Gravitee in prod
## Getting started

- get up from bed
- brush teeth

</details>

&nbsp;&nbsp;&nbsp; api@spec: write OpenAPI specifications

---

<details>
<summary>UNFINISHED</summary>

**PERSO**

&nbsp;&nbsp;&nbsp; shopping: buy milk

**WORK**

&nbsp;&nbsp;&nbsp; projectA@taskA: study solutions

&nbsp;&nbsp;&nbsp; projectB@taskB: implement solution

</details>

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
25 18 * * 1-5 /home/llin/.asdf/shims/trello-daily-logs -f file 1> /dev/null 2> /tmp/trello-daily-logs.log
26 18 * * 1-5 /home/llin/perso/trello-daily-logs/commit.sh 1> /dev/null 2> /tmp/trello-daily-logs.commit.log
```

