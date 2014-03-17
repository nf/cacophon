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
    speed: 'integer',
    scale: 'integer',
    perm: 'integer',
    slew: 'integer',
    root: 'integer',
    square: 'integer',
    amount: 'integer',
    offset: 'integer',
    attack: 'integer',
    decay: 'integer',
    time: 'integer',
    feedback: 'integer',
    mp3: 'string'
  },

  beforeCreate: function(v, next) {
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
        console.error('mp3 encoding failed:', err);
      } else {
        v.mp3 = body.toString('base64');
      }
      next();
    });
  }
};
