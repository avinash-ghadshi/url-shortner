# URL Shortener (# url-shortner)

**Note:** The backend is completed, UI is under development.

## Prerequisites
- Node > 14
- Golang
- npx

## How to Use?

### Clone the Project
Clone the repository to your local machine.

### Go to Project Directory
Navigate to the project directory:
```sh
cd url-shortener
```

### To Run Server
1. Navigate to the server directory:
    ```sh
    cd bin/server
    ```
2. Run the server:
    ```sh
    ./urlshortener
    ```

### To Run Client
1. Navigate to the client directory:
    ```sh
    cd client
    ```
2. Install dependencies and start the client:
    ```sh
    npm install
    npm start
    ```

## Functionality

1. **Login**
    - You can login on the UI with the default username (`admin`) and password (`admin`).

2. **Backend API**
    - The backend is built on REST and can be used directly.

### Available APIs

#### Get Short URL
- **Endpoint:** `GET /api/urlshort/getshort/{url}/{expiry}`
- **Description:** Replace `{url}` with the base64 encoded URL and `{expiry}` with the URL TTL in minutes.
- **Example:**
    ```sh
    /api/urlshort/getshort/aHR0cDovL3Rlc3QuY29t/5
    ```
    - Here, `aHR0cDovL3Rlc3QuY29t` is the base64 encoding of `http://test.com`.

#### Get Original URL
- **Endpoint:** `GET /api/urlshort/getorigin/{url}`
- **Description:** This URL is received in the response of the previous endpoint.
- **Example:**
    ```sh
    /api/urlshort/getorigin/j3lbs8
    ```

#### Delete URL
- **Endpoint:** `DELETE /api/urlshort/delete/{url}`
- **Description:** Replace `{url}` with the last part of the URI from the `getorigin` API.
- **Example:**
    ```sh
    /api/urlshort/delete/j3lbs8
    ```
