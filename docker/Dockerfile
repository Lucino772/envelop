# syntax=docker/dockerfile:1.3-labs
ARG BASE_IMAGE="debian-base"
ARG HATCH_VERSION="1.9.1"
ARG CARGO="cargo"
ARG PYAPP_REPO="pyapp"
ARG PYAPP_VERSION="0.15.1"

# Base Image for Debian Based build
FROM debian:buster-slim as debian-base
RUN apt update && apt upgrade -y --fix-missing && apt install build-essential wget curl -y
ENV CARGO_BUILD_TARGET="x86_64-unknown-linux-gnu"

# Base Image for Alpine Base build
FROM alpine:latest as alpine-build
RUN apk -U upgrade && apk add --update alpine-sdk curl wget
ENV CARGO_BUILD_TARGET="x86_64-unknown-linux-musl"

# Build Image
FROM ${BASE_IMAGE} as build
WORKDIR /build

# Install PyApp
RUN mkdir -p ${PYAPP_REPO} \
    && curl -L https://github.com/ofek/pyapp/releases/download/v${PYAPP_VERSION}/source.tar.gz \
    | tar --strip-components=1 -xzf - -C ${PYAPP_REPO}

# Install Hatch
RUN wget -O hatch-${HATCH_VERSION}-${CARGO_BUILD_TARGET}.tar.gz https://github.com/pypa/hatch/releases/download/hatch-v${HATCH_VERSION}/hatch-${HATCH_VERSION}-${CARGO_BUILD_TARGET}.tar.gz \
    && tar -xzf hatch-${HATCH_VERSION}-${CARGO_BUILD_TARGET}.tar.gz -C /usr/local/bin \
    && rm hatch-${HATCH_VERSION}-${CARGO_BUILD_TARGET}.tar.gz \
    && mv /usr/local/bin/hatch-${HATCH_VERSION}-${CARGO_BUILD_TARGET} /usr/local/bin/hatch \
    && hatch self restore

# Install Rust
RUN curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y
ENV PATH="/root/.cargo/bin:${PATH}" CARGO_TERM_COLOR="always" CARGO_INCREMENTAL="0"
RUN rustup toolchain install stable \
    --target ${CARGO_BUILD_TARGET} \
    --no-self-update \
    --profile minimal \
    && rustup default stable  

# Configure PyApp
RUN cat <<EOF >> ./script.sh
config_file="$PYAPP_REPO/.cargo/config_$CARGO_BUILD_TARGET.toml"
if [ -f "$config_file" ]; then
    mv "$config_file" "$PYAPP_REPO/.cargo/config.toml"
fi
EOF

# Build Envelop
COPY . .
RUN hatch build --target wheel
RUN export PYAPP_PROJECT_PATH="$(pwd)/dist/$(cd dist && echo *.whl)" \
    && hatch build --target app \
    && mv "$(pwd)/dist/app/envelop-$(hatch version)-$CARGO_BUILD_TARGET" /bin/envelop

FROM scratch AS envelop
COPY --from=build /bin/envelop /bin/envelop