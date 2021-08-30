FROM ubuntu:21.04

# Create a platform user
RUN useradd -s /bin/bash -d /home/platform -m platform
RUN mkdir /data && chown -R platform:platform /data

# Only for root
COPY install-packages /usr/local/sbin/install-packages
RUN chmod +x /usr/local/sbin/install-packages

# For all users
COPY getfile /usr/local/bin/getfile
RUN chmod +x /usr/local/bin/getfile

RUN install-packages curl file unzip ca-certificates

# platform user id
USER 1000

WORKDIR /data
