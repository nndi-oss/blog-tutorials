FROM golang:1.19-alpine 
RUN apk add --no-cache git
WORKDIR /go/nginx-cpanel
COPY . .
RUN go generate -x -v
RUN go build -v -o /bin/cpanel main.go && chmod +x /bin/cpanel
EXPOSE 2083
ENTRYPOINT "/bin/cpanel"