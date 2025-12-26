# Day-5:

### 🐳 Dockerfile Build Process (Layer-wise)
```
FROM ubuntu:24.04

RUN apt update

WORKDIR /app

RUN apt install -y golang 

COPY . .

CMD [ "go","run", "main.go"]
```

### Step-1: `FROM ubuntu:24.04`

1. Check local cache
2. Pull image from Docker Hub (if not cached)
3. Load base image

**Layer 0**

---

### Step-2: `RUN apt update`

1. Create temporary container
2. Run `apt update`
3. Commit container changes
4. Remove temporary container

**Layer 1**

---

### Step-3: `RUN apt install -y golang`

1. Create temporary container
2. Run `apt install -y golang`
3. Commit container changes
4. Remove temporary container

**Layer 2**

---

### Step-4: `WORKDIR /app`

1. Create temporary container
2. Set working directory to `/app`
3. Commit container changes
4. Remove temporary container

**Layer 3**

---

### Step-5: `COPY ./server.go ./server.go`

1. Create temporary container
2. Copy `server.go` from host to `/app`
3. Commit container changes
4. Remove temporary container

**Layer 4** which is the final image 

---

### Step-6: `CMD ["go", "run", "server.go"]`

* CMD bind with last layer image(final image) as metadata
* No temporary container
* No filesystem change

**No Layer Created**

---

### 📌 Final Layer Order

```
Layer 0 → Base Image (ubuntu:24.04)
Layer 1 → apt update
Layer 2 → golang installation
Layer 3 → WORKDIR /app
Layer 4 → COPY server.go
CMD → Runtime instruction (no layer)
```
#### Why Docker Stores Intermediate Images?

Docker stores intermediate images to **speed up builds, reuse layers, save disk space, and ensure reliable, repeatable builds**.

```
Intermediate Images = Cache + Reuse + Speed + Reliability
```
#### How Docker Rebuild Works

1. Docker reads the Dockerfile top to bottom
2. For each instruction, Docker checks the cache
3. If the instruction and its context did not change → reuse cached layer
4. If something changes → rebuild that layer
5. **All next layers are rebuilt** automatically

#### Example

```dockerfile
RUN apt install -y golang   # cached
COPY server.go .           # changed
CMD ["go", "run", "server.go"]
```

- Only `COPY` and layers after it are rebuilt
- Go installation remains cached

#### One-line Summary

- Docker rebuilds only the changed layer and everything after it using cache.

[Next: Day 6](../Day_6/README.md)


