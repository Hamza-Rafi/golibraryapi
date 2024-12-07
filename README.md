# golibraryapi

## Endpoints:
- `GET /books`
- `POST /books`
    - Request Body eg:
    ```json
    {
        "title": "example title",
        "author": "example author"
    }
    ```
- `PUT /books`
    - Request body eg:
    ```json
    {
        "id": "id",
        "title": "new title",
        "author": "new author"
    }
    ```

- `GET /books/{id}`
- `DELETE /books/{id}`
