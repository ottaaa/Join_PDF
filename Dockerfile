FROM golang:1.16

WORKDIR /go/src/app
COPY . .

RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends git libgconf-2-4 nodejs npm libgtk-3-0 libgtk-3-dev libxss1 libnss3 libasound2-dev

#RUN go get github.com/pdfcpu/pdfcpu/cmd/...

# RUN go get -d -v ./...
# RUN go install -v ./...

#CMD ["app"]