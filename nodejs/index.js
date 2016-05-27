var app = require('koa')();
var router = require('koa-router')();

// database
var mongoose = require('mongoose');
var uri = 'mongodb://localhost/ruby2go2node2elixir';
var options = {
  server: { poolSize: 20 },
}
global.db = app.context.db = mongoose.createConnection(uri, options);

// controller
var users = require('./controllers/users');

// routes
router.get('/users', users.index);
router.get('/users/:id', users.show);
router.get('/static', users.static);

app.use(router.routes());

const cluster = require('cluster');
const http = require('http');
const numCPUs = require('os').cpus().length;

if (cluster.isMaster) {
  for (var i = 0; i < numCPUs; i++) {
    cluster.fork();
  }
  cluster.on('exit', (worker, code, signal) => {
    console.log(`worker ${worker.process.pid} died`);
  });
} else {
  http.createServer(app.callback())
    .listen(3001, function(){
      console.log('app listening on port 3001!');
    });
}

// var server = require('http').Server(app.callback());
// server.listen(3001, function () {
//   console.log('app listening on port 3001 with single process mode!');
// });
