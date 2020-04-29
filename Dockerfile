FROM centos:7

COPY ./demogo /demogo

RUN chmod 0777 /demogo
EXPOSE 80
CMD ["/demogo"]
