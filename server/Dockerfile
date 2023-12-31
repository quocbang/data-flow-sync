ARG golang_version=1.20.8
ARG alpine_version=3.18

FROM golang:${golang_version}-alpine${alpine_version} as server_builder
ARG go_swagger_version=0.30.5

RUN apk add --no-cache git
RUN apk add --update --no-cache curl build-base
RUN curl -sSLo /usr/local/bin/swagger https://github.com/go-swagger/go-swagger/releases/download/v${go_swagger_version}/swagger_linux_amd64
RUN chmod +x /usr/local/bin/swagger
RUN apk update && apk add openssh

# test env
ARG SMTP_SERVER_TEST
ARG SMTP_PORT_TEST
ARG SMTP_SENDER_TEST
ARG SMTP_PASSWORD_TEST
ENV SMTP_SERVER_TEST=${SMTP_SERVER_TEST}
ENV SMTP_PORT_TEST=${SMTP_PORT_TEST}
ENV SMTP_SENDER_TEST=${SMTP_SENDER_TEST}
ENV SMTP_PASSWORD_TEST=${SMTP_PASSWORD_TEST}

ENV SERVER_DIR ${GOPATH}/src/github.com/quocbang/data-flow-sync/server
ENV APP_NAME data-flow-sync

COPY . ${SERVER_DIR}/

WORKDIR ${SERVER_DIR}

RUN go generate .
RUN go mod download
RUN go vet ./...
RUN go test -race -gcflags -l -coverprofile .testCoverage.txt ./...
RUN go tool cover -func .testCoverage.txt
RUN go build -race -ldflags "-extldflags '-static'" -o /opt/data-flow-sync/server ./swagger/cmd/${APP_NAME}-server

CMD ["/bin/sh"]

# copy exec file to new server 
FROM golang:${golang_version}-alpine${alpine_version}

WORKDIR /root/

COPY --from=server_builder /opt/data-flow-sync/server /root/server
