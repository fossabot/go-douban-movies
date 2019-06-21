FROM golang:latest
COPY . /app
WORKDIR /app
ENV GOPROXY=https://goproxy.io
RUN [ "go", "build"]
EXPOSE 8888
CMD [ "./go-douban-movies" ]