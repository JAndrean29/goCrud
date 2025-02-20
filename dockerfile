#New Scripts attempting for a multi-stage Building
#Building Stage
#Running go build since I run from Windows machine
FROM golang:alpine3.21 AS build

#requires to add libc-dev to use gcc to able to compile the program with CGO_ENABLED=1
RUN apk --no-cache add tzdata gcc libc-dev

ENV TZ=asia/jakarta

COPY go.mod go.sum ./
RUN go mod download

#copy config
COPY config ./config

#copy di
COPY di ./di

#copy handler
COPY handler ./handler

#copy infrastructure
COPY infrastructure ./infrastructure

#copy mode
COPY model ./model

#copy repository
COPY repository ./repository

#copy usecase
COPY usecase ./usecase

#copy main.go
COPY main.go .

ENV CGO_ENABLED=1

RUN go build -o /bin/main main.go

RUN ls -la

#Running Stage
FROM alpine AS running

WORKDIR /app

RUN apk --no-cache add sqlite

COPY --from=build bin .
COPY --from=build usr/share/zoneinfo .

COPY .env .

RUN ls -la

ENTRYPOINT [ "./main", "-p", "8080" ]

EXPOSE 8080

#use FROM golang because just running golang
#FROM golang:latest

#BEST PRACTICE USE COPY
#ANOTHER NOTE COPY ALL NECESSARY FILES FIRST
#PREFREABLY THOSE THAT ARE CONTAINING DEPENDENCIES
#THIS WILL SAVE BUILDING TIME, ASSUMING THERE IS NO CHANGES TO DEPENDENCIES

#Add workdir for easier directory reading
#WORKDIR /app

#copy go mod and go sum for caching purpose
#COPY go.mod go.sum ./
#initiate go mod download for dependency changes
#RUN go mod download

#copy api_docs
#COPY api_docs ./api_docs

#copy di
#COPY di ./di

#copy handler
#COPY handler ./handler

#copy infrastructure
#COPY infrastructure ./infrastructure

#copy mode
#COPY model ./model

#copy repository
#COPY repository ./repository

#copy usecase
#COPY usecase ./usecase

#copy db
#COPY users.db .

#copy main.go
#COPY main.go .

#build go project
#RUN go build -o main ./main.go

#execute command on run finish
#CMD [ "./main" ]

#expose to desireable port, still needs to use -p 8080:8080 command on docker run
#EXPOSE 8080

#THIS IS A FILE THAT WILL EXECUTES THE LISTED CODES
#BUILDING AND RUNNING A DOCKER IS BASICALLY STARTING A FRESH VMWARE
#NO FILE NO CONFIG
#RECOMMENDED TO RUN INITIATION COMMANDS LIKE NPM INSTALL
#OR GO MOD DOWNLOAD IF SHOULD THERE'S A DEPENDENCY CHANGESM
