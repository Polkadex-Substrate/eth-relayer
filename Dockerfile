FROM golang:1.14
WORKDIR /opt/relayer
ADD . .
RUN go build -v -o build/polkadex-eth-relay main.go

FROM parity/subkey:2.0.0
COPY --from=0 /opt/relayer/build/polkadex-eth-relay /usr/local/bin/
ENTRYPOINT ["/usr/local/bin/polkadex-eth-relay"]
