/* global chroma */
'use strict';

/* Controllers */

angular.module('myApp.controllers', ['colorpicker.module'])
    .controller('Cpick', ['$scope','$http','lstorage', 'presets', function($scope,$http,lstorage,presets) {
        $scope.languages = ["ruby","go","python","haskell","javascript"];
        $scope.prevlang = "ruby";
        $scope.adjustbg = false;
        $scope.presettheme = "preset";

        var initialfaces = function (item) {
            var faces = 
                    {
                        deffacefg : "#303030", deffacebg : "#ffffff",
                        keywordface : "#000000", builtinface : "#000000",
                        stringface : "#000000", functionnameface : "#000000",
                        typeface : "#000000", constantface : "#000000",
                        variableface : "#000000", warningface : "#ff0000",
                        commentface : "#606060"
                    };
            return faces;
        };
        $scope.changePreset = function() {
            var newtheme = $scope.presettheme;
            switch (newtheme) {
            case "warmnight":
                $scope.faces = presets.warmnight();
                break;
            case "oldlace":
                $scope.faces = presets.oldlace();
                break;
            case "lightsoap":
                $scope.faces = presets.lightsoap();
                break;
            case "greymatters":
                $scope.faces = presets.greymatters();
                break;
            case "softcharcoal":
                $scope.faces = presets.softcharcoal();
                break;
            case "softmorning":
                $scope.faces = presets.softmorning();
                break;
            case "softstone":
                $scope.faces = presets.softstone();
                break;
            default:
                $scope.faces = initialfaces();
            }
        };


        // if faces are not stored in localStorage
        // set $scope.faces to initialfaces.
        $scope.faces = lstorage.loadFaces() || initialfaces();
        
        // store faces in localStorage.
        $scope.saveTheme = function() {
            lstorage.setFaces($scope.faces);
        };

        // delete localStorage key, set $scope.faces to initialfaces.
        $scope.resetTheme = function () {
            lstorage.deleteKey();
            $scope.faces = initialfaces();

        };

        // setRandomFaces - helper function for ajax calls
        // to ranomColor handlers.
        $scope.setRandomFaces = function(data) {
            $scope.faces.keywordface = data.randkey;
            $scope.faces.builtinface= data.randbuiltin;
            $scope.faces.stringface = data.randstring;
            $scope.faces.functionnameface = data.randfuncname;
            $scope.faces.typeface= data.randtype;
            $scope.faces.constantface = data.randconst;
            $scope.faces.variableface = data.randvariable;
        };


        $scope.getRandomColWarm = function() {
            // xmlhttprequest to get a palette of 7 distinct warm colors
            // using go-colorful's WarmPalette method.
            $http.get('/randomcolorswarm').
                success(function(data) {
                    $scope.setRandomFaces(data);
                });
        };
        $scope.getRandomColSoft = function() {
            $http.get('/randomcolorssoft').
                success(function(data) {
                    $scope.setRandomFaces(data);
                });
        };
        $scope.getRandomColHappy = function() {
            $http.get('/randomcolorshappy').
                success(function(data) {
                    $scope.setRandomFaces(data);
                });
        };
        // darkBg - return True if supplied color is dark.
        $scope.darkBg = function(color) {
            return chroma(color).luminance() <=0.5;
        };

        $scope.facenames = ["deffacefg","commentface","keywordface",
                            "builtinface","stringface","functionnameface",
                            "typeface","constantface","variableface"];
        $scope.decContrast = function() {
            if ($scope.darkBg($scope.faces.deffacebg)) {
                if ($scope.adjustbg) {
                    $scope.faces.deffacebg = chroma($scope.faces.deffacebg).brighten(1).hex();
                }
                for (var i = 0;i<$scope.facenames.length;i++) {
                    $scope.faces[$scope.facenames[i]] = chroma($scope.faces[$scope.facenames[i]]).darken(1).hex();
                }
            } else {
                if ($scope.adjustbg) {
                    $scope.faces.deffacebg = chroma($scope.faces.deffacebg).darken(1).hex();
                }
                for ( i=0;i<$scope.facenames.length;i++) {
                    $scope.faces[$scope.facenames[i]] = chroma($scope.faces[$scope.facenames[i]]).brighten(1).hex();
                }
            }

        };

        $scope.incContrast = function() {
            if ($scope.darkBg($scope.faces.deffacebg)) {
                if ($scope.adjustbg) {
                    $scope.faces.deffacebg = chroma($scope.faces.deffacebg).darken(1).hex();
                }

                for (var i=0;i<$scope.facenames.length;i++) {
                    $scope.faces[$scope.facenames[i]] = chroma($scope.faces[$scope.facenames[i]]).brighten(1).hex();
                }
            } else {
                if ($scope.adjustbg) {
                    $scope.faces.deffacebg = chroma($scope.faces.deffacebg).brighten(1).hex();
                }
                for (i=0;i<$scope.facenames.length;i++) {
                    $scope.faces[$scope.facenames[i]] = chroma($scope.faces[$scope.facenames[i]]).darken(1).hex();
                }
            }
        };
    }]);
