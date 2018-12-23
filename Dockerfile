# BUILD
FROM golang:1.11.4-alpine as build-stage
ARG version
ARG revision
RUN touch /${version} && touch /${revision}
COPY . /go/src/ehw2018/greetings-server
RUN go install -ldflags="-s -w -X \"main.Version=${version}\" -X \"main.Revision=${revision}\" -extldflags \"-static\"" ehw2018/greetings-server

# Production
FROM alpine as production-stage
COPY --from=build-stage /go/bin/greetings-server .
ENV PORT 8080
ENV GIN_MODE=release
CMD ["./greetings-server"]

