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

(MIT)

Copyright (c) 2015 Frederic Gingras (frederic@gingras.cc)

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
