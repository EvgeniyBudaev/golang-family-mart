#  Golang family mart


## Database migrations
https://github.com/golang-migrate/migrate
### Install 
```bash
make install_migrate
```
or 
```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@$TAG
```
      

### Migrate
```bash
make migrate 
```