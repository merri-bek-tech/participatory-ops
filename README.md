# ParOps - Participatory Ops

ParOps is a framework for managing web hosting infrastructure with a focus on a large participatory community of volunteers. Originally developed for use by [Merri-bek Tech](https://merri-bek.tech).

## Overview

ParOps consists of the following parts:

**Front End** - A single page React app, packaged using [vite](https://vitejs.dev/).
**API** - A go web server, using the [echo](https://echo.labstack.com/) web framework.

## Development

### Running Locally

To run both frontend and api simultaneously, we use `nodemon` to execute both. This is a nodejs based tool, but you should already have that installed to run the frontend.

To execute, run `npm run dev`.
