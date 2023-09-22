# Documentation for logger handler

### HTTPBodyDumpResponse function
`HTTPBodyDumpResponse` is a function that logs the request and response body for each HTTP request made to the app. 
It takes in an Echo context object `c`, and the request and response body as bytes (`reqBody` and `resBody` respectively).

The function creates a `RequestLog` object that contains the `method`, `path`, `headers`, and `query` parameters of the incoming request.
If there is an error in `Unmarshaling` the request body into the `RequestLog` object, the function stores the request body as a string. 
The function also creates a `ResponseLog` object that contains the status code and headers of the HTTP response. 
If there is an error in `Unmarshaling` the response body into the `ResponseLog` object, the function stores the response body as a string.

The function then creates a log object that contains both the request and response logs, 
and marshals it into a JSON string using the `json.MarshalIndent` method, and it prints the `log`string to the console.

### LoggerMiddlewareConfig variable
`LoggerMiddlewareConfig` is a middleware that logs each incoming HTTP request to the console using the **Echo** `middleware.LoggerWithConfig` function.

#### Following instructions
It takes in a `LoggerConfig` object that specifies the `format`, `output`, and `custom time format` of the log message.
In this case, the format is set to 
 - `time=${time_rfc3339}`, 
 - `method=${method}`,
 - `uri=${uri}`, 
 - `status=${status}`, 
 - `latency=${latency_human}\n`, which logs the `timestamp`, `HTTP method`, `URI`, `response status code`, and `response time` in a human-readable format.

### MultiWriter method
The output is set to a `MultiWriter` that writes the log message to both the console and the `LoggerConfig` file specified in the config package. 
The custom time format is set to `2006-01-02 15:04:05`.

This middleware is added to an Echo server instance using the `e.Use(LoggerMiddlewareConfig)` method and logs are written into `logs/request.log`.