# My Songs

> Simple choral song tracker using Nuxt and Pocketbase.

## Development

> **NOTE:** Ensure [Volar Take-Over](https://vuejs.org/guide/typescript/overview.html#volar-takeover-mode) mode is enabled!

### Startup

```sh
make api-serve

make app-serve
```

### Ports

Occasionally a port may stay open and refuse to close, which will prevent connecting/hosting to it again. This can be resolved by killing the task attached and restarting.

```
❯ lsof -t -i:5001
####

❯ kill -9 ####
```

## Web (Vue)

> **TODO**

## API (Pocketbase)

Pocketbase can be started with `make api-serve`. Upon initial configuration, user executable permissions may need to be added to the `pocketbase` binary.

- `http://localhost:8090` - Serve static content from `pb_public`
- `http://localhost:8090/_/` - Admin dashboard UI
- `http://localhost:8090/api/` - REST API

### Auth

Add the following rule to Pocketbase collection/resource rules to limit resources to the current user.

```
@request.auth.id = user.id
```

Tokens are valid for 14 days (by default) but can be refreshed if a valid token is provided (there is no refresh workflow). If the token is invalid (ie expired), the auth state can be cleared and a login workflow initiated.

### Admin UI

> **Username:** `kendall@kendallroth.ca`
> **Password:** `Passw0rd!@#`

