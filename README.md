# GO,Gin,GORM,HTMX,TailwindCSS Template


This is Template.

## Tech Stack Involved

**GO** I used go for this project <br/>
**Gin** A web framework I used for this project <br/>
**GORM** An ORM library for to intract with the postgresql database <br/>
**HTMX** Frontend library used along with templ for to manage the UI <br/>
**TailwindCSS** A CSS framework used along with the htmx for to desing user interfaces <br/>
**Docker** Making the application containerized.
**AWS** Deploy docker containers to cloud service Planning to use __AWS ECS, EC2, RDS, S3, Cloud Watch__ and __Github Actions__, 

### Steps to install templ 

```console
npx tailwindcss -o ./static/css/style.css
templ generate
go build -o htmx_app main.go
./htmx_app
```
