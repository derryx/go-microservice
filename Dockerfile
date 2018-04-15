FROM scratch
EXPOSE 8080
ENTRYPOINT ["/go-microservice"]
COPY ./bin/ /