# La Trappe Melder

Get notified when a new batch of La Trappe Quadrupel Oak Aged is released! 🍻

## Usage

**Configuration:**

Edit your configuration in `.env` (for env variables can be found in `config.go`).

**Run the la-trappe-melder:**

Then either use the docker-compose.yml file to run it:

```bash
docker-compose up -d
```

Or run it with Go:

```bash
go run cmd/latrappemelder/main.go
```


## Acknowledgements

- [labstack/echo](https://github.com/labstack/echo)
- [sirupsen/logrus](https://github.com/sirupsen/logrus)
- [xo/dburl](https://github.com/xo/dburl)
- [go-gorm/gorm](https://github.com/go-gorm/gorm)
- [PuerkitoBio/goquery](https://github.com/PuerkitoBio/goquery)
- [google/uuid](https://github.com/google/uuid)
- [go-mail/mail](https://github.com/go-mail/mail)



## Author

[Mathias Beke](https://denbeke.be)