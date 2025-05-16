# ---------- Stage 1: Build ----------
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy source
COPY server/go.mod ./
COPY server/ ./server/
COPY frontend/ ./frontend/

# Build static binary
RUN cd server && go build -o /app/app .

# ---------- Stage 2: Runtime ----------
FROM gcr.io/distroless/static:nonroot

WORKDIR /

COPY --from=builder /app/app .
COPY --from=builder /app/frontend /frontend

USER nonroot:nonroot

EXPOSE 8080
ENTRYPOINT ["/app"]