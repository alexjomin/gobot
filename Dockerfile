# minimal linux distribution
FROM golang:1.7.3-wheezy

# set the go path to import the source project
WORKDIR $GOPATH/src/github.com/alexjomin/gobot
ADD . $GOPATH/src/github.com/alexjomin/gobot

# Install Linter tools
RUN go get -u github.com/alecthomas/gometalinter && gometalinter --install

# Disable Host Checking
RUN mkdir /root/.ssh && touch /root/.ssh/config && echo "Host *\n\tStrictHostKeyChecking no\n" >> ~/.ssh/config && chmod 400 ~/.ssh/config

# Build Binary and remove sources
RUN make all && rm -rf $GOPATH/pkg && rm -rf $GOPATH/src