FROM alpine:latest
COPY ./exec/caching-service .
CMD [ "./caching-service" ]