FROM golang:1.16
COPY . /home/xxx/
WORKDIR /home/xxx
RUN ls
ENTRYPOINT [ "./entrypoint.sh","main" ]
