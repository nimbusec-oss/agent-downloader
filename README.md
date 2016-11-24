# agent-downloader
The agent-downloader simplifies the process of downloading and updating the nimbusec server agent.
The API calls used here are described in our KB (https://kb.nimbusec.com/API/API#agents).

## install

```
go get github.com/nimbusec-oss/agent-downloader
go build
```

## run
Run the compiled binary by providing nimbusec API-KEY and nimbusec API-SECRET as defined in the 
nimbusec portal under the section API TOKEN (https://portal.nimbusec.com/einstellungen/serveragent).

```
./agent-downloader -key=<abc> -secret=<abc>
```

You will get a list of available agent versions and will be asked which to download.

## run headless
It is also possible to run this tool headless, for automatic download of the current agent version for example. For that just set the option `headless`. The defaults are:

| param  | value |
|--------|-------|
| arch   | 64bit |
| os     | linux |
| format | bin   |

The following will download an agent binary with default settings:

```
./agent-downloader -key=<abc> -secret=<abc> -headless
```

But you can of course set everything separately as well e.g.:

```
./agent-downloader -key=<abc> -secret=<abc> -headless -arch=32bit -os=windows -format=zip
```
