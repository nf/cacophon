FROM private/vmruntime-python

ENV DEBIAN_FRONTEND noninteractive
RUN apt-get update && apt-get install  -y --no-install-recommends  --force-yes lame 
ADD . /appgo
WORKDIR /appgo

EXPOSE 8080

ENTRYPOINT ["/appgo/_go_app.bin"]
