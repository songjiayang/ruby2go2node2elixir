var User = require('../models/user');

exports.index = function*() {
  var users = yield User.find();
  this.body = users;
};

exports.show = function*() {
  var user = yield User.findById(this.params.id)
  this.body = user;
};

exports.static = function*() {
  user = {
    "_id": "57472be59cf6a4c96b088a17",
    "username": "Josh Williams",
    "created_at": "2016-05-26T17:01:25.532Z",
    "updated_at": "2016-05-26T17:01:25.532Z",
  }
  this.body = user;
};
