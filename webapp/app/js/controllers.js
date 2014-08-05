'use strict';

/* Controllers */

angular.module('myApp.controllers', ['colorpicker.module'])
    .controller('Cpick', ['$scope','$http',function($scope,$http) {
        $scope.languages = ["ruby","go","python","haskell","javascript"];
        $scope.prevlang = "ruby";

        $scope.getRandomColWarm = function() {
            // xmlhttprequest to get a palette of 7 distinct warm colors
            // using go-colorful's WarmPalette method.
            $http.get('/randomcolorswarm').
                success(function(data) {
                    $scope.keywordface = data.randkey;
                    $scope.builtinface = data.randbuiltin;
                    $scope.stringface = data.randstring;
                    $scope.functionnameface = data.randfuncname;
                    $scope.typeface = data.randtype;
                    $scope.constantface = data.randconst;
                    $scope.variableface = data.randvariable;
                });
        };
        $scope.getRandomColHappy = function() {
            $http.get('/randomcolorshappy').
                success(function(data) {
                    $scope.keywordface = data.randkey;
                    $scope.builtinface = data.randbuiltin;
                    $scope.stringface = data.randstring;
                    $scope.functionnameface = data.randfuncname;
                    $scope.typeface = data.randtype;
                    $scope.constantface = data.randconst;
                    $scope.variableface = data.randvariable;
                });
        };


                }]);
