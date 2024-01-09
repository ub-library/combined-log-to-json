# combined-log-to-json

A simple script to convert web server logs from the combined log format
(NCSA Common log format extended with referer and user-agent) into JSON.

Primarily useful to enable filtering and inspection of logs using `jq`
or a similar tool, instead of operating on the plain text data with e.g.
`grep` or `awk`.

## Usage

`combined-log-to-json < example.log`

Reads log lines from `stdin` and writes new line delimited JSON objects
to `stdout`.

## Field mapping

In Apache HTTP Server the combined log format is specified with the
following directive and format string:

``` Apache
LogFormat "%h %l %u %t \"%r\" %>s %b \"%{Referer}i\" \"%{User-agent}i\"" combined
```

Here is the explanation of the format string symbols, based on the
[Apache HTTP Server
documentation](https://httpd.apache.org/docs/2.4/mod/mod_log_config.html#customlog),
and how each element maps to JSON:

| Format String    | Definition                             | JSON field(s)                | Note                     |
|------------------|----------------------------------------|------------------------------|--------------------------|
| `%h`             | Remote hostname                        | ip                           |                          |
| `%l`             | Remote logname                         | remoteLogName                |                          |
| `%u`             | Remote user                            | user                         |                          |
| `%t`             | Time the request was received          | timestamp                    |                          |
| `%r`             | First line of request                  | method, url, protocolVersion | Quoted. Split by spaces. |
| `%>s`            | Status (final)                         | status                       |                          |
| `%b`             | Size of response in bytes              | size                         |                          |
| `%{Referer}i`    | The contents of Referer header line    | referer                      | Quoted.                  |
| `%{User-agent}i` | The contents of User-agent header line | userAgent                    | Quoted.                  |

## Example

Given a file `example.log` with the following content:

``` txt
192.168.1.10 - 0hLNPuxTM7JscWD [08/Jan/2024:11:44:16 +0100] "GET /foo/bar?baz=no HTTP/1.1" 200 33 "http://example.com/" "Mozilla/5.0 (X11; Linux x86_64; rv:102.0) Gecko/20100101 Firefox/102.0"
```

Running `combined-log-to-json < example.log | jq` will produce the
following JSON output (note that it is `jq` that provides the nice
formatting):

``` json
{ 
  "ip": "192.168.1.10",
  "remoteLogName": "-",
  "user": "0hLNPuxTM7JscWD",
  "timestamp": "2024-01-08T11:44:16+01:00",
  "method": "GET",
  "url": "/foo/bar?baz=no",
  "protocolVersion": "HTTP/1.1",
  "status": 200,
  "size": 33,
  "referer": "http://example.com/",
  "userAgent": "Mozilla/5.0 (X11; Linux x86_64; rv:102.0) Gecko/20100101 Firefox/102.0"
}
```

## Building and installing

The script is written in [Go](https://go.dev) and given a Go version of
1.19 or higher should hopefully build and install using standard Go
commands `go build` and `go install`. (It might work with earlier go
versions as well.)

``` sh
# Install in default $GOBIN directory:
go install github.com:ub-library/combined-log-to-json@latest
```

The script is also available as a [nix
flake](https://nixos.org/manual/nix/stable/command-ref/new-cli/nix3-flake).
To use it you should enable nix experimental features `flakes` and
`nix-command`.

``` sh
# Make the executable available in the current shell:
nix shell github:ub-library/combined-log-to-json
```

See nix flake documentation for how to use it in another flake
declaration. This flake uses nixpkgs 23.11 as input by default, but
should work with any nixpkgs release that has a suitable version of Go.
