'use strict';

/* Controllers */

angular.module('myApp.controllers', ['colorpicker.module'])
    .controller('Cpick', ['$scope', function($scope) {
        $scope.message = "hallo";

    }])
    .controller('DorLightCtrl',['$scope',function($scope) {
        $scope.setbg = 'light';
    }]);
