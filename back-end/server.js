const http = require('http');
var LRU = require("lru-cache")
  , options = { max: 500
              , length: function (n, key) { return n * 2 + key.length }
              , maxAge: 1000 * 60 * 60 }
  , cache = new LRU(options)
  , otherCache = new LRU(50) // sets just the max size
const express = require('express')
const app = express()

const bodyParser = require('body-parser');
app.use(bodyParser.json()); // for parsing application/json


app.get('/', function (request, response) {
  response.send('Simple test for liveliness of the application!');
});

app.get('/storage', function (request, response) {
	var res = cache.dump();
	response.writeHead(200);
	response.write(JSON.stringify(res) ) ;
	response.end();
});


app.delete('/storage/:key', function (request, response) {
	cache.del(request.params.key);
	response.writeHead(200);
	response.end();
});


app.get('/storage/:key', function (request, response) {
	var res = cache.get(request.params.key);
	response.writeHead(200);
	response.write(JSON.stringify(res) ) ;
	response.end();
});


app.put('/storage', function (request, response) {
	var jsonData = request.body;
	cache.set(jsonData.key, jsonData.value);
	response.writeHead(201);
	response.write("nodeAppTesting created") ;
	response.end();
});

app.listen(9080, function () {
  console.log('Example app listening on port 9080!')
});