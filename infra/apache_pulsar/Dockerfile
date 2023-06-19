# Base image
FROM openjdk:8-jre-slim

# Set environment variables
ENV PULSAR_VERSION=2.9.1
ENV PULSAR_HOME=/pulsar
ENV PATH=$PATH:$PULSAR_HOME/bin

# Install dependencies
RUN apt-get update && apt-get install -y wget

# Download Apache Pulsar binary
RUN wget https://archive.apache.org/dist/pulsar/pulsar-${PULSAR_VERSION}/apache-pulsar-${PULSAR_VERSION}-bin.tar.gz && \
    tar xvfz apache-pulsar-${PULSAR_VERSION}-bin.tar.gz && \
    rm apache-pulsar-${PULSAR_VERSION}-bin.tar.gz && \
    mv apache-pulsar-${PULSAR_VERSION} ${PULSAR_HOME}

# Expose Pulsar ports
EXPOSE 6650
EXPOSE 8080

# Set working directory
WORKDIR ${PULSAR_HOME}

# Start Pulsar standalone
CMD [ "bin/pulsar", "standalone" ]
