FROM ethznetsec/scion_base:latest

ARG SCION_SRC
ARG SC

# Update
RUN sudo apt-get update && sudo apt-get install --assume-yes libpcap0.8 libpcap0.8-dev wget

# Install SCION
RUN git clone ${SCION_SRC} -b scionlab --depth 1 ${SC}

RUN cd ${SC} && ./env/deps

RUN cd ${SC} && make
RUN mkdir ${SC}/gen-cache

# copy setup directory
COPY setup /home/setup
RUN /home/setup/install_jackal.sh

# Add environment variables for execution
ENV PYTHONPATH ${SC}/python
ENV SC ${SC}