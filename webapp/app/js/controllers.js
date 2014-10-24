/* global chroma */
'use strict';

/* Controllers */

angular.module('myApp.controllers', ['colorpicker.module'])
    .controller('Cpick', ['$scope','$http',function($scope,$http) {
        $scope.languages = ["ruby","go","python","haskell","javascript"];
        $scope.prevlang = "go";
        $scope.adjustbg = false;

        var faces = {
            deffacefg : "#ffffff", deffacebg : "#ffffff",
            keywordface : "#ffffff", builtinface : "#ffffff",
            stringface : "#ffffff", functionnameface : "#ffffff",
            typeface : "#ffffff", constantface : "#ffffff",
            variableface : "#ffffff", warningface : "#ffffff",
            commentface : "#ffffff"
        };
        var setRandomFaces = function(data) {
            faces["keywordface"] = data.randkey;
            faces["builtinface"]= data.randbuiltin;
            faces["stringface"] = data.randstring;
            faces["functionnameface"] = data.randfuncname;
            faces["typeface"]= data.randtype;
            faces["constantface"] = data.randconst;
            faces["variableface"] = data.randvariable;
        };


        $scope.getRandomColWarm = function() {
            // xmlhttprequest to get a palette of 7 distinct warm colors
            // using go-colorful's WarmPalette method.
            $http.get('/randomcolorswarm').
                success(function(data) {
                    faces.keywordface = data.randkey;
                    faces.builtinface = data.randbuiltin;
                    faces.stringface = data.randstring;
                    faces.functionnameface = data.randfuncname;
                    faces.typeface = data.randtype;
                    faces.constantface = data.randconst;
                    faces.variableface = data.randvariable;
                    // setRandomFaces(data);
                });
        };
        $scope.getRandomColSoft = function() {
            $http.get('/randomcolorssoft').
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

        var darkBg = function(color) {
            return chroma(color).luminance() <=0.5;
        };
        $scope.decContrast = function() {
            if (darkBg($scope.deffacebg)) {
                if ($scope.adjustbg) {
                    $scope.deffacebg = chroma($scope.deffacebg).brighten(1).hex();
                }
                $scope.deffacefg = chroma($scope.deffacefg).darken(1).hex();
                $scope.commentface =chroma($scope.commentface).darken(1).hex();
                $scope.keywordface = chroma($scope.keywordface).darken(1).hex();
                $scope.builtinface = chroma($scope.builtinface).darken(1).hex();
                $scope.stringface = chroma($scope.stringface).darken(1).hex();
                $scope.functionnameface = chroma($scope.functionnameface).darken(1).hex();
                $scope.typeface = chroma($scope.typeface).darken(1).hex();
                $scope.constantface = chroma($scope.constantface).darken(1).hex();
                $scope.variableface = chroma($scope.variableface).darken(1).hex();
            } else {
                if ($scope.adjustbg) {
                    $scope.deffacebg = chroma($scope.deffacebg).darken(1).hex();
                }
                $scope.deffacefg = chroma($scope.deffacefg).brighten(1).hex();
                $scope.commentface =chroma($scope.commentface).brighten(1).hex();
                $scope.keywordface = chroma($scope.keywordface).brighten(1).hex();
                $scope.builtinface = chroma($scope.builtinface).brighten(1).hex();
                $scope.stringface = chroma($scope.stringface).brighten(1).hex();
                $scope.functionnameface = chroma($scope.functionnameface).brighten(1).hex();
                $scope.typeface = chroma($scope.typeface).brighten(1).hex();
                $scope.constantface = chroma($scope.constantface).brighten(1).hex();
                $scope.variableface = chroma($scope.variableface).brighten(1).hex();
            }

        };

        $scope.incContrast = function() {
            if (darkBg($scope.deffacebg)) {
                if ($scope.adjustbg) {
                    $scope.deffacebg = chroma($scope.deffacebg).darken(1).hex();
                }
                $scope.deffacefg = chroma($scope.deffacefg).brighten(1).hex();
                $scope.commentface =chroma($scope.commentface).brighten(1).hex();
                $scope.keywordface = chroma($scope.keywordface).brighten(1).hex();
                $scope.builtinface = chroma($scope.builtinface).brighten(1).hex();
                $scope.stringface = chroma($scope.stringface).brighten(1).hex();
                $scope.functionnameface = chroma($scope.functionnameface).brighten(1).hex();
                $scope.typeface = chroma($scope.typeface).brighten(1).hex();
                $scope.constantface = chroma($scope.constantface).brighten(1).hex();
                $scope.variableface = chroma($scope.variableface).brighten(1).hex();
            } else {
                if ($scope.adjustbg) {
                    $scope.deffacebg = chroma($scope.deffacebg).brighten(1).hex();
                }
                $scope.deffacefg = chroma($scope.deffacefg).darken(1).hex();
                $scope.commentface =chroma($scope.commentface).darken(1).hex();
                $scope.keywordface = chroma($scope.keywordface).darken(1).hex();
                $scope.builtinface = chroma($scope.builtinface).darken(1).hex();
                $scope.stringface = chroma($scope.stringface).darken(1).hex();
                $scope.functionnameface = chroma($scope.functionnameface).darken(1).hex();
                $scope.typeface = chroma($scope.typeface).darken(1).hex();
                $scope.constantface = chroma($scope.constantface).darken(1).hex();
                $scope.variableface = chroma($scope.variableface).darken(1).hex();
            }
        };
    }]);
