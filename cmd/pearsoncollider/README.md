# pearsoncollider

[![Docker Repository on Quay](https://quay.io/repository/fahlke/pearsoncollider/status "Docker Repository on Quay")](https://quay.io/repository/fahlke/pearsoncollider)

## Here for using the tool?

Simply run the pre-build image from [Quay.io](https://quay.io) container registry.

```bash
docker run \
  --rm \
  --env TEST_STRING=foo \
  quay.io/fahlke/pearsoncollider:latest
```

## Test local

```bash
docker build --tag pearsoncollider:latest .
```

```bash
docker run \
  --rm \
  --env TEST_STRING=foo \
  pearsoncollider:latest                      
```

## Example output

```bash
Trying to find 10 hash collisions with Pearson-16...

digest `0xad5c` is the same for `foo` and `5usxiI9UR`
digest `0xad5c` is the same for `foo` and `KC0KäP`
digest `0xad5c` is the same for `foo` and `P51Tyo2NrnwIÖyn`
...
```