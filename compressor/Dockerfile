FROM gcc:4.9 AS builder

COPY . /huff

WORKDIR /huff

RUN make

#FROM alpine:3.7

#COPY --from=builder /huff/app .

#ENTRYPOINT ["./app"]

#######

FROM ubuntu

COPY --from=builder /huff/app .

CMD ["./app"]

