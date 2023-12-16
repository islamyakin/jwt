# JWT

### How to generate a token

```bash
curl localhost:4000/login
```
```json
{
    "username": "kanaya",
    "password": "rainbowdrinker"
}
```

### How to use the token

```bash
curl localhost:4000/api/v1/whoiam
```