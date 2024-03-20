###################
#      Wheel      #
###################
FROM debian:buster-slim as wheel

RUN apt update \
    && apt upgrade -y \
    && apt install wget -y \
    && apt clean

# Install Hatch
ARG HATCH_VERSION="1.9.4"
ARG HATCH_TARGET="x86_64-unknown-linux-gnu"

WORKDIR /tmp/hatch
RUN wget \
        -O hatch-${HATCH_VERSION}-${HATCH_TARGET}.tar.gz \
        https://github.com/pypa/hatch/releases/download/hatch-v${HATCH_VERSION}/hatch-${HATCH_VERSION}-${HATCH_TARGET}.tar.gz \
    && tar -xzf hatch-${HATCH_VERSION}-${HATCH_TARGET}.tar.gz -C . \
    && rm hatch-${HATCH_VERSION}-${HATCH_TARGET}.tar.gz \
    && mv hatch-${HATCH_VERSION}-${HATCH_TARGET} /bin/hatch

RUN hatch self restore

# Build wheel
WORKDIR /tmp/envelop

COPY . .
RUN hatch build --target wheel \
    && mkdir -p /opt/envelop/ \
    && mv "$(pwd)/dist/$(cd dist && echo *.whl)" /opt/envelop/

#####################
#      Envelop      #
#####################
FROM rust:slim-buster as envelop

RUN apt update \
    && apt upgrade -y \
    && apt install --no-install-recommends -y ca-certificates curl build-essential \
    && apt clean

# Install docker-ce-cli
RUN \
    for pkg in docker.io docker-doc docker-compose podman-docker containerd runc; do apt-get remove $pkg; done \
    && apt update \
    && install -m 0755 -d /etc/apt/keyrings \
    && curl -fsSL https://download.docker.com/linux/debian/gpg -o /etc/apt/keyrings/docker.asc \
    && chmod a+r /etc/apt/keyrings/docker.asc \
    && echo \
        "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/debian $(. /etc/os-release && echo "$VERSION_CODENAME") stable" \
        | tee /etc/apt/sources.list.d/docker.list > /dev/null \
    && apt update \
    && apt install --no-install-recommends -y docker-ce-cli \
    && apt clean

# Install cross-rs
ENV CROSS_CONTAINER_IN_CONTAINER=true
RUN cargo install cross

# Install PyApp
ARG PYAPP_REPO="/tmp/pyapp"
ARG PYAPP_VERSION="0.15.1"

RUN mkdir -p ${PYAPP_REPO} \
    && curl -L https://github.com/ofek/pyapp/releases/download/v${PYAPP_VERSION}/source.tar.gz \
    | tar --strip-components=1 -xzf - -C ${PYAPP_REPO}

# Build App
WORKDIR ${PYAPP_REPO}
COPY --from=wheel /opt/envelop/ ./

ENV TARGET_DIR="/tmp/envelop/build"
ENV ENVELOP_OUT_DIR="/opt/envelop/bin"
ENV CARGO_BUILD_TARGET="x86_64-unknown-linux-gnu"
ENV PYAPP_EXEC_SPEC="envelop.cli:cli"

RUN mkdir -p ${TARGET_DIR} \
    && mkdir -p ${ENVELOP_OUT_DIR}

RUN cat <<-"EOT" | tr -d '\r' > ./build.sh
config_file="$PYAPP_REPO/.cargo/config_$CARGO_BUILD_TARGET.toml" \
&& if [ -f "$config_file" ]; then mv "$config_file" "$PYAPP_REPO/.cargo/config.toml"; fi \
&& PYAPP_PROJECT_PATH="$(echo *.whl)" cross build --release --target-dir ${TARGET_DIR} \
&& filename=$(find -O3 "${TARGET_DIR}/${CARGO_BUILD_TARGET}/release/" -maxdepth 1 -name "pyapp.exe") \
&& extension=$([[ "$filename" = *.* ]] && echo ".${filename##*.}" || echo '') \
&& mv ${TARGET_DIR}/${CARGO_BUILD_TARGET}/release/pyapp${extension} ${ENVELOP_OUT_DIR}/envelop-${CARGO_BUILD_TARGET}${extension}
EOT

ENTRYPOINT [ "bash", "./build.sh" ]