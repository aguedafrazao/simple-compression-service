# Getting GCC to compile executable
FROM gcc:4.9 AS builder

# Var to inform program option
ENV NAME OPTION

# Var to inform input path
ENV NAME INPUT

# Var to inform output path
ENV NAME OUTPUT

# Copying all code to /huff
COPY . /huff

# Setting /huff as work dir
WORKDIR /huff

# Build executable
RUN make

# Getting running environemtn
FROM ubuntu

# Copying executable file
COPY --from=builder /huff/app .

# Copying bash file to pass arguments to executable
COPY --from=builder /huff/runner.sh .

# Creating a directory to bind with host directory
RUN mkdir  output

# Running script to pass args to executable
CMD ["/bin/bash", "./runner.sh"]

