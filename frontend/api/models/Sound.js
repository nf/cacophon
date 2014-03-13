/**
 * Sound
 *
 * @module      :: Model
 * @description :: Store a cacophon sound.
 * @docs		:: http://sailsjs.org/#!documentation/models
 */
var request = require('request');
var qs = require('querystring');

module.exports = {
  attributes: {

  	/* e.g.
  	nickname: 'string'
  	*/
        a: 'integer',
        b: 'integer',
        c: 'integer',
        mp3: 'string'
  },

  beforeCreate: function(values, next) {
    var params = qs.stringify({
      a: values.a / 100.0,
      b: values.b / 100.0,
      c: values.c / 100.0
    });
    var url = 'http://localhost:8080/audio?' + params;
    request({
      method: 'GET',
      url: url,
      encoding: null
    }, function(err, response, body) {
      if (err && response.statusCode == 200) {
        console.error('mp3 encoding failed:', err);
      } else {
        values.mp3 = body.toString('base64');
      }
      next();
    });
  }
};
