FROM golang:1.19.3-alpine3.16

ENV ROOT=/go/src/github.com/ha-zu/learn-clean-dev
ENV MAIN=${ROOT}/src/cmd/cleanarchitecture
ENV BUILD=${ROOT}/src/build
WORKDIR ${ROOT}

RUN apk update && \
	apk add bash && \
	apk add git

COPY ./src ./src
COPY go.mod .
COPY go.sum .
RUN go mod tidy

RUN go build -o ${BUILD} ${MAIN}/main.go

CMD [ "./src/build/main" ]
