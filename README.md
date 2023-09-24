# go-proxy-mermaid

HTTP Proxy server used for demos

## Usage

### Env Vars

- `PORT`: Port to listen on (default: 8080)
- `PROXY_URL`: URL to proxy to (default: http://localhost:8081). Can be a comma-separated list of URLs to forward the request to multiple servers.
- `NAME`: Name of the server (default: proxy)


### Responses

#### JSON



#### HTML

Mermaid diagram with the chain of requests/responses:

```mermaid
graph LR
    df25bc50("df25bc50
    service1")
    a5da30cd("a5da30cd
    service2")
    bef70617("bef70617
    service3")
    6dbb8570("6dbb8570
    service4")
    df25bc50 -- service2:8081 --> a5da30cd
    df25bc50 -- service3:8082 --> bef70617
    bef70617 -- service4:8083 --> 6dbb8570
    classDef Level0 stroke-width:2px,stroke:#1382c3,fill:#88cef7,color:#1382c3;
    classDef Level1 stroke-width:1px,stroke:#325c77,fill:#79bae1,color:#325c77;
    classDef Level2 stroke-width:1px,stroke:#06f,fill:#6aa6cb;
    classDef Level3 stroke-width:1px,stroke:#04f,fill:#5c93b5;
    classDef Level4 stroke-width:1px,stroke:#00f,fill:#4d80a0;
    classDef Level5 stroke-width:1px,stroke:#082a3e,fill:#1591da,color:#082a3e;
    class df25bc50 Level0
    class a5da30cd Level1
    class bef70617 Level2
    class 6dbb8570 Level5
```