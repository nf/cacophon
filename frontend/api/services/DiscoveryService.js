var os = require('os');
var request = require('request');
var netroute = require('netroute');
var backend_name = 'backend_0';

module.exports = {
  discoverWebsocket: function(cb) {
    if (process.env.GAE_VM) {
      MetadataService.get('/instance/network-interfaces/0/access-configs/0/external-ip', cb);
    } else {
      cb(null, os.networkInterfaces().eth0[0].address)
    }
  },

  discoverBackend: function(cb) {
    if (process.env.GAE_VM) {
      MetadataService.get('/project/attributes/cacophon-backend', cb);
    } else {
      var gateway = netroute.getGateway();
      request({
        url: 'http://'+gateway+':4243/containers/'+backend_name+'/json',
        json: true
      }, function(err, response, body) {
        if (err || response.statusCode != 200) {
          var error = {
            message: 'container introspection error',
            container: backend_name,
            err: err
          };
          if (response) {
            error.statusCode = response.statusCode;
          }
          console.error('discoverBackend', error);
          cb(error, null);
        } else {
          cb(null, 'http://' + body.NetworkSettings.IPAddress + ':8080');
        }
      });
    }
  }
}
