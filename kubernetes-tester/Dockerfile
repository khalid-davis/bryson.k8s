FROM ubuntu:14.04

COPY kubernetes-tester /app/kuberentes-tester
COPY conf /app/conf

EXPOSE 8080

WORKDIR /app

CMD ./kubernetes-tester
