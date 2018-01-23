### KISS FileServer
  
This is a Web based FileServer extremely simple on K.I.S.S philosophy.  
  
It is writed in go as a study case.
  
### Compiling

```
go get github.com/mgutz/ansi
go build kissfs.go -tags netgo -a -v
```

### Run manualy
  
```
./kissfs
```

_By default it will run on port 3000 and export everything on `/www` folder._   
  
So basically you need to place everything you want on fileserver on `/www` folder and run  
`kissfs`  
  
### Run from Docker
  
```
docker run -p 3000:3000 -v somefolder/www kissfs
```
  
### Docker build
  
_This is a multi-stage Dockerfile, you will need docker 17.05 or higher_   

```
docker built -t kissfs .
```
  
### Run from Docker compose
  
This compose use a custom network called trae. If you dont want use custom network remove `networks`  
from this compose file.  
  
```
docker network create trae
export VOL="/somefolder";export HOST=127.0.0.1; docker-compose up
```
  
Enjoy :wink:  
  

