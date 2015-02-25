# Marks

#### Simplistic Markdown notebook

## Introduction

**Marks** is a simple notebook/wiki style app that allows you to sketch out, keep
track of, document what's keeping you awake and you're lugging around in your head.

It's simple and it works, feel free to clone and deploy on _Heroku_.

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/kiasaki/marks)

## Technologies

This app is the most simple and vanilla Goland app, no fancy framework, libs
as needed, and snippets copied over when small.

Your notes/thoughts/documentation is secured using super advanced Basic Auth.
Although you could chose to simply let your darkest secret open to the public,
by not setting a user and password, it's a choice you have...

Main packages used:

- `database/sql` as way to store notes and retrive them
- `github.com/lib/pq` driver for postgresql
- `html/template` for rendering pages
- `net/http` for all http server needs
- `flag` for configuration
- `github.com/kiasaki/batbelt`

## Getting started

Use:

```bash
go get github.com/kiasaki/marks
marks -postgr...
```

Develop:

```bash
go get github.com/kiasaki/marks
cd $GOPATH/src/github.com/kiasaki/marks
go run *.go -postgre...
```

An before that you should have started postgresql and ran the `db.sql` file
on the database you wish to use so you have the necessaty tables created.

```bash
postgres&
psql marks -f db.sql
```

## Look and feel

Still WIP but gives an idea...

![Screenshot](https://raw.githubusercontent.com/kiasaki/marks/master/screenshot.png)

## License

MIT
