FROM busybox


RUN mkdir -p /rpcx-service
COPY ./server /rpcx-service/server

EXPOSE 8972

WORKDIR /rpcx-service

ENTRYPOINT [ "/rpcx-service/server" ]
