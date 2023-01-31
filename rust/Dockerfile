FROM rust:1.66 as builder
WORKDIR app
COPY . .
RUN cargo install --path .

FROM debian:buster-slim as runner
ARG APP=/user/src/app
RUN mkdir -p ${APP}

# Copy the compiled binaries into the new container.
COPY --from=builder /usr/local/cargo/bin/hause ${APP}/hause

ENV ROCKET_ADDRESS=0.0.0.0
WORKDIR ${APP}

EXPOSE 8000

CMD ["./hause"]
