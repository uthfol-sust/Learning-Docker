# Day-1:

1. Pull the Ubuntu image.
2. Run a container from the Ubuntu image.
3. Go inside the container.
4. Install Golang.
5. Create a directory `app`.
6. Copy the `server.go` file from the host machine to the Docker container at `/app/server.go`.
7. Run the Go app.

---

# Day-2:

## Deep Dive of Docker CMD Command

### What is CMD in Docker?
- **Purpose:** Specifies the default command for a container.  
- **Execution:** Runs when you start a container with `docker run` unless overridden.

```dockerfile
FROM ubuntu:24.04

RUN apt update

WORKDIR /app

RUN apt install -y golang

COPY ./main.go ./main.go

CMD [ "go", "run", "main.go" ]
````

### Difference from RUN:

1. **RUN** executes at build time (creates image layers).
2. **CMD** executes at runtime (when container starts).
---
# Day-3: 
## WORKDIR - Deep Dive
### What is WORKDIR?
- **Purpose:** Sets the current working directory for any subsequent Docker instructions like `RUN`, `CMD`, `ENTRYPOINT`, `COPY`, and `ADD`.
- **Execution:** If the specified directory does not exist, Docker automatically creates it.
- **Syntax:**
```dockerfile
WORKDIR /app
```
---
# Day-4:

### What is Detach Mode?
- **Purpose:** Runs a Docker container in the background, allowing the terminal to be free for other tasks.
- **Command:** Use `-d` flag with `docker run`.

### Syntax:
```bash
docker run -d <image_name>
```
