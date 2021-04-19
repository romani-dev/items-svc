FROM alpine:3
WORKDIR /items
COPY bin/app ./app
CMD [ "/items/app" ]