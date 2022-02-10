FROM alpine:latest
COPY exec/caching-service .
RUN chmod u=rwx caching-service
CMD [ "./caching-service", "watch" ]