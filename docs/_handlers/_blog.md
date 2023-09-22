# Documentation for blog handler

### GetSingleBlog:
**HTTP method:** `GET` </br>
**Route:** `/blogs/:id`</br>
**Description:** Retrieves a single blog post with the given ID from the database. The response is cached to improve performance.</br>
**Request parameters:** `id`: the ID of the blog post to retrieve.</br>
**Response:**
HTTP status code `200 (OK)` and the requested blog post as a JSON object on success.</br>
HTTP status code `404 (Not Found)` and an error message if the requested blog post does not exist.

### GetAllBlogs:
**HTTP method:** `GET`</br>
**Route:** `/blogs`</br>
**Description:** Retrieves all blog posts from the database. The response is cached to improve performance.</br>
**Response:**
HTTP status code `200 (OK)` and an array of all blog posts as JSON objects on success.</br>
HTTP status code `500 (Internal Server Error)` and an error message if an error occurs while querying the database.

### CreateBlog:
**HTTP method:** `POST`</br>
**Route:** `/blog/create`</br>
**Description:** Creates a new blog post in the database.</br>
**Request body:** a JSON object representing the new blog post.</br>
**Response:**
HTTP status code `201 (Created)` and the newly created blog post as a JSON object on success.</br>
HTTP status code `400 (Bad Request)` and an error message if the request body is not a valid JSON object.</br>
HTTP status code `500 (Internal Server Error)` and an error message if an error occurs while creating the new blog post in the database.

### UpdateBlog:
**HTTP method:** `PUT`</br>
**Route:** `/blog/:id`</br>
**Description:** Updates an existing blog post with the given ID in the database.</br>
**Request parameters:** `id`: the ID of the blog post to update.</br>
**Request body:** a JSON object representing the updated blog post.</br>
**Response:**
HTTP status code `200 (OK)` and the updated blog post as a JSON object on success.</br>
HTTP status code `400 (Bad Request)` and an error message if the request parameters or body are not valid.</br>
HTTP status code `500 (Internal Server Error)` and an error message if an error occurs while updating the blog post in the database.

### DeleteBlog:
**HTTP method:** `DELETE`</br>
**Route:** `/blog/:id`</br>
**Description:** Deletes an existing blog post with the given ID from the database.</br>
**Request parameters:** `id`: the ID of the blog post to delete </br>
**Response:**
HTTP status code `204 (No Content)` on success.</br>
HTTP status code `400 (Bad Request)` and an error message if the request parameters are not valid.</br>
HTTP status code `500 (Internal Server Error)` and an error message if an error occurs while deleting the blog post from the database.</br>

**Blog Handler** also utilizes the `go-cache` library to cache blog posts and arrays of blog posts for better performance. </br>
**The cache keys used are:**
`blog:<id>` for a single blog post with the given ID.</br>
`blogs` for an array of all blog posts.

The cache expiration time is set to the default value of 5 minutes but will be updated in the release.

In order to connect to the database, `migrate` package is used to call the `ConnectDatabase` function, 
which returns a `database connection object`. This connection object is then used to perform CRUD Operations.