'use strict';

/* Controllers */

angular.module('myApp.controllers', ['colorpicker.module'])
    .controller('Cpick', ['$scope', function($scope) {

        $scope.languages = ["ruby","go","python","haskell","javascript"];
        $scope.prevlang = "ruby";

    }]);
