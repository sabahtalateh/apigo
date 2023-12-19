FROM ubuntu:latest

# Install protoc & Golang
RUN apt-get update
RUN apt install -y protobuf-compiler
RUN apt install -y golang-go

# Add certificate to install protoc-gen-go
# https://stackoverflow.com/questions/68333944/docker-go-image-cannot-go-get-x509-certificate-signed-by-unknown-authority
RUN apt install -y wget
RUN wget http://www.cisco.com/security/pki/certs/ciscoumbrellaroot.cer
RUN openssl x509 -inform DER -in ciscoumbrellaroot.cer -out ciscoumbrellaroot.crt
RUN cp ciscoumbrellaroot.crt /usr/local/share/ca-certificates/ciscoumbrellaroot.crt
RUN update-ca-certificates

# Install Golang generators
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31 && cp "$(go env GOPATH)/bin/protoc-gen-go" /usr/bin/protoc-gen-go
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3 && cp "$(go env GOPATH)/bin/protoc-gen-go-grpc" /usr/bin/protoc-gen-go-grpc
