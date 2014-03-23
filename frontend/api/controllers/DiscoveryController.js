module.exports = {
  discover: function(req, res) {
    DiscoveryService.discoverExternalIP(function(err, ip) {
      if (err) {
        return res.send(500, err);
      }
      res.send({
        websocket: '//'+ip+':'+sails.config.port
      });
    });
  },
  cookie: function(req, res) {
    res.set('Content-Type', 'text/javascript');
    res.send(200, req.query.callback+'()');
  }
}
