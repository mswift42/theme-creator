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



});
