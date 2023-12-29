# Console

[![Pull Request](https://github.com/streamdal/streamdal/actions/workflows/apps-console-pr.yml/badge.svg)](https://github.com/streamdal/streamdal/actions/workflows/apps-console-pr.yml)
[![Publish Docker](https://github.com/streamdal/streamdal/actions/workflows/apps-console-release.yml/badge.svg)](https://github.com/streamdal/streamdal/actions/workflows/apps-console-release.yml)
[![Discord](https://img.shields.io/badge/Community-Discord-4c57e8.svg)](https://discord.gg/streamdal)

Streamdal's console allows you visualize interact with your services, build and
attach pipelines to consumers and producers and monitor their operations.

![Console](./console-screenshot.png)

<sub>For more details, see the main
[streamdal repo](https://github.com/streamdal/streamdal).</sub>

### Development

[![Made with Fresh](https://fresh.deno.dev/fresh-badge.svg)](https://fresh.deno.dev)

The Console is a Deno + Fresh project that uses Preact, ReactFlow and Twind:
https://fresh.deno.dev/docs/getting-started

Make sure to install Deno: https://deno.land/manual/getting_started/installation

Optionally, copy `example.env` -> `.env` and set environment variables as
needed. By default the console will access the GRPC WEB API running on
`http://localhost:8083`

Then start the project:

```
deno task start
```

This will watch the project directory and restart as necessary.

If you make any significant Deno lib and/or configuration changes and your IDE
gets confused, you can force update the Deno lib cache with
`deno cache --reload main.ts`

**By default, the console will be viewable at: `http://localhost:8080`**

### Server

This console needs a streamdal-server to run against. See
[streamdal/server](https://github.com/streamdal/streamdal/blob/main/apps/server/README.md)
for instructions on running it locally.

### Running (non-development)

If you just want to run the console and the server together for non-development
purposes, you can use the install to bring them both up, see:
https://github.com/streamdal/streamdal#getting-started

### Dependencies

The console depends on the following dependencies (direct and indirect):

- [streamdal/server](https://github.com/streamdal/streamdal/tree/main/apps/server)
- Envoy (used by Console for gRPC-Web)
- Redis (used by `streamdal/server`)

### Releasing

Any push or merge to main with changes on `/apps/console` will automatically tag
and release a new console version with `apps/console/vX.Y.Z`. The VERSION file
will also be bumped for display in the console. If you'd like to skip release on
the push/merge to main, include "norelease" anywhere in the commit message.
