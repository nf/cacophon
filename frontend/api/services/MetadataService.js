var request = require('request');

module.exports = {
  get: function(key, cb) {
    var url = 'http://metadata/computeMetadata/v1' + key;
    request({
      method: 'GET',
      url: url,
      headers: {
        'X-Google-Metadata-Request': 'True'
      }
    }, function(err, response, body) {
      if (err || response.statusCode != 200) {
        var error = {
          message: 'metadata request failed',
          key: key,
          err: err
        };
        if (response) {
          error.statusCode = response.statusCode;
        }
        console.error('MetadataService.get', error);
        cb(error, null);
      } else {
        cb(null, body);
      }
    });
  }
}
