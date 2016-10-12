FROM alpine:latest

MAINTAINER Edward Muller <edward@heroku.com>

WORKDIR "/opt"

ADD .docker_build/line-go-bot /opt/bin/line-go-bot
ADD ./templates /opt/templates
ADD ./static /opt/static

CMD ["/opt/bin/line-go-bot"]
