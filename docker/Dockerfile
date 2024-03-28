###################
#      Hatch      #
###################
FROM debian:buster-slim as hatch

RUN apt update \
    && apt upgrade -y \
    && apt install --no-install-recommends -y ca-certificates wget \
    && apt clean

ARG HATCH_VERSION="1.9.4"
ARG HATCH_TARGET="x86_64-unknown-linux-gnu"

RUN mkdir -p /tmp/hatch \
    && wget \
        -O /tmp/hatch/hatch-${HATCH_VERSION}-${HATCH_TARGET}.tar.gz \
        https://github.com/pypa/hatch/releases/download/hatch-v${HATCH_VERSION}/hatch-${HATCH_VERSION}-${HATCH_TARGET}.tar.gz \
    && tar -xzf /tmp/hatch/hatch-${HATCH_VERSION}-${HATCH_TARGET}.tar.gz -C /tmp/hatch/ \
    && rm /tmp/hatch/hatch-${HATCH_VERSION}-${HATCH_TARGET}.tar.gz \
    && mv /tmp/hatch/hatch-${HATCH_VERSION}-${HATCH_TARGET} /bin/hatch \
    && rm -r /tmp/hatch

##################
#      Base      #
##################
FROM rust:slim-buster as base

RUN apt update \
    && apt upgrade -y \
    && apt install --no-install-recommends -y ca-certificates curl build-essential wget \
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
ENV CARGO=cross

# Install PyApp
ENV PYAPP_REPO="/opt/pyapp"
ARG PYAPP_VERSION="0.15.1"

RUN mkdir -p ${PYAPP_REPO} \
    && curl -L https://github.com/ofek/pyapp/releases/download/v${PYAPP_VERSION}/source.tar.gz \
    | tar --strip-components=1 -xzf - -C ${PYAPP_REPO}

###################
#      Build      #
###################
FROM base as envelop

COPY --from=hatch /bin/hatch /bin/hatch

ENV CARGO_BUILD_TARGET="x86_64-unknown-linux-gnu"
ENV PYAPP_EXEC_SPEC="envelop.cli:cli"

RUN cat <<-"EOT" | tr -d '\r' > ./build.sh
target_dir="/tmp/envelop/build"
build_dir="/opt/build"

config_file="$PYAPP_REPO/.cargo/config_$CARGO_BUILD_TARGET.toml"
if [ -f "$config_file" ]; then
    mv "$config_file" "$PYAPP_REPO/.cargo/config.toml";
fi

if [ ! -d "$build_dir/envelop" ]; then
    echo "Mount the project source code to '$build_dir/envelop'";
    exit 1;
fi

# Build wheel using hatch
cd "$build_dir/envelop"
hatch build --target wheel
wheel="$build_dir/envelop/dist/$(cd $build_dir/envelop/dist && echo *.whl)"
if [ ! -f "$wheel" ]; then
    echo "Could not find wheel in '$build_dir/envelop/dist'";
    exit 1;
fi
mv "$wheel" "$PYAPP_REPO/"

# Build executable using PyApp
cd "$PYAPP_REPO"
PYAPP_PROJECT_PATH="$(cd $PYAPP_REPO && echo *.whl)" cross build --release --target-dir "$target_dir"
filename=$(find -O3 "$target_dir/$CARGO_BUILD_TARGET/release/" -maxdepth 1 -name "pyapp.exe")
extension=$([[ "$filename" = *.* ]] && echo ".${filename##*.}" || echo '')
temp_exe="$target_dir/$CARGO_BUILD_TARGET/release/pyapp$extension"

if [ ! -f "$temp_exe" ]; then
    echo "Failed to find pyapp executable";
    exit 1;
fi

exe_path="$build_dir/envelop/dist/binary/envelop-${CARGO_BUILD_TARGET}${extension}"
if [ ! -d "$build_dir/envelop/dist/binary" ]; then
    mkdir -p "$build_dir/envelop/dist/binary";
fi
mv -f "$temp_exe" "$exe_path"

EOT
ENTRYPOINT [ "bash", "./build.sh" ]