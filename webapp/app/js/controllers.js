'use strict';

/* Controllers */

angular.module('myApp.controllers', ['colorpicker.module','ngResource'])
    .controller('Cpick', ['$scope', function($scope) {

        $scope.message = "hallo";
        $scope.languages = ["ruby","python","haskell","javascript"];
        $scope.prevlang = "ruby";


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




    }]);
