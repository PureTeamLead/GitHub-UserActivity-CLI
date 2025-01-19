# GitHub activity Tracker - CLI application

## How to Run?

```bash
https://github.com/PureTeamLead/GitHub-UserActivity-CLI
cd main
```

## Usage

To start the tool, at first build the application:

```bash
go build -o github-activity
```

## What program can do?

### Get user's info

```bash
# ./github-activity get_info username
./github-activity get_info PureTeamLead
```

Also app has some flags implemented for get_info:

-l: for location
-f: for followers

```bash
# ./github-activity get_info username -l -f
./github-activity get_info PureTeamLead -l -f
```

Credentials flag:

-a

```bash
./github-activity -a
```

### Get user's events

```bash
# ./github-activity get_events username
./github-activity get_events PureTeamLead
```
