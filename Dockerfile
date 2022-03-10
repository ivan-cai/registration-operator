FROM golang:1.17 AS builder

ARG ARCH

WORKDIR /go/src/open-cluster-management.io/registration-operator
COPY . .
ENV GO_PACKAGE open-cluster-management.io/registration-operator
RUN CGO_ENABLED=0 GOOS=linux GOARCH=$ARCH go build -v --ldflags ' -extldflags "-static"' -o registration-operator  ./cmd/registration-operator/main.go

FROM alpine:latest
ENV USER_UID=10001

COPY --from=builder /go/src/open-cluster-management.io/registration-operator/registration-operator /
#RUN microdnf update && microdnf clean all

USER ${USER_UID}
