'use strict';

/* jasmine specs for controllers go here */

describe('controllers', function(){
    beforeEach(module('myApp.controllers'));



    it('Cpick controller should exist.', inject(function($controller) {
        //spec body
        var Cpick = $controller('Cpick', { $scope: {} });
        expect(Cpick).toBeDefined();
    }));
    it('$scope.adjustbg should be set to false.', inject(function($controller, $rootScope) {

        var scope = $rootScope.$new();
        var Cpick = $controller('Cpick', { $scope: scope});
        scope.$digest();
        expect(scope.adjustbg === false);
        expect(scope.prevlang === 'ruby');
        expect(scope.faces.deffacefg === "#303030");
    }));
    it('darkBg function should return true if supplied color is dark',
       inject(function($controller, $rootScope) {
           var scope = $rootScope.$new();
           var Cpick = $controller('Cpick', {$scope: scope});
           scope.$digest();
           expect(scope.darkBg('#000000') === true);
           expect(scope.darkBg('#ffffff') === false);
           expect(scope.darkBg('#002b36') === true);
           expect(scope.darkBg('#586e75') === true);
           expect(scope.darkBg('#fdf6e3') === false);
           expect(scope.darkBg('#87e37e') === false);
       }));



});
