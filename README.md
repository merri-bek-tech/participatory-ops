# ParOps - Participatory Ops

ParOps is a framework for managing web hosting infrastructure with a focus on a large participatory community of volunteers. Originally developed for use by [Merri-bek Tech](https://merri-bek.tech).

## Overview

ParOps consists of the following parts:

**Front End** - A single page React app, packaged using [vite](https://vitejs.dev/).
**API** - A go web server, using the [echo](https://echo.labstack.com/) web framework. For more details, see [the Api README](paropd/README.md)
**ParopD** - A go daemon which runs on each device to be observed. For more details, see [the ParopD README](paropd/README.md)

## Development

### Dependences

Paropd and Parops API need to be run on linux.

Paropd requires nmap. Install with `apt install nmap` on ubuntu/debian systems if not present.

### Running Parops Locally

To run both frontend and api simultaneously, we use `nodemon` to execute both. This is a nodejs based tool, but you should already have that installed to run the frontend.

To execute, run `npm run dev`.

The frontend can be accessed in the browser at port `5173`. If this port doesn't seem to work, scroll up in the output to find the section where VITE starts up (the front end packager). It should look something like this, and contains the full link with the port:

```
[0]   VITE v5.2.13  ready in 403 ms
[0]
[0]   ➜  Local:   http://localhost:5173/
```

### Running ParopD Locally

From the paropd directory (`cd paropd`), run the go program using `go run .`.
