### cli build stage ###
FROM ubuntu AS http_builder

ARG asmttpd_version=0.4.5
ARG elfkickers_version=3.1a

RUN apt-get update && apt-get install -y build-essential yasm

ADD https://github.com/nemasu/asmttpd/archive/${asmttpd_version}.tar.gz /tmp
RUN tar -C /tmp -xf /tmp/${asmttpd_version}.tar.gz
RUN sed -i -r '/;-----Simple request logging/,/;-----End Simple logging/s/;([^-])/\1/' /tmp/asmttpd-${asmttpd_version}/main.asm
RUN make -C /tmp/asmttpd-${asmttpd_version} release
RUN cp /tmp/asmttpd-${asmttpd_version}/asmttpd /tmp

ADD http://www.muppetlabs.com/~breadbox/pub/software/ELFkickers-${elfkickers_version}.tar.gz /tmp
RUN tar -C /tmp -xf /tmp/ELFkickers-${elfkickers_version}.tar.gz
RUN make -C /tmp/ELFkickers-${elfkickers_version}
RUN /tmp/ELFkickers-${elfkickers_version}/sstrip/sstrip -z /tmp/asmttpd


### cli build stage ###
FROM golang:1.14 as cli_builder
WORKDIR /app
COPY . .
RUN go build \
  -ldflags "-linkmode external -extldflags -static" \
  -a main.go

### runtime stage ###
FROM scratch

COPY --from=http_builder /tmp/asmttpd ./asmttpd
COPY --from=cli_builder /app/main ./cli

EXPOSE 8080

CMD [ "./asmttpd", "/data", "8080"]