#use FROM golang because just running golang
FROM golang:latest

#BEST PRACTICE USE COPY
#ANOTHER NOTE COPY ALL NECESSARY FILES FIRST
#PREFREABLY THOSE THAT ARE CONTAINING DEPENDENCIES
#THIS WILL SAVE BUILDING TIME, ASSUMING THERE IS NO CHANGES TO DEPENDENCIES

#Add workdir for easier directory reading
WORKDIR /app

#copy go mod and go sum for caching purpose
COPY go.mod go.sum ./
#initiate go mod download for dependency changes
RUN go mod download

#copy api_docs
COPY api_docs ./api_docs

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

#copy db
COPY users.db .

#copy main.go
COPY main.go .

#build go project
RUN go build -o main ./main.go

#execute command on run finish
CMD [ "./main" ]

#expose to desireable port, still needs to use -p 8080:8080 command on docker run
EXPOSE 8080

#THIS IS A FILE THAT WILL EXECUTES THE LISTED CODES
#BUILDING AND RUNNING A DOCKER IS BASICALLY STARTING A FRESH VMWARE
#NO FILE NO CONFIG
#RECOMMENDED TO RUN INITIATION COMMANDS LIKE NPM INSTALL
#OR GO MOD DOWNLOAD IF SHOULD THERE'S A DEPENDENCY CHANGESM
