/**
 * Sound
 *
 * @module      :: Model
 * @description :: Store a cacophon sound.
 * @docs		:: http://sailsjs.org/#!documentation/models
 */

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
    SynthService.encode(v, function(err, mp3) {
      if (mp3) {
        v.mp3 = mp3;
      }
      next();
    });
  }
};
