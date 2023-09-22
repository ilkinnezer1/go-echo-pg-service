# Documentation for partner handler

### GetSinglePartner:
**HTTP method:** `GET` </br>
**Route:** `/partners/:id`</br>
**Description:** retrieves the details of a single partner identified by its id parameter. It checks if the partner is cached and returns it if found. Otherwise, it fetches the partner from the database,
adds it to the cache, and returns it.</br>
**Request parameters:** `id`: the ID of the partner to retrieve.</br>
**Response:**
HTTP status code `200 (OK)` and the requested partner as a JSON object on success.</br>
HTTP status code `404 (Not Found)` and an error message if the requested partner does not exist.

### GetPartners:
**HTTP method:** `GET`</br>
**Route:** `/partners`</br>
**Description:** retrieves all the partners available. It checks if the partners are cached and returns them if found. Otherwise, it fetches the partners from the database, maps the `ImagePath` property to include the full file URL, adds them to the cache, and returns them.</br>
**Response:**
HTTP status code `200 (OK)` and an array of all partners as JSON objects on success.</br>
HTTP status code `500 (Internal Server Error)` and an error message if an error occurs while querying the database.

### CreatePartner:
**HTTP method:** `POST`</br>
**Route:** `/partner/create`</br>
**Description:** Creates a new partner instance and migrate it to the database.</br>
**Request body:** a JSON object representing the new partner.</br>
**Response:**
HTTP status code `201 (Created)` and the newly created partner as a JSON object on success.</br>
HTTP status code `400 (Bad Request)` and an error message if the request body is not a valid JSON object.</br>
HTTP status code `500 (Internal Server Error)` and an error message if an error occurs while creating the new partner in the database.

### UpdatePartner:
**HTTP method:** `PUT`</br>
**Route:** `/partner/:id`</br>
**Description:** Updates an existing partner with the given ID in the database.</br>
**Request parameters:** `id`: the ID of the partners to update.</br>
**Request body:** a JSON object representing the updated partner.</br>
**Response:**
HTTP status code `200 (OK)` and the updated partner as a JSON object on success.</br>
HTTP status code `400 (Bad Request)` and an error message if the request parameters or body are not valid.</br>
HTTP status code `500 (Internal Server Error)` and an error message if an error occurs while updating the partners in the database.

### DeletePartner:
**HTTP method:** `DELETE`</br>
**Route:** `/partner/:id`</br>
**Description:** Deletes an existing partner with the given ID from the database.</br>
**Request parameters:** `id`: the ID of the partner to delete </br>
**Response:**
HTTP status code `204 (No Content)` on success.</br>
HTTP status code `400 (Bad Request)` and an error message if the request parameters are not valid.</br>
HTTP status code `500 (Internal Server Error)` and an error message if an error occurs while deleting the partner from the database.</br>

**Partner Handler** also utilizes the `go-cache` library to cache partner and arrays of partners for better performance. </br>
**The cache keys used are:**
`partner:<id>` for a single partner with the given ID.</br>
`partners` for an array of all partners posts.
