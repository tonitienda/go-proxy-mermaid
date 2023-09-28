build-and-deploy:
    docker build --build-arg GO_ARCH=x86_64 --target prod -t tonitienda/go-proxy-mermaid:0.2.0-x86 --platform linux/x86_64 .
    docker build --target prod -t tonitienda/go-proxy-mermaid:0.2.0-arm64 . --no-cache

    docker push tonitienda/go-proxy-mermaid:0.2.0-x86
    docker push tonitienda/go-proxy-mermaid:0.2.0-arm64