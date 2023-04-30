FROM mcr.microsoft.com/vscode/devcontainers/base:debian

RUN apt-get update \
    && sudo apt-get install gnupg2 -y \
    && apt-get install jq build-essential -y \
    # Clean up
    && apt-get autoremove -y && apt-get clean -y && rm -rf /var/lib/apt/lists/*
