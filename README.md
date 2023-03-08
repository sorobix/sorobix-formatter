# Sorobix-formatter

> A simple API to format rust code, used for sorobix


## API Docs

## POST
- Endpoint: `/`
- Request Sample:
  - ```json
    {
    "code": "Zm4gbWFpbigpIHsgcHJpbnRsbiEoIkhlbGxvIFdvcmxkISIpOyB9"
    }
    ```
- Response Sample:
  - ```json
    {
    "formatted_code": "Zm4gbWFpbigpIHsKICAgIHByaW50bG4hKCJIZWxsbyBXb3JsZCEiKTsKfQo="
    }
    ```
#### Status Codes

- `200`: formatter_code
- `400`: Invalid/Bad Request
- `406`: Bad Rust Code

## Deployment

- Build the Image
```bash
docker build . -t sorobix-formatter
```

- Run the Image

```bash
docker run -p 3000:3000 sorobix-formatter:latest 
```