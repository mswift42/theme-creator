'use strict';

/* Directives */


angular.module('myApp.directives', []).
  directive('appVersion', ['version', function(version) {
    return function(scope, elm, attrs) {
      elm.text(version);
    };
  }]).
    directive('themeCreator',function() {
        return {
            restrict: 'EA',
            templateUrl: 'partials/themecreator.html',
            controller: 'Cpick'
        };
    }).
    directive('themePreview',function() {
        return {
            restrict: 'EA',
            templateUrl: 'partials/preview.html',
            controller: 'Cpick'
        };
    }).
    directive('pythonView',function() {
        return {
            restrict: 'EA',
            templateUrl: 'partials/python.html',
            controller: 'Cpick'
        };
    }).
    directive('rubyView',function() {
        return {
            restrict: 'EA',
            templateUrl: 'partials/ruby.html',
            controller: 'Cpick'
        };
    }).
    directive('haskellView',function() {
        return {
            restrict: 'EA',
            templateUrl: 'partials/haskell.html',
            controller: 'Cpick'
        };
    }).
    directive('jsView',function() {
        return {
            restrict: 'EA',
            templateUrl: 'partials/js.html',
            controller: 'Cpick'
        };
    });
