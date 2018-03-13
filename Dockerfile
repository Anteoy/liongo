FROM hub.c.163.com/library/centos:7
MAINTAINER anteoy "anteoy@gmail.com"

ADD appServer /home/app/server/
ADD resources /home/app/resources/
WORKDIR  /home/app/server/
EXPOSE 8080
ENTRYPOINT ["./appServer","run","--note"]
