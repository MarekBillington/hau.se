FROM rust:1.66 as builder
WORKDIR /app
COPY . .
RUN cargo install --path .

FROM debian:buster-slim as runner
COPY --from=builder /usr/local/cargo/bin/hause /usr/local/cargo/bin/hause
ENV ROCKET_ADDRESS=0.0.0.0
EXPOSE 8000
CMD ["hause"]
