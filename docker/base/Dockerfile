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

RUN install-packages curl file unzip ca-certificates gnupg lsb-release

# Add docker repo
RUN echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null
RUN curl -fsSL https://download.docker.com/linux/ubuntu/gpg | gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg

# Add the platform user to the docker group for access
RUN groupadd docker && usermod -a -G docker platform

# platform user id
USER 1000

WORKDIR /data
