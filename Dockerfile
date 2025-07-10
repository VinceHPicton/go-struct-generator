FROM golang:1.22 AS project

WORKDIR /app

COPY go.mod main.go ./
ADD buildstructs buildstructs

RUN go build -o ./go-struct-generator


FROM alpine:latest AS prod

COPY --from=project /app/go-struct-generator /usr/bin

# Make a dir for the user's present working dir to be mounted to, and a place for the go binary to live so it wont be overwritten by the volume mount
RUN mkdir /usr/userpwd

WORKDIR /appspace/userpwd

ENTRYPOINT ["/usr/bin/go-struct-generator", "-file"]
