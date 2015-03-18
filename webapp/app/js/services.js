'use strict';

/* Services */


// Demonstrate how to register services
// In this case it is a simple value service.
angular.module('myApp.services',[])
    .factory('lstorage', function() {
        
        var storageService = {};
        storageService.name = 'mswift42themecreator';
        storageService.loadFaces = function() {
            localStorage.getItem(storageService.name);
        };
        storageService.setFaces = function(faces) {
            localStorage.setItem(storageService.name, faces);
        };
        return storageService;
    });
