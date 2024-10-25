Q - Why do we say a stream of data and not just data
    Referring to the response body as a "stream of data" rather than just "data" emphasizes the nature of how the data is being transmitted and processed. Here's why this distinction is important:

    - Continuous Flow: A stream of data implies that the data is being sent and received in a continuous flow, which is crucial for handling large datasets or real-time data. This contrasts with a single, static chunk of data that is transmitted all at once.
    - Efficiency: Streaming data can be more efficient because it allows for processing of data as it is being received, rather than waiting for the entire dataset to be transmitted. This can reduce latency and memory usage.
    - Real-Time Processing

Q - What is data buffer
    A data buffer is a temporary storage area in memory used to hold data while it is being transferred from one place to another. Buffers are used to manage the differences in the rate at which data is received and processed.

Q - What is close in golang while sending a http request
    When you get the data (response body) from a website, you need to "close" it to free up resources and avoid wasting them.
    The "resources" refer to system resources such as memory and network connections that are used to manage the incoming data (the response body) from the HTTP request. eg. 

    Memory: When you make an HTTP request, Go allocates memory to store the response data. If this memory is not released (by closing the response body), it can accumulate and eventually exhaust the available memory, leading to performance issues or crashes.
    File Descriptors: Operating systems use file descriptors to manage open files and network connections. Each HTTP response consumes a file descriptor. If you don't close the response body, these file descriptors remain in use, and the system can run out of them, preventing new network connections or file accesses.
    Network Connections: HTTP responses involve network connections. If these connections are not properly closed, they remain open and can exhaust the available network resources, leading to issues like being unable to make new network requests.

    so if we don't close, the connection between our laptop and that side remain open ?
    Yes, that's correct! If you don't close the response body in your Go program, the connection between your laptop (or server) and the remote site can remain open.

Q - I read somewhere regarding if we don't close, It can also cause issues with the underlying connection as transport may not reuse HTTP/1.x “keep-alive” TCP connections if the Body is not read to completion and closed.
    HTTP/1.x supports "keep-alive" connections, which allow a single TCP connection to be reused for multiple HTTP requests/responses.
    If the response body is not fully read and closed, the HTTP transport in Go may not be able to reuse the TCP connection. This is because the connection might still have unread data, making it unsafe to use for another request.