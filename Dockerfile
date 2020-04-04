FROM scratch

ENV PORT 8000
EXPOSE $PORT

COPY simple-tds /
CMD ["/simple-tds"]
