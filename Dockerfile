FROM golang:1.20 AS project

WORKDIR /app

COPY go.mod go.sum main.go ./
ADD buildstructs buildstructs

RUN go build -o ./go-struct-generator


FROM alpine:latest AS prod

# Make an appspace dir for everything needed, as binary cant be placed in the default /bin from alpine due to permissions
WORKDIR /appspace
# Make a dir for the user's present working dir to be mounted to, and a place for the go binary to live so it wont be overwritten by the volume mount
RUN mkdir bin userpwd

COPY --from=project /app/go-struct-generator ./bin

WORKDIR /appspace/userpwd

ENTRYPOINT ["../bin/go-struct-generator", "-file"]
