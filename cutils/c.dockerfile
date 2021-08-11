FROM golang:1.16
COPY . /home/xxx/
WORKDIR /home/xxx
RUN ls
RUN gcc -o copy ./copy.c
RUN echo 1111 >> a.txt 
RUN touch b.txt
RUN ./copy a.txt b.txt
