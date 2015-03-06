FROM google/golang:1.4

WORKDIR /gopath/src/github.com/kiasaki/marks
ADD . /gopath/src/github.com/kiadaki/marks/
RUN go get github.com/kiasaki/marks

CMD []
ENTRYPOINT ["/gopath/bin/marks"]
