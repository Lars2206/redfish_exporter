FROM docker.io/library/centos:7

LABEL maintainer="Jennings Liu <jenningsloy318@gmail.com>"

ARG ARCH=amd64

ENV GOROOT /usr/local/go
ENV GOPATH /go
ENV PATH "$GOROOT/bin:$GOPATH/bin:$PATH"
ENV GO_VERSION 1.14.2
ENV GO111MODULE=on 
ENV GOPROXY=https://goproxy.cn


# Build dependencies

RUN yum install -y  rpm-build make  git && curl -SL https://dl.google.com/go/go${GO_VERSION}.linux-${ARCH}.tar.gz | tar -xzC /usr/local

#Own Code
RUN git clone https://github.com/Lars2206/redfish_exporter.git /app

#RUN cd /app
#RUN ls /app
RUN make build --directory /app
#RUN ls /app/scripts
RUN rm -rf /etc/yum*
RUN rm -rf /var/lib/yum/*
EXPOSE 9610 

CMD /app/build/redfish_exporter --config.file=/etc/redfish_exporter/redfish_exporter.yml
#--config.file=/app/scripts/redfish_exporter.yml
#--config.file=/etc/redfish_exporter/redfish_exporter.yml
