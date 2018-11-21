FROM alpine:3.7

ENV HUGO_VERSION=0.51 \
    HUGO_SITE=/srv/hugo

RUN apk update
RUN apk upgrade
RUN apk add bash

RUN apk --no-cache add \
        curl \
        git \
    && curl -SL https://github.com/gohugoio/hugo/releases/download/v${HUGO_VERSION}/hugo_${HUGO_VERSION}_Linux-64bit.tar.gz \
        -o /tmp/hugo.tar.gz \
    && tar -xzf /tmp/hugo.tar.gz -C /tmp \
    && mv /tmp/hugo /usr/local/bin/ \
    && apk del curl \
    && mkdir -p ${HUGO_SITE} \
    && rm -rf /tmp/*

ARG git_user_email
ARG git_user_name

RUN git config --global user.email $git_user_email
RUN git config --global user.name $git_user_name

RUN mkdir /dovetail

COPY build-docs.sh /dovetail

WORKDIR /dovetail

ENTRYPOINT ["/dovetail/build-docs.sh", "magic"]
