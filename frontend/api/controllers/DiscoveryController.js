module.exports = {
  discover: function(req, res) {
    if (process.env.GAE_VM) {
      MetadataService.get('/instance/network-interfaces/0/access-configs/0/external-ip', function(err, data) {
        if (err) {
          return res.send(500, err);
        }
        res.send({
          websocket: '//'+data+':8080'
        });
      });
    } else {
      res.send({
        websocket: '//localhost:1337'
      });
    }
  },
  cookie: function(req, res) {
    res.set('Content-Type', 'text/javascript');
    res.send(200, req.query.callback+'()');
  }
}
