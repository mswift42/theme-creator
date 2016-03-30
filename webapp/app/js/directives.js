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
        };
    }).
    directive('themePreview',function() {
        return {
            restrict: 'EA',
            templateUrl: 'partials/preview.html',
        };
    }).
    directive('cView',function() {
        return {
            restrict: 'EA',
            templateUrl: 'partials/c.html',
        };
    }).
    directive('cppView',function() {
        return {
            restrict: 'EA',
            templateUrl: 'partials/cpp.html',
        };
    }).
    directive('pythonView',function() {
        return {
            restrict: 'EA',
            templateUrl: 'partials/python.html',
        };
    }).
    directive('rubyView',function() {
        return {
            restrict: 'EA',
            templateUrl: 'partials/ruby.html',
        };
    }).
    directive('haskellView',function() {
        return {
            restrict: 'EA',
            templateUrl: 'partials/haskell.html',
        };
    }).
    directive('jsView',function() {
        return {
            restrict: 'EA',
            templateUrl: 'partials/js.html',
        };
    }).
    directive('goView',function() {
        return {
            restrict: 'EA',
            templateUrl: 'partials/go.html',
        };
    });
