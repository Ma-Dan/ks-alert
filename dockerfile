FROM kubesphere/alertmanager-builder:0.2 as builder

COPY .  /root/
WORKDIR /root/
COPY . .

#RUN go get -u "github.com/go-openapi/spec" \
#go get -u "github.com/emicklei/go-restful/log" \
#go get -u "github.com/emicklei/go-restful-openapi" \
#go get -u "github.com/emicklei/go-restful"
WORKDIR /root/restapi/api/
RUN go build alertmanagerserver.go
RUN chmod 775 alertmanagerserver

FROM alpine:3.6
RUN apk add --update ca-certificates && update-ca-certificates
COPY --from=builder /root/restapi/api/alertmanagerserver /usr/local/bin/

expose 8080
ENTRYPOINT  /usr/local/bin/alertmanagerserver
CMD ["sh"]
