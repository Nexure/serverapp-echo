# Use the official go docker image built on debian.
FROM golang:1.8

ENV PATHWORK=/go/src/github.com/rewiko/app/
WORKDIR $PATHWORK

# reload code
RUN go get github.com/pilu/fresh

# install glide to manage dependencies
ENV GLIDEVERSION=0.12.3
RUN wget https://github.com/Masterminds/glide/releases/download/v${GLIDEVERSION}/glide-v${GLIDEVERSION}-linux-amd64.tar.gz
RUN mkdir glide-install ; tar xzf glide-v${GLIDEVERSION}-linux-amd64.tar.gz -C glide-install
RUN mv glide-install/linux-amd64/glide /usr/local/bin/ ; rm -rf glide-install

# Grab the source code and add it to the workspace.
ADD ./src/ $PATHWORK

RUN glide install -v

ADD ./entrypoint.sh /
ADD ./fresh.conf /
RUN chmod 755 /entrypoint.sh
CMD /entrypoint.sh

# Open up the port where the app is running.
EXPOSE 8080
