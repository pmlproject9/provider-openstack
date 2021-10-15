FROM alpine:3.7
WORKDIR /
COPY ./bin/openstack-provider-controller  .

EXPOSE 8080
ENTRYPOINT ["/openstack-provider-controller"]
