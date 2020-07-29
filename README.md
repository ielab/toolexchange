# toolexchange

This is a service for exchanging data between systematic review automation tools.

Two tools may communicate data between each other through a shared token. 
To create a token, first a tool submits the data they would like to share to this service. 
If the request is valid, this service will respond with a token. 
This token is then sent to the second tool, and the second tool is then able to access the data stored on this service as it sees fit.

Currently, data is stored on the exchange service for 5 minutes.

## Automation tools with exchange implementations

 - [SRA Polyglot](https://sr-accelerator.com/#/polyglot)
 - [searchrefiner](https://ielab.io/searchrefiner/)

## Usage

The requests that can be made to the exchange server are as follows:

### Create token

request:

```bash
curl -X POST '/exchange' \
     -H 'Content-Type: application/json' \
     -d '{
          "data": {
              "query": "a and b"
          }, 
          "referrer": "searchrefiner"
         }'
```

response:

```
4321example1234
```

### Request data with token

request:

```bash
curl -X GET '/exchange?token=4321example1234'
```

response:

```json
{
 "data": {
     "query": "a and b"
 }, 
 "referrer": "searchrefiner",
 "expiration": "2020-07-29T03:23:26.74893582Z"
}
```
