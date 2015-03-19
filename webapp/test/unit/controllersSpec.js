'use strict';

/* jasmine specs for controllers go here */

describe('controllers', function(){
    beforeEach(module('myApp.controllers'));
    beforeEach(module('myApp.services'));



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
           var darkfaces = ["#000000","#002b36","#586e75",
                            "#932ad7"];
           var lightfaces = ["#ffffff","#fdf6e3","#d2d2d2",
                             "#d7cf47","#87e37e"];
           for (var i = 0;i<darkfaces.length;i++) {
               expect(scope.darkBg(darkfaces[i])).toEqual(true);
           }
           for (i = 0;i<lightfaces.length;i++) {
               expect(scope.darkBg(lightfaces[i])).toEqual(false);
           }
       }));
    it('setRandomFaces should assign to $scope.faces the supplied data.',
       inject(function($controller, $rootScope) {
           var scope = $rootScope.$new();
           var Cpick = $controller('Cpick', {$scope: scope});
           scope.$digest();
           var data = {"randkey": "#333333","randbuiltin":"#444444",
                       "randstring":"#555555","randfuncname":"#666666"};
           scope.setRandomFaces(data);
           expect(scope.faces.keywordface).toEqual("#333333");
           expect(scope.faces.builtinface).toEqual("#444444");

       }));
    it('facenames should contain the correct names.',
       inject(function($controller, $rootScope) {
           var scope = $rootScope.$new();
           var Cpick = $controller('Cpick', {$scope: scope});
           scope.$digest();

           expect(scope.facenames[0]).toEqual("deffacefg");
           expect(scope.facenames[1]).toEqual("commentface");
           expect(scope.facenames[4]).toEqual("stringface");
       }));

    it('decContrast should only affect background if adjustbg is set to true.',
      inject(function($controller,$rootScope) {
        var scope = $rootScope.$new();
        var Cpick = $controller('Cpick', {$scope: scope});
        scope.$digest();
        scope.adjustbg = false;
        scope.decContrast();
        expect(scope.faces.deffacebg).toEqual("#ffffff");
        scope.adjustbg = true;
        scope.decContrast();
        expect(scope.faces.deffacebg).toEqual("#fcfcfc");

      }));

    it('saveTheme fn should store faces object into localStorage',
       inject(function($controller, $rootScope) {
           var scope = $rootScope.$new();
           var Cpick = $controller('Cpick',{$scope:scope});
           scope.$digest();
           scope.saveTheme();
           expect(localStorage.getItem('mswift42themecreator')).toBeDefined();
           scope.resetTheme();
           expect(localStorage.getItem('mswift42themecreator')).toBe(null);
       }));
                                       





});
