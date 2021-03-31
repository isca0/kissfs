### KISS FileServer
  
This is a Web based FileServer extremely simple on K.I.S.S philosophy.  
  
It is writed in go as a study case.
  
### Compiling

```bash
go get github.com/mgutz/ansi
go build kissfs.go -tags netgo -a -v
```

### Run manualy
  
```bash
./kissfs
```

_By default it will run on port 3000 and export everything on current directory._   

To customize the default options use the environment variables:

* `KISSFS_DIR` - `string` - Set the directory to be used in the file server. `default: current directory`  
* `KISSFS_PORT` - `string` - Set the default port where the server will listen. `default 3000`.  

Examples:
```bash
export KISSFS_DIR="/home/user/www"
export KISSFS_PORT="8080"
kissfs
```
_Start the fileserver on the port 8080 and serve everything from the `/home/user/www` directory._  
  
One liner execution:
```bash
KISSFS_DIR="/opt/www" KISSFS_PORT="8081" kissfs
```
_Start the fileserver in the port `8081` and expose everything from the `/opt/www` directory._  
  
### Run from Docker
  
```bash
docker run -p 3000:3000 -e KISSFS_DIR=/var kissfs
```
  
### Docker build
  
_This is a multi-stage Dockerfile, you will need docker 17.05 or higher_   

```bash
docker built -t kissfs .
```
  
### Run from Docker compose
  
This compose use a custom network called web. If you dont want use custom network remove `networks`  
from this compose file.  
  
```bash
KISSFS_DIR="/somefolder" docker-compose up
```
  
Enjoy :wink:  
  

