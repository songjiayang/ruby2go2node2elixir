var Schema = require('mongoose').Schema;

var UserSchema = new Schema({
  name: { type: String }
});

/* global db */
module.exports = db.model('User', UserSchema);
