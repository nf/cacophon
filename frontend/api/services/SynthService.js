var qs = require('querystring');
var request = require('request');

module.exports = {
  encode: function(v, cb) {
    var params = qs.stringify({
      speed: v.speed / 100.0,
      scale: v.scale,
      perm: v.perm,
      slew: v.slew / 100.0,
      root: v.root / 100.0,
      square: v.square / 100.0,
      amount: v.amount / 100.0,
      offset: v.offset / 100.0,
      attack: v.attack / 100.0,
      decay: v.decay / 100.0,
      time: v.time / 100.0,
      feedback: v.feedback / 100.0
    });
    var url = 'http://localhost:8080/audio?' + params;
    request({
      method: 'GET',
      url: url,
      encoding: null
    }, function(err, response, body) {
      if (err || response.statusCode != 200) {
        var error = {
          message: 'mp3 encoding failed',
          statusCode: response.statusCode,
          err: err
        };
        console.error('SynthService.encode', error);
        cb(error, null);
      } else {
        cb(null, body.toString('base64'));
      }
    });
  }
}
