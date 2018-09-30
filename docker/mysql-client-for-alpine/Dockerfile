FROM alpine:latest

RUN apk add --no-cache git cmake make gcc musl-dev g++ boost-dev openssl-dev ncurses-dev rpcgen libtirpc-dev bison

RUN git clone --branch 8.0 --depth 1 https://github.com/mysql/mysql-server.git \
    && cd mysql-server \
    && cmake . -DDOWNLOAD_BOOST=1 -DWITH_BOOST=/usr/include/boost \
    && cd client \
    && make \
    && make install
