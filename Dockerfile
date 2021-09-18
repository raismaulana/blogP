FROM golang:alpine

WORKDIR /app

COPY build/package/bin3 bin
COPY build/package/config.yaml .
COPY infrastructure/auth/casbin_model.conf infrastructure/auth/casbin_model.conf
COPY infrastructure/auth/casbin_policy.csv infrastructure/auth/casbin_policy.csv
COPY public/images/l60Hf.png public/images/l60Hf.png

ENTRYPOINT ["/app/bin", "usingdb"]
