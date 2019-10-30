# Tools
gotip download

# Folder structure
 - Presentation layer, startup, shutdown, validation - `./cmd`
 - Binary folder, the web service we're building - `./cmd/sales-api`
 - Internal packages for `sales-api` - `./cmd/sales-api/internal`
 - Business logic - `./internal`
 - Reusable packages, can't log, highest level of usability - `./internal/platform`

# Notes
Don't create a common package of types. A type system serves one purpose, get data in and out of an API.

Concrete types, interface types

Compiler will not import `internal` folder from another project

goimport
gocode
goplease - Langauge server. Editor sends changes to goplease. Also a push protocol, language server can always push to client. Soon will be able to watch files.

go clean -modcache - Clears the local module cache

What is the purpose of the logs? - If it's simply for debugging then logs shouldn't contain data. They don't need structured logging.

Where is the logging happening? Who is allowed to write to the logs?

Don't log to `stderr`. Safe this for tracing

It's ok to have to restart a service if the config changes

Services shouldn't crash on startup becuase of missing config. There should be defaults.

Config should be read one time at startup and passed through the app. Different parts of the code should not be allowed to get the config themselves.

--help should be implemented

go mod tidy - get the latest version of a package

go get - Good for updated to a specific version

use the `config` package to auto generate `--help 

Only use package when, init doesn't matter, init doesn't depend on config, only code that should touch it is the code where it is defined

Nano second = 12 instructions by the processor

HTTPTreeMux / GorillaMux

Don't create a go routing uunless you know how it will shutdown when the service terminates. Do load shedding, may need to reconfi kubernetes to allow this

Functions should always be returning the concrete type

Storing values in a context, only store values in a contact that will be used for debugging or tracing a request. If you put values in a context that are required for that requies=t to succesd, the code base is in a dangerous sitation

Capturing time is useful

# Questions

- How do I get up that debug dachboard?
- What is a mux? A router
- What is a channel?
