FROM golang:1.13

#WORKDIR /
COPY . /go/src/github.com/DimsFromDergachy/quiz

RUN go get -d -v /go/src/github.com/DimsFromDergachy/quiz \
    && \
    go install -v /go/src/github.com/DimsFromDergachy/quiz

ENTRYPOINT [ "/go/bin/quiz" ]
