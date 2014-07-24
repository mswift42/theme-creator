'use strict';

/* Controllers */

angular.module('myApp.controllers', ['colorpicker.module'])
    .controller('Cpick', ['$scope', function($scope) {

        $scope.languages = ["ruby","go","python","haskell","javascript"];
        $scope.prevlang = "ruby";
        $scope.increaseCont = function() {
            var lum = chroma.luminance($scope.deffacefg);
            if (lum > 0.5) {
            $scope.deffacefg = chroma($scope.deffacefg).darken(2).hex();
            $scope.builtinface = chroma($scope.builtinface).darken(2).hex();
            $scope.stringface = chroma($scope.stringface).darken(2).hex();
            $scope.keywordface = chroma($scope.stringface).darken(2).hex();
            $scope.functionnameface = chroma($scope.functionnameface).darken(2).hex();
            $scope.commentface = chroma($scope.commentface).darken(2).hex();
            $scope.variableface = chroma($scope.variableface).darken(2).hex();
            $scope.typeface = chroma($scope.variableface).darken(2).hex();
                $scope.constantface = chroma($scope.constantface).darken(2).hex();
            } else {

                $scope.deffacefg = chroma($scope.deffacefg).lighten(2).hex();
                $scope.builtinface = chroma($scope.builtinface).lighten(2).hex();
                $scope.stringface = chroma($scope.stringface).lighten(2).hex();
                $scope.keywordface = chroma($scope.stringface).lighten(2).hex();
                $scope.functionnameface = chroma($scope.functionnameface).lighten(2).hex();
                $scope.commentface = chroma($scope.commentface).lighten(2).hex();
                $scope.variableface = chroma($scope.variableface).lighten(2).hex();
                $scope.typeface = chroma($scope.variableface).lighten(2).hex();
                $scope.constantface = chroma($scope.constantface).lighten(2).hex();

            }

        };
    }]);
