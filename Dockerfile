### Build binary from official Go image
FROM golang:stretch as build
COPY . /app
WORKDIR /app
RUN apt-get update && go get github.com/aofei/air && go get github.com/aofei/cameron
RUN go build -o /avatar-generator-tool .

### Put the binary onto Heroku image
FROM heroku/heroku:16
COPY --from=build /avatar-generator-tool /avatar-generator-tool
CMD ["/avatar-generator-tool"]