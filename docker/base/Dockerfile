FROM ubuntu:21.10

# Create a platform user
RUN useradd -s /bin/bash -d /home/platform -m platform
RUN mkdir /data && chown -R platform:platform /data

# Only for root
COPY install-packages /usr/local/sbin/install-packages
RUN chmod +x /usr/local/sbin/install-packages

# For all users
COPY getfile /usr/local/bin/getfile
RUN chmod +x /usr/local/bin/getfile

RUN install-packages curl file unzip ca-certificates gnupg lsb-release software-properties-common

# Add docker repo
RUN curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -
RUN add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"

# Add the platform user to the docker group for access
RUN groupadd docker && usermod -a -G docker platform

# platform user id
USER 1000

WORKDIR /data
