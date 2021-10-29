FROM golang:1.16-alpine

#Defines you application repository
ENV APPLICATION_PACKAGE=./cmd/internal

ENV IGNORE_GO_GET=true
