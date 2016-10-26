# Marks

_A feature-less markdown notebook_

## Introduction

**Marks** is a simple note taking app that mostly gets out of the way. It's not really feature full but it might be just what you are looking for.

Feel free to try it out on Heroku using the deploy button below:

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/kiasaki/marks)

## Getting started

**Running:**

```bash
go get github.com/kiasaki/marks
marks -help # for help
marks -port 8080 -postgres-url postgres://localhost/marks?sslmode=disable
```

**Developing:**

```bash
go get github.com/kiasaki/marks
cd $GOPATH/src/github.com/kiasaki/marks
go run *.go -postgres-url ...
```

**Creating the database tables before first use:**

_Make sure your PostgreSQL server is started and that you are pointing to
an existing database._

```bash
marks -create-db -postgres-url ...
```

## Screenshot

![Screenshot](https://raw.githubusercontent.com/kiasaki/marks/master/screenshot.png)

## License

MIT. See `LICENSE` file.
