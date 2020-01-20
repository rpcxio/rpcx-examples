FROM busybox


RUN mkdir -p /rpcx-service
COPY ./client /rpcx-service/client

WORKDIR /rpcx-service

ENTRYPOINT [ "/rpcx-service/client" ]
