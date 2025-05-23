---
title: Contributing
---

Thanks for your interest in helping developping this app.

Here the guidelines and requirements to do it the best way :rocket:

# Project overview

This project is a `Go API`, its role is to be the _brain_ for the perfect Pokemon player ligue by getting every usefull data.

# Structure

The API is split in 2 parts :
- The framework : The main structure is init as an entire framework from scratch and its design as fondation.
- The modules : In the folder [`modules`](modules/README.md), you can find any modules who create the endpoints.

From those parts, the idea is to have a collection of modules and being able to call all them as module galaxy (factory module).

# Guidelines

If you want to create a new module, please referred yourself at the [`module documentation`](modules/README.md) and/or send [me](https://github.com/mickahell) a message.

If you want to contribute in the framework, check the issue and/or send [me](https://github.com/mickahell) a message.

## How to start

The project if based on `Go=1.23.8` (currently the latest on date 2022/11). And its running using `Docker`, for testing developpment, it can also be run locally. All it need is setting up a local `database`. And some variables from the `.envrc` file.

```
export PASS_DB=root
export TOKEN_API_TEST=test
```

Then :

- **Build with Docker**

```bash
make docker-build
```

- **Run with Docker**

```bash
make docker-run
```

- **Run locally**

```bash
make start-app
```

You can now call your api using :
- Swagger : `http://127.0.0.1:8080/swagger`

## API documentation

The API documentation is building automatically from code using `swagger` and `makefile` :

```bash
make swagger
```

## Configuration files

To run, the app need a configuration file, thoses files can be found in the [`test/` folder](test) and named [`conf_local.yaml`](test/conf_local.yaml) and [`conf_docker.yaml`](test/conf_docker.yaml). They are only for testing, **never for production !**

## Units tests

The goal is to have every part of the API tested `>=80%`, currently some modules doesn't have this minimun testing requirement and so have to be boost up.

To run the unit tests, locally you can run :

```bash
make run-test
```
