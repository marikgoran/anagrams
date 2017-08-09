FROM ubuntu
MAINTAINER Goran Marik

ADD anagrams /opt/anagrams
EXPOSE 8080

ENTRYPOINT /opt/anagrams
