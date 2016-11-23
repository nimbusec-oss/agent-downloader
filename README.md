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