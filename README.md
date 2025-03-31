# Google's OpenID Connect (OIDC) Authentication with Go

> Note: This repository uses Hexagonal Architecture (Port & Adapters pattern)

> I'm trying to makes things as simple as possible, while still following best practices

## Tools used

- Docker (Docker Compose)
- Go (Language, Version 1.24+)
- Postgres (Database)
- Valkey (Cache, better alternative for Redis)
<!-- - OpenSSL (self-signed TLS certificate creation) -->

> Cache service (Valkey) is used for stores and compares the `State` in the process of OIDC 

## Requirements

1. Create Oauth Consent Page on google cloud, and configure the `Authorized redirect URIs` to `http://localhost:8080/login/openid/google/callback` and `Authorized JavaScript origins` to `http://localhost:8080`, you might leaves `authorized domains` blank. And copy the `client id` and `client secret`

2. Copy .env.example to .env. And fills the environment needed for example `GOOGLE_OPEN_ID_CLIENT_ID` and `SECRET`. You also need to specify the certificates path to postgres database connection, in step 3

```bash
cp .env.example .env
```

3. Copy the example of postgres env file on the `secrets` directory and fill/change necessary part

4. You need to create self-signed certificates in `certs` directory and it should look like this

```bash
certs
|-- database
|   |-- priv.key
|   |-- server.csr
|   `-- server.pem
`-- root
    |-- priv.key
    `-- root.pem
```

**Generate local certificates**

1. Create private key and certificates root/ca server
2. Create server private key and Certificate Signing Request (CSR) to the root server
3. Do the same from step 2 but for Database Service and Caching Service

## Run locally (Docker Compose)

**Run the server**

```bash
docker compose up -d
```

**Login**

```bash
curl http://localhost:8080/login/openid/google -i
```

Now, you will see the `Location` header, you could open that on the browser.

After you login with your google account, you will see JSON response with JWT token, and your account is now stores on database.