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
        storageService.deleteKey = function() {
            localStorage.removeItem(storageService.name);
        };
        return storageService;
    })
    .factory('presets', function() {

        var presetService = {};
        presetService.warmNight = function () {
            return {deffacefg : "#b1b1b1", deffacebg : "#292424",
                        keywordface : "#96905f", builtinface : "#71a46c",
                        stringface : "#71a19", functionnameface : "#b680b1",
                        typeface : "#8b8fc6", constantface : "#bd845f",
                        variableface : "#c27d7b", warningface : "#e81050",
                    commentface : "#5d5a58"
                   };
        };
        presetService.oldlace = function() {
            return {deffacefg : "#525252", deffacebg : "#fdf5e6",
                    keywordface : "#3f567b", builtinface : "#7b4135",
                        stringface : "#305f5e", functionnameface : "#714355",
                        typeface : "#634575", constantface : "#64502f",
                        variableface : "#3f5b32", warningface : "#fa0c0c",
                    commentface : "#949494"
                   };
        };
        presetService.
    });
