# Initial preparations
FROM golang:1.11-alpine3.8 as builder

ENV EKSCTL_VERSION "0.1.29"
ENV KUBECTL_VERSION v1.11.5
ENV AWS_AUTH_VER "1.12.7/2019-03-27"

RUN apk update && apk add openssh ca-certificates make git curl

# Let's add the code in
ADD . /go/src/github.com/samsung-cnct/cma-aws

# install eksctl
RUN mkdir -p /usr/local/bin
RUN curl --silent --location "https://github.com/weaveworks/eksctl/releases/download/${EKSCTL_VERSION}/eksctl_Linux_amd64.tar.gz" | tar xz -C /tmp \
  && mv /tmp/eksctl /usr/local/bin && chmod +x /usr/local/bin/eksctl
# install kubectl
RUN curl --silent --location "https://dl.k8s.io/${KUBECTL_VERSION}/bin/linux/amd64/kubectl" -o /usr/local/bin/kubectl \
    && chmod +x /usr/local/bin/kubectl
# install aws-iam-authenticator
RUN curl -L https://amazon-eks.s3-us-west-2.amazonaws.com/${AWS_AUTH_VER}/bin/linux/amd64/aws-iam-authenticator \
  -o /usr/local/bin/aws-iam-authenticator && \
  chmod +x /usr/local/bin/aws-iam-authenticator

# Build package in container
WORKDIR /go/src/github.com/samsung-cnct/cma-aws
RUN make -f build/Makefile container-build

# Now that we have build a static binary, we can stuff it in a bare-bones alpine image
# But we will need bash and openssh right now
# ca-certificates is needed for aws library
FROM alpine:3.8
RUN apk update && apk add bash openssh ca-certificates jq
COPY --from=builder /go/src/github.com/samsung-cnct/cma-aws/cma-aws /cma-aws
COPY --from=builder /usr/local/bin/kubectl /usr/local/bin/kubectl
COPY --from=builder /usr/local/bin/aws-iam-authenticator /usr/local/bin/aws-iam-authenticator
COPY --from=builder /usr/local/bin/eksctl /usr/local/bin/eksctl
ENTRYPOINT ["/cma-aws"]