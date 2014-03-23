var os = require('os');
var request = require('request');

module.exports = {
  discoverExternalIP: function(cb) {
    if (process.env.GAE_VM) {
      MetadataService.get('/instance/network-interfaces/0/access-configs/0/external-ip', cb);
    } else {
      cb(null, os.networkInterfaces().eth0[0].address)
    }
  },

  discoverDispatcher: function() {
    if (process.env.GAE_VM) {
      return process.env.APP_ID + '.appspot.com';
    } else {
      return process.env.DISPATCHER_HOST + ':8080';
    }
  }
}
