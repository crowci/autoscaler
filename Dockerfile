FROM --platform=$BUILDPLATFORM golang:1.23 AS build

RUN groupadd -g 1000 crow && \
  useradd -u 1000 -g 1000 crow && \
  mkdir -p /etc/crow && \
  chown -R crow:crow /etc/crow

WORKDIR /src
COPY . .
ARG TARGETOS TARGETARCH
# install just FIXME: 'apt install -y just' from Debian >= 13
# hadolint ignore=DL4006
RUN curl --proto '=https' --tlsv1.2 -sSf https://just.systems/install.sh | bash -s -- --to /usr/bin/
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg \
    just build

FROM --platform=$BUILDPLATFORM scratch
ENV GODEBUG=netdns=go

# copy certs from build image
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
# copy agent binary
COPY --from=build /src/dist/crow-autoscaler /bin/

USER crow

ENTRYPOINT ["/bin/crow-autoscaler"]
