# Documentation for AuthMiddleware function

### Short Description
`AuthMiddleware` is a middleware function that checks if the incoming HTTP request contains a **valid access token** in the `Authorization` header. </br>
If the `access token` is valid, the function sets the `admin` token in the context and allows the request to pass through to the next middleware or handler. </br>
If the token is invalid or missing, the function returns an `error response`.</br>

#### Input parameters:
 - `next` - the next handler function in the Echo middleware chain.

#### Output parameters:
 - `echo.HandlerFunc` - the modified handler function that includes the authentication middleware.

#### Return values:
 - `error` - an error response if the `token` is missing or invalid, otherwise `nil`.

#### Instructions:
 **The middleware function performs the following:**

 - Extract the` access token` from the `Authorization` header of the incoming HTTP request.
 - Check if the token is missing, and return an `error` response with a JSON message `"Token is missing"` and HTTP status code `http.StatusUnauthorized`.
 - Remove the `"Bearer"` prefix from the token, if present to validate.
 - Parse the token using the `jwt.Parse()` method from the `golang-jwt` library, and validate the token signing method and key.
 - Set the admin token in the Echo context using the `c.Set("admin", adminToken)` method.
 - Retrieve the admin token from the context and verify that it is not `nil` and has the expected `claims`.
 - If the `admin token` is valid, allow the `request` to pass through to the `next middleware` or handler, otherwise return an error response with HTTP status code `echo.ErrUnauthorized`.

