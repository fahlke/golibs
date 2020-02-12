# pearsoncollider

## Build image locally

```bash
docker build --tag pearsoncollider:latest .
```

## Run the tool

```bash
docker run \
  --rm \
  --env TEST_STRING=foo \
  pearsoncollider:latest                      
```

## Example output

```bash
Trying to find a maximum of 10 hash collisions with Pearson-16...

digest `0xad5c` is the same for `foo` and `5usxiI9UR`
digest `0xad5c` is the same for `foo` and `KC0KäP`
digest `0xad5c` is the same for `foo` and `P51Tyo2NrnwIÖyn`
...
```