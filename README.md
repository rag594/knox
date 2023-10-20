# knox

Tool to generate the spec template for kong with support of different plguins. Different plugins at service/route level are needed to setup our service on kong.This can be used to generate a template and playaround with it.

Now you can generate template with plugins at global, service and route level.You take the template and extend it for other routes/services.

![render1697425879741-min](https://github.com/rag594/knox/assets/8286518/6e1f2a31-c0e8-4ca2-ae5d-d289498abba9)



## Supported Plugins(More plugins to come)

- [x] key-auth 
- [x] cors 
- [x] request-transformer-advanced
- [x] rate-limiting-advanced
- [x] exit-transformer
- [x] pre-function
- [x] acl

## Features

- [x] Generate the template with plugins at service,route or global level
- [x] Summarise the kong-spec by loading the existing kong-spec and visualise them



## Install

#### For Mac
```
brew tap rag594/tap
brew install knox
```

## Visualise kong spec via graphviz

#### Install Graphviz

```
brew install graphviz
```

#### Render the graph into DOT

```
knox visualise --spec <spec-file> --output <output-file>
```

##### Generate the SVG

```
dot -Tsvg -O <dot-output>
```

##### Sample SVG Generated

<img width="1465" alt="Screenshot 2023-10-21 at 2 49 59 AM" src="https://github.com/rag594/knox/assets/8286518/708cc159-223c-471c-a23d-29395d9f16c4">



## Usage

#### Generate Spec for playground

```console
foo@bar:~$ knox generate --help
NAME:
   CLI for generating kong-spec template generate - generate kong-spec template

USAGE:
   CLI for generating kong-spec template generate [command options] [arguments...]

OPTIONS:
   --service-plugins value  add service-plugins in comma-separated format. Example --service-plugins key-auth,cors
   --route-plugins value    add route-plugins in comma-separated format. Example --route-plugins key-auth,cors
   --global-plugins value   add global-plugins in comma-separated format. Example --global-plugins key-auth,cors
   --help, -h               show help
```

#### Visualise your kong spec via graphviz

```console
foo@bar:~$ knox visualise --help
NAME:
   CLI for generating/visualising kong-spec visualise - generate the graphviz

USAGE:
   CLI for generating/visualising kong-spec visualise [command options] [arguments...]

OPTIONS:
   visualise

   --output value  provide your output file using --output <file-path>
   --spec value    load the file using --spec <file-path>
```

## Sample

```console
foo@bar:~$ knox generate --service-plugins "key-auth"
_format_version: "3.0"
_info:
  defaults: {}
  select_tags: []
services:
  - name: DefaultService
    host: xyz.com
    port: 80
    protocol: ""
    path: /
    retries: 0
    enabled: true
    read_timeout: 60000
    connect_timeout: 60000
    write_timeout: 60000
    plugin:
      - config:
          key_names:
            - apiKey
          hide_credentials: true
          key_in_header: true
          key_in_query: false
          key_in_body: false
          run_on_preflight: false
        enabled: true
        name: key-auth
        protocols:
          - http
          - https
    routes:
      - name: route1
        protocols:
          - http
          - https
        methods:
          - POST
        hosts:
          - abc.com
        paths:
          - /abc
        https_redirect_status_code: 426
        regex_priority: 0
        strip_path: false
        path_handling: v0
        preserve_host: true
        request_buffering: true
        response_buffering: true
```
