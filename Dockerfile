FROM golang:1.22-alpine3.18 as builder
WORKDIR /app
COPY . .
RUN go build -tags musl -mod=vendor -o /app/build/user /app/cmd/exercise

FROM alpine:3.18
RUN mkdir /app
EXPOSE 9090
COPY --from=builder /app/build/exercise /app
ARG version
RUN echo $version > /app/version.txt
CMD /app/exercise
