FROM gcc:4.9 AS builder

ENV NAME VAR1

ENV NAME VAR2

ENV NAME VAR3

COPY . /huff

WORKDIR /huff

RUN make

RUN ls

FROM ubuntu

COPY --from=builder /huff/app .

COPY --from=builder /huff/runner.sh .

CMD ["/bin/bash", "./runner.sh"]

