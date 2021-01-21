# La Trappe Melder

## Usage

First time configuration:

Edit `.env`.

Then run installation on Listmonk.

```bash
docker-compose up db
docker-compose run listmonk ./listmonk --install
docker-compose up
```

Then go to `http://localhost:9000/settings` and configure listmonk.

Next create the `la-trappe-melder`list. And put the new Id in the `.env` file.



Listmonk requests:
```bash
curl 'http://localhost:9000/api/campaigns' \
-X 'POST' \
-H 'Accept: application/json, text/plain, */*' \
-H 'Content-Type: application/json;charset=utf-8' \
-H 'Origin: http://localhost:9000' \
-H 'Referer: http://localhost:9000/campaigns/new' \
-H 'Authorization: Basic dGVzdDp0ZXN0' \
-H 'Content-Length: 195' \
-H 'Host: localhost:9000' \
-H 'Accept-Language: nl-be' \
-H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0.2 Safari/605.1.15' \
-H 'Accept-Encoding: gzip, deflate' \
-H 'Connection: keep-alive' \
--data-binary '{"name":"Batch 37","subject":"Batch 37","lists":[3],"from_email":"La Trappe Melder <no-reply@denbeke.be>","content_type":"richtext","messenger":"email","type":"regular","tags":[],"template_id":1}'



curl 'http://localhost:9000/api/campaigns/2' \
-X 'PUT' \
-H 'Accept: application/json, text/plain, */*' \
-H 'Content-Type: application/json;charset=utf-8' \
-H 'Origin: http://localhost:9000' \
-H 'Referer: http://localhost:9000/campaigns/2' \
-H 'Authorization: Basic dGVzdDp0ZXN0' \
-H 'Content-Length: 263' \
-H 'Host: localhost:9000' \
-H 'Accept-Language: nl-be' \
-H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0.2 Safari/605.1.15' \
-H 'Accept-Encoding: gzip, deflate' \
-H 'Connection: keep-alive' \
--data-binary '{"name":"Batch 37","subject":"Batch 37","lists":[3],"from_email":"La Trappe Melder <no-reply@denbeke.be>","messenger":"email","type":"regular","tags":[],"send_later":false,"send_at":null,"template_id":1,"content_type":"richtext","body":"<p>Batch 37 is daar!</p>"}'


curl 'http://localhost:9000/api/campaigns/4/status' \
-X 'PUT' \
-H 'Accept: application/json, text/plain, */*' \
-H 'Content-Type: application/json;charset=utf-8' \
-H 'Origin: http://localhost:9000' \
-H 'Referer: http://localhost:9000/campaigns' \
-H 'Authorization: Basic dGVzdDp0ZXN0' \
-H 'Content-Length: 20' \
-H 'Host: localhost:9000' \
-H 'Accept-Language: nl-be' \
-H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0.2 Safari/605.1.15' \
-H 'Accept-Encoding: gzip, deflate' \
-H 'Connection: keep-alive' \
--data-binary '{"status":"running"}'
```