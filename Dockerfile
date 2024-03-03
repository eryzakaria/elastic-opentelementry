FROM golang:1.17.0

WORKDIR /usr/src/app
COPY golang/go.mod .
COPY golang/main.go .
RUN go mod tidy

RUN go golang/build -o hello-app
ENTRYPOINT [ "/usr/src/app/hello-app" ]
