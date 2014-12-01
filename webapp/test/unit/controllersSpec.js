'use strict';

/* jasmine specs for controllers go here */

describe('controllers', function(){
    beforeEach(module('myApp.controllers'));


  it('should ....', inject(function($controller) {
    //spec body
    var Cpick = $controller('Cpick', { $scope: {} });
      expect(Cpick).toBeDefined();
  }));

  // it('should ....', inject(function($controller) {
  //   //spec body
  //   var myCtrl2 = $controller('MyCtrl2', { $scope: {} });
  //   expect(myCtrl2).toBeDefined();
  // }));
});
