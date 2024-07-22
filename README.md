
# TMDb Proxy API Server

This repository provides a proxy API server for The Movie Database (TMDb). **Due to IP bans from ISPs in India, TMDb's API services may be inaccessible to Indian users**. This proxy server allows you to access TMDb data seamlessly.


## API Reference

#### Call to themoviedb API

```http
  GET /v1/tmdb
```

`/v1/tmdb` maps to =====> `https://api.themoviedb.org/`


| Header | Description                       |
| :-------- | :------------------------------- |
| `X-WPROXY-KEY`       `string` | **Required**. All requests made via proxy must contain a X-WPROXY-KEY |


> `X-WPROXY-KEY` can be configured in `.env` file.



## Run Locally


### Prerequisites
Go: Make sure you have Go installed. You can download it from golang.org.

Clone the project

```bash
  git clone https://github.com/thewolmer/tmdb-proxy-api-server
```

Go to the project directory

```bash
  cd tmdb-proxy-api-server
```

Install dependencies

```bash
  go mod tidy
```

Create a `.env` file and put `WPROXY-KEY` in it

```bash
  echo 'X-WPROXY-KEY=secret_key' > .env
```

Run the server
```bash
  go run cmd/main.go
```
---

*Note: Running this proxy API server locally might not work if you're in India, as your IP address could still be subject to the same restrictions imposed on TMDb. For optimal results, deploy this server to a cloud provider or hosting service outside of India. This will help ensure that your requests are routed through an IP address that is not affected by the IP ban.*

---

## Adding Proxy Routes
To add additional routes for TMDb, add a function in the `controllers/proxy.go` file: 
```go
func ExampleProxy(c *gin.Context) {
	proxy := utils.CreateReverseProxy("https://api.example.com", "/v1/example")
	proxy.ServeHTTP(c.Writer, c.Request)
}
```
next, modify the `main()` function in `cmd/main.go` to add our controller
```go
func main() {
    // Previous Code
    ...
    // Add proxy route
	r.Any("/v1/example/*proxyPath", controllers.ExampleProxy)
    ...
    // Previous Code
}
```

This will add new route `/v1/example` mapped to `api/example.com`

## Deployment

Deploy this server for free on [Render](https://render.com/)

 - Create an account on [Render](https://render.com/)

- Add a new [Web Service](https://dashboard.render.com/web/new)

- Choose your repo from the list

- Add build command
```bash
  go build -tags netgo -ldflags '-s -w' -o bin/main ./cmd

```
- Add run command 
```bash
./bin/main
```
- Add environment variables 
- Deploy 🚀


## Tech Stack

**Client:** //

**Server:** Go, Gin


## Roadmap

- Retelimits



## Contributing

Feel free to contribute to this project by opening issues, submitting pull requests, or providing feedback.


## Related

Here are some related links

[Reddit Thread](https://www.reddit.com/r/NovaVideoPlayer/comments/1beft2v/information_tmdb_is_not_accessible_in_india/)
[![image.png](https://i.postimg.cc/SQthfH4V/image.png)](https://www.themoviedb.org/talk/65ebdf52b7d352017be51b55)

