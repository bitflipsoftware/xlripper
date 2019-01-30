FROM golang:latest
ENV ROOTDIR="/wrk"
ENV GOPATH="${ROOTDIR}"
ENV REPODIR="${GOPATH}/src/github.com/bitflip-software/xlripper"
COPY . "${REPODIR}"
WORKDIR "${REPODIR}"
ENTRYPOINT go test -v -coverprofile=cover.out -timeout 10m && go tool cover -func=cover.out