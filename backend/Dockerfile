FROM private/vmruntime-go

ENV DEBIAN_FRONTEND noninteractive

ADD app.yaml /appgo/
ADD install_deps.sh /appgo/
RUN /appgo/install_deps.sh /appgo/app.yaml

ADD . /appgo
WORKDIR /appgo

EXPOSE 8080

ENTRYPOINT ["/appgo/_go_app.bin"]
