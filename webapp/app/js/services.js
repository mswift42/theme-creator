'use strict';

/* Services */


// Demonstrate how to register services
// In this case it is a simple value service.
angular.module('myApp.services',[])
    .factory('lstorage', function() {
        
        var storageService = {};
        storageService.name = 'mswift42themecreator';
        storageService.containsKey = function() {
            return localStorage.getItem(storageService.name) != null;
        };
        storageService.loadFaces = function() {
            return JSON.parse(localStorage.getItem(storageService.name));
        };
        storageService.setFaces = function(faces) {
            localStorage.setItem(storageService.name, JSON.stringify(faces));
        };
        return storageService;
    });
