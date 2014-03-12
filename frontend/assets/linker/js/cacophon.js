function value(id) {
  var el = document.getElementById(id);
  return el.value / 100;
}

var el;

function CacophonCtrl($scope) {
  $scope.sounds = [
    {url: 'http://foo'}
  ];
  $scope.listen = function () {
    var a = value('knobA');
    var b = value('knobB');
    var c = value('knobC');
    if (el) document.body.removeChild(el)
    el = document.createElement('audio');
    el.src = '/audio?a='+a+'&b='+b+'&c='+c;
    el.autoplay = true;
    document.body.appendChild(el);
    $scope.sounds.push({url: el.src});
  }
}
