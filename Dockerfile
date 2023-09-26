ARG golang_version=1.20.8
ARG alpine_version=3.18
ARG node_version=18

FROM golang:${golang_version}-alpine${alpine_version} as server_builder
ARG go_swagger_version=0.27.0

RUN apk add --no-cache git
RUN apk add --update --no-cache curl build-base
RUN curl -sSLo /usr/local/bin/swagger https://github.com/go-swagger/go-swagger/releases/download/v${go_swagger_version}/swagger_linux_amd64
RUN chmod +x /usr/local/bin/swagger
RUN apk update && apk add openssh

ENV REPO_DIR ${GOPATH}/src/github.com/quocbang/data-flow-sync
ENV SERVER_DIR ${REPO_DIR}/server
ENV APP_NAME data-flow-sync

COPY ./swagger.yml ${REPO_DIR}/
COPY ./server ${SERVER_DIR}/

WORKDIR ${SERVER_DIR}

RUN go generate .
RUN go mod download
RUN go vet ./...
# RUN go test -race -gcflags -l -coverprofile .testCoverage.txt ./...
# RUN go tool cover -func .testCoverage.txt
RUN go build -race -ldflags "-extldflags '-static'" -o /opt/data-flow-sync ./swagger/cmd/${APP_NAME}-server

CMD ["/bin/sh"]

FROM node:${node_version}-alpine${alpine_version} as ui_builder
ENV UI_DIR /app
COPY ./ui ${UI_DIR}

WORKDIR ${UI_DIR}

RUN rm -rf node_models
RUN npm install
RUN npm run build

CMD ["/bin/sh"]

# deploy
FROM alpine:${alpine_version}

WORKDIR /root/

COPY --from=server_builder /opt/data-flow-sync/server /root/server
COPY --from=ui_builder /app/dist /root/ui
