FROM alpine:3.9
RUN apk add --no-cache libc6-compat curl

ENV PORT 8000
EXPOSE $PORT

COPY simple-tds /

CMD ["./simple-tds"]
