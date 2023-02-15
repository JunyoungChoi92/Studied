# gRPC.md

====

## What is gRPC

[한국어 참조](https://appmaster.io/ko/blog/grpcneun-mueosibnigga)

- gRPC uses Protocol Buffers (Protobuf), Google’s mature open source mechanism for serializing structured data (although it can be used with other data formats such as JSON). Since HTTP/2 relies on the transferred data being binary encoded protobuf plays a very important role for gRPC.

- gRPC supports code-generation for many of the most popular programming languages. This can be done as we can easily define services using Protobuf and generate client-side and server-side code based on that definition. In microservices, we may have one service written in Java (the server) and the other one in Python (the client), with gRPC’s code generation we can define our service in Protobuf and generate the stub for both Python and Java easily and we do not have to worry about the contract and the communication between the two application, we can just make use of the generated code and access the methods on server-side as if it was a method defined in client’s codebase. This feature makes gRPC a great option to consider for microservices architecture where polyglot architectures are usually chosen.

- gRPC does not expose HTTP to the API designer or the API user. gRPC made all the decisions for you to map the RPC layer on top of HTTP and gRPC-generated stubs and skeletons hide HTTP from the client and server too, so nobody has to worry how the RPC concepts are mapped to HTTP — they just have to learn gRPC (I am not quite sure if it is a good thing though).

- gRPC brings the two powerful technologies together to increase the performance — HTTP/2 and Protocol Buffers. HTTP/2 offers more performant data transfer models compared to HTTP/1.1 by relying on the data being transferred is in binary format, and that is enabled by using Protocol Buffers. In addition to that, using Protocol Buffers compared to JSON (or any other text-based data format) plays an important role as it is more performant to encode/decode binary data.

  - Pros:

  1. High performance: gRPC uses binary serialization and HTTP/2 for transport, which enables efficient communication and reduces latency, resulting in higher performance.
  2. Language agnostic: gRPC supports multiple programming languages, including C++, Java, Python, and Go, making it a versatile option for developing distributed systems.
  3. Simplified development: gRPC generates client and server code automatically, making it easy to develop and maintain distributed systems.
  4. Interoperability: gRPC can be used with other frameworks and services, such as RESTful APIs, making it a flexible option for building microservices.
  5. Strongly typed contracts: Protocol Buffers defines the interface between the client and server, making it easier to understand and maintain the API.

  - Cons:

  1. Learning curve: gRPC can have a steep learning curve, especially for developers who are not familiar with Protocol Buffers or RPC.
  2. Limited browser support: gRPC uses HTTP/2 for transport, which is not yet widely supported in browsers, limiting its use for web applications.
  3. Potential for versioning issues: Changing a service's contract can break existing clients, requiring careful management of versioning to prevent compatibility issues.
  4. Debugging can be difficult: Debugging gRPC can be more challenging than debugging REST APIs due to its binary data format.
  5. Limited support for dynamic discovery: gRPC's service discovery is limited, making it difficult to use with dynamically scaling environments such as Kubernetes.

  ### dynamically scaling environments

  - Dynamically scaling environments, such as Kubernetes, are environments where the number of containers running an application can change dynamically based on demand or resource usage. These environments can be challenging to manage because the number of containers can change rapidly, and the location of the containers may not be fixed.

  - gRPC has a con in this area because its service discovery mechanism is limited, making it challenging to use in dynamically scaling environments. <u>gRPC relies on service discovery via DNS or a static configuration file, which may not be sufficient for dynamically scaling environments.</u> In these environments, services may be created or deleted frequently, and their IP addresses may change rapidly, making it difficult for gRPC clients to discover and connect to the appropriate services.

  - To address this limitation, gRPC supports integrating with service mesh technologies, such as Istio or Linkerd, which provide more advanced service discovery and load balancing capabilities that are better suited for dynamically scaling environments.

  - [참조 - 쿠버네티스 환경에서의 gRPC](https://blog.cloudflare.com/moving-k8s-communication-to-grpc/)

## REST vs. gRPC

- [참조](https://medium.com/sahibinden-technology/benchmarking-rest-vs-grpc-5d4b34360911)

## Structure

![gRPC structure](https://grpc.io/img/landing-2.svg)

### Protocol Buffer(Ptotobuf)

- Protocol Buffers is a language-agnostic data serialization format that is used to efficiently encode structured data for communication between systems. Data serialization is the process of converting data objects or data structures into a byte stream that can be sent over a network or stored in a file.

- Protocol Buffers use a compact binary format to represent data objects, making them smaller and faster to transmit than other serialization formats such as XML or JSON. <u>The structure of the data is defined in a schema file using a language called Protocol Buffer Language (Proto), which provides a way to define messages, enums, and services</u>.

- The schema file defines the data structure in a language-agnostic way, making it easy to generate code for multiple programming languages to encode and decode the data. The generated code provides a simple API for accessing the data in the messages, which makes it easy to use the Protocol Buffers format in distributed systems.

- When data is serialized using Protocol Buffers, it is encoded into a binary format that is both compact and efficient. The encoded data can be sent over a network or stored in a file, and then decoded back into the original data structure using the generated code. This process of serialization and deserialization is essential for communication between distributed systems and enables applications to exchange data in a standardized and efficient way.

```javascript
// these are same!

{
    "id":42 // JSON type (9 Bytes)
}

0x08 0x2a // Protobuf type (2 Bytes)
```

- Because gRPC uses serialization to send and receive byte-by-byte messages, it has a very small granularity. This makes it effective in low power or bandwidth networks. This makes it ideal for use in the IOT, for example.

### Interface Definition Language

#### Data Serialization
