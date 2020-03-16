FROM gcc:4.9 AS builder

ENV NAME OPTION

ENV NAME INPUT

ENV NAME OUTPUT

COPY . /huff

WORKDIR /huff

RUN make

FROM ubuntu

COPY --from=builder /huff/app .

COPY --from=builder /huff/runner.sh .

RUN mkdir  output

CMD ["/bin/bash", "./runner.sh"]

