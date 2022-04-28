FROM alpine
MAINTAINER "OCEANGZY"
WORKDIR /go/web/oceam_im
COPY . /go/web/oceam_im
#RUN go env -w GO111MODULE=on
#RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN chmod -R 777 /go/web/oceam_im/ocean_im_server
RUN chmod -R 777 /go/web/oceam_im/resources/conf.yaml
EXPOSE 8080
#CMD ["go run main.go"]
CMD ["./ocean_im_server"]