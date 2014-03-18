/**
    * SoundController
    *
    * @module      :: Controller
    * @description	:: A set of functions called `actions`.
    *
    *                 Actions contain code telling Sails how to respond to a certain type of request.
    *                 (i.e. do stuff, then send some JSON, show an HTML page, or redirect to another URL)
    *
    *                 You can configure the blueprint URLs which trigger these actions (`config/controllers.js`)
    *                 and/or override them with custom routes (`config/routes.js`)
    *
    *                 NOTE: The code you write here supports both HTTP and Socket.io automatically.
    *
    * @docs        :: http://sailsjs.org/#!documentation/controllers
    */
var base64 = require('base64-stream');
var stream = require('stream');

module.exports = {


  /**
      * Overrides for the settings in `config/controllers.js`
      * (specific to SoundController)
  */
  _config: {},

  render: function(req, res) {
    Sound.findOne(req.param('id')).exec(function (err, sound) {
      if (err) {
        return res.send(err, 500);
      }
      if (!sound.mp3) {
        return res.send('sound #' + req.param('id') + ' has no mp3 data', 415);
      }
      res.set('Content-Type', 'audio/mp3');
      var s = new stream.Readable();
      s.push(sound.mp3);
      s.push(null);
      s.pipe(base64.decode()).pipe(res);
    });
  },

  subscribe: function(req, res) {
    Sound.subscribe(req.socket);
    res.send(200);
  },
};
