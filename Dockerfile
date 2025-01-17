FROM --platform=$BUILDPLATFORM golang:1.23-alpine3.21 AS build

RUN addgroup -g 1000 -S crow && \
  adduser -u 1000 -G crow -S crow

WORKDIR /src
COPY . .
ARG TARGETOS TARGETARCH
RUN apk add --no-cache -q just curl bash git gcc musl-dev
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg \
    just build

FROM --platform=$BUILDPLATFORM scratch
ENV GODEBUG=netdns=go

# copy certs from build image
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
# copy binary
COPY --from=build /src/dist/crow-autoscale[r] /bin/crow-autoscaler

COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /etc/group /etc/group

USER crow

ENTRYPOINT ["/bin/crow-autoscaler"]
