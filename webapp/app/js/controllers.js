'use strict';

/* Controllers */

angular.module('myApp.controllers', ['colorpicker.module','ngResource'])
    .controller('Cpick', ['$scope', function($scope) {

        $scope.message = "hallo";
        $scope.languages = ["python","ruby","haskell","javascript"];
        $scope.prevlang = "ruby";
        $scope.keywordface = "#a52a2a";
        $scope.deffacefg = "#2e3436";
        $scope.deffacebg = "#ededed";
        $scope.builtinface = "#a020f0";
        $scope.constantface = "#f5666d";
        $scope.commentface = "#204a87";
        $scope.functionnameface = "#00578e";
        $scope.stringface = "#4e9a06";
        $scope.typeface = "#2f8b58";
        $scope.variableface = "#0084c8";
        $scope.warningface = "#f5666d";
        // $scope.saveTheme = function(){
        //     var theme = new Theme();
        //     theme.keywordface = $scope.keywordface;
        //     theme.deffacebg = $scope.deffacebg;
        //     theme.deffacefg = $scope.deffacefg;
        //     theme.builtinface = $scope.builtinface;
        //     theme.constantface = $scope.constantface;
        //     theme.commentface = $scope.commentface;
        //     theme.functionnameface = $scope.functionnameface;
        //     theme.stringface = $scope.stringface;
        //     theme.typeface = $scope.typeface;
        //     theme.variableface = $scope.variableface;
        //     theme.warningface = $scope.warningface;
        //     theme.$save();




    }])
    .controller('DorLightCtrl',['$scope',function($scope) {
        $scope.setbg = 'light';
    }]);
