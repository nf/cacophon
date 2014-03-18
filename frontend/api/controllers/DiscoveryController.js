module.exports = {
  discover: function(req, res) {
    if (process.env.GAE_VM) {
      MetadataService.get('/instance/network-interfaces/0/access-configs/0/external-ip', function(err, data) {
        if (err) {
          return res.send(500, err);
        }
        res.send({
          https: 'https://'+data':8080',
          websocket: 'wss://'+data+':8080'
        });
      });
    } else {
      res.send({
        https: 'https://localhost::8080',
        websocket: 'wss://localhost:1337'
      });
    }
  },
  cookie: function(req, res) {
    res.send(200);
  }
}
