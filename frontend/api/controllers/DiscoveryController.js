module.exports = {
  discover: function(req, res) {
    res.send({
      websocket: 'http://localhost:1337'
    });
  }
}
