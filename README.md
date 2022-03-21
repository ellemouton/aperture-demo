# Aperture Dynamic Price Demo

### Video Tutorial:

https://www.youtube.com/watch?v=Y2ZG-qcw7Sw

### Install the app:
`make install`

### Diagrams:
*Basic App*
![](./images/app.jpg?raw=true "Basic App")

*Static Prices*
![](./images/aperture-static-prices.jpg?raw=true "Static Prices")

*Dynamic Prices*
![](./images/aperture-dynamic-prices.jpg?raw=true "Static Prices")


### CLI commands:

Add an article:

`appcli addarticle --title="The Blocksize War" --author="Jonathan Bier" --content="All the Bitcoin gossip"`


Add a quote:

`appcli addquote --author="Shakespeare" --content="To be or not to be? That is the question"`


### TODO:

- dockerize the proto generation process. Need go plugins and then run the
  following:    
        `protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative contentrpc/content.proto``

