# Copyright (c) 2024 Joshua Rich <joshua.rich@gmail.com>
# 
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT
FROM docker.io/golang:1.22

WORKDIR /usr/src/go-hass-agent

# copy the src to the workdir
ADD . .

# add dpkg filters
COPY <<EOF /etc/dpkg/dpkg.cfg.d/excludes
# Drop all man pages
path-exclude=/usr/share/man/*
# Drop all translations
path-exclude=/usr/share/locale/*/LC_MESSAGES/*.mo
# Drop all documentation ...
path-exclude=/usr/share/doc/*
# ... except copyright files ...
path-include=/usr/share/doc/*/copyright
# ... and Debian changelogs for native & non-native packages
path-include=/usr/share/doc/*/changelog.*
EOF

# https://developer.fyne.io/started/#prerequisites
ARG DEBIAN_FRONTEND=noninteractive
RUN apt-get update && apt-get -y install gcc pkg-config libgl1-mesa-dev xorg-dev && rm -rf /var/lib/apt/lists/* /var/cache/apt/*

# install build dependencies
RUN go install github.com/matryer/moq@latest && \
  go install golang.org/x/tools/cmd/stringer@latest && \
  go install golang.org/x/text/cmd/gotext@latest

# build the binary
RUN go generate ./... && \
  go build -o /go/bin/go-hass-agent && \
  go clean -cache -modcache && \
  rm -fr /usr/src/go-hass-agent

# remove fyne build dependencies
RUN apt-get -y remove gcc pkg-config libgl1-mesa-dev xorg-dev && apt-get -y autoremove

# reinstall minimum libraries for running
RUN apt-get -y update && apt-get -y install libx11-6 libgl1-mesa-glx dbus-x11 && rm -rf /var/lib/apt/lists/* /var/cache/apt/*

# create a user to run the agent
RUN useradd -ms /bin/bash gouser
USER gouser
WORKDIR /home/gouser

ENTRYPOINT ["go-hass-agent"]
CMD ["--terminal"]
