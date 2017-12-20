# Web API Definition
> Definition of mock web API server using custom .wad file

## .wad definition file
[.wad](./app.wad) file contains definiton of your API

```
-port 3000

GET "/api/hello"
-response
    {"message": "hello"}

GET "/api/world"
-response "world"
```

### What this example does?
1) Creates localhost http server on port 3000
2) Creates ``HTTP GET`` enpoint ``/api/hello`` which returns ``application/json`` data
```
{
    "message": "hello"
}
```
3) Creates ``HTTP GET`` endpoint ``/api/world`` which returns string ``world`` as ``text/plain``
4) Responds to every other request with ``404 page not found``

## Run .wad file
### Build and run
1) ``go build`` contents of [./scr](./src)
2) place builded executable file in same folder as [your.wad](./app.wad) file
3) ... execute

### Windows
1) place [./wad.exe](./wad.exe) in same folder as [your.wad](./app.wad) file
2) ... execute