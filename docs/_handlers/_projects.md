# Documentation for project handler

### GetSingleProjectDetails:
**HTTP method:** `GET` </br>
**Route:** `/projects/:id`</br>
**Description:** t retrieves the details of a single project identified by its id parameter. It checks if the project is cached and returns it if found. Otherwise, it fetches the project from the database,
adds it to the cache, and returns it.</br>
**Request parameters:** `id`: the ID of the project to retrieve.</br>
**Response:**
HTTP status code `200 (OK)` and the requested project as a JSON object on success.</br>
HTTP status code `404 (Not Found)` and an error message if the requested project does not exist.

### GetProjectsDetails:
**HTTP method:** `GET`</br>
**Route:** `/projects`</br>
**Description:** retrieves all the projects available. It checks if the projects are cached and returns them if found. Otherwise, it fetches the projects from the database, maps the `ImagePath` property to include the full file URL, adds them to the cache, and returns them.</br>
**Response:**
HTTP status code `200 (OK)` and an array of all projects as JSON objects on success.</br>
HTTP status code `500 (Internal Server Error)` and an error message if an error occurs while querying the database.

### CreateNewProject:
**HTTP method:** `POST`</br>
**Route:** `/project/create`</br>
**Description:** Creates a new project instance and migrate it to the database.</br>
**Request body:** a JSON object representing the new project.</br>
**Response:**
HTTP status code `201 (Created)` and the newly created project as a JSON object on success.</br>
HTTP status code `400 (Bad Request)` and an error message if the request body is not a valid JSON object.</br>
HTTP status code `500 (Internal Server Error)` and an error message if an error occurs while creating the new project in the database.

### UpdateProjectDetail:
**HTTP method:** `PUT`</br>
**Route:** `/project/:id`</br>
**Description:** Updates an existing project with the given ID in the database.</br>
**Request parameters:** `id`: the ID of the project to update.</br>
**Request body:** a JSON object representing the updated project.</br>
**Response:**
HTTP status code `200 (OK)` and the updated project as a JSON object on success.</br>
HTTP status code `400 (Bad Request)` and an error message if the request parameters or body are not valid.</br>
HTTP status code `500 (Internal Server Error)` and an error message if an error occurs while updating the project in the database.

### DeleteProject:
**HTTP method:** `DELETE`</br>
**Route:** `/project/:id`</br>
**Description:** Deletes an existing project with the given ID from the database.</br>
**Request parameters:** `id`: the ID of the project to delete </br>
**Response:**
HTTP status code `204 (No Content)` on success.</br>
HTTP status code `400 (Bad Request)` and an error message if the request parameters are not valid.</br>
HTTP status code `500 (Internal Server Error)` and an error message if an error occurs while deleting the project from the database.</br>

**Project Handler** also utilizes the `go-cache` library to cache project and arrays of project for better performance. </br>
**The cache keys used are:**
`project:<id>` for a single project with the given ID.</br>
`project` for an array of all project posts.
