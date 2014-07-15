'use strict';

/* Directives */


angular.module('myApp.directives', []).
  directive('appVersion', ['version', function(version) {
    return function(scope, elm, attrs) {
      elm.text(version);
    };
  }]).
    directive('themePreview',function() {
        return {
            restrict: 'EA',
            templateUrl: 'partials/preview.html'
        };
    }).
    directive('pythonView',function() {
        return {
            restrict: 'EA',
            templateUrl: 'partials/python.html'
        };
    }).
    directive('rubyView',function() {
        return {
            restrict: 'EA',
            templateUrl: 'partials/ruby.html'
        };
    });
