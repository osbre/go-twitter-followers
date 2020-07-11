A simple script to get Twitter account followers.

## Installation

```bash
git clone git@github.com:osbre/go-twitter-followers.git
go get -u github.com/gocolly/colly/
```

## Usage

With one username:

```bash
go run main.go USERNAME
```

With many usernames:

```bash
go run main.go USERNAME USERNAME USERNAME USERNAME...
```