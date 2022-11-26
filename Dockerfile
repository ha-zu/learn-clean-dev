FROM golang:1.19.3-alpine3.16

ENV ROOT=/go/src/github.com/ha-zu/learn-clean-dev
ENV MAIN=${ROOT}/src/cmd
ENV BUILD=${ROOT}/src/build/cleanarchitecture
WORKDIR ${ROOT}

RUN apk update && \
	apk add bash && \
	apk add git

COPY go.mod go.sum ./
RUN go mod download
COPY ./src ./src

RUN go build -o ${BUILD} ${MAIN}/main.go

CMD ["./src/build/cleanarchitecture"]
