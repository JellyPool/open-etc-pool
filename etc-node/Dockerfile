# Build Geth in a stock Go builder container
FROM golang:1.17-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git

RUN git clone https://github.com/etclabscore/core-geth.git getc && \ 
    cp ./getc /go-ethereum -r

RUN cd /go-ethereum && make geth

# Pull Geth into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /go-ethereum/build/bin/geth /usr/local/bin/

EXPOSE 8545 8546 30303 30303/udp
# ENTRYPOINT ["geth", "--classic", "--rpc", "--http.addr 0.0.0.0", "console"] 
# --rpccorsdomain "*" --rpcvhosts=* has been removed since an attack from this address: 0x6Ef57BE1168628A2bD6c5788322A41265084408a
CMD geth --classic --nousb --http --http.port "8545" --http.addr "0.0.0.0" --http.corsdomain "*" --http.vhosts=* -http.api personal,web3,eth --miner.etherbase "0x2b96eF48ab692BDB2C6A4734554f93e174EB4497" --mine
