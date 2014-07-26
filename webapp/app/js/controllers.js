'use strict';

/* Controllers */

angular.module('myApp.controllers', ['colorpicker.module','vr.directives.slider'])
    .controller('Cpick', ['$scope', function($scope) {

        $scope.languages = ["ruby","go","python","haskell","javascript"];
        $scope.prevlang = "ruby";
        $scope.increaseCont = function() {
            var lum = chroma.luminance($scope.deffacebg);
            if (lum > 0.5) {
                console.log("bright");

            } else {
                console.log("dark");
            }

        };
    }]);
