FROM debian:bookworm AS base

RUN --mount=type=cache,target=/var/cache/apt,sharing=locked \
    --mount=type=cache,target=/var/lib/apt,sharing=locked \
    apt-get update && \
    apt-get --no-install-recommends install -y \
      build-essential \
      cmake  \
      ca-certificates  \
      curl

FROM base AS wla-dx

ARG WLA_DX_VERISON="v10.6"

RUN curl --location --output "/tmp/wla-dx.tar.gz" "https://github.com/vhelin/wla-dx/archive/${WLA_DX_VERISON}.tar.gz"

RUN --mount=type=cache,target=/opt/wla-dx/CMakeFiles \
    mkdir -p /opt/wla-dx && \
    tar -zxvf "/tmp/wla-dx.tar.gz" --strip-components=1 -C /opt/wla-dx && \
    cd /opt/wla-dx && \
    cmake . && \
    make

ENV PATH=${PATH}:/opt/wla-dx/binaries

FROM wla-dx AS mooneye-test-suite

ARG MOONEYE_TEST_SUITE_VERSION=main

RUN curl --location --output "/tmp/mooneye-test-suite.tar.gz" "https://github.com/Gekkio/mooneye-test-suite/archive/${MOONEYE_TEST_SUITE_VERSION}.tar.gz"

RUN mkdir -p /opt/mooneye && \
    tar -zxvf "/tmp/mooneye-test-suite.tar.gz" --strip-components=1 -C /opt/mooneye && \
    cd /opt/mooneye && \
    ls -lah . && \
    make clean all

COPY mooneye_tests.txt /
COPY copy_mooneye_tests.sh /
